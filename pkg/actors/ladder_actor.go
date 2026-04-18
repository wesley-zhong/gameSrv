package actors

import "gameSrv/pkg/scene"

// LadderActor is a ladder actor
type LadderActor struct {
	SimpleActor
}

func (l *LadderActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewLadderActor creates a new LadderActor
func NewLadderActor() *LadderActor {
	l := &LadderActor{}
	l.SimpleActor = *NewSimpleActor()
	return l
}

// GetActorType returns actor type
func (l *LadderActor) GetActorType() int32 {
	return 13 // EActorType_Ladder
}
