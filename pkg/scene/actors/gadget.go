package actors

import (
	"gameSrv/pkg/scene/actors/entity"
)

// Gadget is a gadget actor
type Gadget struct {
	entity.SimpleActor
}

// NewGadget creates a new Gadget
func NewGadget() *Gadget {
	g := &Gadget{}
	g.SimpleActor = *entity.NewSimpleActor()
	return g
}

// GetActorType returns the actor type
func (g *Gadget) GetActorType() int32 {
	return 9 // EActorType_Gadget
}

// HandleInteract handles interaction
func (g *Gadget) HandleInteract(interactMan *entity.Entity, optionType int32, optionParams []int64) int32 {
	// TODO: implement
	return 0 // SUCCESS
}