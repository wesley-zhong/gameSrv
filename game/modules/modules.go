package modules

import (
	"gameSrv/pkg/orm"
)

type ModuleTypeId int32

const (
	ROLE_MODULE ModuleTypeId = 1
	ITEM_MODULE ModuleTypeId = 2
)

type IModule interface {
	ModuleId() ModuleTypeId
	initFromDB() bool
	ToDO() any
	Destroy()
}

type IGameModule[DOType any] interface {
	IModule
	SetDAO(dao *orm.MongodbDAO[DOType])
	FromDO(do *DOType) error
}

// GameModule =================================  GameModule implement===========================
type GameModule[DOType any] struct {
	Pid    int64
	Dao    *orm.MongodbDAO[DOType]
	DataDO *DOType
}

func (gm *GameModule[DOType]) initFromDB() bool {
	moduleDO := gm.Dao.FindOneById(gm.Pid)
	gm.FromDO(moduleDO)
	return true
}
func (gm *GameModule[DOType]) SetDAO(dao *orm.MongodbDAO[DOType]) {
	gm.Dao = dao
}
func (gm *GameModule[DOType]) ModuleId() ModuleTypeId {
	return 0
}

func (gm *GameModule[DOType]) FromDO(dataDO *DOType) error {
	gm.DataDO = dataDO
	return nil
}

func (gm *GameModule[DOType]) ToDO() any {
	return gm.DataDO
}

func (gm *GameModule[DOType]) Destroy() {
}

// ModuleContainer ==========================
type ModuleContainer struct {
	Pid      int64
	IModules map[ModuleTypeId]IModule
}

func NewModuleContainer(pid int64) *ModuleContainer {
	moduleContainer := &ModuleContainer{}
	moduleContainer.Pid = pid
	return moduleContainer
}

func RegisterNewModule[DOType any](aresModule IGameModule[DOType], container *ModuleContainer, dao *orm.MongodbDAO[DOType], id int64) {
	// 必须先转为 any，再断言为接口
	if am, ok := any(aresModule).(IGameModule[DOType]); ok {
		am.SetDAO(dao)
		container.IModules[am.ModuleId()] = am
	} else {
		panic("ModuleType does not implement IGameModule interface")
	}
}

func (mc *ModuleContainer) InitModules() error {
	for _, itemModule := range mc.IModules {
		itemModule.initFromDB()
	}
	return nil
}

func (mc *ModuleContainer) Destroy() {
	for _, itemModule := range mc.IModules {
		itemModule.Destroy()
	}
}
