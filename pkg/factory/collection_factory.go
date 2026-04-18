package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// CollectionFactory creates collection actors
type CollectionFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewCollectionFactory creates a new CollectionFactory
func NewCollectionFactory() *CollectionFactory {
	return &CollectionFactory{}
}

// CreateActor creates a collection actor
func (f *CollectionFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *CollectionFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *CollectionFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *CollectionFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Collection
}
