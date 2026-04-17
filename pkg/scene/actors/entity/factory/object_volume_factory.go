package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors/entity"
)

// ObjectVolumeFactory creates object volume actors
type ObjectVolumeFactory struct {
	EntityFactory[entity.SimpleActor]
}

// NewObjectVolumeFactory creates a new ObjectVolumeFactory
func NewObjectVolumeFactory() *ObjectVolumeFactory {
	return &ObjectVolumeFactory{}
}

// CreateActor creates an object volume actor
func (f *ObjectVolumeFactory) createEntity() *entity.SimpleActor {
	return entity.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *ObjectVolumeFactory) initFromConfig(actor *entity.SimpleActor, confId int32) {

}

func (f *ObjectVolumeFactory) initFromDO(actor *entity.SimpleActor) {
}

func (f *ObjectVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_ObjectVolume
}