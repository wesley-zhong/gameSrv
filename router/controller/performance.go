package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func performanceTest(roleId int64, req proto.Message) {
	//testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	//log.Infof("========== router performanceTest %d  roleId=%d", testReq.SomeId, roleId)
	//client.GetInnerClient(client.GAME).SendInnerMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES, 0, res)
}

func onDirectToWorld(roleId int64, req proto.Message) {
	echoMsg := req.(*protoGen.EchoReq)
	client.GetInnerClient(global.GAME).SendMsg(protoGen.ProtoCode_DIRECT_FROM_WORLD_CLIENT, roleId, echoMsg)
	log.Infof("-------------  on direct to router")
}
