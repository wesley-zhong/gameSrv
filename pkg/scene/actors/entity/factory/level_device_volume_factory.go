package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors/entity"
)

// LevelDeviceVolumeFactory creates level device volume actors
type LevelDeviceVolumeFactory struct {
	EntityFactory[entity.SimpleActor]
}

// NewLevelDeviceVolumeFactory creates a new LevelDeviceVolumeFactory
func NewLevelDeviceVolumeFactory() *LevelDeviceVolumeFactory {
	return &LevelDeviceVolumeFactory{}
}

// CreateActor creates a level device volume actor
func (f *LevelDeviceVolumeFactory) createEntity() *entity.SimpleActor {
	return entity.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *LevelDeviceVolumeFactory) initFromConfig(actor *entity.SimpleActor, confId int32) {

}

func (f *LevelDeviceVolumeFactory) initFromDO(actor *entity.SimpleActor) {
}

func (f *LevelDeviceVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_LevelDeviceVolume
}