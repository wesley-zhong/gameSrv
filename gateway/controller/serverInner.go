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

func handShakeResp(ctx tcp.Channel, request proto.Message) {
	handShake := request.(*protoGen.InnerServerHandShakeRes)
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	log.Infof("handShake response finished from serverId=%s  addr =%s handshake finished  ",
		handShake.FromServerSid, validInnerClient.Ctx.RemoteAddr())
}

func loginResponseFromGameServer(pid int64, request proto.Message) {
	//	context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	player := player.PlayerMgr.GetByRoleId(pid)
	if player == nil {
		log.Infof("pid = %d have disconnected", pid)
		return
	}
	sid := player.Context.Sid
	if sid != innerLoginResponse.GetSid() {
		log.Infof("pid =%d have reconnected now sid =%d longRes sid =%d", pid, sid, innerLoginResponse.GetSid())
		return
	}
	log.Infof("login loginResponseFromGameServer =pid =%d siId= %slogin success =", pid, innerLoginResponse)
	player.SetValid()

	response := &protoGen.LoginResponse{
		ErrorCode:  0,
		ServerTime: time.Now().UnixMilli(),
		RoleId:     pid,
	}
	player.Context.SendMsg(protoGen.MsgId_LOGIN_RESPONSE, response)
}
