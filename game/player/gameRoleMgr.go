package player

import (
	"gameSrv/game/do"
	"sync"
)

type GameRole struct {
	Sid      int64
	RoleId   int64
	RoleName string
	ServerId int64
	RoleDo   *do.RoleDO
}

var RoleOlineMgr = NewRoleMgr()

func NewRole(roleId int64, roleDO *do.RoleDO) *GameRole {
	return &GameRole{
		RoleId: roleId,
		RoleDo: roleDO,
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
	roles  map[int64]*GameRole
	rwLock sync.RWMutex
}

func NewRoleMgr() *RoleMgrWrap {
	return &RoleMgrWrap{
		roles:  make(map[int64]*GameRole),
		rwLock: sync.RWMutex{},
	}
}

func (roleMgr *RoleMgrWrap) AddRole(role *GameRole) {
	roleMgr.rwLock.Lock()
	defer roleMgr.rwLock.Unlock()
	roleMgr.roles[role.RoleId] = role
}

func (roleMgr *RoleMgrWrap) GetByRoleId(roleId int64) *GameRole {
	roleMgr.rwLock.RLock()
	defer roleMgr.rwLock.RUnlock()
	return roleMgr.roles[roleId]
}
