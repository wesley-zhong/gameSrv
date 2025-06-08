package controller

import (
	"gameSrv/pkg/handshake"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func init() {

	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_REQ), &protoGen.InnerServerHandShakeReq{}, handshake.HandShakeReq)
	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_RES), &protoGen.InnerServerHandShakeRes{}, handshake.HandShakeRes)

	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES), &protoGen.HeartBeatResponse{}, handshake.HeartBeatRes)

	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)

	//	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromWorldServer)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)

	//performance test  msg
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_DIRECT_TO_GAME), &protoGen.EchoReq{}, directToGame)

	//add  msg  to game server
}
