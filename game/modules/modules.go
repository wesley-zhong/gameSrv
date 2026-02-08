package modules

import (
	"gameSrv/game/dal"
	"gameSrv/pkg/log"
)

type AresModule[DOType any] interface {
	FromDO(do *DOType) error
	ToDO() *DOType
}

type ModuleContainer struct {
	pid        int64
	ItemModule ItemModule
}

func FromDO[DOTYpe any](doObj *DOTYpe, modlue AresModule[DOTYpe]) {
	modlue.FromDO(doObj)
}
func ToDO[DOType any](module AresModule[DOType]) *DOType {
	return module.ToDO()
}

func NewModuleContainer(pid int64) *ModuleContainer {
	moduleContainer := &ModuleContainer{}
	moduleContainer.pid = pid
	return moduleContainer
}

func (moduleContainer *ModuleContainer) initModules() {
	itemDO := dal.ItemDAO.FindOneById(moduleContainer.pid)
	itemModue := &ItemModule{}
	FromDO(itemDO, itemModue)
	newItemDo := ToDO(itemModue)
	log.Infof("=-----%d", newItemDo.Pid)
}
