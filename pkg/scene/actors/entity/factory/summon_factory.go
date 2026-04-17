package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// SummonFactory creates summon actors
type SummonFactory struct {
	EntityFactory[actors.SummonActor]
}

// NewSummonFactory creates a new SummonFactory
func NewSummonFactory() *SummonFactory {
	return &SummonFactory{}
}

// CreateActor creates a summon actor
func (f *SummonFactory) createEntity() *actors.SummonActor {
	return actors.NewSummonActor()
}

// InitFromConfigId initializes from config ID
func (f *SummonFactory) initFromConfig(actor *actors.SummonActor, confId int32) {

}

func (f *SummonFactory) initFromDO(actor *actors.SummonActor) {

}

func (f *SummonFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Summon
}