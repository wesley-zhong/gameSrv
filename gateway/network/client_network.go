package network

import (
	"encoding/binary"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

type ClientEventHandler struct {
	codec tcp.ICodec
}

func (clientNetwork *ClientEventHandler) OnOpened(c tcp.Channel) (out []byte, action int) {
	//	context := client.NewClientContext(c)
	log.Infof("----------  client opened  addr=%s", c.RemoteAddr())
	clientContext := client.NewInnerClientContext(c)
	c.SetContext(clientContext)
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (clientNetwork *ClientEventHandler) OnClosed(c tcp.Channel, err error) (action int) {
	log.Infof("XXXXXXXXXXXXXXXXXXXX  client closed addr =%s ", c.RemoteAddr())
	if c.Context() == nil {
		return 0
	}

	return 1

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (clientNetwork *ClientEventHandler) PreWrite(c tcp.Channel) {
	//log.Infof("pppppppppppppppppp")
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (clientNetwork *ClientEventHandler) AfterWrite(c tcp.Channel, b []byte) {

}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (clientNetwork *ClientEventHandler) React(packet []byte, ctx tcp.Channel) (action int) {
	// Packet format: [4 bytes len] [2 bytes msgId] [2 bytes headerLen] [header] [payload]
	const (
		packetLenHeaderSize = 4
		msgIdOffset         = packetLenHeaderSize
		headerLenOffset     = msgIdOffset + 2
		headerOffset        = headerLenOffset + 2
	)

	// Validate packet length
	if len(packet) < headerOffset {
		log.Errorf("packet too short: len=%d addr=%s", len(packet), ctx.RemoteAddr())
		return 0
	}

	totalMsgLen := int32(binary.BigEndian.Uint32(packet[:packetLenHeaderSize]))
	if totalMsgLen < 0 || int(totalMsgLen) > len(packet) {
		log.Errorf("invalid totalMsgLen: totalMsgLen=%d packetLen=%d addr=%s", totalMsgLen, len(packet), ctx.RemoteAddr())
		return 0
	}

	msgId := int16(binary.BigEndian.Uint16(packet[msgIdOffset:headerLenOffset]))
	headerLen := int16(binary.BigEndian.Uint16(packet[headerLenOffset:headerOffset]))

	headerEnd := headerOffset + int(headerLen)
	if headerEnd > len(packet) {
		log.Errorf("invalid headerLen: msgId=%d headerLen=%d packetLen=%d addr=%s",
			msgId, headerLen, len(packet), ctx.RemoteAddr())
		return 0
	}

	// Unmarshal inner header directly from packet (zero-copy)
	innerMsg := &protoGen.InnerHead{}
	headerBytes := packet[headerOffset:headerEnd]
	if err := proto.Unmarshal(headerBytes, innerMsg); err != nil {
		log.Errorf("unmarshal header failed: msgId=%d addr=%s err=%v", msgId, ctx.RemoteAddr(), err)
		return 0
	}

	// Payload is the remaining bytes after header (zero-copy)
	payload := packet[headerEnd:]

	if !tcp.HasMethod(msgId) {
		// Forward to player (zero-copy: modify packet in-place and move payload)
		p := player.PlayerMgr.GetByRoleId(innerMsg.Id)
		if p == nil {
			log.Warnf("player not found: playerId=%d", innerMsg.Id)
			return 0
		}

		// msgLen = payload length + msgId(2)
		// Original: totalMsgLen = msgId(2) + headerLen(2) + header + payload
		// So payload length = totalMsgLen - 4 - headerLen
		payloadLen := totalMsgLen - 4 - int32(headerLen)
		if payloadLen < 0 {
			log.Errorf("invalid payloadLen: msgId=%d totalMsgLen=%d headerLen=%d addr=%s",
				msgId, totalMsgLen, headerLen, ctx.RemoteAddr())
			return 0
		}

		// Zero-copy: move payload to position after msgId
		// Original: [4B len] [2B msgId] [2B headerLen] [header] [payload]
		// After:  [4B len] [2B msgId] [payload]
		// Move payload from headerEnd to position 6
		copy(packet[6:6+payloadLen], payload)

		// Update length field
		newLen := 2 + payloadLen // msgId + payload
		binary.BigEndian.PutUint32(packet[:4], uint32(newLen))

		// Send modified packet (no new allocation)
		p.Context.Send(packet[:4+newLen])
		return 0
	}

	// Dispatch message to handlers
	if tcp.CallMethodWithChannelContext(msgId, ctx, payload) {
		return 0
	}
	tcp.CallMethodWithPlayerId(msgId, innerMsg.Id, payload)
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientEventHandler) Tick() (delay time.Duration, action int) {
	heartBeat := &protoGen.InnerHeartBeatRequest{}
	client.SendInnerToGameServer(0, protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ, heartBeat)
	return 5000 * time.Millisecond, 0
}
