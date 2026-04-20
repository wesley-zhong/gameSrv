package player

import (
	"gameSrv/game/dal"
	"gameSrv/game/modules"
	"gameSrv/pkg/actor_module"
	"gameSrv/pkg/event"
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

type GamePlayer struct {
	Id         int64
	Sid        int64
	Modules    *modules.ModuleContainer
	asyncActor *actor_module.Actor
}

func NewGamePlayer(pid int64, sid int64) *GamePlayer {
	player := &GamePlayer{
		Id:         pid,
		Sid:        sid,
		Modules:    modules.NewModuleContainer(pid),
		asyncActor: actor_module.NewActor(pid),
	}
	modules.RegisterNewModule(&modules.RoleModule{}, player, dal.RoleDAO, func(module modules.IModule) {
		player.Modules.IModules[module.ModuleId()] = module
	})

	modules.RegisterNewModule(&modules.ItemModule{}, player, dal.ItemDAO, func(module modules.IModule) {
		player.Modules.IModules[module.ModuleId()] = module
	})
	modules.RegisterNewModule(&modules.QuestModule{}, player, dal.QuestDAO, func(module modules.IModule) {
		player.Modules.IModules[module.ModuleId()] = module
	})
	modules.RegisterNewModule(&modules.WorldModule{}, player, dal.WorldDAO, func(module modules.IModule) {
		player.Modules.IModules[module.ModuleId()] = module
	})
	return player
}

func (gp *GamePlayer) LoadDataFromDB() error {
	return gp.Modules.InitModules()
}

func (gp *GamePlayer) SaveData() {
	gp.asyncActor.Call(func() {
		gp.Modules.AsyncSave()
	})
}

func (gp *GamePlayer) OnLogin() {
	for _, module := range gp.Modules.IModules {
		module.OnLogin()
	}
}

func (gp *GamePlayer) OnDisconnect() {
	for _, module := range gp.Modules.IModules {
		module.OnDisconnect()
	}
}

func (gp *GamePlayer) DispatchEvent(ev event.Event) {
	event.Dispatcher.Dispatch(ev)
}

func GetModule[T any](gp *GamePlayer, moduleId modules.ModuleTypeId) *T {
	module, ok := gp.Modules.IModules[moduleId]
	if !ok {
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
