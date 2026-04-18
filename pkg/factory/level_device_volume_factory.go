package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// LevelDeviceVolumeFactory creates level device volume actors
type LevelDeviceVolumeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewLevelDeviceVolumeFactory creates a new LevelDeviceVolumeFactory
func NewLevelDeviceVolumeFactory() *LevelDeviceVolumeFactory {
	return &LevelDeviceVolumeFactory{}
}

// CreateActor creates a level device volume actor
func (f *LevelDeviceVolumeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *LevelDeviceVolumeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *LevelDeviceVolumeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *LevelDeviceVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_LevelDeviceVolume
}
