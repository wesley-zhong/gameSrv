package dal

import (
	"gameSrv/game/modules"
	"gameSrv/pkg/orm"
)

// ------------------mongodb----------------

var RoleDAO *orm.MongodbDAO[modules.RoleDO]
var ItemDAO *orm.MongodbDAO[modules.ItemDO]
var QuestDAO *orm.MongodbDAO[modules.QuestDO]

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	RoleDAO = &orm.MongodbDAO[modules.RoleDO]{Collection: orm.GetDBConnTable("game", "role")}
	ItemDAO = &orm.MongodbDAO[modules.ItemDO]{Collection: orm.GetDBConnTable("game", "item")}
	QuestDAO = &orm.MongodbDAO[modules.QuestDO]{Collection: orm.GetDBConnTable("game", "quest")}
	return nil
}

//------------------------ redis-------------

var RoleRedisDAO *orm.RedisDAO[modules.RoleDO]

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	RoleRedisDAO = &orm.RedisDAO[modules.RoleDO]{Rdb: orm.Rdb}
	return nil
}
