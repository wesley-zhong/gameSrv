package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"gameSrv/world/message"
	"google.golang.org/protobuf/proto"
)

func Init() {
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.LoginRequest{}, playerLogin)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_REQ), &protoGen.InnerHeartBeatRequest{}, innerHeartBeat)
	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)

}

var PlayerMgr = player.NewPlayerMgr() //make(map[int64]network.ChannelContext)
//var innerClientMap = make(map[int32]*client.ConnInnerClientContext)

func handShake(ctx network.ChannelContext, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	//innerClientMap[client.InnerClientType_GAME] = validInnerClient
	client.AddInnerClientConnect(client.InnerClientType_GAME, validInnerClient)
	log.Infof("client id =%d  addr =%s handshake finished", validInnerClient.Sid, validInnerClient.Ctx.RemoteAddr())
}

func innerHeartBeat(ctx network.ChannelContext, request proto.Message) {
	innerClient := ctx.Context().(*client.ConnInnerClientContext)
	log.Infof(" addr =%s  heartbeat time", ctx.RemoteAddr())

	response := &protoGen.InnerHeartBeatResponse{}
	//	PlayerMgr.Get()
	//PlayerMgr.GetByContext(context).Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
	innerClient.SendInnerMsg(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

func playerLogin(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	loginRequest := request.(*protoGen.LoginRequest)
	log.Infof("login token = %s id = %d", loginRequest.LoginToken, loginRequest.RoleId)
	innerLoginReq := &protoGen.InnerLoginRequest{
		SessionId: context.Sid,
		AccountId: loginRequest.AccountId,
		RoleId:    loginRequest.RoleId,
	}
	msgHeader := &protoGen.InnerHead{
		FromServerUid:    message.BuildServerUid(message.TypeGateway, 35),
		ToServerUid:      message.BuildServerUid(message.TypeGame, 35),
		ReceiveServerUid: 0,
		Id:               loginRequest.RoleId,
		SendType:         0,
		ProtoCode:        message.INNER_PROTO_LOGIN_REQUEST,
		CallbackId:       0,
	}

	innerMsg := &client.InnerMessage{
		InnerHeader: msgHeader,
		Body:        innerLoginReq,
	}
	//innerClientMap[client.InnerClientType_GAME].Send(innerMsg)
	client.GetInnerClient(client.InnerClientType_GAME).Send(innerMsg)
	//player := player.NewPlayer(loginRequest.GetRoleId(), context)
	//	PlayerMgr.Add(player)
}

func performanceTest(ctx network.ChannelContext, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	res := &protoGen.PerformanceTestRes{
		SomeId:    testReq.SomeId,
		ResBody:   testReq.SomeBody,
		SomeIdAdd: testReq.SomeId + 1,
	}
	log.Infof("========== world performanceTest %d  remomoteAddr=%s", testReq.SomeId, ctx.RemoteAddr())
	//ctx.Context().(*player.Player).Context.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	client.GetInnerClient(client.InnerClientType_GAME).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
}
