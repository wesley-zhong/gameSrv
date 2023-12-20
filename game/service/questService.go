package service

import (
	"gameSrv/gateway/player"
)

const (
	IVALID        = 0
	GOT_ITE       = 1
	KILL_MONSTE   = 2
	WIN_PVP_COUNT = 3
)

type KillMonster struct {
	MonsterId int64
	Count     int32
}
type BattleResult struct {
	Score        int32
	KillMonsters []KillMonster
}

type EventId int32

type QuestCond struct {
	EventId     EventId
	ExtData     string
	TargetCount int32
	CurCount    int32
}
type Quest struct {
	Qid       int64
	QuestCnds []QuestCond
}

type QuestProcess func(player *player.Player, quest *Quest)

func QuestKillMonster(player *player.Player, quest *Quest) {

}

func DispatchEvent(player *player.Player, eventId EventId, result BattleResult) []*Quest {
	quests := getQuestByEventId(eventId)
	if len(quests) == 0 {
		return nil
	}
	process := QuestProcessMap[eventId]
	if process == nil {
		return nil
	}
	for _, quest := range quests {
		process(player, quest)
	}

	return quests
}

var QuestProcessMap map[EventId]QuestProcess = make(map[EventId]QuestProcess)

func registerQuestProcessFun(id EventId, process QuestProcess) {
	QuestProcessMap[id] = process
}

func QuestServiceInit() {
	registerQuestProcessFun(KILL_MONSTE, QuestKillMonster)
}

func getQuestByEventId(id EventId) []*Quest {
	return nil
}
