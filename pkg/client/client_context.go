package client

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/pkg/utils"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

// ConnContext ========================== user client -==========================================================
type ConnContext struct {
	Ctx tcp.Channel
	Sid int64
}

func NewClientContext(ctx tcp.Channel) *ConnContext {
	return &ConnContext{ctx, utils.NextId()}
}

func ClientConnect(addr string) tcp.Channel {
	channel, err := tcp.Dial("tcp", addr)
	if err != nil {
		log.Error(err)
		return nil
	}
	return channel
}

func (client *ConnContext) SendMsg(code protoGen.ProtoCode, body proto.Message) {
	packet := &tcp.MsgPacket{
		MsgId: int16(code),
		Body:  body,
	}
	encode, err := ICodec.Encode(packet)
	if err != nil {
		log.Error(err)
		return
	}
	client.Ctx.SendTo(encode)
}

func (client *ConnContext) Send(body []byte) {
	client.Ctx.SendTo(body)
}
