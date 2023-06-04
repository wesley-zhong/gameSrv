package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func Init() {
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ), &protoGen.InnerHeartBeatRequest{}, innerHeartBeat)

	core.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)
	core.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)
	core.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
}

var PlayerMgr = player.NewPlayerMgr()

func innerPlayerLogin(roleId int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("world inner login sessionId = %d = %d from  finished", loginRequest.Sid, loginRequest.RoleId)

	res := &protoGen.InnerLoginResponse{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), loginRequest.RoleId, res)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	log.Infof("---roleId = %d disconnected", roleId)
}

func performanceTest(roleId int64, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	res := &protoGen.PerformanceTestRes{
		SomeId:    testReq.SomeId,
		ResBody:   testReq.SomeBody,
		SomeIdAdd: testReq.SomeId + 1,
	}
	log.Infof("========== world performanceTest %d  roleId=%d", testReq.SomeId, roleId)
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), 0, res)
}
