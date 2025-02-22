package modules

import "gameSrv/game/do"

type ItemModule struct {
	AresModuleBase[do.ItemDO]
	Items map[int64]*do.Item
}
