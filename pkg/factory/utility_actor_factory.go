package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// UtilityActorFactory creates utility actors
type UtilityActorFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewUtilityActorFactory creates a new UtilityActorFactory
func NewUtilityActorFactory() *UtilityActorFactory {
	return &UtilityActorFactory{}
}

// CreateActor creates a utility actor
func (f *UtilityActorFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *UtilityActorFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *UtilityActorFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *UtilityActorFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Utility
}
