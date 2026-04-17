package actors

import (
	"gameSrv/pkg/scene/actors/entity"
)

// Chest is a chest class
type Chest struct {
	entity.SimpleActor
	DropID int32
}

// NewChest creates a new Chest
func NewChest() *Chest {
	c := &Chest{}
	c.SimpleActor = *entity.NewSimpleActor()
	return c
}

// GetActorType returns actor type
func (c *Chest) GetActorType() int32 {
	return 8 // EActorType_Chest
}

// GetDropID returns drop ID
func (c *Chest) GetDropID() int32 {
	return c.DropID
}

// SetDropID sets drop ID
func (c *Chest) SetDropID(dropID int32) {
	c.DropID = dropID
}