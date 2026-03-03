package player

var PlayerMgr *MgrWrap = NewRoleMgr()

func PlayerLogin(pid int64, sid int64) *GamePlayer {
	existPlayer := PlayerMgr.GetPlayerById(pid)
	if existPlayer != nil {
		return existPlayer
	}

	existPlayer = NewGamePlayer(pid, sid)
	PlayerMgr.AddPlayer(existPlayer)
	return existPlayer
}
