package modules

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gameevent"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

type QuestData struct {
	EvId    int32
	CurData []int64
}
type Quest struct {
	Id       int64
	CurData  []*QuestData
	Finished bool
}
type QuestDO struct {
	Id     int64            //  player id
	CurQId []int64          //当前任务 ids
	Quests map[int64]*Quest // quest cur data
}

type QuestModule struct {
	GameModule[QuestDO]
}

func (questModule *QuestModule) ModuleId() ModuleTypeId {
	return QUEUE_MODULE
}

func (questModule *QuestModule) OnDataLoaded() error {
	return nil
}

func (questModule *QuestModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if questModule.DataDO == nil {
		questModule.DataDO = &QuestDO{
			Id:     questModule.Pid,
			Quests: make(map[int64]*Quest),
			CurQId: make([]int64, 0),
		}
		event.Dispatcher.Dispatch(gameevent.NewEvent[gameevent.GameEvent](questModule.Pid, gameevent.QuestInitEventID))
		questModule.MarkDirty()
	}
	return
}

func (questModule *QuestModule) AddQuest(questCnf *cfg.QuestQuestCnf) *Quest {
	newQuest := &Quest{
		Id:       questCnf.Id,
		Finished: false,
	}
	questModule.DataDO.Quests[questCnf.Id] = newQuest
	questModule.MarkDirty()
	return newQuest
}

func (questModule *QuestModule) AddQuestStep(questStepCnf *cfg.QuestQuestStepCnf) *Quest {
	newQuest := &Quest{
		Id:       questStepCnf.Id,
		Finished: false,
	}
	conds := questStepCnf.FinishCond
	for _, cond := range conds {
		questData := &QuestData{
			EvId:    cond.Type,
			CurData: []int64{0},
		}
		newQuest.CurData = append(newQuest.CurData, questData)
	}
	questModule.DataDO.Quests[questStepCnf.Id] = newQuest
	questModule.MarkDirty()
	return newQuest
}

func (questModule *QuestModule) FindQuest(questId int64) *Quest {
	return questModule.DataDO.Quests[questId]
}

func (questModule *QuestModule) FindQuestByEventId(evId int32) []*Quest {
	quests := questModule.DataDO.Quests
	questList := make([]*Quest, 0, len(quests))
	for _, quest := range quests {
		if quest.Finished {
			continue
		}
		for _, curData := range quest.CurData {
			if curData.EvId == evId {
				questList = append(questList, quest)
			}
		}
	}
	return questList
}

func (questModule *QuestModule) RemoveQuest(questId int64) {

}

func (questModule *QuestModule) OnLogout() {

}

func (questModule *QuestModule) OnDisconnect() {

}
