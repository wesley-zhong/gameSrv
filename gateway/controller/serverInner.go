package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

func loginResponseFromGameServer(roleId int64, request proto.Message) {
	//	context := ctx.Context().(*aresTcpClient.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	player := player.PlayerMgr.GetByRoleId(roleId)
	if player == nil {
		log.Infof("roleId = %d have disconnected", roleId)
		return
	}
	sid := player.Context.Sid
	if sid != innerLoginResponse.GetSid() {
		log.Infof("roleId =%d have reconnected now sid =%d longRes sid =%d", roleId, sid, innerLoginResponse.GetSid())
		return
	}
	log.Infof("login loginResponseFromGameServer =roleId =%d siId= %slogin success =", roleId, innerLoginResponse)
	player.SetValid()

	response := &protoGen.LoginResponse{
		ErrorCode:  0,
		ServerTime: time.Now().UnixMilli(),
		RoleId:     roleId,
	}
	player.Context.SendMsg(protoGen.ProtoCode_LOGIN_RESPONSE, response)
}
