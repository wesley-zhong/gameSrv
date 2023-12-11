package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
)

func Init() {

	tcp.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	tcp.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	tcp.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES), &protoGen.HeartBeatResponse{}, heartBeatResponse)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)

	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromWorldServer)
	tcp.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)

	//performance test  msg
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	tcp.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	//
	//core.RegisterMethods[protoGen.InnerServerHandShake](int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), handShake1)

	//add  msg  to game server
}

var RoleOlineMgr = role.NewRoleMgr()
