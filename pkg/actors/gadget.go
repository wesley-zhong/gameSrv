package actors

import "gameSrv/pkg/scene"

// Gadget is a gadget actor
type Gadget struct {
	SimpleActor
}

func (g *Gadget) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewGadget creates a new Gadget
func NewGadget() *Gadget {
	g := &Gadget{}
	g.SimpleActor = *NewSimpleActor()
	return g
}

// GetActorType returns the actor type
func (g *Gadget) GetActorType() int32 {
	return 9 // EActorType_Gadget
}

// HandleInteract handles interaction
func (g *Gadget) HandleInteract(interactMan *Entity, optionType int32, optionParams []int64) int32 {
	// TODO: implement
	return 0 // SUCCESS
}
