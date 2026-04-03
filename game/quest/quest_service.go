package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/modules"

	"gameSrv/game/gameevent"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

const (
	IVALID        = 0
	GOT_ITE       = 1
	KILL_MONSTE   = 2
	WIN_PVP_COUNT = 3
)

type EventId int32

type QData struct {
	EvId    EventId
	CurData []int64
}
type Quest struct {
	Qid       int64   // quest cnf id
	QuestData []QData // quest cur data
}

func init() {
	event.Dispatcher.Register(gameevent.RoleLvlUpdateEventId, onRoleLevelUpdateEvent)
	event.Dispatcher.Register(gameevent.KillMonsterEventID, onKillMonsterEvent)
	event.Dispatcher.Register(gameevent.ObtainItemEventID, onObtainItemEvent)
	event.Dispatcher.Register(gameevent.QuestStepFinishEventID, onQuestFinished)
}

func processQuestByEvent(evId int, ev event.Event) {
	gameEvent := ev.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	//process accept
	readyToAcceptQuestList := getAcceptedQuestByEventId(evId)
	for _, questCnf := range readyToAcceptQuestList {
		baccepted := ProcessQuestAcceptByEvent(gamePlayer, questCnf, ev)
		if baccepted {
			questModule := player.GetModule[modules.QuestModule](gamePlayer, modules.QUEUE_MODULE)
			questModule.AddQuest(questCnf)
		}
	}

	ownQuestList := getOwnQuestByEventId(evId)
	for _, quest := range ownQuestList {
		bfnished := ProcessQuestContentByEvent(gamePlayer, quest, ev)
		if bfnished == FINISH {
			event.Dispatcher.Dispatch(gameevent.NewEvent[gameevent.MainQuestFinishEvent](gamePlayer.Id, gameevent.MainQuestUnlockEventID))
		}
	}
	// process own quest
}

func getAcceptedQuestByEventId(evId int) []*cfg.QuestQuestCnf {

	return nil
}
func getOwnQuestByEventId(id int) []*Quest {
	return nil
}

func onRoleLevelUpdateEvent(event event.Event) {
	roleLvlUp := event.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role level update event id={}", roleLvlUp.CurLvl)
}

func onKillMonsterEvent(ev event.Event) {
	processQuestByEvent(cfg.QuestAcceptConditionType_KILL_MONSTER, ev)
}

func onObtainItemEvent(ev event.Event) {
	processQuestByEvent(cfg.QuestAcceptConditionType_OBTAIN_ITEM, ev)
}

func onQuestFinished(ev event.Event) {
	processQuestByEvent(cfg.QuestAcceptConditionType_QUEST_FINISHED, ev)
}
