package player

import (
	"sync"
)

var RoleOlineMgr = NewRoleMgr()

type MgrWrap struct {
	players map[int64]*GamePlayer
	rwLock  sync.RWMutex
}

func NewRoleMgr() *MgrWrap {
	return &MgrWrap{
		players: make(map[int64]*GamePlayer),
		rwLock:  sync.RWMutex{},
	}
}

func (roleMgr *MgrWrap) AddPlayer(player *GamePlayer) {
	roleMgr.rwLock.Lock()
	defer roleMgr.rwLock.Unlock()
	roleMgr.players[player.Id] = player
}

func (roleMgr *MgrWrap) GetPlayerById(pid int64) *GamePlayer {
	roleMgr.rwLock.RLock()
	defer roleMgr.rwLock.RUnlock()
	return roleMgr.players[pid]
}
