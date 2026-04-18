package factory

import (
	"gameSrv/pkg/actors"
)

// MonsterSquadFactory creates monster squad actors
type MonsterSquadFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewMonsterSquadFactory creates a new MonsterSquadFactory
func NewMonsterSquadFactory() *MonsterSquadFactory {
	return &MonsterSquadFactory{}
}

// CreateActor creates a monster squad actor
func (f *MonsterSquadFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *MonsterSquadFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *MonsterSquadFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *MonsterSquadFactory) GetEntityType() int32 {
	return 22 // EActorType_MonsterSquad
}
