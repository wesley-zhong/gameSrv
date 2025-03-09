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
	module.Init(id, player)
	module.DAO = dal.RoleDAO
	//module.dataObj = nil
	module.onFromDO = module.FromDO
	module.toDO = module.ToDO
}

func (module *RoleModule) FromDO(do *DO.RoleDO) {

	// fill data from here
	//module.Items = make(map[int64]*DO.Item)
	if do == nil {
		do = &DO.RoleDO{
			Id:   module.Player.GetPlayerId(),
			Name: "haha",
		}
		module.dataObj = do
		module.AsynSaveDB()
	}
	module.dataObj = do
}

func (module *RoleModule) ToDO() *DO.RoleDO {
	return module.dataObj
}
