package core

import (
	"gameSrv/pkg/gopool"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"google.golang.org/protobuf/proto"
	"runtime/debug"
	"unsafe"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

var msgIdContextMap = make(map[int32]*protoMethod[network.ChannelContext])
var msgIdRoleIdMap = make(map[int32]*protoMethod[int64])

var msgIdMap = make(map[int32]unsafe.Pointer)

var msgGoPool = gopool.StartNewWorkerPool(1, 1024)

type protoMethod[T1 any] struct {
	methodFuc MsgIdFuc[T1, proto.Message]
	param     proto.Message
}

type protoMethod1[T1 any, T2 any] struct {
	param     *T2
	methodCmm func(channelContext network.ChannelContext, body []byte)
}

func (method *protoMethod[T1]) execute(any T1, body []byte) {
	param := method.param.ProtoReflect().New().Interface()
	if body == nil {
		method.methodFuc(any, nil)
	}
	proto.Unmarshal(body, param)
	msgGoPool.SubmitTask(func() {
		method.methodFuc(any, param)
	})
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

func CallMethodWitheRoleId(msgId int32, roleId int64, body []byte) {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Infof("err=%v, stack=%s", r, s)
		}
	}()
	method := msgIdRoleIdMap[msgId]
	if method == nil {
		log.Infof(" CallMethodWitheRoleId msgId = %d not found method", msgId)
		return
	}
	method.execute(roleId, body)
}

func CallMethod(msgId int32, body []byte, ctx network.ChannelContext) {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Infof("err=%v, stack=%s", r, s)
		}
	}()
	method := msgIdContextMap[msgId]
	if method == nil {
		log.Infof("CallMethod  msgId = %d not found method", msgId)
		return
	}
	method.execute(ctx, body)
}

//func CallMethod1(msgId int32, body []byte, ctx network.ChannelContext) {
//	defer func() {
//		if r := recover(); r != nil {
//			s := string(debug.Stack())
//			log.Infof("err=%v, stack=%s", r, s)
//		}
//	}()
//	methodPtr := msgIdMap[msgId]
//	if methodPtr == nil {
//		log.Infof("CallMethod  msgId = %d not found method", msgId)
//		return
//	}
//
//	method1 := (*protoMethod1[interface{}, interface{}])(methodPtr)
//	method1.methodCmm(ctx, body)
//
//	log.Infof(" value = %s")
//}
