package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

func Init() {
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES), &protoGen.HeartBeatResponse{}, heartBeatResponse)

	core.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)

	core.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromWorldServer)
	core.RegisterCallPlayerMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)

	//performance test
	core.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	core.RegisterCallPlayerMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	//connect world server
	client.InnerClientConnect(client.WORLD, viper.GetString("worldServerAddr"), client.GAME)

	//add  msg  to game server to add me
}

var RoleMgr = role.NewRoleMgr() //make(map[int64]network.ChannelContext)

func performanceTest(roleId int64, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	log.Infof("========== game performanceTest %d  roleId=%d", testReq.SomeId, roleId)
	//ctx.Context().(*player.Player).Context.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	client.GetInnerClient(client.WORLD).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), 0, req)
}

func performanceTestResFromWorld(roleId int64, res proto.Message) {
	testRes := res.(*protoGen.PerformanceTestRes)
	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), 0, testRes)

}
