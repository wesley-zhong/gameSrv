package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// SummonSpawnerFactory creates summon spawner actors
type SummonSpawnerFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewSummonSpawnerFactory creates a new SummonSpawnerFactory
func NewSummonSpawnerFactory() *SummonSpawnerFactory {
	return &SummonSpawnerFactory{}
}

// CreateActor creates a summon spawner actor
func (f *SummonSpawnerFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *SummonSpawnerFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *SummonSpawnerFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *SummonSpawnerFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_SummonSpawner
}
