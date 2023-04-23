package controller

import (
	"gameSrv/gateway/message"
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	protoGen "gameSrv/proto"
	"time"

	"google.golang.org/protobuf/proto"
)

var gameInnerClient *client.ConnInnerClientContext

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
}

var PlayerMgr = player.NewPlayerMgr() //make(map[int64]network.ChannelContext)

func login(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
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
	gameInnerClient.Send(innerMsg)
	//PlayerContext[loginRequest.RoleId] = ctx
	player := player.NewPlayer(loginRequest.GetRoleId(), context)
	PlayerMgr.Add(player)
	context.Ctx.SetContext(player)
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
