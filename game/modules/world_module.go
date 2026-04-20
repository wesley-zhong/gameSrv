package modules

import (
	"gameSrv/pkg/log"
	"gameSrv/pkg/scene"
)

type WorldDO struct {
	Id            int64
	CurSceneCnfId int64
}

type WorldModule struct {
	GameModule[WorldDO]
	CurScene scene.IScene
}

func (worldModule *WorldModule) ModuleId() ModuleTypeId {
	return WORLE_MODULE
}

func (worldModule *WorldModule) OnDataLoaded() error {
	return nil
}

func (worldModule *WorldModule) OnLogin() {
	log.Infof("worldModule OnLogin")
	if worldModule.DataDO == nil {
		worldModule.DataDO = &WorldDO{
			Id:            worldModule.Uid(),
			CurSceneCnfId: 0, //set cur scene configId from cfg
		}
		worldModule.MarkDirty()
	}
	return
}

func (worldModule *WorldModule) OnLogout() {

}
