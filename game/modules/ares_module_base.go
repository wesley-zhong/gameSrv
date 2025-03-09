package modules

import (
	"errors"
	"gameSrv/pkg/gopool"
	"gameSrv/pkg/log"
	"gameSrv/pkg/orm"
)

/**
base module define  and  modules container define
*/

type ModuleId int

const (
	ROLE_MODULE ModuleId = 1
	ITEM_MODULE ModuleId = 2

	MAX_ITEM_MODULES = 5
)

type AresModule interface {
	LoadFromDB()
	Destroy()
}

type AresModuleBase[DOType any] struct {
	ModuleId ModuleId
	Player   IGmePlayer
	DAO      *orm.MongodbDAOInterface[DOType]
	dataObj  *DOType
	onFromDO func(*DOType)
	toDO     func() *DOType
}

var DbWriteGoPool = gopool.StartNewWorkerPool(1, 8192)

func (module *AresModuleBase[DOType]) LoadFromDB() {
	do := module.DAO.FindOneById(module.Player.GetPlayerId())
	module.onFromDO(do)
}

func (module *AresModuleBase[DOType]) Init(id ModuleId, player IGmePlayer) {
	module.ModuleId = id
	module.Player = player
}

func (module *AresModuleBase[DOType]) SaveDB() {
	err := DbWriteGoPool.SubmitTaskByHashCode((int)(module.Player.GetPlayerId()), func() {
		module.DAO.Save(module.Player.GetPlayerId(), module.toDO())
	})
	if err != nil {
		log.Error(err)
	}
}

func (module *AresModuleBase[DOType]) Destroy() {
	// DO some error log
	log.Error(errors.New(" sub class not implement"))
}

func (module *AresModuleBase[DOType]) InitModule(id ModuleId, player IGmePlayer) {
	log.Error(errors.New(" sub class not implement"))

}
