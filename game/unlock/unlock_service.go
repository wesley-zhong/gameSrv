package unlock

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/game/gameevent"
	"gameSrv/game/modules"
	"gameSrv/game/player"
	"gameSrv/pkg/event"
)

type CheckFunc func(event event.Event, cond *cfg.UnLockCondBean) bool

var unlockCheckId2Func map[int32]CheckFunc
var unlockCheckCndId2ConfigData map[int32][]int32

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
func mainQuestUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) bool {
	mainQuestFinish := event.(*gameevent.MainQuestFinishEvent)
	if mainQuestFinish.MainQuestId == cond.Value {
		//unlocked logic process
	}
	return true
}

func questStepUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) bool {
	stepFinish := event.(*gameevent.QuestStepFinishEvent)
	if stepFinish.StepQuestId == cond.Value {
		//unlocked logic process
	}
	return true
}

func getAvatarUnlockEventCheck(event event.Event, cond *cfg.UnLockCondBean) bool {
	avatarEvent := event.(*gameevent.GetAvatarEvent)
	if avatarEvent.AvatarCnfId == cond.Value {
		//unlocked logic process
	}
	return true
}

func processCndCnfigDataEvents() {
	unlockCheckCndId2ConfigData = make(map[int32][]int32)
	for _, v := range gamedata.Tables.TbCommonUnlock.GetDataList() {
		for _, cn := range v.UnlockCnds {
			unlockCheckCndId2ConfigData[cn.CndId] = append(unlockCheckCndId2ConfigData[cn.CndId], v.Id)
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

	for _, cdataId := range cndConfigDatas {
		ftnUnlockData := gamedata.Tables.TbCommonUnlock.Get(cdataId)
		if ftnUnlockData == nil {
			return
		}
		if checkCndValid(checkFunc, ftnUnlockData, event) {
			// check success
			gmEv := event.(*gameevent.GameEvent)
			gp := player.RoleOlineMgr.GetPlayerById(gmEv.PlayerId)
			if gp == nil {
				return
			}
			module := gp.ModuleContainer.IModules[modules.UNLOCK_MODULE]
			if module == nil {
				return
			}
		}
	}
}

func checkCndValid(checkFunc CheckFunc, unlockData *cfg.SysCommonUnlock, event event.Event) bool {
	logicAnd := unlockData.UnLockOptType == cfg.UnLockOpt_AND
	for _, cn := range unlockData.UnlockCnds {
		if logicAnd {
			if !checkFunc(event, cn) {
				return false
			}
		} else {
			if checkFunc(event, cn) {
				return true
			}
		}
	}
	if logicAnd {
		return true
	}
	return false
}
