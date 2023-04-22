package player

import (
	"gameSvr/internal/client"
)

//------player

type Player struct {
	Context *client.ConnClientContext
	Pid     int64
}

func NewPlayer(pid int64, context *client.ConnClientContext) *Player {
	return &Player{Context: context, Pid: pid}
}

var PlayerMgr *PlayerMgrWrap

// player mgr----

func NewPlayerMgr() *PlayerMgrWrap {
	PlayerMgr = &PlayerMgrWrap{
		playerIdMap:  make(map[int64]*Player),
		playerSidMap: make(map[int64]*Player),
	}
	return PlayerMgr
}

type PlayerMgrWrap struct {
	playerIdMap  map[int64]*Player
	playerSidMap map[int64]*Player
}

func (playerMgr *PlayerMgrWrap) Add(player *Player) {
	playerMgr.playerIdMap[int64(player.Pid)] = player
	playerMgr.playerSidMap[player.Context.Sid] = player
}

func (playerMgr *PlayerMgrWrap) GetByRoleId(pid int64) *Player {
	return playerMgr.playerIdMap[pid]
}

func (playerMgr *PlayerMgrWrap) GetByContext(context *client.ConnClientContext) *Player {
	return playerMgr.playerSidMap[context.Sid]
}

func (playerMgr *PlayerMgrWrap) GetBySid(sid int64) *Player {
	return playerMgr.playerSidMap[sid]
}
