package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/scene/actors/entity"
)

// QuestVolumeFactory creates quest volume actors
type QuestVolumeFactory struct {
	EntityFactory[entity.SimpleActor]
}

// NewQuestVolumeFactory creates a new QuestVolumeFactory
func NewQuestVolumeFactory() *QuestVolumeFactory {
	return &QuestVolumeFactory{}
}

// CreateActor creates a quest volume actor
func (f *QuestVolumeFactory) createEntity() *entity.SimpleActor {
	return entity.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *QuestVolumeFactory) initFromConfig(actor *entity.SimpleActor, confId int32) {

}

func (f *QuestVolumeFactory) initFromDO(actor *entity.SimpleActor) {
}

func (f *QuestVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_QuestVolume
}