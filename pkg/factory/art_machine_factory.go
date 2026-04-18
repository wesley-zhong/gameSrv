package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// ArtMachineFactory creates art machine actors
type ArtMachineFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewArtMachineFactory creates a new ArtMachineFactory
func NewArtMachineFactory() *ArtMachineFactory {
	return &ArtMachineFactory{}
}

// CreateActor creates an art machine actor
func (f *ArtMachineFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *ArtMachineFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *ArtMachineFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *ArtMachineFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_ArtMachine
}
