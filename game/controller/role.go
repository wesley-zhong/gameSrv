package controller

import (
	"gameSrv/game/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
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
	//roleDO := quest.FindRoleData(loginRequest.RoleId)
	//if roleDO == nil {
	//	log.Errorf("not found roleId=%d ", roleId)
	//	return
	//}

	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_REQ, loginRequest.RoleId, innerLoginReq)
	gameRole := player.NewGamePlayer(loginRequest.RoleId, loginRequest.Sid)
	gameRole.Sid = loginRequest.Sid
	player.RoleOlineMgr.AddPlayer(gameRole)
}

func loginResponseFromWorldServer(roleId int64, request proto.Message) {
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	log.Infof("------loginResponseFromWorldServer 11111  response = %d   =%s", roleId, innerLoginResponse)
	player := player.RoleOlineMgr.GetPlayerById(roleId)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", roleId)
		return
	}
	client.GetInnerClient(global.GATE_WAY).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, roleId, innerLoginResponse)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	gamePlayer := player.RoleOlineMgr.GetPlayerById(roleId)
	if gamePlayer == nil {
		log.Infof("roleId =%d not found", roleId)
		return
	}
	playerDisconnectRequest := request.(*protoGen.InnerPlayerDisconnectRequest)
	if playerDisconnectRequest.Sid != gamePlayer.Sid {
		log.Infof("roleId =%d have reconnected ", roleId)
		return
	}
	log.Infof("roleId =%d logout", roleId)
	client.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, roleId, playerDisconnectRequest)
}
