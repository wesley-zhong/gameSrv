package modules

type ItemDO struct {
	Id      int64
	Account string
	Pid     int64
}

type ItemModule struct {
	GameModule[ItemDO]
}

func (itemModule *ItemModule) ModuleId() ModuleTypeId {
	return ITEM_MODULE
}

func (itemModule *ItemModule) FromDO(itemDo *ItemDO) error {
	itemModule.DataDO = itemDo
	return nil
}
