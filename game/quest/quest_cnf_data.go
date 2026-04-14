package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/modules"
	"gameSrv/game/player"
)

var acceptCondQuestMap = make(map[int32][]*cfg.QuestQuestCnf)
var unlockCndMap = make(map[int32][]*cfg.QuestQuestCnf)
var questMap = make(map[int64]*cfg.QuestQuestCnf)
var questContentEventQuest = make(map[int32][]*cfg.QuestQuestStepCnf)

func refactorQuestCnfData() {
	refactorQuestCnf()
	refactorQuestStepCnf()
}

func refactorQuestCnf() {
	for _, val := range gamedata.Tables.TbQuest.GetDataList() {
		conds := val.AcceptCond
		for _, cond := range conds {
			acceptCondQuestMap[cond.Type] = append(acceptCondQuestMap[cond.Type], val)
		}

		for _, unlock := range val.UnlockCond {
			unlockCndMap[unlock.Type] = append(unlockCndMap[unlock.Type], val)
		}
		questMap[val.Id] = val
	}
}

func refactorQuestStepCnf() {
	for _, val := range gamedata.Tables.TbQuestStep.GetDataList() {
		conds := val.FinishCond
		for _, cond := range conds {
			questContentEventQuest[cond.Type] = append(questContentEventQuest[cond.Type], val)
		}
		//set parent and child quest
		questMap[val.Id].ChildQuestMap[val.Id] = val
	}
}

func findQuestStepWithEvent(evId int32) []*cfg.QuestQuestStepCnf {
	return questContentEventQuest[evId]
}

func findQuestWithAcceptEvent(gp *player.GamePlayer, evId int32) []*cfg.QuestQuestCnf {
	readyToAcceptQuestCnfList := acceptCondQuestMap[evId]
	questModule := player.GetModule[modules.QuestModule](gp, modules.QUEUE_MODULE)
	readyToAcceptList := make([]*cfg.QuestQuestCnf, len(readyToAcceptQuestCnfList))
	for _, readyToAccept := range readyToAcceptQuestCnfList {
		quest := questModule.FindQuest(readyToAccept.Id)
		if quest == nil {
			readyToAcceptList = append(readyToAcceptList, readyToAccept)
		}
	}
	return readyToAcceptList
}
