package dal

import (
	"gameSrv/game/module"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------
var AccountDAO *orm.MongodbDAOInterface[module.AccountDO]
var RoleDAO *orm.MongodbDAOInterface[module.RoleDO]
var ItemDAO *orm.MongodbDAOInterface[module.ItemDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	AccountDAO = &orm.MongodbDAOInterface[module.AccountDO]{Collection: orm.GetDBConnTable("game", "account")}
	RoleDAO = &orm.MongodbDAOInterface[module.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAOInterface[module.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}

	return nil
}

//------------------------ redis-------------

var AccountRedisDAO *orm.RedisDAOInterface[module.AccountDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	AccountRedisDAO = &orm.RedisDAOInterface[module.AccountDO]{Rdb: orm.Rdb}

	return nil

}
