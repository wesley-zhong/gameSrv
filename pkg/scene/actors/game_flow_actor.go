package actors

import "gameSrv/pkg/scene/actors/entity"

// GameFlowActor is a game flow actor
type GameFlowActor struct {
	entity.SimpleActor
	SpawnMode int32
}

// NewGameFlowActor creates a new GameFlowActor
func NewGameFlowActor() *GameFlowActor {
	g := &GameFlowActor{
		SpawnMode: 0,
	}
	g.SimpleActor = *entity.NewSimpleActor()
	return g
}

// GetActorType returns actor type
func (g *GameFlowActor) GetActorType() int32 {
	return 20 // EActorType_GameFlow
}

// GetSpawnMode returns spawn mode
func (g *GameFlowActor) GetSpawnMode() int32 {
	return g.SpawnMode
}

// SetSpawnMode sets spawn mode
func (g *GameFlowActor) SetSpawnMode(spawnMode int32) {
	g.SpawnMode = spawnMode
}