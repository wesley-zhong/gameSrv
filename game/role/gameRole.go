package role

type GameRole struct {
	RoleId   uint64
	RoleName string
	ServerId uint64
}

func (role *GameRole) GetRoleId() uint64 {
	return role.RoleId
}
func (role *GameRole) GetServerId() uint64 {
	return role.ServerId
}

type IGameRole interface {
	GetRoleId() uint64
	GetServerId() uint64
}
