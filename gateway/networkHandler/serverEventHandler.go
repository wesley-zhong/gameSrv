package networkHandler

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gameSrv/gateway/controller"
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
	clientContext := client.NewClientContext(c)
	c.SetContext(clientContext)
	//log.Infof("new connect addr =%s  id=%d", clientContext.Ctx.RemoteAddr(), clientContext.Sid)
	//test for worker pool
	//workerPool := gopool.StartNewWorkerPool(2, 4)
	//workerPool.SubmitTask(func() {
	//	log.Infof("XXXXXXXXXXX  execute task come from remoteAddr=%s", clientContext.Ctx.RemoteAddr())
	//})
	//	log.Infof("pppppppppppppppp sid=%d", clientContext.Sid)
	log.Infof("new connect addr =%s ", c.RemoteAddr())
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
		controller.ClientDisConnect(c)
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
	bytebuffer := bytes.NewBuffer(packet)
	var length uint32
	binary.Read(bytebuffer, binary.BigEndian, &length)

	var msgId int16
	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	//	log.Infof("------receive msgId = %d length =%d", msgId, length)
	hasMethod := tcp.HasMethod(msgId)
	if !hasMethod {
		// direct to game server
		if ctx.Context() == nil {
			log.Error(errors.New(fmt.Sprintf("msgId = %d error", msgId)))
			return 0
		}

		player := ctx.Context().(*player.Player)
		headMsg := &protoGen.InnerHead{Id: player.Pid}
		headerBytes, _ := proto.Marshal(headMsg)
		totalLen := 4 + 2 + 2 + len(headerBytes) + int(length) - 2

		directSendByte := make([]byte, totalLen)
		directSendBuff := bytes.NewBuffer(directSendByte)
		directSendBuff.Reset()
		binary.Write(directSendBuff, binary.BigEndian, int32(totalLen-4))
		binary.Write(directSendBuff, binary.BigEndian, msgId)
		binary.Write(directSendBuff, binary.BigEndian, int16(len(headerBytes)))
		binary.Write(directSendBuff, binary.BigEndian, headerBytes)
		binary.Write(directSendBuff, binary.BigEndian, packet[6:])
		client.GetInnerClient(client.GAME).SendBytesMsg(directSendBuff.Bytes())

		//log.Infof("-------- msgId =%d direct to game server  total len =%d bytes =%d", msgId, int32(totalLen-4), len(directSendBuff.Bytes()))
		return 0
	}

	bodyLen := bytebuffer.Len()
	if bodyLen == 0 {
		tcp.CallMethodWithChannelContext(msgId, ctx, nil)
		return 0
	}
	log.Infof("------#########receive msgId = %d length =%d", msgId, bodyLen)
	tcp.CallMethodWithChannelContext(msgId, ctx, packet[6:])
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {
	if player.PlayerMgr.GetSize() == 0 {
		return 5000 * time.Millisecond, 0
	}
	//sample to do something
	player.PlayerMgr.Range(func(player *player.Player) {
		response := &protoGen.HeartBeatResponse{
			ClientTime: time.Now().UnixMilli(),
			ServerTime: time.Now().UnixMilli(),
		}
		player.Context.SendMsg(protoGen.ProtoCode_HEART_BEAT_RESPONSE, response)
	})
	return 5000 * time.Millisecond, 0
}
