package client

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

// ConnContext ========================== user client -==========================================================
type ConnContext struct {
	Ctx tcp.Channel
	Sid int64
}

// NewClientContext - ------ user client -------------------
func NewClientContext(context tcp.Channel) *ConnContext {
	return &ConnContext{Ctx: context, Sid: genSid()}
}

func ClientConnect(addr string) *ConnContext {
	context, err := tcp.Dial("tcp", addr)
	if err != nil {
		log.Error(err)
		return nil
	}
	clientContext := NewClientContext(context)
	context.SetContext(clientContext)
	return clientContext
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
