package modules

import "gameSrv/game/DO"

type RoleModule struct {
	AresModuleBase[DO.RoleDO]
	Roles map[int64]*DO.Role
}
