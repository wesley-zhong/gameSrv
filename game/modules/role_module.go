package modules

import "gameSrv/pkg/log"

type RoleDO struct {
	Id   int64
	Name string
}

type RoleModule struct {
	GameModule[RoleDO]
}

func (roleModule *RoleModule) ModuleId() ModuleTypeId {
	return ROLE_MODULE
}

func (roleModule *RoleModule) OnDataLoaded() error {
	return nil
}

func (roleModule *RoleModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if roleModule.DataDO == nil {
		roleModule.DataDO = &RoleDO{
			Id:   roleModule.Pid,
			Name: "name",
		}
		roleModule.MarkDirty()
	}
	return
}
