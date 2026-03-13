package player

import (
	"gameSrv/pkg/sync"
)

var RoleOlineMgr = NewRoleMgr()

type MgrWrap struct {
	// 使用分段锁的并发安全map
	players *sync.SyncRWMap[int64, *GamePlayer]
}

func NewRoleMgr() *MgrWrap {
	return &MgrWrap{
		players: sync.NewSyncRWMap[int64, *GamePlayer](),
	}
}

func (roleMgr *MgrWrap) AddPlayer(player *GamePlayer) {
	roleMgr.players.Store(player.Id, player)
}

func (roleMgr *MgrWrap) GetPlayerById(pid int64) *GamePlayer {
	player, ok := roleMgr.players.Load(pid)
	if !ok {
		return nil
	}
	return player
}

func (roleMgr *MgrWrap) Size() int {
	return roleMgr.players.Size()
}

func (roleMgr *MgrWrap) Remove(pid int64) *GamePlayer {
	player, ok := roleMgr.players.LoadAndDelete(pid)
	if !ok {
		return nil
	}
	return player
}

// Range 遍历所有玩家，fn返回false时停止遍历
func (roleMgr *MgrWrap) Range(fn func(pid int64, player *GamePlayer) bool) {
	roleMgr.players.Range(fn)
}