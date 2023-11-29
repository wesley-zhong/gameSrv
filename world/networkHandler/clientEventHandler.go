package networkHandler

import (
	"bytes"
	"encoding/binary"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	msg "gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

type ClientEventHandler struct {
}

func (clientNetwork *ClientEventHandler) OnOpened(c network.ChannelContext) (out []byte, action int) {
	context := client.NewClientContext(c)
	log.Infof("----------inner  client opened  addr=%s, id=%d", context.Ctx.RemoteAddr(), context.Sid)
	return nil, 0
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (clientNetwork *ClientEventHandler) OnClosed(c network.ChannelContext, err error) (action int) {
	context := c.Context().(*client.ConnInnerClientContext)
	log.Infof("XXXXXXXXXXXXXXXXXXXX  client closed addr =%s id =%d", c.RemoteAddr(), context)
	return 1

}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (clientNetwork *ClientEventHandler) PreWrite(c network.ChannelContext) {
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (clientNetwork *ClientEventHandler) AfterWrite(c network.ChannelContext, b []byte) {

}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (clientNetwork *ClientEventHandler) React(packet []byte, ctx network.ChannelContext) (out []byte, action int) {
	log.Infof("  client React receive addr =%s", ctx.RemoteAddr())
	var innerHeaderLen int32
	bytebuffer := bytes.NewBuffer(packet)
	binary.Read(bytebuffer, binary.BigEndian, &innerHeaderLen)
	innerMsg := &msg.InnerHead{}
	innerBody := make([]byte, innerHeaderLen)
	binary.Read(bytebuffer, binary.BigEndian, innerBody)

	proto.Unmarshal(innerBody, innerMsg)

	//body := make([]byte, bytebuffer.Len())
	//binary.Read(bytebuffer, binary.BigEndian, body)
	bodyLen := bytebuffer.Len()
	if bodyLen > 0 {
		body := make([]byte, bodyLen)
		binary.Read(bytebuffer, binary.BigEndian, body)
	}

	//servers internal  system call
	if innerMsg.Id == 0 {
		if bodyLen == 0 {
			core.CallMethod(innerMsg.ProtoCode, nil, ctx)
			return nil, 0
		}
		log.Infof("------#########receive msgId = %d length =%d", innerMsg.ProtoCode, bodyLen)
		core.CallMethod(innerMsg.ProtoCode, packet[innerHeaderLen+4:], ctx)
		return nil, 0
	}
	// server send player msg
	if bodyLen == 0 {
		//core.CallMethod(innerMsg.ProtoCode, nil, ctx)
		core.CallMethodWitheRoleId(innerMsg.ProtoCode, innerMsg.Id, nil)
		return nil, 0
	}
	log.Infof("------#########receive msgId = %d length =%d", innerMsg.ProtoCode, bodyLen)
	core.CallMethodWitheRoleId(innerMsg.ProtoCode, innerMsg.Id, packet[innerHeaderLen+4:])
	return nil, 0
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (clientNetwork *ClientEventHandler) Tick() (delay time.Duration, action int) {
	return 1000 * time.Millisecond, 0
}
