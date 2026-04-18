package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// PickupActorFactory creates pickup actors
type PickupActorFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewPickupActorFactory creates a new PickupActorFactory
func NewPickupActorFactory() *PickupActorFactory {
	return &PickupActorFactory{}
}

// CreateActor creates a pickup actor
func (f *PickupActorFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *PickupActorFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *PickupActorFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *PickupActorFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Pickup
}
