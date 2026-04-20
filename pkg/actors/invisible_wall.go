package actors

import "gameSrv/pkg/scene"

// InvisibleWall is an invisible wall actor
type InvisibleWall struct {
	SimpleActor
}

func (i *InvisibleWall) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewInvisibleWall creates a new InvisibleWall
func NewInvisibleWall() *InvisibleWall {
	i := &InvisibleWall{}
	i.SimpleActor = *NewSimpleActor()
	return i
}

// GetActorType returns actor type
func (i *InvisibleWall) GetActorType() int32 {
	return 14 // EActorType_InvisibleWall
}
