package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// SummonMachineFactory creates summon machine actors
type SummonMachineFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewSummonMachineFactory creates a new SummonMachineFactory
func NewSummonMachineFactory() *SummonMachineFactory {
	return &SummonMachineFactory{}
}

// CreateActor creates a summon machine actor
func (f *SummonMachineFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *SummonMachineFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *SummonMachineFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *SummonMachineFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_SummonMachine
}
