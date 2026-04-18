package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// CommonSwitchFactory creates common switch actors
type CommonSwitchFactory struct {
	EntityFactory[actors.CommonSwitchActor]
}

// NewCommonSwitchFactory creates a new CommonSwitchFactory
func NewCommonSwitchFactory() *CommonSwitchFactory {
	return &CommonSwitchFactory{}
}

// CreateActor creates a common switch actor
func (f *CommonSwitchFactory) createEntity() *actors.CommonSwitchActor {
	return actors.NewCommonSwitchActor()
}

// InitFromConfigId initializes from config ID
func (f *CommonSwitchFactory) initFromConfig(actor *actors.CommonSwitchActor, confId int32) {

}

func (f *CommonSwitchFactory) initFromDO(actor *actors.CommonSwitchActor) {
}

func (f *CommonSwitchFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_CommonSwitch
}
