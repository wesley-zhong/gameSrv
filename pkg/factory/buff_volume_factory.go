package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// BuffVolumeFactory creates buff volume actors
type BuffVolumeFactory struct {
	EntityFactory[actors.BuffVolumeActor]
}

// NewBuffVolumeFactory creates a new BuffVolumeFactory
func NewBuffVolumeFactory() *BuffVolumeFactory {
	return &BuffVolumeFactory{}
}

// CreateActor creates a buff volume actor
func (f *BuffVolumeFactory) createEntity() *actors.BuffVolumeActor {
	return actors.NewBuffVolumeActor()
}

// InitFromConfigId initializes from config ID
func (f *BuffVolumeFactory) initFromConfig(actor *actors.BuffVolumeActor, confId int32) {

}

func (f *BuffVolumeFactory) initFromDO(actor *actors.BuffVolumeActor) {

}

func (f *BuffVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_BuffVolume
}
