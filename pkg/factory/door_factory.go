package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// DoorFactory creates door actors
type DoorFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewDoorFactory creates a new DoorFactory
func NewDoorFactory() *DoorFactory {
	return &DoorFactory{}
}

// CreateActor creates a door actor
func (f *DoorFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *DoorFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *DoorFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *DoorFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Door
}
