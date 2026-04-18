package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// TeleporterSceneFactory creates teleporter scene actors
type TeleporterSceneFactory struct {
	EntityFactory[actors.SimpleActor]
}

// NewTeleporterSceneFactory creates a new TeleporterSceneFactory
func NewTeleporterSceneFactory() *TeleporterSceneFactory {
	return &TeleporterSceneFactory{}
}

// CreateActor creates a teleporter scene actor
func (f *TeleporterSceneFactory) createEntity() *actors.SimpleActor {
	return actors.NewSimpleActor()
}

// InitFromConfigId initializes from config ID
func (f *TeleporterSceneFactory) initFromConfig(actor *actors.SimpleActor, confId int32) {

}

func (f *TeleporterSceneFactory) initFromDO(actor *actors.SimpleActor) {
}

func (f *TeleporterSceneFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_TeleporterScene
}
