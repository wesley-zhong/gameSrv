package modules

import (
	"gameSrv/pkg/actors"
	"gameSrv/pkg/log"
)

type AvatarDO struct {
	cnfId int32
}
type RoleDO struct {
	Id   int64
	Name string
}

type RoleModule struct {
	GameModule[RoleDO]
	HeroAvatars map[int64]*actors.HeroAvatarActor
}

func (roleModule *RoleModule) ModuleId() ModuleTypeId {
	return ROLE_MODULE
}

func (roleModule *RoleModule) OnDataLoaded() error {
	return nil
}

func (roleModule *RoleModule) OnLogin() {
	log.Infof("itemModule OnLogin")
	if roleModule.DataDO == nil {
		roleModule.DataDO = &RoleDO{
			Id:   roleModule.GamePlayer.GetUid(),
			Name: "name",
		}

		roleModule.HeroAvatars = make(map[int64]*actors.HeroAvatarActor)
		roleModule.MarkDirty()
	}
	return
}

func (roleModule *RoleModule) AddAvatar(avatar *actors.HeroAvatarActor) {
	roleModule.HeroAvatars[avatar.GetConfigId()] = avatar

}
