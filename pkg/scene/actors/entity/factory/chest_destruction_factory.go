package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// ChestDestructionFactory creates chest destruction actors
type ChestDestructionFactory struct {
	EntityFactory[actors.ChestDestructionActor]
}

// NewChestDestructionFactory creates a new ChestDestructionFactory
func NewChestDestructionFactory() *ChestDestructionFactory {
	return &ChestDestructionFactory{}
}

// CreateActor creates a chest destruction actor
func (f *ChestDestructionFactory) createEntity() *actors.ChestDestructionActor {
	return actors.NewChestDestructionActor()
}

// InitFromConfigId initializes from config ID
func (f *ChestDestructionFactory) initFromConfig(actor *actors.ChestDestructionActor, confId int32) {

}

func (f *ChestDestructionFactory) initFromDO(actor *actors.ChestDestructionActor) {

}

func (f *ChestDestructionFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_ChestDestruction
}