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
	itemModule := &ItemModule{}
	itemModule.InitModule(ITEM_MODULE, moduleContainer.GamePlayer)
	itemModule.LoadFromDB()
	moduleContainer.Modules[0] = itemModule

	roleModule := &RoleModule{}
	roleModule.InitModule(ROLE_MODULE, moduleContainer.GamePlayer)
	roleModule.LoadFromDB()
	moduleContainer.Modules = append(moduleContainer.Modules, roleModule)
}
func (moduleContainer *ModuleContainer) FindModule(moduleId ModuleId) AresModule {
	return moduleContainer.Modules[moduleId]
}

func (moduleContainer *ModuleContainer) DestroyModules() {
	for _, modul := range moduleContainer.Modules {
		modul.Destroy()
	}
}
