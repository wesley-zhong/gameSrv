package dal

import (
	"gameSrv/pkg/orm"
)

func InitMongoDB(addr string, userName string, pwd string) error {
	err := orm.InitMongoDB(addr, userName, pwd)
	if err != nil {
		return err
	}
	initMongoAllDAO()
	return nil
}

func InitRedisDB(addr string, password string) error {
	err := orm.InitRedis(addr, password)
	if err != nil {
		return err
	}
	initRedisAllDAO()
	return nil
}
