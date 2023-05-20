package core

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"google.golang.org/protobuf/proto"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

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

func CallMethod(msgId int32, body []byte, ctx network.ChannelContext) {
	method := msgIdMap[msgId]
	if method == nil {
		log.Infof("msgId = %d not found method", msgId)
		return
	}
	param := method.param.ProtoReflect().New().Interface()
	proto.Unmarshal(body, param)
	defer func() {
		if r := recover(); r != nil {
			log.Infof("=======Recovered:", r)
		}
	}()
	method.methodFuc(ctx, param)
}
