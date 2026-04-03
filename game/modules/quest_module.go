package modules

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/log"
)

type QuestDO struct {
	Id   int64
	Name string
}

type QuestModule struct {
	GameModule[QuestDO]
}

func (QuestModule *QuestModule) ModuleId() ModuleTypeId {
	return QUEUE_MODULE
}

func (QuestModule *QuestModule) OnDataLoaded() error {
	return nil
}

func (QuestModule *QuestModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if QuestModule.DataDO == nil {
		QuestModule.DataDO = &QuestDO{
			Id:   QuestModule.Pid,
			Name: "name",
		}
		QuestModule.MarkDirty()
	}
	return
}

func (QuestModule *QuestModule) AddQuest(questCnf *cfg.QuestQuestCnf) {

}
