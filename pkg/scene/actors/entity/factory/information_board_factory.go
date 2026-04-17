package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors/entity"
)

// InformationBoardFactory creates information board actors
type InformationBoardFactory struct {
	EntityFactory[entity.SimpleActor]
}

// NewInformationBoardFactory creates a new InformationBoardFactory
func NewInformationBoardFactory() *InformationBoardFactory {
	return &InformationBoardFactory{}
}

// CreateActor creates an information board actor
func (f *InformationBoardFactory) createEntity() *entity.SimpleActor {
	return entity.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *InformationBoardFactory) initFromConfig(actor *entity.SimpleActor, confId int32) {

}

func (f *InformationBoardFactory) initFromDO(actor *entity.SimpleActor) {
}

func (f *InformationBoardFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_InformationBoard
}