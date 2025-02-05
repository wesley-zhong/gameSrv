package controller

import (
	"gameSrv/game/player"
	"gameSrv/pkg/web"
)

func Init(methodInterface web.HttpMethodInterface) {
	login := new(Login)
	methodInterface.RegisterController(login)
	role := new(Role)
	methodInterface.RegisterController(role)
}

var RoleOlineMgr = player.NewRoleMgr()
