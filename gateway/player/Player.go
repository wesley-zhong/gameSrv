package player

import (
	"gameSrv/pkg/client"
	"sync"
	"sync/atomic"
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
		playerIdMap: &sync.Map{},
	}
	return PlayerMgr
}

type PlayerMgrWrap struct {
	playerIdMap *sync.Map
	size        atomic.Int32
}

func (playerMgr *PlayerMgrWrap) Add(player *Player) {
	_, loaded := playerMgr.playerIdMap.LoadOrStore(player.Pid, player)
	if loaded {
		playerMgr.playerIdMap.Store(player.Pid, player)
		return
	}
	playerMgr.size.Add(1)
}

func (playerMgr *PlayerMgrWrap) Remove(player *Player) {
	_, ok := playerMgr.playerIdMap.LoadAndDelete(player.Pid)
	if ok {
		playerMgr.size.Add(-1)
	}
}

func (playerMgr *PlayerMgrWrap) GetByRoleId(pid int64) *Player {
	value, ok := playerMgr.playerIdMap.Load(pid)
	if ok {
		return value.(*Player)
	}
	return nil
}

func (playerMgr *PlayerMgrWrap) GetSize() int32 {
	return playerMgr.size.Load()
}

func (playerMgr *PlayerMgrWrap) Range(iter func(player *Player)) {
	playerMgr.playerIdMap.Range(func(key, value any) bool {
		iter(value.(*Player))
		return true
	})
}
