package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// KillZVolumeFactory creates kill zone volume actors
type KillZVolumeFactory struct {
	EntityFactory[actors.KillZVolumeActor]
}

// NewKillZVolumeFactory creates a new KillZVolumeFactory
func NewKillZVolumeFactory() *KillZVolumeFactory {
	return &KillZVolumeFactory{}
}

// CreateActor creates a kill zone volume actor
func (f *KillZVolumeFactory) createEntity() *actors.KillZVolumeActor {
	return actors.NewKillZVolumeActor()
}

// InitFromConfigId initializes from config ID
func (f *KillZVolumeFactory) initFromConfig(actor *actors.KillZVolumeActor, confId int32) {

}

func (f *KillZVolumeFactory) initFromDO(actor *actors.KillZVolumeActor) {

}

func (f *KillZVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_KillZVolume
}
