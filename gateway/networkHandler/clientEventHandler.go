package networkHandler

import (
	"bytes"
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

func (clientNetwork *ClientEventHandler) OnOpened(c tcp.ChannelContext) (out []byte, action int) {
	//	context := client.NewClientContext(c)
	log.Infof("----------  client opened  addr=%s", c.RemoteAddr())
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (clientNetwork *ClientEventHandler) OnClosed(c tcp.ChannelContext, err error) (action int) {
	context := c.Context().(*client.ConnInnerClientContext)
	log.Infof("XXXXXXXXXXXXXXXXXXXX  client closed addr =%s id =%d", c.RemoteAddr(), context.Sid)
	return 1

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (clientNetwork *ClientEventHandler) PreWrite(c tcp.ChannelContext) {
	//log.Infof("pppppppppppppppppp")
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (clientNetwork *ClientEventHandler) AfterWrite(c tcp.ChannelContext, b []byte) {

}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (clientNetwork *ClientEventHandler) React(packet []byte, ctx tcp.ChannelContext) (action int) {
	bytebuffer := bytes.NewBuffer(packet)
	var totalMsgLen int32
	binary.Read(bytebuffer, binary.BigEndian, &totalMsgLen)
	var msgId int16
	binary.Read(bytebuffer, binary.BigEndian, &msgId)

	var innerHeaderLen int16
	binary.Read(bytebuffer, binary.BigEndian, &innerHeaderLen)
	innerMsg := &protoGen.InnerHead{}
	innerBody := make([]byte, innerHeaderLen)
	binary.Read(bytebuffer, binary.BigEndian, innerBody)
	proto.Unmarshal(innerBody, innerMsg)

	exist := tcp.HasMethod(msgId)
	if !exist {
		//direct to send client
		existPlayer := player.PlayerMgr.GetByRoleId(innerMsg.Id)
		if existPlayer == nil {
			log.Warnf("XXXXXXXX pid = %d not found", innerMsg.Id)
		}
		skipLen := 2 + int32(innerHeaderLen)

		msgLen := totalMsgLen - skipLen
		toPlayerBody := make([]byte, msgLen+4)
		toPlayerBodyBuf := bytes.NewBuffer(toPlayerBody)
		toPlayerBodyBuf.Reset()

		binary.Write(toPlayerBodyBuf, binary.BigEndian, msgLen)
		binary.Write(toPlayerBodyBuf, binary.BigEndian, msgId)
		binary.Write(toPlayerBodyBuf, binary.BigEndian, bytebuffer.Bytes())
		existPlayer.Context.Send(toPlayerBodyBuf.Bytes())
		log.Infof("roleId %d msgId =%d not found method direct to send client ", innerMsg.Id, msgId)
		return 0
	}

	processed := tcp.CallMethodWithChannelContext(msgId, ctx, bytebuffer.Bytes())
	if processed {
		return 0
	}

	tcp.CallMethodWithRoleId(msgId, innerMsg.Id, bytebuffer.Bytes())
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientEventHandler) Tick() (delay time.Duration, action int) {
	innerClient := client.GetInnerClient(client.GAME)
	if innerClient == nil {
		//	log.Infof("no found connect type =%d", client.GAME)
		return 5000 * time.Millisecond, 0
	}
	heartBeat := &protoGen.InnerHeartBeatRequest{}
	innerClient.SendInnerMsg(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ, 0, heartBeat)
	return 5000 * time.Millisecond, 0
}
