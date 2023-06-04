package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func performanceTest(ctx network.ChannelContext, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	res := &protoGen.PerformanceTestRes{
		SomeId:    testReq.SomeId,
		ResBody:   testReq.SomeBody,
		SomeIdAdd: testReq.SomeId + 1,
	}
	log.Infof("==========  performanceTest %d  remoteAddr=%s", testReq.SomeId, ctx.RemoteAddr())
	ctx.Context().(*client.ConnClientContext).SendMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES, res)
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), 0, req)
}

func performanceTestResFromWorld(roleId int64, res proto.Message) {
	//	testRes := res.(*protoGen.PerformanceTestRes)
	perFormancePlayer := PlayerMgr.GetByRoleId(roleId)
	if perFormancePlayer == nil {
		return
	}
	perFormancePlayer.Context.SendMsg(protoGen.ProtoCode_PERFORMANCE_TEST_RES, res)
	//client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), testRes)
}
