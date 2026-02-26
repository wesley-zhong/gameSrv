package dal

import (
	"gameSrv/game/do"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------

var AccountDAO *orm.MongodbDAO[do.AccountDO]
var RoleDAO *orm.MongodbDAO[do.RoleDO]
var ItemDAO *orm.MongodbDAO[do.ItemDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAO[do.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAO[do.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAO[do.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}

	return nil
}

//------------------------ redis-------------

var AccountRedisDAO *orm.RedisDAO[do.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAO[do.AccountDO]{Rdb: orm.Rdb}

	return nil
}
