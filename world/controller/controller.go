package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func Init() {
	tcp.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	tcp.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ), &protoGen.InnerHeartBeatRequest{}, innerHeartBeat)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
}

var PlayerMgr = player.NewPlayerMgr()
