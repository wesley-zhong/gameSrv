package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

func performanceTest(pid int64, req proto.Message) {
	//testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	//log.Infof("========== router performanceTest %d  pid=%d", testReq.SomeId, pid)
	//client.getInnerClient(client.GAME).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES, 0, res)
}

func onDirectToWorld(pid int64, req proto.Message) {
	echoMsg := req.(*protoGen.EchoReq)
	client.SendToGameServer(pid, protoGen.ProtoCode_DIRECT_FROM_WORLD_CLIENT, echoMsg)
	log.Infof("-------------  on direct to router")
}
