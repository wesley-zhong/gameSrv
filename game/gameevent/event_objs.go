package gameevent

import (
	"gameSrv/pkg/event"
	"gameSrv/pkg/scene"
)

const (
	LoginEventID           event.GameEventID = 1
	LogoutEventID          event.GameEventID = 2
	DisconnectEventID      event.GameEventID = 3
	QuestStepFinishEventID event.GameEventID = 4
	MainQuestFinishEventID event.GameEventID = 5
	GetAvatarEventID       event.GameEventID = 6
	RoleLvlUpdateEventId   event.GameEventID = 7
	ObtainItemEventID      event.GameEventID = 8
	KillMonsterEventID     event.GameEventID = 9
	QuestInitEventID       event.GameEventID = 10
)

type GameEvent struct {
	//PlayerId int64
	eventId event.GameEventID
	player  scene.IGamePlayer
}

// 给基础结构体实现这个方法
func (ge *GameEvent) Init(gm scene.IGamePlayer, eid event.GameEventID) {
	ge.player = gm
	ge.eventId = eid
}

func (ge *GameEvent) EventId() event.GameEventID {
	return ge.eventId
}

func (ge *GameEvent) Player() scene.IGamePlayer {
	return ge.player
}

func NewEvent[T any, PT interface {
	*T
	event.Event
}](gm scene.IGamePlayer, evId event.GameEventID) PT {
	ev := PT(new(T))  // 分配内存并转为接口类型
	ev.Init(gm, evId) // 初始化基础字段
	return ev
}

type MainQuestFinishEvent struct {
	GameEvent
	MainQuestId int64
}

type QuestStepFinishEvent struct {
	GameEvent
	StepQuestId int64
}

type GetAvatarEvent struct {
	GameEvent
	AvatarCnfId int64
}

type RoleLvlUpEvent struct {
	GameEvent
	CurLvl int32
}

type KillMonsterEvent struct {
	GameEvent
	MonsterId int64
}

type ObtainItemEvent struct {
	GameEvent
	ItemId int64
	CnfId  int32
	Num    int32
}
