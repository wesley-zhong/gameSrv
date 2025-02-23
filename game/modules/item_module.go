package modules

import "gameSrv/game/DO"

type ItemModule struct {
	AresModuleBase[DO.ItemDO]
	Items map[int64]*DO.Item
}

func (module *ItemModule) FromDO(do *DO.ItemDO) {

}

func (module *ItemModule) ToDO() {

}
