package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(roleId int64, request proto.Message) {
	//context := ctx.Context().(*client.ConnInnerClientContext)
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof(" innerPlayerLogin login pid = %d s = %d", loginRequest.RoleId, loginRequest.GetSid())

	existRole := RoleMgr.GetByRoleId(loginRequest.GetRoleId())
	if existRole != nil {
		log.Infof("roleId =%d have login no need process", existRole.RoleId)
		return
	}
	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(client.WORLD).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_LOGIN_REQ), loginRequest.RoleId, innerLoginReq)
	gameRole := role.NewRole(loginRequest.RoleId)
	RoleMgr.AddRole(gameRole)
}

func loginResponseFromWorldServer(roleId int64, request proto.Message) {
	//context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	//roleId := innerLoginResponse.RoleId
	log.Infof("login response = %d  sid =%d", roleId, innerLoginResponse.Sid)
	player := RoleMgr.GetByRoleId(roleId)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", roleId)
		return
	}
	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_LOGIN_RES), roleId, innerLoginResponse)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	gameRole := RoleMgr.GetByRoleId(roleId)
	if gameRole == nil {
		log.Infof("roleId =%d not found", roleId)
		return
	}
	playerDisconnectRequest := request.(*protoGen.InnerPlayerDisconnectRequest)
	if playerDisconnectRequest.Sid != gameRole.Sid {
		log.Infof("roleId =%d have reconnected ", roleId)
		return
	}
	log.Infof("roleId =%d logout", roleId)
	client.GetInnerClient(client.GAME).SendInnerMsg(int32(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ), roleId, playerDisconnectRequest)
}
