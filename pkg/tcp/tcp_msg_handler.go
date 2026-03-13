package tcp

import (
	"gameSrv/pkg/log"
	"runtime/debug"

	"google.golang.org/protobuf/proto"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

var (
	msgIdContextMap = make(map[int16]*protoMethod[Channel])
	msgPlayerIdMap  = make(map[int16]*protoMethod[int64])
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

func (method *protoMethod[T1]) Execute(any T1, param proto.Message) {
	method.methodFuc(any, param)
}

func (method *protoMethod[T1]) Parse(body []byte) (proto.Message, error) {
	param := method.param.ProtoReflect().New().Interface()
	if body == nil {
		return nil, nil
	}
	if err := proto.Unmarshal(body, param); err != nil {
		log.Infof("Protobuf unmarshal error: %v", err)
		return nil, err // Don't submit to pool if data is corrupted
	}
	return param, nil
}

func RegisterMethod(msgId int16, param proto.Message, fuc MsgIdFuc[Channel, proto.Message]) {
	method := &protoMethod[Channel]{
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
	msgPlayerIdMap[int16(msgId)] = method
}

func CallMethodWithPlayerId(msgId int16, pid int64, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Errorf("panic in CallMethodWithPlayerId: err=%v, stack=%s", r, s)
		}
	}()
	method := msgPlayerIdMap[msgId]

	if method == nil {
		log.Infof("CallMethodWithPlayerId msgId = %d not found method", msgId)
		return false
	}
	method.execute(pid, body)
	return true
}

func CallMethodWithChannelContext(msgId int16, context Channel, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Errorf("panic in CallMethodWithChannelContext: err=%v, stack=%s", r, s)
		}
	}()
	method := msgIdContextMap[msgId]

	if method == nil {
		return false
	}
	method.execute(context, body)
	return true
}

func GetCallMethodWithId(msgId int16) *protoMethod[int64] {
	return msgPlayerIdMap[msgId]
}

func GetMethodWithChannel(msgId int16) *protoMethod[Channel] {
	return msgIdContextMap[msgId]
}

func HasMethod(msgId int16) bool {
	return msgPlayerIdMap[msgId] != nil || msgIdContextMap[msgId] != nil
}
