package actors

import (
	"gameSrv/pkg/scene"
)

// Creature is a combatable entity in scene with abilities and properties
type Creature struct {
	Entity
	LifeState int32
	Level     int32
	CurState  int32
	CampType  int32
}

func (c *Creature) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

func (c *Creature) GetOnlyVisionUid() int64 {
	return 0
}

func (c *Creature) GetPhasingId() int64 {
	return c.PhasingID
}

// NewCreature creates a new Creature
func NewCreature() *Creature {
	c := &Creature{
		LifeState: 0, // LIFE_NONE
		Level:     1,
	}
	c.Entity = *NewEntity()
	return c
}

// Init initializes the creature
func (c *Creature) Init() {
	// TODO: implement
}

// Reset resets the creature
func (c *Creature) Reset() bool {
	// TODO: implement
	return true
}

// GetExtData returns extended data
func (c *Creature) GetExtData() interface{} {
	return nil
}

// SetExtData sets extended data
func (c *Creature) SetExtData(object interface{}) {
	// Default implementation does nothing
}

// HandleDead handles death
func (c *Creature) HandleDead(killerActor *Entity) bool {
	// TODO: implement
	return true
}

// OnBeforeLeaveScene is called before leaving scene
func (c *Creature) OnBeforeLeaveScene(context interface{}) {
	c.Entity.OnBeforeLeaveScene(context)
	// TODO: implement buff clearing and summons
}

// OnBeforeEnterScene is called before entering scene
func (c *Creature) OnBeforeEnterScene(scn scene.IScene, context interface{}) {
	c.Entity.OnBeforeEnterScene(scn, context)
}

// OnAfterEnterScene is called after entering scene
func (c *Creature) OnAfterEnterScene(scn scene.IScene, context interface{}) {
	c.Entity.OnAfterEnterScene(scn, context)
}

// OnAfterLeaveScene is called after leaving scene
func (c *Creature) OnAfterLeaveScene(context interface{}) {
	c.Entity.OnAfterLeaveScene(context)
}

// OnEnterPlayerView is called when entering a player's view
func (c *Creature) OnEnterPlayerView(p scene.IGamePlayer) {
	c.Entity.OnEnterPlayerView(p)
}

// OnExitPlayerView is called when leaving a player's view
func (c *Creature) OnExitPlayerView(p scene.IGamePlayer) {
	c.Entity.OnExitPlayerView(p)
}

// ChangeObjectState changes the object state
func (c *Creature) ChangeObjectState(srcEntity scene.IEntity, state int32) {
	c.InteractInfo.ObjectState = state
}
