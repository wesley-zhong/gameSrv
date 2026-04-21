package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

func login(ctx tcp.Channel, request proto.Message) {
	context := ctx.Context().(*client.ConnContext)
	loginRequest := request.(*protoGen.LoginRequest)
	pid := loginRequest.RoleId
	existPlayer := player.PlayerMgr.GetByRoleId(pid)
	if existPlayer != nil {
		//close exist player
		existPlayer.Context.Ctx.Close()
		existPlayer.SetContext(context)
	} else {
		existPlayer = player.NewPlayer(loginRequest.GetRoleId(), context)
		player.PlayerMgr.Add(existPlayer)
	}
	context.Ctx.SetContext(existPlayer)
	innerRequest := &protoGen.InnerLoginRequest{
		Sid:    context.Sid,
		RoleId: existPlayer.Pid,
	}
	log.Infof("====== loginAddr=%s now loginCount =%d  content= %s", ctx.RemoteAddr(), player.PlayerMgr.GetSize(), innerRequest)
	client.SendInnerToGameServer(existPlayer.Pid, protoGen.InnerProtoCode_INNER_LOGIN_REQ, innerRequest)
}

func heartBeat(ctx tcp.Channel, request proto.Message) {
	player := ctx.Context().(*player.Player)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" context= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	player.Context.SendMsg(protoGen.MsgId_HEART_BEAT_RESPONSE, response)
}

func ClientDisConnect(ctx tcp.Channel) {
	disConnPlayer := ctx.Context().(*player.Player)
	//check right?
	if disConnPlayer.Context.Ctx.GetId() != ctx.GetId() {
		log.Infof("context =%s disconnected but playerId =%d have reconnected", ctx.RemoteAddr(), disConnPlayer.Pid)
		return
	}
	player.PlayerMgr.Remove(disConnPlayer)
	log.Infof("conn =%s  sid=%d pid=%d  closed now playerCount=%d",
		ctx.RemoteAddr(), disConnPlayer.Context.Sid, disConnPlayer.Pid, player.PlayerMgr.GetSize())

	disconnectRequest := &protoGen.InnerPlayerDisconnectRequest{
		Sid:    disConnPlayer.Context.Sid,
		RoleId: disConnPlayer.Pid,
	}
	client.SendInnerToGameServer(disConnPlayer.Pid, protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, disconnectRequest)
}

func innerServerKickout(pid int64, request proto.Message) {
	kickOut := request.(*protoGen.KickOutRequest)
	log.Infof("kickout role= %d  sid =%d reason=%d", kickOut.GetRoleId(), kickOut.Sid, kickOut.GetReason())
}
