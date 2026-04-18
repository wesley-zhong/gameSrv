package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// ArtVolumeFactory creates art volume actors
type ArtVolumeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewArtVolumeFactory creates a new ArtVolumeFactory
func NewArtVolumeFactory() *ArtVolumeFactory {
	return &ArtVolumeFactory{}
}

// CreateActor creates an art volume actor
func (f *ArtVolumeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *ArtVolumeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *ArtVolumeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *ArtVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_ArtVolume
}
