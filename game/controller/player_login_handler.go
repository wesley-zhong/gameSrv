package controller

import (
	"gameSrv/game/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(pid int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("innerPlayerLogin login pid = %d pid = %d", loginRequest.RoleId, pid)
	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	player.OnPlayerLogin(loginRequest.RoleId, loginRequest.Sid)
	client.SendMsgToRouterServer(loginRequest.RoleId, protoGen.InnerProtoCode_INNER_LOGIN_REQ, innerLoginReq)
}

func loginResponseFromWorldServer(pid int64, request proto.Message) {
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	log.Infof("------loginResponseFromWorldServer   response=%s", innerLoginResponse)
	player := player.RoleOlineMgr.GetPlayerById(pid)
	if player == nil {
		log.Infof(" role id = %d not found or have disconnected", pid)
		return
	}
	client.SendInnerToGateway(pid, protoGen.InnerProtoCode_INNER_LOGIN_RES, innerLoginResponse)
}

func innerPlayerDisconnect(pid int64, request proto.Message) {
	playerDisconnectRequest := request.(*protoGen.InnerPlayerDisconnectRequest)
	player.OnPlayerDisconnected(pid, playerDisconnectRequest.Sid)
	log.Infof("#####  pid =%d  disconnected finished ", pid)
}
