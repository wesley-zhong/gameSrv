package player

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/sync"
)

//------player

type Player struct {
	Context *client.ConnContext
	Pid     int64
	valid   bool
}

func (player *Player) SetValid() {
	player.valid = true
}
func (player *Player) SetContext(context *client.ConnContext) {
	player.Context = context
}

func NewPlayer(pid int64, context *client.ConnContext) *Player {
	return &Player{Context: context, Pid: pid}
}

var PlayerMgr *PlayerMgrWrap = NewPlayerMgr()

// player mgr----
func NewPlayerMgr() *PlayerMgrWrap {
	return &PlayerMgrWrap{
		playerIdMap: sync.NewSyncRWMap[int64, *Player](),
	}
}

type PlayerMgrWrap struct {
	playerIdMap *sync.SyncRWMap[int64, *Player]
}

func (playerMgr *PlayerMgrWrap) Add(player *Player) {
	playerMgr.playerIdMap.Store(player.Pid, player)
}

func (playerMgr *PlayerMgrWrap) Remove(player *Player) {
	playerMgr.playerIdMap.Delete(player.Pid)
}

func (playerMgr *PlayerMgrWrap) GetByRoleId(pid int64) *Player {
	player, ok := playerMgr.playerIdMap.Load(pid)
	if !ok {
		return nil
	}
	return player
}

func (playerMgr *PlayerMgrWrap) GetSize() int32 {
	return int32(playerMgr.playerIdMap.Size())
}

func (playerMgr *PlayerMgrWrap) Range(iter func(player *Player)) {
	playerMgr.playerIdMap.Range(func(pid int64, player *Player) bool {
		iter(player)
		return true
	})
}