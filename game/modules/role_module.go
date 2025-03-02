package modules

import (
	"gameSrv/game/DO"
	"gameSrv/game/dal"
)

type RoleModule struct {
	AresModuleBase[DO.RoleDO]
	Roles map[int64]*DO.Role
}

func (module *RoleModule) InitModule(id ModuleId, player IGmePlayer) {
	module.ModuleId = id
	module.Player = player
	module.DAO = dal.RoleDAO
	//module.dataObj = nil
	module.onFromDO = module.FromDO
	module.toDO = module.ToDO
}

func (module *RoleModule) FromDO(do *DO.RoleDO) {

	// fill data from here
	//module.Items = make(map[int64]*DO.Item)
}

func (module *RoleModule) ToDO() *DO.RoleDO {

	return module.dataObj
}
