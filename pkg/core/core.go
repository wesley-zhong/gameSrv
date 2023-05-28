package core

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"google.golang.org/protobuf/proto"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

var msgIdContextMap = make(map[int32]*protoMethod[network.ChannelContext])
var msgIdRoleIdMap = make(map[int32]*protoMethod[int64])

type protoMethod[T1 any] struct {
	methodFuc MsgIdFuc[T1, proto.Message]
	param     proto.Message
}

func RegisterMethod(msgId int32, param proto.Message, fuc MsgIdFuc[network.ChannelContext, proto.Message]) {
	method := &protoMethod[network.ChannelContext]{
		methodFuc: fuc,
		param:     param,
	}
	msgIdContextMap[msgId] = method
}

func RegisterCallPlayerMethod(msgId int32, param proto.Message, fuc MsgIdFuc[int64, proto.Message]) {
	method := &protoMethod[int64]{
		methodFuc: fuc,
		param:     param,
	}
	msgIdRoleIdMap[msgId] = method
}

func CallMethodWitheRoleId(msgId int32, roleId int64, body []byte, ctx network.ChannelContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Infof("=======Recovered:", r)
		}
	}()
	method := msgIdRoleIdMap[msgId]
	if method == nil {
		log.Infof("msgId = %d not found method", msgId)
		return
	}
	param := method.param.ProtoReflect().New().Interface()
	if body != nil {
		proto.Unmarshal(body, param)
	}
	method.methodFuc(roleId, param)
}

func CallMethod(msgId int32, body []byte, ctx network.ChannelContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Infof("=======Recovered:", r)
		}
	}()
	method := msgIdContextMap[msgId]
	if method == nil {
		log.Infof("msgId = %d not found method", msgId)
		return
	}
	param := method.param.ProtoReflect().New().Interface()
	if body != nil {
		proto.Unmarshal(body, param)
	}
	method.methodFuc(ctx, param)
}
