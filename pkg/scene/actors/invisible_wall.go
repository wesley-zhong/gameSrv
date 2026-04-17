package actors

import "gameSrv/pkg/scene/actors/entity"

// InvisibleWall is an invisible wall actor
type InvisibleWall struct {
	entity.SimpleActor
}

// NewInvisibleWall creates a new InvisibleWall
func NewInvisibleWall() *InvisibleWall {
	i := &InvisibleWall{}
	i.SimpleActor = *entity.NewSimpleActor()
	return i
}

// GetActorType returns actor type
func (i *InvisibleWall) GetActorType() int32 {
	return 14 // EActorType_InvisibleWall
}