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

type protoMethod[T any] struct {
	methodFuc MsgIdFuc[T, proto.Message]
	param     proto.Message
}

func (m *protoMethod[T]) execute(ctx T, body []byte) {
	param := m.param.ProtoReflect().New().Interface()
	if body == nil {
		m.methodFuc(ctx, param)
		return
	}
	if err := proto.Unmarshal(body, param); err != nil {
		log.Infof("protobuf unmarshal error: %v", err)
		return
	}
	m.methodFuc(ctx, param)
}

func (m *protoMethod[T]) Execute(ctx T, param proto.Message) {
	m.methodFuc(ctx, param)
}

func (m *protoMethod[T]) Parse(body []byte) (proto.Message, error) {
	if body == nil {
		return nil, nil
	}
	param := m.param.ProtoReflect().New().Interface()
	if err := proto.Unmarshal(body, param); err != nil {
		log.Infof("protobuf unmarshal error: %v", err)
		return nil, err
	}
	return param, nil
}

func RegisterMethod(msgId int16, param proto.Message, fuc MsgIdFuc[Channel, proto.Message]) {
	msgIdContextMap[msgId] = &protoMethod[Channel]{
		methodFuc: fuc,
		param:     param,
	}
}

func RegisterCallPlayerMethod(msgId int32, param proto.Message, fuc MsgIdFuc[int64, proto.Message]) {
	msgPlayerIdMap[int16(msgId)] = &protoMethod[int64]{
		methodFuc: fuc,
		param:     param,
	}
}

func CallMethodWithPlayerId(msgId int16, pid int64, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic in CallMethodWithPlayerId: err=%v, stack=%s", r, string(debug.Stack()))
		}
	}()

	method := msgPlayerIdMap[msgId]
	if method == nil {
		log.Infof("CallMethodWithPlayerId msgId=%d not found", msgId)
		return false
	}
	method.execute(pid, body)
	return true
}

func CallMethodWithChannelContext(msgId int16, context Channel, body []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic in CallMethodWithChannelContext: err=%v, stack=%s", r, string(debug.Stack()))
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
