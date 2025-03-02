package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func performanceTest(roleId int64, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	log.Infof("========== game performanceTest %d  roleId=%d", testReq.SomeId, roleId)
	//ctx.Context().(*player.GamePlayer).Context.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	//client.GetInnerClient(client.ROUTER).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_REQ, 0, req)
}

func performanceTestResFromWorld(roleId int64, res proto.Message) {
	//testRes := res.(*protoGen.PerformanceTestRes)
	//client.GetInnerClient(client.GATE_WAY).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES) 0, testRes)
}

func directToGame(roleId int64, req proto.Message) {
	echoMsg := req.(*protoGen.EchoReq)
	log.Infof("--------- directToGame  body =%s", echoMsg)
	client.GetInnerClient(global.GATE_WAY).SendMsg(protoGen.ProtoCode_DIRECT_FROM_GAME_CLIENT, roleId, echoMsg)
}
