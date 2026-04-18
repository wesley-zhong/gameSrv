package actors

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

// PlayerCreature is a creature controlled by a player
type PlayerCreature struct {
	Creature
}

func (p *PlayerCreature) GetOnlyVisionUid() int64 {
	return 0
}

func (p *PlayerCreature) GetPhasingId() int64 {
	return p.PhasingID
}

// NewPlayerCreature creates a new PlayerCreature
func NewPlayerCreature(p scene.IGamePlayer) *PlayerCreature {
	pc := &PlayerCreature{}
	pc.Creature = *NewCreature()
	pc.SetOwner(p)
	pc.SetNeedSaveToDb(false)
	return pc
}

// GetObjectState returns the object state
func (p *PlayerCreature) GetObjectState() int32 {
	return 0
}

// CanEnterAoi returns whether this creature enters AOI
func (p *PlayerCreature) CanEnterAoi() bool {
	return false
}

// EnterScene enters the scene
func (p *PlayerCreature) EnterScene(scn scene.IScene, context interface{}) {
	p.OnBeforeEnterScene(scn, context)
	p.OnAfterEnterScene(scn, context)
}

// LeaveScene leaves the scene
func (p *PlayerCreature) LeaveScene(context scene.IScene, deadClearTime int64) {
	p.OnBeforeLeaveScene(context)
	p.Scene = nil
	p.OnAfterLeaveScene(context)
}

// OnBeforeEnterScene is called before entering scene
func (p *PlayerCreature) OnBeforeEnterScene(scn scene.IScene, context interface{}) {
	// TODO: implement
}

// ClearSpeed clears the speed
func (p *PlayerCreature) ClearSpeed() {
	p.SetSpeed(math.ZeroVector3())
}

// ClearMotionState clears the motion state
func (p *PlayerCreature) ClearMotionState() {
	// TODO: implement
}
