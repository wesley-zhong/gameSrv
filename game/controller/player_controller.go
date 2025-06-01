package controller

import (
	"gameSrv/game/player"
	"gameSrv/pkg/aresTcpClient"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(ctx tcp.Channel, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("innerPlayerLogin login pid = %d player id = %d", loginRequest.GetPlayerId(), loginRequest.PlayerId)

	//get from memory
	existPlayer := player.PlayerOlineMgr.GetPlayerById(loginRequest.PlayerId)
	if existPlayer != nil {
		log.Infof("roleId =%d have login no need process", existPlayer.Pid)
		return
	}

	gamePlayer := player.NewGamePlayer(loginRequest.PlayerId, ctx.Context().(*aresTcpClient.ConnInnerClientContext))

	innerLoginReq := &protoGen.InnerLoginRequest{
		Sid:      loginRequest.Sid,
		PlayerId: loginRequest.PlayerId,
	}
	gamePlayer.Sid = loginRequest.Sid
	aresTcpClient.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_REQ, loginRequest.PlayerId, innerLoginReq)
	player.PlayerOlineMgr.AddPlayer(gamePlayer)
}

//func loginResponseFromWorldServer(roleId int64, request proto.Message) {
//	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
//	log.Infof("------loginResponseFromWorldServer  response = %d   =%s", roleId, innerLoginResponse)
//	player := player.PlayerOlineMgr.GetPlayerById(roleId)
//	if player == nil {
//		log.Infof(" role id = %d not found or have disconnected", roleId)
//		return
//	}
//	aresTcpClient.GetInnerClient(global.GATE_WAY).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, roleId, innerLoginResponse)
//}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	gameRole := player.PlayerOlineMgr.GetPlayerById(roleId)
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
	aresTcpClient.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, roleId, playerDisconnectRequest)
}
