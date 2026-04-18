package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// MovingPlatformFactory creates moving platform actors
type MovingPlatformFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewMovingPlatformFactory creates a new MovingPlatformFactory
func NewMovingPlatformFactory() *MovingPlatformFactory {
	return &MovingPlatformFactory{}
}

// CreateActor creates a moving platform actor
func (f *MovingPlatformFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *MovingPlatformFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *MovingPlatformFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *MovingPlatformFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_MovingPlatform
}
