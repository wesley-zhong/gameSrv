package dal

import (
	"gameSrv/game/module"
	"gameSrv/pkg/orm"
)

var PlayerDAO *orm.DAOIterface
var ItemDAO *orm.DAOIterface
var AccountDAO *orm.DAOIterface

func Init(addr string, userName string, pwd string) {
	orm.Init(addr, userName, pwd)
	AccountDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "account"), Object: &module.AccountDO{}}
	PlayerDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "player"), Object: &module.PlayerDO{}}
	ItemDAO = &orm.DAOIterface{Collection: orm.GetDBConnTable("game", "item"), Object: &module.ItemDO{}}
	//may be called like this
	//ItemDAO.FindOneById(1110010)
}
