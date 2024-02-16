package networkHandler

import (
	"bytes"
	"encoding/binary"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"time"
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
	case *client.ConnInnerClientContext:
		log.Infof(" OnClosed-----addr =%s not login", c.RemoteAddr())
		return 0
	case *player.Player:
		player := c.Context().(*player.Player)
		log.Infof("conn =%s  sid=%d pid=%d  closed", c.RemoteAddr(), player.Context.Sid, player.Pid)
		return 0
	default:
		return 0

	}

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (serverNetWork *ServerEventHandler) PreWrite(c tcp.Channel) {
	//	log.Infof("conn =%s PreWrite", c.RemoteAddr())

}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (serverNetWork *ServerEventHandler) AfterWrite(c tcp.Channel, b []byte) {
	//	log.Infof("conn =%s AfterWrite", c.RemoteAddr())
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
	bytebuffer := bytes.NewBuffer(packet[4:])
	var msgId int16
	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	if msgId == int16(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ) {
		return 0
	}

	var innerHeaderLen int16
	binary.Read(bytebuffer, binary.BigEndian, &innerHeaderLen)
	innerMsg := &protoGen.InnerHead{}
	innerBody := make([]byte, innerHeaderLen)
	binary.Read(bytebuffer, binary.BigEndian, innerBody)
	proto.Unmarshal(innerBody, innerMsg)

	processed := tcp.CallMethodWithChannelContext(msgId, ctx, bytebuffer.Bytes())
	if processed {
		return 0
	}

	tcp.CallMethodWithRoleId(msgId, innerMsg.Id, bytebuffer.Bytes())
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {
	return 1000 * time.Millisecond, 0
}
