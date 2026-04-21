package modules

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/orm"
	"gameSrv/pkg/scene"
)

type ModuleTypeId int32

const (
	ROLE_MODULE  ModuleTypeId = 1
	ITEM_MODULE  ModuleTypeId = 2
	QUEUE_MODULE ModuleTypeId = 3
	WORLE_MODULE ModuleTypeId = 4
	MAX_MODULE   ModuleTypeId = 5
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
	Uid() int64
	SetGamePlayer(gp scene.IGamePlayer)
}

// GameModule =================================  GameModule implement===========================
type GameModule[DOType any] struct {
	GamePlayer scene.IGamePlayer
	Dao        *orm.MongodbDAO[DOType]
	DataDO     *DOType
	isDirty    bool
}

func (gm *GameModule[DOType]) InitFromDB() error {
	moduleDO, err := gm.Dao.FindOneById(gm.GetPid())
	if err != nil {
		return err
	}
	gm.DataDO = moduleDO
	return nil
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
	return gm.Dao.Save(gm.GetPid(), gm.ToDO())
}
func (gm *GameModule[DOType]) GetPid() int64 {
	return gm.GamePlayer.GetUid()
}

func (gm *GameModule[DOType]) OnLogin() {
	log.Infof("game module %d login", gm.GetPid())
}
func (gm *GameModule[DOType]) OnDisconnect() {
	log.Infof("game module %d disconnect", gm.GetPid())
}

func (gm *GameModule[DOType]) SetGamePlayer(gp scene.IGamePlayer) {
	gm.GamePlayer = gp
}

func (gm *GameModule[DOType]) Uid() int64 {
	return gm.GamePlayer.GetUid()
}

// ModuleContainer ==========================
type ModuleContainer struct {
	Pid      int64
	IModules []IModule
}

func NewModuleContainer(pid int64) *ModuleContainer {
	moduleContainer := &ModuleContainer{
		Pid:      pid,
		IModules: make([]IModule, MAX_MODULE),
	}
	return moduleContainer
}

func RegisterNewModule[DOType any](aresModule IGameModule[DOType], gamePlayer scene.IGamePlayer, dao *orm.MongodbDAO[DOType], onCreated func(module IModule)) {
	// 必须先转为 any，再断言为接口
	if am, ok := any(aresModule).(IGameModule[DOType]); ok {
		am.SetDAO(dao)
		am.SetGamePlayer(gamePlayer)
		onCreated(am)
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
