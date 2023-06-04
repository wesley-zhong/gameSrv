package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

var playerConn = make(map[int64]*client.ConnClientContext)

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), &protoGen.HeartBeatResponse{}, hearBeatResponse)
	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceRes)
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_RESPONSE), &protoGen.LoginResponse{}, loginResponse)

	go startConnection(1)
}

func hearBeatResponse(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
	response := request.(*protoGen.HeartBeatResponse)
	//kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("pigd =%d heatbeat response = %d  ", context.Sid, response.ServerTime)
}

func performanceRes(ctx network.ChannelContext, res proto.Message) {
	performanceRes := res.(*protoGen.PerformanceTestRes)
	log.Infof("------response id =%d", performanceRes.SomeIdAdd)
}

func loginResponse(ctx network.ChannelContext, msg proto.Message) {
	res := msg.(*protoGen.LoginResponse)
	log.Infof("------login response roleId=%d", res.RoleId)

}

func startConnection(count int) {
	for i := 0; i < count; i++ {
		//	client := client.ClientConnect("124.222.26.216:9101")
		client := client.ClientConnect("127.0.0.1:9101")
		//add  msg  to game server to add me
		playerConn[client.Sid] = client
		request := &protoGen.LoginRequest{
			AccountId:  int64(i + 100),
			RoleId:     int64(i + 100000),
			LoginToken: "abc",
			GameTicket: 0,
			ServerId:   0,
		}
		client.SendMsg(protoGen.ProtoCode_LOGIN_REQUEST, request)
	}

	//req := &protoGen.PerformanceTestReq{
	//	SomeId:   2,
	//	SomeBody: "hello",
	//}
	//timer := time.NewTimer(3 * time.Second)

	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		for {
	//			timer.Reset(100 * time.Millisecond)
	//			select {
	//			case <-timer.C:
	//				for _, v := range playerConn {
	//					v.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), req)
	//				}
	//			}
	//		}
	//	}
	//}()

}
