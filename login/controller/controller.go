package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/web"
)

func Init(methodInterface web.HttpMethodInterface) {
	login := new(LoginController)
	methodInterface.RegisterController(login)
	role := new(RoleController)
	methodInterface.RegisterController(role)

}

var RoleOlineMgr = role.NewRoleMgr()
