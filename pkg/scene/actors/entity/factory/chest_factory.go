package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// ChestFactory creates chest actors
type ChestFactory struct {
	EntityFactory[actors.Chest]
}

// NewChestFactory creates a new ChestFactory
func NewChestFactory() *ChestFactory {
	return &ChestFactory{}
}

// CreateActor creates a chest actor
func (f *ChestFactory) createEntity() *actors.Chest {
	return actors.NewChest()
}

// InitFromConfigId initializes from config ID
func (f *ChestFactory) initFromConfig(chest *actors.Chest, confId int32) {

}

func (f *ChestFactory) initFromDO(chest *actors.Chest) {

}

func (f *ChestFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Chest
}