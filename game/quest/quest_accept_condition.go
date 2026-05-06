package quest

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gameevent"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

var acceptCndProcess = make(map[int]func(*player.GamePlayer, *cfg.QuestAcceptCondition, event.Event) bool)

func acceptConditionInit() {
	acceptCndProcess[cfg.QuestAcceptConditionType_ROLE_LEVEL_UP] = RoleLvlUpAcceptCnd
	acceptCndProcess[cfg.QuestAcceptConditionType_QUEST_FINISHED] = MainQuestFinishAcceptCnd
	acceptCndProcess[cfg.QuestAcceptConditionType_KILL_MONSTER] = KillMonsterAcceptCnd
	acceptCndProcess[cfg.QuestAcceptConditionType_OBTAIN_ITEM] = ObtainItemAcceptCnd
}

func processQuestAcceptByEvent(player *player.GamePlayer, cnd *cfg.QuestQuestCnf, ev event.Event) bool {
	ret := true
	for _, cond := range cnd.AcceptCond {
		processFuc := acceptCndProcess[int(cond.Type)]
		if processFuc != nil {
			ret = ret && processFuc(player, cond, ev)
			if cnd.AcceptCondComb == cfg.LogicType_AND {
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

func RoleLvlUpAcceptCnd(player *player.GamePlayer, cnd *cfg.QuestAcceptCondition, ev event.Event) bool {
	roleLvlUp := ev.(*gameevent.RoleLvlUpEvent)
	log.Infof("on role lvl up  cur lvl ={}", roleLvlUp.CurLvl)

	return roleLvlUp.CurLvl >= int32(cnd.Param[0])
}

// 完成主线任务 触发
func MainQuestFinishAcceptCnd(player *player.GamePlayer, cnd *cfg.QuestAcceptCondition, ev event.Event) bool {
	questFinsih := ev.(*gameevent.MainQuestFinishEvent)
	log.Infof("on role lvl up  cur lvl ={}", questFinsih.MainQuestId)
	return false
}

// 杀死 触发
func KillMonsterAcceptCnd(player *player.GamePlayer, cnd *cfg.QuestAcceptCondition, ev event.Event) bool {

	return false
}

// 获取道具 触发
func ObtainItemAcceptCnd(player *player.GamePlayer, cnd *cfg.QuestAcceptCondition, ev event.Event) bool {

	return false
}
