package network

import (
	"bytes"
	"encoding/binary"
	"gameSrv/game/executor"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

type ServerEventHandler struct {
}

func (serverNetWork *ServerEventHandler) OnOpened(c tcp.Channel) (out []byte, action int) {
	clientContext := client.NewInnerClientContext(c)
	c.SetContext(clientContext)
	log.Infof("new connect addr =%s  id=%d", clientContext.Ctx.RemoteAddr(), clientContext.Sid)
	//test for worker pool
	//workerPool := gopool.StartNewWorkerPool(2, 4)
	//workerPool.SubmitTask(func() {
	//	log.Infof("XXXXXXXXXXX  execute task come from remoteAddr=%s", clientContext.Ctx.RemoteAddr())
	//})
	log.Infof("pppppppppppppppp sid=%d", clientContext.Sid)
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (serverNetWork *ServerEventHandler) OnClosed(c tcp.Channel, err error) (action int) {
	switch c.Context().(type) {
	case *client.ConnContext:
		log.Infof("addr =%s not login", c.RemoteAddr())
		return 1
	case *player.Player:
		player := c.Context().(*player.Player)
		log.Infof("conn =%s  sid=%d pid=%d  closed", c.RemoteAddr(), player.Context.Sid, player.Pid)
		return 1
	case *client.ConnInnerClientContext:
		log.Infof("addr =%s innerClient disconnected", c.RemoteAddr())
		return 1
	default:
		return 1
	}
}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (serverNetWork *ServerEventHandler) PreWrite(c tcp.Channel) {
	//log.Infof("conn =%s PreWrite", c.RemoteAddr())

}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (serverNetWork *ServerEventHandler) AfterWrite(c tcp.Channel, b []byte) {
	//log.Infof("conn =%s AfterWrite", c.RemoteAddr())
}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (serverNetWork *ServerEventHandler) React(packet []byte, ctx tcp.Channel) (action int) {
	// Packet format: [4 bytes len] [2 bytes msgId] [2 bytes headerLen] [header] [payload]
	const msgIdOffset = 4
	const headerOffset = 8

	buf := bytes.NewBuffer(packet[msgIdOffset:])

	var msgId int16
	if err := binary.Read(buf, binary.BigEndian, &msgId); err != nil {
		log.Errorf("read msgId failed: addr=%s len=%d err=%v", ctx.RemoteAddr(), len(packet), err)
		return 0
	}

	// Ignore heartbeat messages
	if msgId == int16(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ) {
		return 0
	}

	//log.Infof("receive msgId=%d addr=%s len=%d", msgId, ctx.RemoteAddr(), len(packet))

	// Direct unhandled messages to world server
	if !tcp.HasMethod(msgId) {
		log.Infof("msgId=%d direct to world server", msgId)
		client.SendPckToRouterServer(packet)
		return 0
	}

	var headerLen int16
	if err := binary.Read(buf, binary.BigEndian, &headerLen); err != nil {
		log.Errorf("read headerLen failed: msgId=%d addr=%s err=%v", msgId, ctx.RemoteAddr(), err)
		return 0
	}

	headerEnd := headerOffset + int(headerLen)
	if headerEnd > len(packet) {
		log.Errorf("invalid packet: msgId=%d headerEnd=%d exceeds packetLen=%d", msgId, headerEnd, len(packet))
		return 0
	}

	innerMsg := &protoGen.InnerHead{}
	headerBytes := packet[headerOffset:headerEnd]
	if err := proto.Unmarshal(headerBytes, innerMsg); err != nil {
		log.Errorf("unmarshal header failed: msgId=%d addr=%s err=%v", msgId, ctx.RemoteAddr(), err)
		return 0
	}

	playerId := innerMsg.GetId()
	log.Infof("process msgId=%d playerId=%d", msgId, playerId)

	hashCode := callPlayerMsgHashCode(msgId, playerId)
	payload := packet[headerEnd:]
	dispatchMessage(msgId, playerId, ctx, payload, hashCode)
	return 0
}

// dispatchMessage routes the message to the appropriate handler
func dispatchMessage(msgId int16, playerId int64, ctx tcp.Channel, payload []byte, hashCode int64) {
	pidMethod := tcp.GetCallMethodWithId(msgId)
	if pidMethod != nil {
		body, err := pidMethod.Parse(payload)
		if err != nil {
			log.Errorf("unmarshal header failed: msgId=%d addr=%s err=%v", msgId, ctx.RemoteAddr(), err)
			return
		}
		executor.AsyncNetMsgExecutor.SubmitTask(hashCode, func() {
			pidMethod.Execute(playerId, body)
		})
		return
	}
	channelMethod := tcp.GetMethodWithChannel(msgId)
	if channelMethod != nil {
		body, err := channelMethod.Parse(payload)
		if err != nil {
			log.Errorf("unmarshal header failed: msgId=%d addr=%s err=%v", msgId, ctx.RemoteAddr(), err)
			return
		}
		executor.AsyncNetMsgExecutor.SubmitTask(hashCode, func() {
			channelMethod.Execute(ctx, body)
		})
	}
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {
	return 1000 * time.Millisecond, 0
}

func callPlayerMsgHashCode(msgId int16, playerId int64) int64 {
	return playerId
}
