package player

import "sync"

var PlayerOlineMgr = NewRoleMgr()

type RoleMgrWrap struct {
	players map[int64]*GamePlayer
	rwLock  sync.RWMutex
}

func NewRoleMgr() *RoleMgrWrap {
	return &RoleMgrWrap{
		players: make(map[int64]*GamePlayer),
		rwLock:  sync.RWMutex{},
	}
}

func (roleMgr *RoleMgrWrap) AddPlayer(player *GamePlayer) {
	roleMgr.rwLock.Lock()
	defer roleMgr.rwLock.Unlock()
	roleMgr.players[player.Pid] = player
}

func (roleMgr *RoleMgrWrap) GetPlayerById(roleId int64) *GamePlayer {
	roleMgr.rwLock.RLock()
	defer roleMgr.rwLock.RUnlock()
	return roleMgr.players[roleId]
}
