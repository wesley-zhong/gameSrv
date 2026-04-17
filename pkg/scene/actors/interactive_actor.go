package actors

import "gameSrv/pkg/scene/actors/entity"

// InteractiveActor is an interactive actor
type InteractiveActor struct {
	entity.SimpleActor
}

// NewInteractiveActor creates a new InteractiveActor
func NewInteractiveActor() *InteractiveActor {
	i := &InteractiveActor{}
	i.SimpleActor = *entity.NewSimpleActor()
	return i
}

// GetActorType returns actor type
func (i *InteractiveActor) GetActorType() int32 {
	return 15 // EActorType_InteractiveActor
}