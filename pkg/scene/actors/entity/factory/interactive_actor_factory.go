package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// InteractiveActorFactory creates interactive actors
type InteractiveActorFactory struct {
	EntityFactory[actors.InteractiveActor]
}

// NewInteractiveActorFactory creates a new InteractiveActorFactory
func NewInteractiveActorFactory() *InteractiveActorFactory {
	return &InteractiveActorFactory{}
}

// CreateActor creates an interactive actor
func (f *InteractiveActorFactory) createEntity() *actors.InteractiveActor {
	return actors.NewInteractiveActor()
}

// InitFromConfigId initializes from config ID
func (f *InteractiveActorFactory) initFromConfig(actor *actors.InteractiveActor, confId int32) {

}

func (f *InteractiveActorFactory) initFromDO(actor *actors.InteractiveActor) {

}

func (f *InteractiveActorFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_InteractiveActor
}