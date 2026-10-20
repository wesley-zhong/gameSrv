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

	// 先把玩家绑定到新连接并放入 PlayerMgr，再关闭旧连接。
	// 旧连接的 OnClosed 看 GetId() 不匹配只会打日志，
	// 不会在 Close 与 SetContext 之间把新登录的玩家从 PlayerMgr 中移除。
	existPlayer := player.PlayerMgr.GetByRoleId(pid)
	if existPlayer != nil {
		oldConn := existPlayer.Context.Ctx
		existPlayer.SetContext(context)
		ctx.SetContext(existPlayer)
		if oldConn.GetId() != ctx.GetId() {
			oldConn.Close()
		}
	} else {
		existPlayer = player.NewPlayer(loginRequest.GetRoleId(), context)
		player.PlayerMgr.Add(existPlayer)
		ctx.SetContext(existPlayer)
	}
	innerRequest := &protoGen.InnerLoginRequest{
		Sid:    context.Sid,
		RoleId: existPlayer.Pid,
	}
	log.Infof("====== loginAddr=%s now loginCount =%d  content= %s", ctx.RemoteAddr(), player.PlayerMgr.GetSize(), innerRequest)
	client.SendInnerToGameServer(existPlayer.Pid, protoGen.InnerProtoCode_INNER_LOGIN_REQ, innerRequest)
}

func heartBeat(ctx tcp.Channel, request proto.Message) {
	p, ok := ctx.Context().(*player.Player)
	if !ok {
		return
	}
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" context= %d  heartbeat time = %d", p.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	p.Context.SendMsg(protoGen.MsgId_HEART_BEAT_RESPONSE, response)
}

func ClientDisConnect(ctx tcp.Channel) {
	disConnPlayer, ok := ctx.Context().(*player.Player)
	if !ok {
		return
	}
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
