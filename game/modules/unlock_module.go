package modules

import (
	"gameSrv/pkg/log"
)

type UnlockDO struct {
	fntId []int32 // function id
}

type UnlockModule struct {
	GameModule[UnlockDO]
	fntIds []int32
}

func (unlock *UnlockModule) ModuleId() ModuleTypeId {
	return UNLOCK_MODULE
}

func (unlock *UnlockModule) OnDataLoaded() error {
	return nil
}

func (unlock *UnlockModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if unlock.DataDO == nil {
		unlock.DataDO = &UnlockDO{}
	}

	//unlock.GamePlayer.SendMsg()
}

func (unlock *UnlockModule) AddFntId(fntId int32) {
	unlock.fntIds = append(unlock.fntIds, fntId)
	unlock.DataDO.fntId = unlock.fntIds
	unlock.MarkDirty()
}
