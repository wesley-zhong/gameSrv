package modules

type RoleDO struct {
	Id   int64
	Name string
}

type RoleModule struct {
	GameModule[RoleDO]
}

func (roleModule *RoleModule) ModuleId() ModuleTypeId {
	return ITEM_MODULE
}

func (roleModule *RoleModule) FromDO(roleDO *RoleDO) error {
	roleModule.DataDO = roleDO
	return nil
}
