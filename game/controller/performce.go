package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

func performanceTest(pid int64, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	log.Infof("========== game performanceTest %d  pid=%d", testReq.SomeId, pid)
	//ctx.Context().(*player.Player).Context.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	//client.getInnerClient(client.ROUTER).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_REQ, 0, req)
}

func performanceTestResFromWorld(pid int64, res proto.Message) {
	//testRes := res.(*protoGen.PerformanceTestRes)
	//client.getInnerClient(client.GATE_WAY).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES) 0, testRes)
}

func directToGame(pid int64, req proto.Message) {
	echoMsg := req.(*protoGen.EchoReq)
	log.Infof("--------- directToGame  body =%s", echoMsg)
	client.SendToGateway(pid, protoGen.MsgId_DIRECT_FROM_GAME_CLIENT, echoMsg)
}
