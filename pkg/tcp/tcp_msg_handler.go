package tcp

import (
	"gameSrv/pkg/log"
	"runtime/debug"
	"sync"

	"google.golang.org/protobuf/proto"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

var (
	msgIdContextMap = make(map[int16]*protoMethod[Channel])
	msgPlayerIdMap  = make(map[int16]*protoMethod[int64])
	msgMapMutex     sync.RWMutex
)

//var msgGoPool = gopool.StartNewWorkerPool(1, 1024)

type protoMethod[T1 any] struct {
	methodFuc MsgIdFuc[T1, proto.Message]
	param     proto.Message
}

func (method *protoMethod[T1]) execute(any T1, body []byte) {
	param := method.param.ProtoReflect().New().Interface()
	//if body == nil {
	//	method.methodFuc(any, nil)
	//}
	if err := proto.Unmarshal(body, param); err != nil {
		log.Infof("Protobuf unmarshal error: %v", err)
		return // Don't submit to pool if data is corrupted
	}
	method.methodFuc(any, param)
	//msgGoPool.SubmitTask(func() {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			s := string(debug.Stack())
	//			log.Infof("err=%v, stack=%s", r, s)
	//		}
	//	}()
	//	method.methodFuc(any, param)
	//})
}

func RegisterMethod(msgId int16, param proto.Message, fuc MsgIdFuc[Channel, proto.Message]) {
	method := &protoMethod[Channel]{
		methodFuc: fuc,
		param:     param,
	}
	msgMapMutex.Lock()
	msgIdContextMap[msgId] = method
	msgMapMutex.Unlock()
}

func RegisterCallPlayerMethod(msgId int32, param proto.Message, fuc MsgIdFuc[int64, proto.Message]) {
	method := &protoMethod[int64]{
		methodFuc: fuc,
		param:     param,
	}
	msgMapMutex.Lock()
	msgPlayerIdMap[int16(msgId)] = method
	msgMapMutex.Unlock()
}

func CallMethodWithRoleId(msgId int16, roleId int64, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Errorf("panic in CallMethodWithRoleId: err=%v, stack=%s", r, s)
		}
	}()
	msgMapMutex.RLock()
	method := msgPlayerIdMap[msgId]
	msgMapMutex.RUnlock()

	if method == nil {
		log.Infof("CallMethodWithRoleId msgId = %d not found method", msgId)
		return false
	}
	method.execute(roleId, body)
	return true
}

func CallMethodWithChannelContext(msgId int16, context Channel, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Errorf("panic in CallMethodWithChannelContext: err=%v, stack=%s", r, s)
		}
	}()
	msgMapMutex.RLock()
	method := msgIdContextMap[msgId]
	msgMapMutex.RUnlock()

	if method == nil {
		return false
	}
	method.execute(context, body)
	return true
}

func GetCallMethodById(msgId int16) *protoMethod[int64] {
	msgMapMutex.RLock()
	defer msgMapMutex.RUnlock()
	return msgPlayerIdMap[msgId]
}

func CallMethod(roleId int64, body []byte, method *protoMethod[int64]) {
	method.execute(roleId, body)
}

func HasMethod(msgId int16) bool {
	msgMapMutex.RLock()
	defer msgMapMutex.RUnlock()
	return msgPlayerIdMap[msgId] != nil || msgIdContextMap[msgId] != nil
}
