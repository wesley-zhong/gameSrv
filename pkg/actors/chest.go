package actors

import "gameSrv/pkg/scene"

// Chest is a chest class
type Chest struct {
	SimpleActor
	DropID int32
}

func (c *Chest) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewChest creates a new Chest
func NewChest() *Chest {
	c := &Chest{}
	c.SimpleActor = *NewSimpleActor()
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
