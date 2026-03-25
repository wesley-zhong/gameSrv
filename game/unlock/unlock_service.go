package unlock

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/gameevent"
	"gameSrv/pkg/event"
)

type Event struct {
}
type UnlockFunc func(event event.Event, cond *cfg.UnLockCondBean)

var unlockCheckFuncs map[int32]UnlockFunc
var unlockCheckData map[int32][]*cfg.UnLockCondBean

func InitEvents() {
	event.Dispatcher.Register(gameevent.QuestStepUnlockEventID, questStepUnlockEvent)
	event.Dispatcher.Register(gameevent.MainQuestUnlockEventID, mainQuestUnlockEvent)
	//unlock check functions
	unlockCheckFuncs = make(map[int32]UnlockFunc)
	unlockCheckFuncs[cfg.UnLockCnd_MAIN_TASK] = mainQuestUnlockEventCheck
	unlockCheckFuncs[cfg.UnLockCnd_GUIDE_STEP] = questStepUnlockEventCheck
	unlockCheckFuncs[cfg.UnLockCnd_GET_AVATAR] = getAvatarUnlockEventCheck

	//unlock data
	registerCnfDataEventHandler()
}

func questStepUnlockEvent(event event.Event) {
	dispatch(cfg.UnLockCnd_GUIDE_STEP, event)
}

func mainQuestUnlockEvent(event event.Event) {
	dispatch(cfg.UnLockCnd_MAIN_TASK, event)
}

// unlock check
func mainQuestUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) {
	mainQuestFinish := event.(*gameevent.MainQuestFinishEvent)
	if mainQuestFinish.MainQuestId == cond.Value {
		//unlocked logic process

	}
}

func questStepUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) {
	stepFinish := event.(*gameevent.QuestStepFinishEvent)
	if stepFinish.StepQuestId == cond.Value {
		//unlocked logic process
	}
}

func getAvatarUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) {
	avatarEvent := event.(*gameevent.GetAvatarEvent)
	if avatarEvent.AvatarCnfId == cond.Value {
		//unlocked logic process
	}
}

func registerCnfDataEventHandler() {
	unlockCheckData = make(map[int32][]*cfg.UnLockCondBean)
	for _, v := range gamedata.Tables.TbCommonUnlock.GetDataList() {
		for _, cn := range v.UnlockCnds {
			unlockCheckData[cn.CndId] = append(unlockCheckData[cn.CndId], cn)
		}
	}
}

func dispatch(unlockCndId int32, event event.Event) {
	checkFunc := unlockCheckFuncs[unlockCndId]
	if checkFunc == nil {
		return
	}
	cndDatas := unlockCheckData[unlockCndId]
	if cndDatas == nil || len(cndDatas) == 0 {
		return
	}

	for _, cn := range cndDatas {
		checkFunc(event, cn)
	}
}
