package tcp

import (
	"gameSrv/pkg/gopool"
	"gameSrv/pkg/log"
	"google.golang.org/protobuf/proto"
	"reflect"
	"runtime"
	"runtime/debug"
)

type MsgIdFuc[T1 any, T2 any] func(T1, T2)

var msgIdContextMap = make(map[int16]*protoMethod[Channel])
var msgIdRoleIdMap = make(map[int16]*protoMethod[int64])

var msgGoPool = gopool.StartNewWorkerPool(1, 1024)

type protoMethod[T1 any] struct {
	methodFuc MsgIdFuc[T1, proto.Message]
	param     proto.Message
}

func (method *protoMethod[T1]) execute(any T1, body []byte) {
	param := method.param.ProtoReflect().New().Interface()
	//if body == nil {
	//	method.methodFuc(any, nil)
	//}
	err := proto.Unmarshal(body, param)
	if err != nil {
		log.Error(err)
		return
	}
	err = msgGoPool.SubmitTask(func() {
		defer func() {
			if r := recover(); r != nil {
				s := string(debug.Stack())
				log.Infof("err=%v, stack=%s", r, s)
			}
		}()
		method.methodFuc(any, param)
	})
	if err != nil {
		log.Error(err)
		return
	}
}

func RegisterMethod(msgId int16, param proto.Message, fuc MsgIdFuc[Channel, proto.Message]) {
	method := &protoMethod[Channel]{
		methodFuc: fuc,
		param:     param,
	}
	msgIdContextMap[msgId] = method
	log.Infof("======  register msgId=%d ##funtion: %s", msgId, runtime.FuncForPC(reflect.ValueOf(fuc).Pointer()).Name())
}

func RegisterCallPlayerMethod(msgId int32, param proto.Message, fuc MsgIdFuc[int64, proto.Message]) {
	method := &protoMethod[int64]{
		methodFuc: fuc,
		param:     param,
	}
	msgIdRoleIdMap[int16(msgId)] = method
}

func CallMethodWithRoleId(msgId int16, roleId int64, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Infof("err=%v, stack=%s", r, s)
		}
	}()
	method := msgIdRoleIdMap[msgId]
	if method == nil {
		log.Infof(" CallMethodWithRoleId msgId = %d not found method", msgId)
		return false
	}
	method.execute(roleId, body)
	return true
}

func CallMethodWithChannelContext(msgId int16, context Channel, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			log.Infof("err=%v, stack=%s", r, s)
		}
	}()
	method := msgIdContextMap[msgId]
	if method == nil {
		return false
	}
	method.execute(context, body)
	return true
}

func GetCallMethodById(msgId int16) *protoMethod[int64] {
	return msgIdRoleIdMap[msgId]
}

func CallMethod(roleId int64, body []byte, method *protoMethod[int64]) {
	method.execute(roleId, body)
}

func HasMethod(msgId int16) bool {
	return msgIdRoleIdMap[msgId] != nil || msgIdContextMap[msgId] != nil
}
