package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

var playerConn = make(map[int64]*client.ConnClientContext)

func Init() {
	tcp.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), &protoGen.HeartBeatResponse{}, hearBeatResponse)
	tcp.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceRes)
	tcp.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_RESPONSE), &protoGen.LoginResponse{}, loginResponse)
}

func hearBeatResponse(ctx tcp.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
	response := request.(*protoGen.HeartBeatResponse)
	//kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("pid =%d heat beat response = %d  ", context.Sid, response.ServerTime)
}

func performanceRes(ctx tcp.ChannelContext, res proto.Message) {
	performanceRes := res.(*protoGen.PerformanceTestRes)
	log.Infof("------response id =%d", performanceRes.SomeIdAdd)
}

func loginResponse(ctx tcp.ChannelContext, msg proto.Message) {
	res := msg.(*protoGen.LoginResponse)
	log.Infof("------login response roleId=%d", res.RoleId)

}

func StartConnection(count int) {
	for i := 0; i < count; i++ {
		//24.222.26.216:9101
		serverAddr := viper.GetString("serverAddr")
		log.Infof("client  connnet addr = %s", serverAddr)

		client := client.ClientConnect(serverAddr)
		//client := client.ClientConnect("127.0.0.1:9101")
		//add  msg  to game server to add me
		playerConn[client.Sid] = client
		request := &protoGen.LoginRequest{
			AccountId:  int64(i + 10011),
			RoleId:     int64(i + 1000001),
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
