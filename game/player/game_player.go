package player

import (
	"gameSrv/game/dal"
	"gameSrv/game/modules"
	"gameSrv/pkg/actor_module"
	"gameSrv/pkg/event"
	"gameSrv/pkg/math"
	"gameSrv/pkg/orm"
	"gameSrv/pkg/scene"
)

type GamePlayer struct {
	Id              int64
	Sid             int64
	ModuleContainer *modules.ModuleContainer
	asyncActor      *actor_module.Actor
}

func NewGamePlayer(pid int64, sid int64) *GamePlayer {
	player := &GamePlayer{
		Id:              pid,
		Sid:             sid,
		ModuleContainer: modules.NewModuleContainer(pid),
		asyncActor:      actor_module.NewActor(pid),
	}
	registerPlayerModules(&modules.RoleModule{}, player, dal.RoleDAO)
	registerPlayerModules(&modules.ItemModule{}, player, dal.ItemDAO)
	registerPlayerModules(&modules.QuestModule{}, player, dal.QuestDAO)
	registerPlayerModules(&modules.WorldModule{}, player, dal.WorldDAO)
	return player
}

func registerPlayerModules[DOType any](aresModule modules.IGameModule[DOType], gamePlayer *GamePlayer, dao *orm.MongodbDAO[DOType]) {
	modules.RegisterNewModule(aresModule, gamePlayer, dao, func(module modules.IModule) {
		gamePlayer.ModuleContainer.IModules[module.ModuleId()] = module
	})
}

func (gp *GamePlayer) LoadDataFromDB() error {
	return gp.ModuleContainer.InitModules()
}

func (gp *GamePlayer) SaveDataOnPlayerRouting() {
	gp.asyncActor.Call(func() {
		gp.ModuleContainer.SaveDirtyDataToDB()
	})
}

func (gp *GamePlayer) OnLogin() {
	for _, module := range gp.ModuleContainer.IModules {
		module.OnLogin()
	}
}

func (gp *GamePlayer) OnDisconnect() {
	for _, module := range gp.ModuleContainer.IModules {
		module.OnDisconnect()
	}
}

func (gp *GamePlayer) DispatchEvent(ev event.Event) {
	event.Dispatcher.Dispatch(ev)
}

func GetModule[T any](gp *GamePlayer, moduleId modules.ModuleTypeId) *T {
	module := gp.ModuleContainer.IModules[moduleId]
	if module == nil {
		return nil
	}

	if target, ok := any(module).(*T); ok {
		return target
	}

	return nil
}
func (gp *GamePlayer) GetUid() int64 {
	return gp.Id
}

func (gp *GamePlayer) GetAvatarTeam() scene.IEntity {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetCurAvatarConfId() int64 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) InPrivatePhasing() bool {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetLevel() int32 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetLifeState() int32 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetExp() int64 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetExceedID() int64 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) GetCampType() int32 {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) SetCachePosRot(pos, rot *math.Vector3) {
	//TODO implement me
	panic("implement me")
}

func (gp *GamePlayer) OnTeamAvatarDead(actor scene.IEntity) {
	//TODO implement me
	panic("implement me")
}
