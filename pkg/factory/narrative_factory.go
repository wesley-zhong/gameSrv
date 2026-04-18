package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// NarrativeFactory creates narrative actors
type NarrativeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewNarrativeFactory creates a new NarrativeFactory
func NewNarrativeFactory() *NarrativeFactory {
	return &NarrativeFactory{}
}

// CreateActor creates a narrative actor
func (f *NarrativeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *NarrativeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *NarrativeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *NarrativeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Narrative
}
