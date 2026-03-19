package network

import (
	"encoding/binary"
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
	const headerSize = 6

	if len(packet) < headerSize {
		log.Errorf("packet too short: len=%d addr=%s", len(packet), ctx.RemoteAddr())
		return 0
	}

	length := int32(binary.BigEndian.Uint32(packet[:4]))
	msgId := int16(binary.BigEndian.Uint16(packet[4:6]))

	if int(length) > len(packet) {
		log.Errorf("invalid length: length=%d packet_len=%d addr=%s", length, len(packet), ctx.RemoteAddr())
		return 0
	}

	if tcp.HasMethod(msgId) {
		body := packet[headerSize:]
		tcp.CallMethodWithChannelContext(msgId, ctx, body)
		if len(body) > 0 {
			log.Infof("receive msgId=%d length=%d", msgId, len(body))
		}
		return 0
	}

	return serverNetWork.forwardToGameServer(msgId, length, packet, ctx)
}

// forwardToGameServer 将消息转发到游戏服务器
func (serverNetWork *ServerEventHandler) forwardToGameServer(msgId int16, length int32, packet []byte, ctx tcp.Channel) int {
	ctxPlayer, ok := ctx.Context().(*player.Player)
	if !ok || ctx.Context() == nil {
		log.Errorf("invalid context for msgId=%d", msgId)
		return 0
	}

	headMsg := &protoGen.InnerHead{Id: ctxPlayer.Pid}
	headerBytes, _ := proto.Marshal(headMsg)
	body := packet[6:]

	// 构造转发数据包: len(4) + msgId(2) + headerLen(2) + header + body
	totalLen := 4 + 2 + 2 + len(headerBytes) + int(length) - 2
	buf := make([]byte, totalLen)
	binary.BigEndian.PutUint32(buf[0:4], uint32(totalLen-4))
	binary.BigEndian.PutUint16(buf[4:6], uint16(msgId))
	binary.BigEndian.PutUint16(buf[6:8], uint16(len(headerBytes)))
	copy(buf[8:8+len(headerBytes)], headerBytes)
	copy(buf[8+len(headerBytes):], body)

	client.SendPckToGameServer(buf)
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (serverNetWork *ServerEventHandler) Tick() (delay time.Duration, action int) {
	if player.PlayerMgr.GetSize() == 0 {
		return 5000 * time.Millisecond, 0
	}
	//sample to dos something
	player.PlayerMgr.Range(func(player *player.Player) {
		response := &protoGen.HeartBeatResponse{
			ClientTime: time.Now().UnixMilli(),
			ServerTime: time.Now().UnixMilli(),
		}
		player.Context.SendMsg(protoGen.ProtoCode_HEART_BEAT_RESPONSE, response)
	})
	return 5000 * time.Millisecond, 0
}
