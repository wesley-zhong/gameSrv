package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// DropableDestructionFactory creates dropable destruction actors
type DropableDestructionFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewDropableDestructionFactory creates a new DropableDestructionFactory
func NewDropableDestructionFactory() *DropableDestructionFactory {
	return &DropableDestructionFactory{}
}

// CreateActor creates a dropable destruction actor
func (f *DropableDestructionFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *DropableDestructionFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *DropableDestructionFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *DropableDestructionFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_DropableDestruction
}
