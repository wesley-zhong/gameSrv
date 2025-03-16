package modules

import (
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

type IGmePlayer interface {
	GetPlayerId() int64
	SendMsg(msgId protoGen.ProtoCode, msg proto.Message)
}

type ModuleContainer struct {
	GamePlayer IGmePlayer
	Modules    []AresModule
}

//func NewModuleContainer(player IGmePlayer) *ModuleContainer {
//	moduleContainer := &ModuleContainer{}
//
//	return moduleContainer
//}

func (moduleContainer *ModuleContainer) InitModules() {
	roleModule := &RoleModule{}
	roleModule.InitModule(ROLE_MODULE, moduleContainer.GamePlayer)
	moduleContainer.Modules = append(moduleContainer.Modules, roleModule)

	itemModule := &ItemModule{}
	itemModule.InitModule(ITEM_MODULE, moduleContainer.GamePlayer)
	moduleContainer.Modules = append(moduleContainer.Modules, itemModule)
}

func (moduleContainer *ModuleContainer) LoadFromDB() {
	for _, module := range moduleContainer.Modules {
		module.LoadFromDB()
	}
}

func (moduleContainer *ModuleContainer) SaveToDB() {
	for _, module := range moduleContainer.Modules {
		module.AsynSaveDB()
	}
}
func (moduleContainer *ModuleContainer) FindModule(moduleId ModuleId) AresModule {
	return moduleContainer.Modules[moduleId]
}

func (moduleContainer *ModuleContainer) DestroyModules() {
	for _, modul := range moduleContainer.Modules {
		modul.Destroy()
	}
}
