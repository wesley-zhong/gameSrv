package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors/entity"
)

// DropableDestructionFactory creates dropable destruction actors
type DropableDestructionFactory struct {
	EntityFactory[entity.SimpleActor]
}

// NewDropableDestructionFactory creates a new DropableDestructionFactory
func NewDropableDestructionFactory() *DropableDestructionFactory {
	return &DropableDestructionFactory{}
}

// CreateActor creates a dropable destruction actor
func (f *DropableDestructionFactory) createEntity() *entity.SimpleActor {
	return entity.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *DropableDestructionFactory) initFromConfig(actor *entity.SimpleActor, confId int32) {

}

func (f *DropableDestructionFactory) initFromDO(actor *entity.SimpleActor) {
}

func (f *DropableDestructionFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_DropableDestruction
}