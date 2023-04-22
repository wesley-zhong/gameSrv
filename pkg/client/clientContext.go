package client

import (
	"bytes"
	"encoding/binary"
	"gameSvr/pkg/log"
	"gameSvr/pkg/network"
	"sync/atomic"

	"google.golang.org/protobuf/proto"
)

var sId int64

func genSid() int64 {
	return atomic.AddInt64(&sId, 1)
}

// ConnInnerClientContext -------------server inner client ---------------
type ConnInnerClientContext struct {
	Ctx network.ChannelContext
	Sid int64
}

func NewInnerClientContext(context network.ChannelContext) *ConnInnerClientContext {
	c := &ConnInnerClientContext{Ctx: context, Sid: genSid()}
	context.SetContext(c)
	return c
}

func (client *ConnInnerClientContext) Send(msg *InnerMessage) {
	header, err := proto.Marshal(msg.InnerHeader)
	if err != nil {
		log.Error(err)
	}

	body, err := proto.Marshal(msg.Body)

	headerLen := len(header)
	bodyLen := 0
	if body != nil {
		bodyLen = len(body)
	}

	msgLen := headerLen + bodyLen + 4
	buffer := bytes.Buffer{}

	buffer.Write(writeInt(msgLen))
	buffer.Write(writeInt(headerLen))
	buffer.Write(header)
	if bodyLen > 0 {
		buffer.Write(body)
	}
	client.Ctx.AsyncWrite(buffer.Bytes())
}

// ConnClientContext ====================================================================================
type ConnClientContext struct {
	Ctx network.ChannelContext
	Sid int64
}

// NewClientContext - ------ user client -------------------
func NewClientContext(context network.ChannelContext) *ConnClientContext {
	return &ConnClientContext{Ctx: context, Sid: genSid()}
}

func (client *ConnClientContext) Send(msgId int32, msg proto.Message) {
	buffer := bytes.Buffer{}
	buffer.Write(writeInt(int(msgId)))
	marshal, err := proto.Marshal(msg)
	if err != nil {
		log.Error(err)
		return
	}
	bodyLen := len(marshal)
	buffer.Write(writeInt(bodyLen))
	buffer.Write(marshal)
	client.Ctx.AsyncWrite(buffer.Bytes())
}

func readInt(byteBuf *bytes.Buffer) int {
	b := make([]byte, 4)
	_, err := byteBuf.Read(b)
	if err != nil {
		return 0
	}
	return int(int32(binary.BigEndian.Uint32(b)))
}

func writeInt(value int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(value))
	return b
}
