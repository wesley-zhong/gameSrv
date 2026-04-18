package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// GadgetFactory creates gadget actors
type GadgetFactory struct {
	EntityFactory[actors.Gadget]
}

// NewGadgetFactory creates a new GadgetFactory
func NewGadgetFactory() *GadgetFactory {
	return &GadgetFactory{}
}

// CreateActor creates a gadget actor
func (f *GadgetFactory) createEntity() *actors.Gadget {
	return actors.NewGadget()
}

// InitFromConfigId initializes from config ID
func (f *GadgetFactory) initFromConfig(gadget *actors.Gadget, confId int32) {

}

func (f *GadgetFactory) initFromDO(gadget *actors.Gadget) {

}

func (f *GadgetFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Gadget
}
