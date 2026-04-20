package actors

import "gameSrv/pkg/scene"

// GameFlowActor is a game flow actor
type GameFlowActor struct {
	SimpleActor
	SpawnMode int32
}

func (g *GameFlowActor) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewGameFlowActor creates a new GameFlowActor
func NewGameFlowActor() *GameFlowActor {
	g := &GameFlowActor{
		SpawnMode: 0,
	}
	g.SimpleActor = *NewSimpleActor()
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
