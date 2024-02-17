package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

var playerConn = make(map[int64]*client.ConnContext)

func Init() {
	tcp.RegisterMethod(int16(protoGen.ProtoCode_HEART_BEAT_RESPONSE), &protoGen.HeartBeatResponse{}, hearBeatResponse)
	tcp.RegisterMethod(int16(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceRes)
	tcp.RegisterMethod(int16(protoGen.ProtoCode_LOGIN_RESPONSE), &protoGen.LoginResponse{}, loginResponse)
	tcp.RegisterMethod(int16(protoGen.ProtoCode_DIRECT_FROM_GAME_CLIENT), &protoGen.EchoReq{}, onDirectFromGame)
	tcp.RegisterMethod(int16(protoGen.ProtoCode_DIRECT_FROM_WORLD_CLIENT), &protoGen.EchoReq{}, onDirectFromWorld)
}

func hearBeatResponse(ctx tcp.Channel, request proto.Message) {
	context := ctx.Context().(*client.ConnContext)
	response := request.(*protoGen.HeartBeatResponse)
	//kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("pid =%d heat beat response = %d  ", context.Sid, response.ServerTime)
}

func performanceRes(ctx tcp.Channel, res proto.Message) {
	performanceRes := res.(*protoGen.PerformanceTestRes)
	log.Infof("------response id =%d", performanceRes.SomeIdAdd)
}

func loginResponse(ctx tcp.Channel, msg proto.Message) {
	res := msg.(*protoGen.LoginResponse)
	log.Infof("------login response roleId=%d  remote addr =%s", res.RoleId, ctx.RemoteAddr())

	context := playerConn[res.RoleId]
	req := &protoGen.EchoReq{
		RequestBody: "uuaauuauauau",
		SomeId:      99999999999999999,
	}
	context.SendMsg(protoGen.ProtoCode_DIRECT_TO_GAME, req)
	ctx.SetContext(context)
}

func onDirectFromGame(ctx tcp.Channel, msg proto.Message) {
	res := msg.(*protoGen.EchoReq)
	context := ctx.Context().(*client.ConnContext)
	log.Infof("-----on   -onDirectFromGame body=%s  ", res)
	context.SendMsg(protoGen.ProtoCode_DIRECT_TO_WORLD, res)
}

func onDirectFromWorld(ctx tcp.Channel, msg proto.Message) {
	res := msg.(*protoGen.EchoReq)
	log.Infof("-----on -onDirectFromWorld body=%s ", res)
}

func StartConnection(count int) {
	for i := 0; i < count; i++ {
		//24.222.26.216:9101
		serverAddr := viper.GetString("serverAddr")
		log.Infof("client  connnet addr = %s", serverAddr)

		channel := client.ClientConnect(serverAddr)
		//client := client.ClientConnect("127.0.0.1:9101")
		//add  msg  to game server to add me

		request := &protoGen.LoginRequest{
			AccountId:  int64(i + 10011),
			RoleId:     int64(i + 1000001),
			LoginToken: "abc",
			GameTicket: 0,
			ServerId:   0,
		}
		playerConn[request.RoleId] = &client.ConnContext{Ctx: channel, Sid: 111}
		playerConn[request.RoleId].SendMsg(protoGen.ProtoCode_LOGIN_REQUEST, request)
	}
}
