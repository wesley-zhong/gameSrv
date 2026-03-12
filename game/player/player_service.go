package player

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
)

// this may be run  on io thread
func OnPlayerLogin(pid int64, sid int64) *GamePlayer {
	existPlayer := RoleOlineMgr.GetPlayerById(pid)
	if existPlayer != nil {
		return existPlayer
	}

	existPlayer = NewGamePlayer(pid, sid)
	err := existPlayer.LoadDataFromDB()
	if err != nil {
		log.Infof("pid = %d, sid = %d login failed", pid, sid)
		return nil
	}
	RoleOlineMgr.AddPlayer(existPlayer)
	OnPlayerLoginLogic(existPlayer)
	return existPlayer
}

// this may be run logic thread
func OnPlayerLoginLogic(player *GamePlayer) {
	player.OnLogin()
	player.SaveData()
}

func OnPlayerDisconnected(pid int64, sid int64) {
	existPlayer := RoleOlineMgr.GetPlayerById(pid)
	if existPlayer == nil {
		log.Infof(" pid = %d, sid = %d not found", pid, sid)
		return
	}

	if existPlayer.Sid != sid {
		log.Infof(" pid = %d, now sid = %d disconnected sid =%d", pid, existPlayer.Sid, sid)
		return
	}
	existPlayer.OnDisconnect()
	playerDisconnectRequest := &protoGen.InnerPlayerDisconnectRequest{
		Sid:    sid,
		RoleId: pid,
	}
	client.GetInnerClient(global.ROUTER).SendInnerMsg(protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, pid, playerDisconnectRequest)
}
