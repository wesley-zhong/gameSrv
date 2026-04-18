package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// WharfFactory creates wharf actors
type WharfFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewWharfFactory creates a new WharfFactory
func NewWharfFactory() *WharfFactory {
	return &WharfFactory{}
}

// CreateActor creates a wharf actor
func (f *WharfFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *WharfFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *WharfFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *WharfFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Wharf
}
