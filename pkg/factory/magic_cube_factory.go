package factory

import (
	"gameSrv/pkg/actors"
)

// MagicCubeFactory creates magic cube actors
type MagicCubeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewMagicCubeFactory creates a new MagicCubeFactory
func NewMagicCubeFactory() *MagicCubeFactory {
	return &MagicCubeFactory{}
}

// CreateActor creates a magic cube actor
func (f *MagicCubeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *MagicCubeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *MagicCubeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *MagicCubeFactory) GetEntityType() int32 {
	return 21 // EActorType_MagicCube
}
