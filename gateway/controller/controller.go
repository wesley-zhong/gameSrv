package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"github.com/spf13/viper"
	"time"

	"google.golang.org/protobuf/proto"
)

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	core.RegisterMethod(int32(-6), &protoGen.InnerLoginResponse{}, loginResponseFromGameServer)

	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)

	core.RegisterMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)

	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), &protoGen.PerformanceTestReq{}, performanceTest)
	//core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceTestResFromWorld)

	client.InnerClientConnect(client.GAME, viper.GetString("gameServerAddr"), client.GATE_WAY)
}

var PlayerMgr = player.NewPlayerMgr()

func login(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
	loginRequest := request.(*protoGen.LoginRequest)
	roleId := loginRequest.RoleId
	existPlayer := PlayerMgr.GetByRoleId(roleId)
	if existPlayer != nil {
		//close exist player
		existPlayer.Context.Ctx.Close()
		existPlayer.SetContext(context)
	} else {
		existPlayer = player.NewPlayer(loginRequest.GetRoleId(), context)
		PlayerMgr.Add(existPlayer)
	}
	context.Ctx.SetContext(existPlayer)
	innerRequest := &protoGen.InnerLoginRequest{
		Sid:    context.Sid,
		RoleId: existPlayer.Pid,
	}
	log.Infof("====== loginAddr=%s now loginCount =%d", ctx.RemoteAddr(), PlayerMgr.GetSize())
	client.GetInnerClient(client.GAME).SendInnerMsgProtoCode(protoGen.InnerProtoCode_INNER_LOGIN_REQ, existPlayer.Pid, innerRequest)
}

func loginResponseFromGameServer(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	roleId := innerLoginResponse.GetRoleId()
	player := PlayerMgr.GetByRoleId(roleId)
	if player == nil {
		log.Infof("roleId = %d have disconnected", roleId)
		return
	}
	sid := player.Context.Sid
	if sid != innerLoginResponse.GetSid() {
		log.Infof("roleId =%d have reconnected now sid =%d longRes sid =%d", roleId, sid, innerLoginResponse.GetSid())
		return
	}
	log.Infof("login response =roleId =%d siId= %d login success =", roleId, context.Sid)
	player.SetValid()

	response := &protoGen.LoginResponse{
		ErrorCode:  0,
		ServerTime: time.Now().UnixMilli(),
	}
	PlayerMgr.GetByRoleId(innerLoginResponse.RoleId).Context.Send(int32(protoGen.ProtoCode_LOGIN_RESPONSE), response)
}

func heartBeat(ctx network.ChannelContext, request proto.Message) {
	player := ctx.Context().(*player.Player)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" context= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

func ClientDisConnect(ctx network.ChannelContext) {
	disConnPlayer := ctx.Context().(*player.Player)
	//check right?
	if disConnPlayer.Context.Ctx != ctx {
		log.Infof("context =%s disconnected but playerId ={} have reconnected", ctx.RemoteAddr(), disConnPlayer.Pid)
		return
	}
	player.PlayerMgr.Remove(disConnPlayer)
	log.Infof("conn =%s  sid=%d pid=%d  closed now playerCount=%d",
		ctx.RemoteAddr(), disConnPlayer.Context.Sid, disConnPlayer.Pid, player.PlayerMgr.GetSize())
}

func innerServerKickout(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("login response = %d  sid =%d", kickOut.Reason, context.Sid)
}

func performanceTest(ctx network.ChannelContext, req proto.Message) {
	testReq := req.(*protoGen.PerformanceTestReq)
	res := &protoGen.PerformanceTestRes{
		SomeId:    testReq.SomeId,
		ResBody:   testReq.SomeBody,
		SomeIdAdd: testReq.SomeId + 1,
	}
	log.Infof("==========  performanceTest %d  remoteAddr=%s", testReq.SomeId, ctx.RemoteAddr())
	ctx.Context().(*client.ConnClientContext).Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), res)
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), 0, req)
}

//func performanceTestResFromWorld(ctx network.ChannelContext, res proto.Message) {
//	PlayerMgr.GetBySid()
//	testRes := res.(*protoGen.PerformanceTestRes)
//	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), testRes)
//}
