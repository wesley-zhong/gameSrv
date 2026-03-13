package dal

import (
	"gameSrv/login/dos"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAO[dos.AccountDO]
var RoleDAO *orm.MongodbDAO[dos.AccountDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAO[dos.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}

	return nil
}

// ------------------------ redis-------------
var AccountRedisDAO *orm.RedisDAO[dos.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAO[dos.AccountDO]{Rdb: orm.Rdb}

	return nil
}
