package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// GameFlowActorFactory creates game flow actors
type GameFlowActorFactory struct {
	EntityFactory[actors.GameFlowActor]
}

// NewGameFlowActorFactory creates a new GameFlowActorFactory
func NewGameFlowActorFactory() *GameFlowActorFactory {
	return &GameFlowActorFactory{}
}

// CreateActor creates a game flow actor
func (f *GameFlowActorFactory) createEntity() *actors.GameFlowActor {
	return actors.NewGameFlowActor()
}

// InitFromConfigId initializes from config ID
func (f *GameFlowActorFactory) initFromConfig(actor *actors.GameFlowActor, confId int32) {

}

func (f *GameFlowActorFactory) initFromDO(actor *actors.GameFlowActor) {

}

func (f *GameFlowActorFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_GameFlow
}