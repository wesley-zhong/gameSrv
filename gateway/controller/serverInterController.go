package controller

import (
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"time"
)

func loginResponseFromGameServer(roleId int64, request proto.Message) {
	//	context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
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
	log.Infof("login loginResponseFromGameServer =roleId =%d siId= %slogin success =", roleId, innerLoginResponse)
	player.SetValid()

	response := &protoGen.LoginResponse{
		ErrorCode:  0,
		ServerTime: time.Now().UnixMilli(),
		RoleId:     roleId,
	}
	PlayerMgr.GetByRoleId(innerLoginResponse.RoleId).Context.SendMsg(protoGen.ProtoCode_LOGIN_RESPONSE, response)
}
