package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/core"
)

func Init(methodInterface core.HttpMethodInterface) {
	login := new(LoginController)
	methodInterface.RegisterController(login)
	role := new(RoleController)
	methodInterface.RegisterController(role)

}

var RoleOlineMgr = role.NewRoleMgr()
