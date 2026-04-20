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

func Init() {
	refactorQuestCnfData()
	acceptConditionInit()
	finishContentInit()
	event.Dispatcher.Register(gameevent.RoleLvlUpdateEventId, onRoleLevelUpdateEvent)
	event.Dispatcher.Register(gameevent.KillMonsterEventID, onKillMonsterEvent)
	event.Dispatcher.Register(gameevent.ObtainItemEventID, onObtainItemEvent)
	event.Dispatcher.Register(gameevent.QuestStepFinishEventID, onQuestFinished)
	event.Dispatcher.Register(gameevent.QuestInitEventID, onPlayerQuestInit)
}

func initPlayerQuest(gamePlayer *player.GamePlayer) {
	readyToAcceptQuestList := getAcceptedQuestByEventId(gamePlayer, 0)
	for _, questCnf := range readyToAcceptQuestList {
		acceptNewQuest(gamePlayer, questCnf)
	}
}

func processQuestByEvent(gamePlayer *player.GamePlayer, evId int32, ev event.Event) {
	//process accept new quest
	readyToAcceptQuestList := getAcceptedQuestByEventId(gamePlayer, evId)
	for _, questCnf := range readyToAcceptQuestList {
		accepted := processQuestAcceptByEvent(gamePlayer, questCnf, ev)
		if accepted {
			acceptNewQuest(gamePlayer, questCnf)
		}
	}
	// process own quest  content
	ownQuestList := getOwnQuestStepByEventId(gamePlayer, evId)
	for _, quest := range ownQuestList {
		finished := ProcessQuestContentByEvent(gamePlayer, quest, ev)
		if finished == FINISH {
			finishQuest(gamePlayer, quest)
		}
	}
}

func acceptNewQuest(gamePlayer *player.GamePlayer, questCnf *cfg.QuestQuestCnf) {
	if questCnf == nil {
		return
	}
	questModule := player.GetModule[modules.QuestModule](gamePlayer, modules.QUEUE_MODULE)
	questModule.AddQuest(questCnf)
	if len(questCnf.ChildQuestList) == 0 {
		log.Warnf("questId =%d not found quest step", questCnf.Id)
		return
	}
	//get the main quest first quest step
	questStep := questCnf.ChildQuestList[0]
	acceptNewQuestStep(gamePlayer, questStep)
}
func acceptNewQuestStep(gamePlayer *player.GamePlayer, questStep *cfg.QuestQuestStepCnf) {
	questModule := player.GetModule[modules.QuestModule](gamePlayer, modules.QUEUE_MODULE)
	questModule.AddQuestStep(questStep)
	exeQuestBeginEvent(gamePlayer, questStep)
}

func finishQuest(gamePlayer *player.GamePlayer, quest *modules.Quest) {
	questStepCnf := findQuestStepCnf(quest.Id)
	if questStepCnf != nil {
		log.Errorf("quest step cnf is %v", questStepCnf)
		return
	}
	if questStepCnf.FinishParent {
		event.Dispatcher.Dispatch(gameevent.NewEvent[gameevent.MainQuestFinishEvent](gamePlayer.Id, gameevent.MainQuestFinishEventID))
	} else {
		event.Dispatcher.Dispatch(gameevent.NewEvent[gameevent.QuestStepFinishEvent](gamePlayer.Id, gameevent.QuestStepFinishEventID))
	}
	exeQuestFinishedEvent(gamePlayer, questStepCnf)
	nextQuestStep := findNextQuestStep(quest)
	if nextQuestStep != nil {
		acceptNewQuestStep(gamePlayer, nextQuestStep)
	}
}

func getAcceptedQuestByEventId(gp *player.GamePlayer, evId int32) []*cfg.QuestQuestCnf {
	return findQuestWithAcceptEvent(gp, evId)
}
func getOwnQuestStepByEventId(gp *player.GamePlayer, evId int32) []*modules.Quest {
	questModule := player.GetModule[modules.QuestModule](gp, modules.QUEUE_MODULE)
	return questModule.FindQuestByEventId(evId)
}

func onRoleLevelUpdateEvent(event event.Event) {
	gameEvent := event.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	roleLvlUp := event.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role level update event id={}", roleLvlUp.CurLvl)
	processQuestByEvent(gamePlayer, cfg.QuestContentType_ROLE_LEVEL_UP, event)
}

func onKillMonsterEvent(ev event.Event) {
	gameEvent := ev.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	processQuestByEvent(gamePlayer, cfg.QuestAcceptConditionType_KILL_MONSTER, ev)
}

func onObtainItemEvent(ev event.Event) {
	gameEvent := ev.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	processQuestByEvent(gamePlayer, cfg.QuestAcceptConditionType_OBTAIN_ITEM, ev)
}

func onQuestFinished(ev event.Event) {
	gameEvent := ev.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	processQuestByEvent(gamePlayer, cfg.QuestAcceptConditionType_QUEST_FINISHED, ev)
}

func onPlayerQuestInit(ev event.Event) {
	gameEvent := ev.(*gameevent.GameEvent)
	gamePlayer := player.RoleOlineMgr.GetPlayerById(gameEvent.PlayerId)
	initPlayerQuest(gamePlayer)
}
