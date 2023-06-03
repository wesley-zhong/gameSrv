package controller

import (
	"gameSrv/game/role"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
	"time"
)

func Init() {
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), &protoGen.InnerServerHandShake{}, handShake)
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_HEART_BEAT_RES), &protoGen.HeartBeatResponse{}, heartBeatResponse)

	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), &protoGen.InnerLoginRequest{}, innerPlayerLogin)
	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), &protoGen.InnerLoginResponse{}, loginResponseFromWorldServer)

	core.RegisterMethod(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), &protoGen.InnerPlayerDisconnectRequest{}, innerPlayerDisconnect)

	core.RegisterMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)
	//performance test
	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	core.RegisterMethod(int32(protoGen.ProtoCode_LOGOUT_REQUEST), &protoGen.LogoutRequest{}, logout)

	//connect world server
	client.InnerClientConnect(client.WORLD, viper.GetString("worldServerAddr"), client.GAME)

	//add  msg  to game server to add me
}

var RoleMgr = role.NewRoleMgr() //make(map[int64]network.ChannelContext)

func innerPlayerLogin(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof(" innerPlayerLogin login pid = %d s = %d", loginRequest.RoleId, loginRequest.GetSid())

	existRole := RoleMgr.GetByRoleId(loginRequest.GetRoleId())
	if existRole != nil {
		log.Infof("roleId =%d have login no need process", existRole.RoleId)
		return
	}
	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    context.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(client.WORLD).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), loginRequest.RoleId, innerLoginReq)
	gameRole := role.NewRole(loginRequest.RoleId)
	RoleMgr.AddRole(gameRole)
}

func loginResponseFromWorldServer(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	roleId := innerLoginResponse.RoleId
	log.Infof("login response = %d  sid =%d", roleId, context.Sid)
	player := RoleMgr.GetByRoleId(roleId)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", roleId)
		return
	}
	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), roleId, innerLoginResponse)
}

func heartBeat(ctx network.ChannelContext, request proto.Message) {
	player := ctx.Context().(*player.Player)
	//context := ctx.Context().(*client.ConnClientContext)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" contex= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	//	PlayerMgr.Get()
	//PlayerMgr.GetByContext(context).Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
	player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

func heartBeatResponse(ctx network.ChannelContext, request proto.Message) {
	//context := ctx.Context().(*client.ConnInnerClientContext)
	//log.Infof("==== receive sid=%d  addr %s ", context.Sid, ctx.RemoteAddr())
}

func innerServerKickout(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("login response = %d  sid =%d", kickOut.Reason, context.Sid)
}

func innerPlayerDisconnect(ctx network.ChannelContext, request proto.Message) {

	//disconnectRequest := &protoGen.InnerPlayerDisconnectRequest{
	//	Sid:    disConnPlayer.Context.Sid,
	//	RoleId: disConnPlayer.Pid,
	//}

	//client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), disConnPlayer.Pid, disconnectRequest)

}

func performanceTest(ctx network.ChannelContext, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	//res := &protoGen.PerformanceTestRes{
	//	SomeId:    testReq.SomeId,
	//	ResBody:   testReq.SomeBody,
	//	SomeIdAdd: testReq.SomeId + 1,
	//}
	log.Infof("========== game performanceTest %d  remoteAddr=%s", testReq.SomeId, ctx.RemoteAddr())
	//ctx.Context().(*player.Player).Context.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	client.GetInnerClient(client.WORLD).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), 0, req)
}

func performanceTestResFromWorld(ctx network.ChannelContext, res proto.Message) {
	testRes := res.(*protoGen.PerformanceTestRes)
	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), 0, testRes)

}

func handShake(ctx network.ChannelContext, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	handShake := request.(*protoGen.InnerServerHandShake)
	validInnerClient.ServerId = handShake.FromServerId
	fromServerType := handShake.FromServerType
	client.AddInnerClientConnect(client.GameServerType(fromServerType), validInnerClient)
	log.Infof("client id =%d from serverId=%d  serverType= %d addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, fromServerType, validInnerClient.Ctx.RemoteAddr())
}
