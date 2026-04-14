package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/player"
	"gameSrv/pkg/log"
)

var questExeEvent = make(map[int32]func(*player.GamePlayer, *cfg.QuestQuestStepCnf))

func initQuestExeEventHandler() {
	questExeEvent[cfg.QuestExecType_ADD_BUFF] = exeAddBuff
	questExeEvent[cfg.QuestExecType_TELEPORT_PLAYER] = exeTeleport

	// some other exec events
}

func exeQuestFinishedEvent(gp *player.GamePlayer, questStep *cfg.QuestQuestStepCnf) {
	for _, event := range questStep.FinishExec {
		eventFun := questExeEvent[event.Type]
		if eventFun == nil {
			log.Errorf("not found questExeEvent[%d]", event.Type)
			continue
		}
		eventFun(gp, questStep)
	}
}

func exeQuestBeginEvent(gp *player.GamePlayer, questStep *cfg.QuestQuestStepCnf) {
	for _, event := range questStep.BeginExec {
		eventFun := questExeEvent[event.Type]
		if eventFun == nil {
			log.Errorf("not found questExeEvent[%d]", event.Type)
			continue
		}
		eventFun(gp, questStep)
	}
}

// ===========================================  some other exec events
func exeTeleport(gamePlayer *player.GamePlayer, cnf *cfg.QuestQuestStepCnf) {

}

func exeAddBuff(gamePlayer *player.GamePlayer, cnf *cfg.QuestQuestStepCnf) {

}
