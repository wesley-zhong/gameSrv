package actors

import "gameSrv/pkg/scene/actors/entity"

// LadderActor is a ladder actor
type LadderActor struct {
	entity.SimpleActor
}

// NewLadderActor creates a new LadderActor
func NewLadderActor() *LadderActor {
	l := &LadderActor{}
	l.SimpleActor = *entity.NewSimpleActor()
	return l
}

// GetActorType returns actor type
func (l *LadderActor) GetActorType() int32 {
	return 13 // EActorType_Ladder
}
