package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// InformationBoardFactory creates information board actors
type InformationBoardFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewInformationBoardFactory creates a new InformationBoardFactory
func NewInformationBoardFactory() *InformationBoardFactory {
	return &InformationBoardFactory{}
}

// CreateActor creates an information board actor
func (f *InformationBoardFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *InformationBoardFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *InformationBoardFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *InformationBoardFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_InformationBoard
}
