package dal

import (
	"gameSrv/game/do"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAOInterface[do.AccountDO]
var RoleDAO *orm.MongodbDAOInterface[do.RoleDO]
var ItemDAO *orm.MongodbDAOInterface[do.ItemDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAOInterface[do.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAOInterface[do.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAOInterface[do.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}

	return nil
}

//------------------------ redis-------------

var AccountRedisDAO *orm.RedisDAOInterface[do.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAOInterface[do.AccountDO]{Rdb: orm.Rdb}

	return nil
}
