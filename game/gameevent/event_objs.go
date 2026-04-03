package gameevent

import (
	"gameSrv/pkg/event"
)

const (
	LoginEventID           event.GameEventID = 1
	LogoutEventID          event.GameEventID = 2
	DisconnectEventID      event.GameEventID = 3
	QuestStepUnlockEventID event.GameEventID = 4
	MainQuestUnlockEventID event.GameEventID = 5
	GetAvatarEventID       event.GameEventID = 6
	RoleLvlUpdateEventId   event.GameEventID = 7
	ObtainItemEventID      event.GameEventID = 8
	KillMonsterEventID     event.GameEventID = 9
	QuestStepFinishEventID event.GameEventID = 10
)

type GameEvent struct {
	PlayerId int64
	eventId  event.GameEventID
}

// 给基础结构体实现这个方法
func (ge *GameEvent) Init(pid int64, eid event.GameEventID) {
	ge.PlayerId = pid
	ge.eventId = eid
}

func NewEvent[T any, PT interface {
	*T
	event.Event
}](pid int64, eid event.GameEventID) PT {
	ev := PT(new(T))  // 分配内存并转为接口类型
	ev.Init(pid, eid) // 初始化基础字段
	return ev
}

func (ge *GameEvent) EventId() event.GameEventID {
	return ge.eventId
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
