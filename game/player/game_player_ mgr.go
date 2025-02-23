package player

import "sync"

var PlayerOlineMgr = NewRoleMgr()

type PlayerMgrWrap struct {
	players map[int64]*GamePlayer
	rwLock  sync.RWMutex
}

func NewRoleMgr() *PlayerMgrWrap {
	return &PlayerMgrWrap{
		players: make(map[int64]*GamePlayer),
		rwLock:  sync.RWMutex{},
	}
}

func (roleMgr *PlayerMgrWrap) AddPlayer(player *GamePlayer) {
	roleMgr.rwLock.Lock()
	defer roleMgr.rwLock.Unlock()
	roleMgr.players[player.Pid] = player
}

func (roleMgr *PlayerMgrWrap) GetPlayerById(roleId int64) *GamePlayer {
	roleMgr.rwLock.RLock()
	defer roleMgr.rwLock.RUnlock()
	return roleMgr.players[roleId]
}
