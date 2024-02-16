package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(roleId int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("innerPlayerLogin login pid = %d roleId = %d", loginRequest.RoleId, roleId)

	//existRole := RoleOlineMgr.GetByRoleId(loginRequest.GetRoleId())
	//if existRole != nil {
	//	log.Infof("roleId =%d have login no need process", existRole.RoleId)
	//	return
	//}
	// load role data form db
	//roleDO := service.FindRoleData(loginRequest.RoleId)
	//if roleDO == nil {
	//	log.Errorf("not found roleId=%d ", roleId)
	//	return
	//}

	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(client.WORLD).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_REQ, loginRequest.RoleId, innerLoginReq)
	gameRole := role.NewRole(loginRequest.RoleId, nil)
	gameRole.Sid = loginRequest.Sid
	RoleOlineMgr.AddRole(gameRole)
}

func loginResponseFromWorldServer(roleId int64, request proto.Message) {
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	log.Infof("------loginResponseFromWorldServer  response = %d   =%s", roleId, innerLoginResponse)
	player := RoleOlineMgr.GetByRoleId(roleId)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", roleId)
		return
	}
	client.GetInnerClient(client.GATE_WAY).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, roleId, innerLoginResponse)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	gameRole := RoleOlineMgr.GetByRoleId(roleId)
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
	client.GetInnerClient(client.WORLD).SendInnerMsg(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, roleId, playerDisconnectRequest)
}
