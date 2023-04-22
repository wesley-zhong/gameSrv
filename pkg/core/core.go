package core

import (
	"gameSvr/pkg/log"
	"gameSvr/pkg/network"
	"google.golang.org/protobuf/proto"
)

type MsgIdFuc[T any, T2 any] func(T, T2)

var msgIdMap = make(map[int32]*protoMethod)

type protoMethod struct {
	methodFuc MsgIdFuc[network.ChannelContext, proto.Message]
	param     proto.Message
}

func RegisterMethod(msgId int32, param proto.Message, fuc MsgIdFuc[network.ChannelContext, proto.Message]) {
	method := &protoMethod{
		methodFuc: fuc,
		param:     param,
	}
	msgIdMap[msgId] = method
}

func Init() {

}
func CallMethod(msgId int32, body []byte, ctx network.ChannelContext) {
	method := msgIdMap[msgId]
	if method == nil {
		log.Infof("msgId = %d not found method", msgId)
		return
	}
	proto.Unmarshal(body, method.param)
	method.methodFuc(ctx, method.param)
}
