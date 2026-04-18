package factory

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// TeleporterFactory creates teleporter actors
type TeleporterFactory struct {
	EntityFactory[actors.TeleporterActor]
}

// NewTeleporterFactory creates a new TeleporterFactory
func NewTeleporterFactory() *TeleporterFactory {
	return &TeleporterFactory{}
}

// CreateActor creates a teleporter actor
func (f *TeleporterFactory) createEntity() *actors.TeleporterActor {
	return actors.NewTeleporterActor()
}

// InitFromConfigId initializes from config ID
func (f *TeleporterFactory) initFromConfig(actor *actors.TeleporterActor, confId int32) {

}

func (f *TeleporterFactory) initFromDO(actor *actors.TeleporterActor) {
}

func (f *TeleporterFactory) GetEntityType() int32 {
	return cfg.ActorType_EActorType_Teleporter
}
