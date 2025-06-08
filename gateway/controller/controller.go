package controller

import (
	"gameSrv/pkg/handshake"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func init() {
	tcp.RegisterMethod(int16(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_RES), &protoGen.InnerServerHandShakeRes{}, handshake.HandShakeRes)

	tcp.RegisterMethod(int16(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	tcp.RegisterMethod(int16(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)

	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromGameServer)
}
