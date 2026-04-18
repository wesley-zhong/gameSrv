package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// InvisibleWallFactory creates invisible wall actors
type InvisibleWallFactory struct {
	EntityFactory[actors.InvisibleWall]
}

// NewInvisibleWallFactory creates a new InvisibleWallFactory
func NewInvisibleWallFactory() *InvisibleWallFactory {
	return &InvisibleWallFactory{}
}

// CreateActor creates an invisible wall actor
func (f *InvisibleWallFactory) createEntity() *actors.InvisibleWall {
	return actors.NewInvisibleWall()
}

// InitFromConfigId initializes from config ID
func (f *InvisibleWallFactory) initFromConfig(actor *actors.InvisibleWall, confId int32) {

}

func (f *InvisibleWallFactory) initFromDO(actor *actors.InvisibleWall) {

}

func (f *InvisibleWallFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_InvisibleWall
}
