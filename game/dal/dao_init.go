package dal

import (
	"gameSrv/game/DO"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAOInterface[DO.AccountDO]
var RoleDAO *orm.MongodbDAOInterface[DO.RoleDO]
var ItemDAO *orm.MongodbDAOInterface[DO.ItemDO]

func initMongoAllDAO() {
	AccountDAO = &orm.MongodbDAOInterface[DO.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAOInterface[DO.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAOInterface[DO.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}
	// add other dao ...

}

// ------------------------ redis-------------
var AccountRedisDAO *orm.RedisDAOInterface[DO.AccountDO]

func initRedisAllDAO() {
	AccountRedisDAO = &orm.RedisDAOInterface[DO.AccountDO]{Rdb: orm.Rdb}

}
