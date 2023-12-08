package networkHandler

import (
	"bytes"
	"encoding/binary"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"time"
)

type ServerEventHandler struct {
}

func (serverNetWork *ServerEventHandler) OnOpened(c network.ChannelContext) (out []byte, action int) {
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
func (serverNetWork *ServerEventHandler) OnClosed(c network.ChannelContext, err error) (action int) {
	switch c.Context().(type) {
	case *client.ConnClientContext:
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
func (serverNetWork *ServerEventHandler) PreWrite(c network.ChannelContext) {
	//log.Infof("conn =%s PreWrite", c.RemoteAddr())

}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (serverNetWork *ServerEventHandler) AfterWrite(c network.ChannelContext, b []byte) {
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
func (serverNetWork *ServerEventHandler) React(packet []byte, ctx network.ChannelContext) (out []byte, action int) {
	var headerSize int32
	bytebuffer := bytes.NewBuffer(packet)
	binary.Read(bytebuffer, binary.BigEndian, &headerSize)
	if headerSize < 0 || headerSize > int32(network.MaxPackageLen) {
		log.Warnf("XXXXXXXXXX headerSize =%d too large addr =%s", headerSize, ctx.RemoteAddr())
		ctx.Close()
		return nil, 0
	}
	headerBody := make([]byte, headerSize)
	binary.Read(bytebuffer, binary.BigEndian, headerBody)
	innerHeader := &protoGen.InnerHead{}
	err := proto.Unmarshal(headerBody, innerHeader)
	if err != nil {
		log.Error(err)
		return nil, 0
	}
	if innerHeader.ProtoCode == int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ) {
		ctx.Context().(*client.ConnInnerClientContext).SendInnerMsgProtoCode(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES, 0, &protoGen.InnerHeartBeatResponse{})
		return nil, 0
	}
	bodyLen := bytebuffer.Len()

	//servers internal  system call
	if innerHeader.Id == 0 {
		if bodyLen == 0 {
			core.CallMethod(innerHeader.ProtoCode, nil, ctx)
			return nil, 0
		}

		log.Infof("------#########receive msgId = %d length =%d", innerHeader.ProtoCode, bodyLen)
		core.CallMethod1(innerHeader.ProtoCode, packet[headerSize+4:], ctx)
		return nil, 0
	}
	// server send player msg
	if bodyLen == 0 {
		//core.CallMethod(innerHeader.ProtoCode, nil, ctx)
		core.CallMethodWitheRoleId(innerHeader.ProtoCode, innerHeader.Id, nil)
		return nil, 0
	}
	log.Infof("------#########receive msgId = %d length =%d", innerHeader.ProtoCode, bodyLen)
	core.CallMethodWitheRoleId(innerHeader.ProtoCode, innerHeader.Id, packet[headerSize+4:])
	return nil, 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {
	return 1000 * time.Millisecond, 0
}
