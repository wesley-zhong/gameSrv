package modules

import (
	"errors"
	"gameSrv/game/dal"
	"gameSrv/game/do"
	"gameSrv/pkg/log"
	"gameSrv/pkg/orm"
)

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
	// do some error log
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
		AresModuleBase: AresModuleBase[do.ItemDO]{
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
