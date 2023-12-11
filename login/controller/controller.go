package controller

import (
	"gameSrv/game/role"
	"gameSrv/pkg/web"
)

func Init(methodInterface web.HttpMethodInterface) {
	login := new(Login)
	methodInterface.RegisterController(login)
	role := new(Role)
	methodInterface.RegisterController(role)

}

var RoleOlineMgr = role.NewRoleMgr()
