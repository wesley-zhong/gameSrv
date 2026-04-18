package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// AIVolumeFactory creates AI volume actors
type AIVolumeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewAIVolumeFactory creates a new AIVolumeFactory
func NewAIVolumeFactory() *AIVolumeFactory {
	return &AIVolumeFactory{}
}

// CreateActor creates an AI volume actor
func (f *AIVolumeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *AIVolumeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *AIVolumeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *AIVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_AIVolume
}
