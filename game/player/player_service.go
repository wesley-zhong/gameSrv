package player

import (
	"gameSrv/game/gameevent"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
)

// this may be run  on io thread
func OnPlayerLogin(pid int64, sid int64) *GamePlayer {
	existPlayer := RoleOlineMgr.GetPlayerById(pid)
	if existPlayer != nil && existPlayer.Sid == sid {
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
	log.Infof("pid = %d, sid = %d login success  now count =%d", pid, sid, RoleOlineMgr.Size())
	return existPlayer
}

// this may be run logic thread
func OnPlayerLoginLogic(player *GamePlayer) {
	player.OnLogin()
	player.DispatchEvent(gameevent.NewEvent[gameevent.GameEvent](player.Id, gameevent.LoginEventID))
	player.SaveDataOnPlayerRouting()
}

func OnPlayerDisconnected(pid int64, sid int64) {
	existPlayer := RoleOlineMgr.Remove(pid)

	if existPlayer == nil {
		log.Infof(" pid = %d, sid = %d not found", pid, sid)
		return
	}

	if existPlayer.Sid != sid {
		log.Infof(" pid = %d, now sid = %d disconnected sid =%d  do not process", pid, existPlayer.Sid, sid)
		return
	}
	existPlayer.DispatchEvent(gameevent.NewEvent[gameevent.GameEvent](existPlayer.Id, gameevent.DisconnectEventID))
	existPlayer.OnDisconnect()
	playerDisconnectRequest := &protoGen.InnerPlayerDisconnectRequest{
		Sid:    sid,
		RoleId: pid,
	}
	log.Infof("pid = %d, sid = %d disconnected now count =%d", pid, sid, RoleOlineMgr.Size())
	client.SendMsgToRouterServer(pid, protoGen.InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ, playerDisconnectRequest)
}
