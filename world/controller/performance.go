package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

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
