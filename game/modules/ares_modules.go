package modules

import (
	"errors"
	"gameSrv/game/DO"
	"gameSrv/game/dal"
	"gameSrv/pkg/log"
	"gameSrv/pkg/orm"
)

/**
base module define  and  modules container define
*/

type AresModule interface {
	LoadFromDB()
	Destroy()
}

type AresModuleBase[DOType any] struct {
	playerId int64
	DAO      *orm.MongodbDAOInterface[DOType]
	dataObj  *DOType
}

func (module *AresModuleBase[DOType]) LoadFromDB() {
	do := module.DAO.FindOneById(module.playerId)
	module.FromDO(do)
}

func (module *AresModuleBase[DOType]) Destroy() {
	// DO some error log
	log.Error(errors.New(" sub class not implement"))
}

func (module *AresModuleBase[DOType]) FromDO(do *DOType) {
	log.Error(errors.New(" sub class not implement"))
}

func (module *AresModuleBase[DOType]) ToDO() *DOType {
	return module.dataObj
}

type ModuleContainer struct {
	Pid     int64
	Modules []AresModule
}

func NewModuleContainer(pid int64) *ModuleContainer {
	moduleContainer := &ModuleContainer{}
	moduleContainer.Pid = pid
	return moduleContainer
}

func (moduleContainer *ModuleContainer) InitModules() {
	itemModule := &ItemModule{
		AresModuleBase: AresModuleBase[DO.ItemDO]{
			playerId: moduleContainer.Pid,
			DAO:      dal.ItemDAO,
		},
		Items: nil,
	}
	itemModule.LoadFromDB()

	moduleContainer.Modules[0] = itemModule
}

func (moduleContainer *ModuleContainer) DestroyModules() {
	for _, modul := range moduleContainer.Modules {
		modul.Destroy()
	}
}
