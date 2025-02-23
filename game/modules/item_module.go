package modules

import "gameSrv/game/DO"

type ItemModule struct {
	AresModuleBase[DO.ItemDO]
	Items map[int64]*DO.Item
}

func (module *ItemModule) FromDO(do *DO.ItemDO) {

	// fill data from here
	//module.Items = make(map[int64]*DO.Item)
}

func (module *ItemModule) ToDO() *DO.ItemDO {

	return module.dataObj
}

func (module *ItemModule) GetItem(id int64) *DO.Item {
	return nil
}

func (module *ItemModule) UseItems(id int64) {

}
