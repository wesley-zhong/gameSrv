package modules

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/orm"
)

type ModuleTypeId int32

const (
	ROLE_MODULE  ModuleTypeId = 1
	ITEM_MODULE  ModuleTypeId = 2
	QUEUE_MODULE ModuleTypeId = 3
)

type IModule interface {
	ModuleId() ModuleTypeId
	InitFromDB() error
	Destroy()
	IsDirty() bool
	SaveToDB() error
	OnDataLoaded() error
	OnLogin()
	OnDisconnect()
}

type IGameModule[DOType any] interface {
	IModule
	SetDAO(dao *orm.MongodbDAO[DOType])
	SetPid(pid int64)
}

// GameModule =================================  GameModule implement===========================
type GameModule[DOType any] struct {
	Pid     int64
	Dao     *orm.MongodbDAO[DOType]
	DataDO  *DOType
	isDirty bool
}

func (gm *GameModule[DOType]) InitFromDB() error {
	moduleDO, err := gm.Dao.FindOneById(gm.Pid)
	if err != nil {
		return err
	}
	gm.DataDO = moduleDO
	return nil
}

func (gm *GameModule[DOType]) SetPid(pid int64) {
	gm.Pid = pid
}

func (gm *GameModule[DOType]) SetDAO(dao *orm.MongodbDAO[DOType]) {
	gm.Dao = dao
}

func (gm *GameModule[DOType]) ToDO() *DOType {
	return gm.DataDO
}

func (gm *GameModule[DOType]) MarkDirty() {
	gm.isDirty = true
}

func (gm *GameModule[DOType]) IsDirty() bool {
	return gm.isDirty
}

func (gm *GameModule[DOType]) Destroy() {
}

func (gm *GameModule[DOType]) SaveToDB() error {
	gm.isDirty = false
	return gm.Dao.Save(gm.Pid, gm.ToDO())
}

func (gm *GameModule[DOType]) OnLogin() {
	log.Infof("game module %d login", gm.Pid)
}
func (gm *GameModule[DOType]) OnDisconnect() {
	log.Infof("game module %d disconnect", gm.Pid)
}

// ModuleContainer ==========================
type ModuleContainer struct {
	Pid      int64
	IModules map[ModuleTypeId]IModule
}

func NewModuleContainer(pid int64) *ModuleContainer {
	moduleContainer := &ModuleContainer{
		Pid:      pid,
		IModules: make(map[ModuleTypeId]IModule, 1024),
	}
	return moduleContainer
}

func RegisterNewModule[DOType any](aresModule IGameModule[DOType], container *ModuleContainer, dao *orm.MongodbDAO[DOType]) {
	// 必须先转为 any，再断言为接口
	if am, ok := any(aresModule).(IGameModule[DOType]); ok {
		am.SetDAO(dao)
		am.SetPid(container.Pid)
		container.IModules[am.ModuleId()] = am
	} else {
		panic("ModuleType does not implement IGameModule interface")
	}
}

func (mc *ModuleContainer) InitModules() error {
	for _, module := range mc.IModules {
		module.InitFromDB()
	}
	// only for just now it maybe called on another thread
	for _, module := range mc.IModules {
		module.OnDataLoaded()
	}
	return nil
}

func (mc *ModuleContainer) Destroy() {
	for _, module := range mc.IModules {
		module.Destroy()
	}
}

func (mc *ModuleContainer) AsyncSave() {
	for _, module := range mc.IModules {
		if module.IsDirty() {
			module.SaveToDB()
		}
	}
}
