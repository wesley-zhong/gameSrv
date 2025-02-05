package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func Init() {
	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_REQ), &protoGen.InnerServerHandShakeReq{}, handShakeReq)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)

	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_DIRECT_TO_WORLD), &protoGen.EchoReq{}, onDirectToWorld)
}

var PlayerMgr = player.NewPlayerMgr()
