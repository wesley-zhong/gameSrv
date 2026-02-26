package dal

import (
	"gameSrv/login/dos"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAO[module.AccountDO]
var RoleDAO *orm.MongodbDAO[module.RoleDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAO[module.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAO[module.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	return nil
}

// ------------------------ redis-------------
var AccountRedisDAO *orm.RedisDAO[module.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAO[module.AccountDO]{Rdb: orm.Rdb}

	return nil
}
