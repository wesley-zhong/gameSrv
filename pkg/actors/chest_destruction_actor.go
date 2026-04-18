package actors

import "gameSrv/pkg/scene"

// ChestDestructionActor is a chest destruction actor
type ChestDestructionActor struct {
	SimpleActor
}

func (c *ChestDestructionActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewChestDestructionActor creates a new ChestDestructionActor
func NewChestDestructionActor() *ChestDestructionActor {
	c := &ChestDestructionActor{}
	c.SimpleActor = *NewSimpleActor()
	return c
}

// GetActorType returns actor type
func (c *ChestDestructionActor) GetActorType() int32 {
	return 16 // EActorType_ChestDestruction
}
