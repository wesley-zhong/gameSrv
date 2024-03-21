package module

type WorldRole struct {
	Sid    int64
	RoleId int64
}

var WorldRoleMgr *WorldRoleMgrWrap = &WorldRoleMgrWrap{}

type WorldRoleMgrWrap struct {
	roleMap map[int64]*WorldRole
}

func (roleIst *WorldRoleMgrWrap) AddRole(role *WorldRole) {
	roleIst.roleMap[role.RoleId] = role
}
