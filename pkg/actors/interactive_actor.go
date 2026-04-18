package actors

import "gameSrv/pkg/scene"

// InteractiveActor is an interactive actor
type InteractiveActor struct {
	SimpleActor
}

func (i *InteractiveActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewInteractiveActor creates a new InteractiveActor
func NewInteractiveActor() *InteractiveActor {
	i := &InteractiveActor{}
	i.SimpleActor = *NewSimpleActor()
	return i
}

// GetActorType returns actor type
func (i *InteractiveActor) GetActorType() int32 {
	return 15 // EActorType_InteractiveActor
}
