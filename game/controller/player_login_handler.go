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
	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	player.OnPlayerLogin(loginRequest.RoleId, loginRequest.Sid)
	client.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_REQ, loginRequest.RoleId, innerLoginReq)
}

func loginResponseFromWorldServer(roleId int64, request proto.Message) {
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	log.Infof("------loginResponseFromWorldServer   response=%s", innerLoginResponse)
	player := player.RoleOlineMgr.GetPlayerById(roleId)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", roleId)
		return
	}
	client.GetInnerClient(global.GATE_WAY).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, roleId, innerLoginResponse)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	playerDisconnectRequest := request.(*protoGen.InnerPlayerDisconnectRequest)
	player.OnPlayerDisconnected(roleId, playerDisconnectRequest.Sid)
	log.Infof("#####  roleId =%d  disconnected finished ", roleId)
}
