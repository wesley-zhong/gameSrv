package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/log"
	"gameSrv/pkg/scene"
)

var questExeEvent = make(map[int32]func(scene.IGamePlayer, *cfg.QuestQuestStepCnf))

func initQuestExeEventHandler() {
	questExeEvent[cfg.QuestExecType_ADD_BUFF] = exeAddBuff
	questExeEvent[cfg.QuestExecType_TELEPORT_PLAYER] = exeTeleport

	// some other exec events
}

func exeQuestFinishedEvent(gp scene.IGamePlayer, questStep *cfg.QuestQuestStepCnf) {
	for _, event := range questStep.FinishExec {
		eventFun := questExeEvent[event.Type]
		if eventFun == nil {
			log.Errorf("not found questExeEvent[%d]", event.Type)
			continue
		}
		eventFun(gp, questStep)
	}
}

func exeQuestBeginEvent(gp scene.IGamePlayer, questStep *cfg.QuestQuestStepCnf) {
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
func exeTeleport(gamePlayer scene.IGamePlayer, cnf *cfg.QuestQuestStepCnf) {

}

func exeAddBuff(gamePlayer scene.IGamePlayer, cnf *cfg.QuestQuestStepCnf) {

}
