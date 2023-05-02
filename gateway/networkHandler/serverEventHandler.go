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
	"time"
)

type ServerEventHandler struct {
}

func (serverNetWork *ServerEventHandler) OnOpened(c network.ChannelContext) (out []byte, action int) {
	clientContext := client.NewClientContext(c)
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
		disConnPlayer := c.Context().(*player.Player)
		player.PlayerMgr.Remove(disConnPlayer)
		log.Infof("conn =%s  sid=%d pid=%d  closed now playerCount=%d",
			c.RemoteAddr(), disConnPlayer.Context.Sid, disConnPlayer.Pid, player.PlayerMgr.GetSize())
		return 1
	default:
		return 1
	}
}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (serverNetWork *ServerEventHandler) PreWrite(c network.ChannelContext) {
	log.Infof("conn =%s PreWrite", c.RemoteAddr())
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (serverNetWork *ServerEventHandler) AfterWrite(c network.ChannelContext, b []byte) {
	log.Infof("conn =%s AfterWrite", c.RemoteAddr())
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
	var msgId int32
	bytebuffer := bytes.NewBuffer(packet)
	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	var length uint32
	binary.Read(bytebuffer, binary.BigEndian, &length)
	log.Infof("------receive msgId = %d length =%d", msgId, length)

	core.CallMethod(msgId, packet[8:], ctx)
	return nil, 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {

	playerList := player.PlayerMgr.GetPlayerList()
	if len(playerList) == 0 {
		return 5000 * time.Millisecond, 0
	}

	response := &protoGen.HeartBeatResponse{
		ClientTime: time.Now().UnixMilli(),
		ServerTime: time.Now().UnixMilli(),
	}
	//	PlayerMgr.Get()
	//PlayerMgr.GetByContext(context).Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)

	for _, player := range playerList {

		player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
	}
	return 5000 * time.Millisecond, 0
}
