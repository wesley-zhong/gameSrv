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

var PlayerMgr = player.NewPlayerMgr()

func handShake(ctx network.ChannelContext, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	client.AddInnerClientConnect(client.GAME, validInnerClient)
	handShake := request.(*protoGen.InnerServerHandShake)
	validInnerClient.ServerId = handShake.FromServerId
	serverType := handShake.ServerType
	client.AddInnerClientConnect(client.GameServerType(serverType), validInnerClient)
	log.Infof("client id =%d from serverId=%d  serverType= %d addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, serverType, validInnerClient.Ctx.RemoteAddr())
}

func innerHeartBeat(ctx network.ChannelContext, request proto.Message) {
	innerClient := ctx.Context().(*client.ConnInnerClientContext)
	response := &protoGen.InnerHeartBeatResponse{}
	innerClient.SendInnerMsg(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
	log.Infof(" addr =%s  heartbeat time", ctx.RemoteAddr())
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
		Id:         loginRequest.RoleId,
		SendType:   0,
		ProtoCode:  message.INNER_PROTO_LOGIN_REQUEST,
		CallbackId: 0,
	}

	innerMsg := &client.InnerMessage{
		InnerHeader: msgHeader,
		Body:        innerLoginReq,
	}
	client.GetInnerClient(client.GAME).Send(innerMsg)
}

func performanceTest(ctx network.ChannelContext, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	res := &protoGen.PerformanceTestRes{
		SomeId:    testReq.SomeId,
		ResBody:   testReq.SomeBody,
		SomeIdAdd: testReq.SomeId + 1,
	}
	log.Infof("========== world performanceTest %d  remoteAddr=%s", testReq.SomeId, ctx.RemoteAddr())
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
}
