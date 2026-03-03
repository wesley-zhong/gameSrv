package player

import (
	"gameSrv/game/dal"
	"gameSrv/game/modules"
)

type GamePlayer struct {
	Id      int64
	Sid     int64
	Modules *modules.ModuleContainer
}

func NewGamePlayer(pid int64, sid int64) *GamePlayer {
	player := &GamePlayer{
		Id:      pid,
		Sid:     sid,
		Modules: modules.NewModuleContainer(pid),
	}
	modules.RegisterNewModule(&modules.RoleModule{}, player.Modules, dal.RoleDAO, pid)

	modules.RegisterNewModule(&modules.ItemModule{}, player.Modules, dal.ItemDAO, pid)
	return player
}

func (gp *GamePlayer) LoadData() {
	gp.Modules.InitModules()
}
