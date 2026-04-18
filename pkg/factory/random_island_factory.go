package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// RandomIslandFactory creates random island actors
type RandomIslandFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewRandomIslandFactory creates a new RandomIslandFactory
func NewRandomIslandFactory() *RandomIslandFactory {
	return &RandomIslandFactory{}
}

// CreateActor creates a random island actor
func (f *RandomIslandFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *RandomIslandFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *RandomIslandFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *RandomIslandFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_RandomIsland
}
