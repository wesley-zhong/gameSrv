package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// LadderFactory creates ladder actors
type LadderFactory struct {
	EntityFactory[actors.LadderActor]
}

// NewLadderFactory creates a new LadderFactory
func NewLadderFactory() *LadderFactory {
	return &LadderFactory{}
}

// CreateActor creates a ladder actor
func (f *LadderFactory) createEntity() *actors.LadderActor {
	return actors.NewLadderActor()
}

// InitFromConfigId initializes from config ID
func (f *LadderFactory) initFromConfig(actor *actors.LadderActor, confId int32) {

}

func (f *LadderFactory) initFromDO(actor *actors.LadderActor) {
}

func (f *LadderFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Ladder
}
