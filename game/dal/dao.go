package dal

import (
	"gameSrv/game/module"
	"gameSrv/pkg/orm"
)

var AccountDAO *orm.DAOIterface
var RoleDAO *orm.DAOIterface
var ItemDAO *orm.DAOIterface

func Init(addr string, userName string, pwd string) {
	orm.Init(addr, userName, pwd)
	AccountDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "account"), Object: &module.AccountDO{}}
	RoleDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "role"), Object: &module.RoleDO{}}
	ItemDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "item"), Object: &module.ItemDO{}}
}
