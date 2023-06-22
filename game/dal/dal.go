package dal

import (
	"gameSrv/game/module"
	"gameSrv/pkg/orm"
)

var AccountDAO *orm.DAOInterface[module.AccountDO]
var RoleDAO *orm.DAOInterface[module.RoleDO]
var ItemDAO *orm.DAOInterface[module.ItemDO]

func Init(addr string, userName string, pwd string) {
	orm.Init(addr, userName, pwd)
	AccountDAO = &orm.DAOInterface[module.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.DAOInterface[module.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.DAOInterface[module.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}
}
