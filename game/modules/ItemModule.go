package modules

import "gameSrv/game/do"

type ItemModule struct {
	ItemDO *do.ItemDO
}

func (itemModule *ItemModule) FromDO(itemDo *do.ItemDO) error {
	itemModule.ItemDO = itemDo
	return nil
}

func (itemModule *ItemModule) ToDO() *do.ItemDO {
	return itemModule.ItemDO
}
