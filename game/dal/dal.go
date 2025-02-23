package dal

import (
	"gameSrv/game/DO"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAOInterface[DO.AccountDO]
var RoleDAO *orm.MongodbDAOInterface[DO.RoleDO]
var ItemDAO *orm.MongodbDAOInterface[DO.ItemDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAOInterface[DO.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAOInterface[DO.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAOInterface[DO.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}

	return nil
}

//------------------------ redis-------------

var AccountRedisDAO *orm.RedisDAOInterface[DO.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAOInterface[DO.AccountDO]{Rdb: orm.Rdb}

	return nil
}
