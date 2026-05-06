package unlock

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/gameevent"
	"gameSrv/pkg/event"
)

type CheckFunc func(event event.Event, cond *cfg.UnLockCondBean)

var unlockCheckId2Func map[int32]CheckFunc
var unlockCheckCndId2ConfigData map[int32][]*cfg.UnLockCondBean

var gameEventId2UnlockCndId = map[event.GameEventID]int32{
	gameevent.QuestStepFinishEventID: cfg.UnLockCnd_GUIDE_STEP,
	gameevent.MainQuestFinishEventID: cfg.UnLockCnd_MAIN_TASK,
	gameevent.GetAvatarEventID:       cfg.UnLockCnd_GET_AVATAR,
}

func InitEvents() {
	for gameEventId := range gameEventId2UnlockCndId {
		event.Dispatcher.Register(gameEventId, gameEvent2UnlockEvent)
	}
	//unlock check functions
	unlockCheckId2Func = make(map[int32]CheckFunc)
	unlockCheckId2Func[cfg.UnLockCnd_MAIN_TASK] = mainQuestUnlockEventCheck
	unlockCheckId2Func[cfg.UnLockCnd_GUIDE_STEP] = questStepUnlockEventCheck
	unlockCheckId2Func[cfg.UnLockCnd_GET_AVATAR] = getAvatarUnlockEventCheck

	//unlock data config
	processCndCnfigDataEvents()
}

func gameEvent2UnlockEvent(event event.Event) {
	unlockEvent := gameEventId2UnlockCndId[event.EventId()]
	if unlockEvent == 0 {
		return
	}
	dispatch(unlockEvent, event)
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

func processCndCnfigDataEvents() {
	unlockCheckCndId2ConfigData = make(map[int32][]*cfg.UnLockCondBean)
	for _, v := range gamedata.Tables.TbCommonUnlock.GetDataList() {
		for _, cn := range v.UnlockCnds {
			unlockCheckCndId2ConfigData[cn.CndId] = append(unlockCheckCndId2ConfigData[cn.CndId], cn)
		}
	}
}

func dispatch(unlockCndId int32, event event.Event) {
	checkFunc := unlockCheckId2Func[unlockCndId]
	if checkFunc == nil {
		return
	}
	cndConfigDatas := unlockCheckCndId2ConfigData[unlockCndId]
	if cndConfigDatas == nil || len(cndConfigDatas) == 0 {
		return
	}

	for _, cn := range cndConfigDatas {
		checkFunc(event, cn)
	}
}
