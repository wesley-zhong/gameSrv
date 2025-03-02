package modules

import "gameSrv/game/player"

type ModuleContainer struct {
	GamePlayer *player.GamePlayer
	Modules    []AresModule
}

func NewModuleContainer(player *player.GamePlayer) *ModuleContainer {
	moduleContainer := &ModuleContainer{}
	return moduleContainer
}

func (moduleContainer *ModuleContainer) InitModules() {
	itemModule := &ItemModule{}
	itemModule.InitModule(ITEM_MODULE, moduleContainer.GamePlayer)
	itemModule.LoadFromDB()

	moduleContainer.Modules[0] = itemModule
}
func (moduleContainer *ModuleContainer) FindModule(moduleId ModuleId) AresModule {
	return moduleContainer.Modules[moduleId]
}

func (moduleContainer *ModuleContainer) DestroyModules() {
	for _, modul := range moduleContainer.Modules {
		modul.Destroy()
	}
}
