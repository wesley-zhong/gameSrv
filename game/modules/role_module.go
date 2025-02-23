package modules

import "gameSrv/game/DO"

type RoleModule struct {
	AresModuleBase[DO.RoleDO]
	Roles map[int64]*DO.Role
}

func (module *RoleModule) FromDO(do *DO.RoleDO) {

	// fill data from here
	//module.Items = make(map[int64]*DO.Item)
}

func (module *RoleModule) ToDO() *DO.RoleDO {

	return module.dataObj
}
