package modules

import (
	"gameSrv/game/DO"
	"gameSrv/game/dal"
	"gameSrv/game/player"
)

type ItemModule struct {
	AresModuleBase[DO.ItemDO]
	Items map[int64]*DO.Item
}

func (module *ItemModule) InitModule(id ModuleId, player *player.GamePlayer) {
	module.ModuleId = id
	module.Player = player
	module.DAO = dal.ItemDAO
	//module.dataObj = nil
	module.onFromDO = module.FromDO
	module.toDO = module.ToDO
}

func (module *ItemModule) FromDO(do *DO.ItemDO) {

	// fill data from here
	module.Items = make(map[int64]*DO.Item)
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
