package client

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/pkg/utils"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

var _ICodec = &tcp.DefaultCodec{}

func ClientConnect(addr string) tcp.Channel {
	channel, err := tcp.Dial("tcp", addr)
	if err != nil {
		log.Error(err)
		return nil
	}
	return channel
}

type ConnContext struct {
	Ctx tcp.Channel
	Sid int64
}

func NewClientContext(ctx tcp.Channel) *ConnContext {
	return &ConnContext{
		Ctx: ctx,
		Sid: utils.NextId(),
	}
}

func (c *ConnContext) SendMsg(code protoGen.MsgId, body proto.Message) {
	packet := &tcp.MsgPacket{
		MsgId: int16(code),
		Body:  body,
	}
	encode, err := _ICodec.Encode(packet)
	if err != nil {
		log.Error(err)
		return
	}
	c.Ctx.SendTo(encode)
}

func (c *ConnContext) Send(data []byte) {
	c.Ctx.SendTo(data)
}
