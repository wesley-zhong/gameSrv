package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors"
)

// MonsterFactory creates monster actors
type MonsterFactory struct {
	EntityFactory[actors.Monster]
}

// NewMonsterFactory creates a new MonsterFactory
func NewMonsterFactory() *MonsterFactory {
	return &MonsterFactory{}
}

// CreateActor creates a monster actor
func (f *MonsterFactory) createEntity() *actors.Monster {
	return actors.NewMonster()
}

// InitFromConfigId initializes from config ID
func (f *MonsterFactory) initFromConfig(monster *actors.Monster, confId int32) {

}

func (f *MonsterFactory) initFromDO(monster *actors.Monster) {

}

func (f *MonsterFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Monster
}