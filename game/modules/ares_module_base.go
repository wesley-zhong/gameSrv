package modules

import (
	"errors"
	"gameSrv/game/player"
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
)

type AresModule interface {
	LoadFromDB()
	Destroy()
}

type AresModuleBase[DOType any] struct {
	ModuleId ModuleId
	Player   *player.GamePlayer
	DAO      *orm.MongodbDAOInterface[DOType]
	dataObj  *DOType
	onFromDO func(*DOType)
	toDO     func() *DOType
}

func (module *AresModuleBase[DOType]) LoadFromDB() {
	do := module.DAO.FindOneById(module.Player.Pid)
	module.onFromDO(do)
}

func (module *AresModuleBase[DOType]) Destroy() {
	// DO some error log
	log.Error(errors.New(" sub class not implement"))
}

func (module *AresModuleBase[DOType]) InitModule(id ModuleId, player *player.GamePlayer) {
	log.Error(errors.New(" sub class not implement"))

}
