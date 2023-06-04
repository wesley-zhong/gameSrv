package role

type GameRole struct {
	Sid      int64
	RoleId   int64
	RoleName string
	ServerId int64
}

func NewRole(roleId int64) *GameRole {
	return &GameRole{
		RoleId: roleId,
	}
}
func (role *GameRole) GetRoleId() int64 {
	return role.RoleId
}
func (role *GameRole) GetServerId() int64 {
	return role.ServerId
}

type IGameRole interface {
	GetRoleId() uint64
	GetServerId() uint64
}

type RoleMgrWrap struct {
	roles map[int64]*GameRole
}

func NewRoleMgr() *RoleMgrWrap {
	return &RoleMgrWrap{roles: make(map[int64]*GameRole)}
}

func (roleMgr *RoleMgrWrap) AddRole(role *GameRole) {
	roleMgr.roles[role.RoleId] = role
}

func (roleMgr *RoleMgrWrap) GetByRoleId(roleId int64) *GameRole {
	return roleMgr.roles[roleId]
}
