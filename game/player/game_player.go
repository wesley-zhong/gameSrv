package player

import (
	"gameSrv/game/dal"
	"gameSrv/game/modules"
	"gameSrv/pkg/actor"
	"gameSrv/pkg/event"
)

type GamePlayer struct {
	Id         int64
	Sid        int64
	Modules    *modules.ModuleContainer
	asyncActor *actor.Actor
}

func NewGamePlayer(pid int64, sid int64) *GamePlayer {
	player := &GamePlayer{
		Id:         pid,
		Sid:        sid,
		Modules:    modules.NewModuleContainer(pid),
		asyncActor: actor.NewActor(pid),
	}
	modules.RegisterNewModule(&modules.RoleModule{}, player.Modules, dal.RoleDAO)

	modules.RegisterNewModule(&modules.ItemModule{}, player.Modules, dal.ItemDAO)
	modules.RegisterNewModule(&modules.QuestModule{}, player.Modules, dal.QuestDAO)
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
