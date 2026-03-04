package modules

import "gameSrv/pkg/log"

type ItemDO struct {
	Id    int64
	Name  string
	Count int64
}

type ItemModule struct {
	GameModule[ItemDO]
}

func (itemModule *ItemModule) OnDataLoaded() error {
	return nil
}

func (itemModule *ItemModule) ModuleId() ModuleTypeId {
	return ITEM_MODULE
}

func (itemModule *ItemModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if itemModule.DataDO == nil {
		itemModule.DataDO = &ItemDO{
			Id:    itemModule.Pid,
			Name:  "itemModule",
			Count: 100,
		}
		itemModule.MarkDirty()
	}
	return
}
