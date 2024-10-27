package controller

import (
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func Init() {

	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_REQ), &protoGen.InnerServerHandShakeReq{}, handShakeReq)
	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_RES), &protoGen.InnerServerHandShakeRes{}, handShakeResp)

	tcp.RegisterMethod(int16(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES), &protoGen.HeartBeatResponse{}, heartBeatResponse)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromWorldServer)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)

	//performance test  msg
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_DIRECT_TO_GAME), &protoGen.EchoReq{}, directToGame)

	//add  msg  to game server
}
