package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// RestPointFactory creates rest point actors
type RestPointFactory struct {
	EntityFactory[actors.RestPoint]
}

// NewRestPointFactory creates a new RestPointFactory
func NewRestPointFactory() *RestPointFactory {
	return &RestPointFactory{}
}

// CreateActor creates a rest point actor
func (f *RestPointFactory) createEntity() *actors.RestPoint {
	return actors.NewRestPoint()
}

// InitFromConfigId initializes from config ID
func (f *RestPointFactory) initFromConfig(actor *actors.RestPoint, confId int32) {

}

func (f *RestPointFactory) initFromDO(actor *actors.RestPoint) {
}

func (f *RestPointFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_RestPoint
}