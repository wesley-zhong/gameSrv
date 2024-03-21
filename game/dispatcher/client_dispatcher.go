package dispatcher

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gameSrv/pkg/client"
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
	//log.Infof("  client React receive addr =%s", c.RemoteAddr())
	bytebuffer := bytes.NewBuffer(packet[4:])
	var msgId int16
	binary.Read(bytebuffer, binary.BigEndian, &msgId)
	exist := tcp.HasMethod(msgId)
	if !exist {
		//direct to send gateway
		client.GetInnerClient(client.GATE_WAY).SendBytesMsg(packet)
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

	processed = tcp.CallMethodWithRoleId(msgId, innerMsg.Id, bytebuffer.Bytes())
	if processed {
		return 0
	}
	log.Error(errors.New(fmt.Sprintf("msgId =%d  process error ", msgId)))
	return 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientEventHandler) Tick() (delay time.Duration, action int) {
	innerClient := client.GetInnerClient(client.ROUTER)
	if innerClient == nil {
		//log.Infof("no found connect type =%d", client.ROUTER)
		return 5000 * time.Millisecond, 0
	}
	heartBeat := &protoGen.InnerHeartBeatRequest{}
	innerClient.SendInnerMsg(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ, 0, heartBeat)
	//.Infof("send inner hear beat = %s", innerClient.Ctx.RemoteAddr())
	return 5000 * time.Millisecond, 0
}
