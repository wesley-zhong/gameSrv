package controller

import (
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func Init() {
	tcp.RegisterMethod(int16(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)

	tcp.RegisterMethod(int16(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)

	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)

	tcp.RegisterMethod(int16(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromGameServer)
}
