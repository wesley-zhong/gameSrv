package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/gameevent"
	"gameSrv/game/modules"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

var questProcess = make(map[int]func(*player.GamePlayer, *modules.Quest, event.Event) ProRet)

type ProRet int32

const (
	NONE     ProRet = 0
	CONTINUE ProRet = 1
	FINISH   ProRet = 2
	FAILED   ProRet = 3
)

func finishContentInit() {
	questProcess[cfg.QuestContentType_ROLE_LEVEL_UP] = RoleLvlUpProcess
	questProcess[cfg.QuestContentType_QUEST_FINISHED] = MainQuestFinishProcess
	questProcess[cfg.QuestContentType_KILL_MONSTER] = KillMonsterProcess
	questProcess[cfg.QuestContentType_OBTAIN_ITEM] = ObtainItemProcess
}

func ProcessQuestContentByEvent(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {
	processRet := NONE
	questStepCnf := gamedata.Tables.TbQuestStep.Get(quest.Id)
	if questStepCnf != nil {
		return processRet
	}

	for _, content := range questStepCnf.FinishCond {
		processFuc := questProcess[int(content.Type)]
		if processFuc == nil {
			log.Errorf(" finish type ={} not found process function", content.Type)
		}
		processRet = processFuc(player, quest, ev)
		quest.Finished = checkQuestFinished(quest, questStepCnf)

		if quest.Finished {
			processRet = FINISH
			return processRet
		}
	}
	return processRet
}

func RoleLvlUpProcess(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {
	roleLvlUp := ev.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role lvl up  cur lvl ={}", roleLvlUp.CurLvl)
	return CONTINUE
}

// 完成主线任务 触发
func MainQuestFinishProcess(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {
	questFinsih := ev.(*gameevent.MainQuestFinishEvent)
	log.Infof("on role lvl up  cur lvl ={}", questFinsih.MainQuestId)
	return CONTINUE
}

// 完成分支任务 触发
func StepQuestFinishProcess(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {
	questFinish := ev.(*gameevent.QuestStepFinishEvent)
	log.Infof("on role lvl up  cur lvl ={}", questFinish.StepQuestId)
	return CONTINUE
}

// 杀死 触发
func KillMonsterProcess(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {

	return CONTINUE
}

// 获取道具 触发
func ObtainItemProcess(player *player.GamePlayer, quest *modules.Quest, ev event.Event) ProRet {

	return CONTINUE
}

// finish check
func checkQuestFinished(quest *modules.Quest, questStepCnf *cfg.QuestQuestStepCnf) bool {
	ret := true
	for _, cnd := range questStepCnf.FinishCond {
		ret = ret && checkFinishCnd(quest, cnd)
		// And logic
		if questStepCnf.FinishCondComb == cfg.LogicType_AND {
			if !ret {
				return false
			}
			continue
		}
		// Or logic
		if ret {
			return false
		}
	}
	return true
}

func checkFinishCnd(quest *modules.Quest, cnd *cfg.QuestContent) bool {
	for _, questData := range quest.CurData {
		if questData.EvId != cnd.Type {
			continue
		}
		for index, val := range questData.CurData {
			if val != cnd.Param[index] {
				return false
			}
		}
		return true
	}
	return false
}
