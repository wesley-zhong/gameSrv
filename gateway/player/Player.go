package player

import (
	"gameSrv/pkg/client"
	"sync"
)

//------player

type Player struct {
	Context *client.ConnClientContext
	Pid     int64
	valid   bool
}

func (player *Player) SetValid() {
	player.valid = true
}
func (player *Player) SetContext(context *client.ConnClientContext) {
	player.Context = context
}

func NewPlayer(pid int64, context *client.ConnClientContext) *Player {
	return &Player{Context: context, Pid: pid}
}

var playerMutex sync.Mutex
var PlayerMgr *PlayerMgrWrap

// player mgr----
func NewPlayerMgr() *PlayerMgrWrap {
	PlayerMgr = &PlayerMgrWrap{
		playerIdMap: make(map[int64]*Player),
		//playerSidMap: make(map[int64]*Player),
	}
	return PlayerMgr
}

type PlayerMgrWrap struct {
	playerIdMap map[int64]*Player
	//playerSidMap map[int64]*Player
}

func (playerMgr *PlayerMgrWrap) Add(player *Player) {
	playerMutex.Lock()
	defer playerMutex.Unlock()
	playerMgr.playerIdMap[player.Pid] = player
	//playerMgr.playerSidMap[player.Context.Sid] = player
}

func (playerMgr *PlayerMgrWrap) Remove(player *Player) {
	playerMutex.Lock()
	defer playerMutex.Unlock()
	delete(playerMgr.playerIdMap, player.Pid)
	//delete(playerMgr.playerSidMap, player.Context.Sid)
}

func (playerMgr *PlayerMgrWrap) GetByRoleId(pid int64) *Player {
	playerMutex.Lock()
	defer playerMutex.Unlock()
	return playerMgr.playerIdMap[pid]
}

func (playerMgr *PlayerMgrWrap) GetSize() int {
	playerMutex.Lock()
	defer playerMutex.Unlock()
	return len(playerMgr.playerIdMap)
}

func (playerMgr *PlayerMgrWrap) GetPlayerList() map[int64]*Player {
	return playerMgr.playerIdMap
}
