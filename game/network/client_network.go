package network

import (
	"encoding/binary"
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

type ClientEventHandler struct {
}

func (clientNetwork *ClientEventHandler) OnOpened(c tcp.Channel) (out []byte, action int) {
	//context := client.NewClientContext(c)
	log.Infof("----------  client opened  addr=%s", c.RemoteAddr())
	clientContext := client.NewInnerClientContext(c)
	c.SetContext(clientContext)
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (clientNetwork *ClientEventHandler) OnClosed(c tcp.Channel, err error) (action int) {
	context := c.Context().(*client.ConnInnerClientContext)
	log.Infof("XXXXXXXXXXXXXXXXXXXX  client closed addr =%s id =%d", c.RemoteAddr(), context.Sid)
	return 1

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (clientNetwork *ClientEventHandler) PreWrite(c tcp.Channel) {
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

	length := binary.BigEndian.Uint32(packet[:packetLenHeaderSize])
	if length > uint32(len(packet)) {
		log.Errorf("invalid packet length: headerLen=%d packetLen=%d addr=%s", length, len(packet), ctx.RemoteAddr())
		return 0
	}

	// Read msgId
	msgId := int16(binary.BigEndian.Uint16(packet[msgIdOffset:headerLenOffset]))

	// Direct unhandled messages to gateway
	if !tcp.HasMethod(msgId) {
		client.GetInnerClient(global.GATE_WAY).SendBytesMsg(packet)
		return 0
	}

	// Read header length
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

	// Dispatch message to handlers
	if !dispatchClientMessage(msgId, innerMsg.Id, ctx, payload) {
		log.Errorf("message not handled: msgId=%d playerId=%d addr=%s", msgId, innerMsg.Id, ctx.RemoteAddr())
	}
	return 0
}

// dispatchClientMessage routes the message to the appropriate handler
func dispatchClientMessage(msgId int16, playerId int64, ctx tcp.Channel, payload []byte) bool {
	if tcp.CallMethodWithChannelContext(msgId, ctx, payload) {
		return true
	}
	return tcp.CallMethodWithRoleId(msgId, playerId, payload)
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientEventHandler) Tick() (delay time.Duration, action int) {
	innerClient := client.GetInnerClient(global.ROUTER)
	if innerClient == nil {
		//log.Infof("no found connect type =%d", client.ROUTER)
		return 5000 * time.Millisecond, 0
	}
	heartBeat := &protoGen.InnerHeartBeatRequest{}
	innerClient.SendInnerMsg(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ, 0, heartBeat)
	//.Infof("send inner hear beat = %s", innerClient.Ctx.RemoteAddr())
	return 5000 * time.Millisecond, 0
}
