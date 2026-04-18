package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// ObjectVolumeFactory creates object volume actors
type ObjectVolumeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewObjectVolumeFactory creates a new ObjectVolumeFactory
func NewObjectVolumeFactory() *ObjectVolumeFactory {
	return &ObjectVolumeFactory{}
}

// CreateActor creates an object volume actor
func (f *ObjectVolumeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *ObjectVolumeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *ObjectVolumeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *ObjectVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_ObjectVolume
}
