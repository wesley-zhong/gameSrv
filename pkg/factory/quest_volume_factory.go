package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// QuestVolumeFactory creates quest volume actors
type QuestVolumeFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewQuestVolumeFactory creates a new QuestVolumeFactory
func NewQuestVolumeFactory() *QuestVolumeFactory {
	return &QuestVolumeFactory{}
}

// CreateActor creates a quest volume actor
func (f *QuestVolumeFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *QuestVolumeFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *QuestVolumeFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *QuestVolumeFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_QuestVolume
}
