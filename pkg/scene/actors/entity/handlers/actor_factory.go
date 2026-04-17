package handlers

import (
	"gameSrv/pkg/scene/actors/entity"
)

// ActorFactory is the actor factory interface
type ActorFactory[T entity.Entity] interface {
	// CreateActorFrom creates an actor from level instance configuration
	CreateActorFrom(levelActorCfg interface{}, isFromDO bool) T
	// CreateActorFromConfig creates an actor from configuration ID
	CreateActorFromConfig(configId int64, isFromDO bool) T
	// CreateActor creates an actor instance
	CreateActor() T
	// InitFromLevelInstance initializes from level instance
	InitFromLevelInstance(act T, levelActorCfg interface{}, isFromDO bool) error
	// InitFromConfigId initializes from configuration ID
	InitFromConfigId(act T, configId int64, isFromDO bool) bool
	// InitSpawnerProperty initializes spawner properties
	InitSpawnerProperty(act T, levelActorCfg interface{})
	// GetActorType gets actor type
	GetActorType() int32
}

// DefaultActorFactory is the default actor factory implementation
type DefaultActorFactory[T entity.Entity] struct{}

// CreateActorFrom creates an actor from level instance configuration
func (d *DefaultActorFactory[T]) CreateActorFrom(levelActorCfg interface{}, isFromDO bool) T {
	_ = d.CreateActor()
	// TODO: implement init from level instance
	var zero T
	return zero
}

// CreateActorFromConfig creates an actor from configuration ID
func (d *DefaultActorFactory[T]) CreateActorFromConfig(configId int64, isFromDO bool) T {
	_ = d.CreateActor()
	// TODO: implement init from config ID
	var zero T
	return zero
}

// CreateActor creates an actor instance
func (d *DefaultActorFactory[T]) CreateActor() T {
	var zero T
	return zero
}

// InitFromLevelInstance initializes from level instance
func (d *DefaultActorFactory[T]) InitFromLevelInstance(act T, levelActorCfg interface{}, isFromDO bool) error {
	// Default implementation, subclasses can override
	return nil
}

// InitFromConfigId initializes from configuration ID
func (d *DefaultActorFactory[T]) InitFromConfigId(act T, configId int64, isFromDO bool) bool {
	// Default implementation, subclasses can override
	return true
}

// InitSpawnerProperty initializes spawner properties
func (d *DefaultActorFactory[T]) InitSpawnerProperty(act T, levelActorCfg interface{}) {
	// Default implementation, subclasses can override
}

// GetActorType gets actor type
func (d *DefaultActorFactory[T]) GetActorType() int32 {
	// Default implementation, subclasses must override
	return 0 // ActorTypeNone
}