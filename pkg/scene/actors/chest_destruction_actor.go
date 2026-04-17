package actors

import "gameSrv/pkg/scene/actors/entity"

// ChestDestructionActor is a chest destruction actor
type ChestDestructionActor struct {
	entity.SimpleActor
}

// NewChestDestructionActor creates a new ChestDestructionActor
func NewChestDestructionActor() *ChestDestructionActor {
	c := &ChestDestructionActor{}
	c.SimpleActor = *entity.NewSimpleActor()
	return c
}

// GetActorType returns actor type
func (c *ChestDestructionActor) GetActorType() int32 {
	return 16 // EActorType_ChestDestruction
}