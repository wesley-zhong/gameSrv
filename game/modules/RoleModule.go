package modules

import "gameSrv/game/do"

type RoleModule struct {
	AresModuleBase[do.RoleDO]
	Roles map[int64]*do.Role
}
