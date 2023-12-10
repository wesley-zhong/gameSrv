package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"time"
)

func login(ctx tcp.ChannelContext, request proto.Message) {
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

func heartBeat(ctx tcp.ChannelContext, request proto.Message) {
	player := ctx.Context().(*player.Player)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" context= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

func ClientDisConnect(ctx tcp.ChannelContext) {
	disConnPlayer := ctx.Context().(*player.Player)
	//check right?
	if disConnPlayer.Context.Ctx != ctx {
		log.Infof("context =%s disconnected but playerId ={} have reconnected", ctx.RemoteAddr(), disConnPlayer.Pid)
		return
	}
	player.PlayerMgr.Remove(disConnPlayer)
	log.Infof("conn =%s  sid=%d pid=%d  closed now playerCount=%d",
		ctx.RemoteAddr(), disConnPlayer.Context.Sid, disConnPlayer.Pid, player.PlayerMgr.GetSize())

	disconnectRequest := &protoGen.InnerPlayerDisconnectRequest{
		Sid:    disConnPlayer.Context.Sid,
		RoleId: disConnPlayer.Pid,
	}
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), disConnPlayer.Pid, disconnectRequest)
}

func innerServerKickout(roleId int64, request proto.Message) {
	kickOut := request.(*protoGen.KickOutRequest)
	log.Infof("kickout role= %d  sid =%d reason=%d", kickOut.GetRoleId(), kickOut.Sid, kickOut.GetReason())
}
