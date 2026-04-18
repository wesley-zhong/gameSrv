package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// LandmarkFactory creates landmark actors
type LandmarkFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewLandmarkFactory creates a new LandmarkFactory
func NewLandmarkFactory() *LandmarkFactory {
	return &LandmarkFactory{}
}

// CreateActor creates a landmark actor
func (f *LandmarkFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *LandmarkFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *LandmarkFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *LandmarkFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Landmark
}
