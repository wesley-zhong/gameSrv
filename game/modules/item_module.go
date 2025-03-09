package modules

import (
	"gameSrv/game/DO"
	"gameSrv/game/dal"
)

type ItemModule struct {
	AresModuleBase[DO.ItemDO]
	Items map[int64]*DO.Item
}

func (module *ItemModule) InitModule(id ModuleId, player IGmePlayer) {
	module.ModuleId = id
	module.Player = player
	module.DAO = dal.ItemDAO
	//module.dataObj = nil
	module.onFromDO = module.FromDO
	module.toDO = module.ToDO
}

func (module *ItemModule) FromDO(do *DO.ItemDO) {

	if do == nil {
		do = &DO.ItemDO{
			BaseDO: DO.BaseDO{Id: module.Player.GetPlayerId()},
		}
		module.Items = make(map[int64]*DO.Item)
		module.dataObj = do
		module.SaveDB()
	}
	// fill data from here
	module.dataObj = do
}

func (module *ItemModule) Destroy() {

	// fill data from here
	//	module.Items = make(map[int64]*DO.Item)
}

func (module *ItemModule) ToDO() *DO.ItemDO {

	return module.dataObj
}

func (module *ItemModule) GetItem(id int64) *DO.Item {
	return nil
}

func (module *ItemModule) UseItems(id int64) {

}
