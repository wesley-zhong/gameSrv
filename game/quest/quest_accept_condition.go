package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gameevent"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

var acceptCnfProcess = make(map[int]func(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool)

func acceptConditionInit() {
	acceptCnfProcess[cfg.QuestAcceptConditionType_ROLE_LEVEL_UP] = RoleLvlUpAcceptCnd
	acceptCnfProcess[cfg.QuestAcceptConditionType_QUEST_FINISHED] = MainQuestFinishAcceptCnd
	acceptCnfProcess[cfg.QuestAcceptConditionType_KILL_MONSTER] = KillMonsterAcceptCnd
	acceptCnfProcess[cfg.QuestAcceptConditionType_OBTAIN_ITEM] = ObtainItemAcceptCnd
}

func processQuestAcceptByEvent(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool {
	ret := true
	for _, cond := range cnf.AcceptCond {
		processFuc := acceptCnfProcess[int(cond.Type)]
		if processFuc != nil {
			ret = ret && processFuc(player, cnf, ev)
			if cnf.AcceptCondComb == cfg.LogicType_AND {
				if !ret {
					return false
				}
				continue
			}
			if ret {
				return true
			}
		}
	}
	return true
}

func RoleLvlUpAcceptCnd(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool {
	roleLvlUp := ev.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role lvl up  cur lvl ={}", roleLvlUp.CurLvl)
	return false
}

// 完成主线任务 触发
func MainQuestFinishAcceptCnd(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool {
	questFinsih := ev.(*gameevent.MainQuestFinishEvent)
	log.Infof("on role lvl up  cur lvl ={}", questFinsih.MainQuestId)
	return false
}

// 杀死 触发
func KillMonsterAcceptCnd(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool {

	return false
}

// 获取道具 触发
func ObtainItemAcceptCnd(player *player.GamePlayer, cnf *cfg.QuestQuestCnf, ev event.Event) bool {

	return false
}
