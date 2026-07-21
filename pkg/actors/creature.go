package actors

import (
	"gameSrv/pkg/scene"
	"gameSrv/pkg/actors/battle_props"
)

// IActorBuffModule manages buffs for an actor (forward declaration)
// The actual type is in game/battle/buff package to avoid circular imports
type IActorBuffModule interface {
	AddBuffTag(inTag int) bool
	RemoveBuffTag(inTag int) bool
	HasBuffTag(tagEnum int) bool
	AddBuff(templateId int, casterActor, holder *Creature, uid int64, exParam int, bSystem bool) interface{}
	RemoveBuffByConfId(cnfId int, reason int) bool
	ClearLeaveBattleBuff()
}

// IActorBattleModule manages battle data for an actor (forward declaration)
// The actual type is in game/battle package to avoid circular imports
type IActorBattleModule interface {
	GetOwner() *Creature
	GetRandomValue() int
	GetBuffUidStart() int64
	GetRandomStartIndex() int
	SetDoAttackStartIndex(idx int64)
	GetDoAttackStartIndex() int64
	ResetBattleInfo()
	ClearAllChanged()
	ClearLeaveBattleBuff()
	HasState(state int32) bool
	AddState(state int32, bSyncStateToTag bool, count int) bool
	RemoveState(state int32, count int, bSyncStateToClient bool)
	HasImmunityAbility(immunityType int32) bool
	GetCurHp() int
	GetMaxHp() int
	GetProperty(prop int32) int
	SetProperty(prop, value int32, isNeedRecalculate bool) int
	AddProperty(prop, value int32) int
	OnAvatarPropsChangFinish()
	OnAvatarPropsRestFinish(reason int32)
	OnBuffUpdate(buff interface{})
	OnBuffRemoved(removedBuff interface{}, reason int)
	OnPlayerLeaveScene()
	HandleDead(killerActor *Entity) bool
	SystemAddBuff(confId int, uid int64) bool
	SystemDelBuff(buffCnfId int, uid int64) bool
	ReCalculateBattleProps(prop int32)
	OnPropertyChange(prop, oldValue, newValue int)
}

// Creature is a combatable entity in scene with abilities and properties
type Creature struct {
	Entity
	LifeState int32
	Level     int32
	CurState  int32
	CampType  int32

	// Battle modules - will be set when the battle system is initialized
	ActorBuffModule  IActorBuffModule
	ActorBattleModule IActorBattleModule

	// Battle properties - manages combat stats and attributes
	BattleProps *battle_props.BattleProps
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
	// Reset life state
	c.LifeState = 0 // LIFE_NONE
	c.Level = 1
	c.CurState = 0

	// Initialize Entity
	c.Entity.Init()

	// Initialize battle modules if present
	if c.ActorBattleModule != nil {
		if module, ok := c.ActorBattleModule.(interface{ Init() }); ok {
			module.Init()
		}
	}
}

// Reset resets the creature
func (c *Creature) Reset() bool {
	// Reset life state
	c.LifeState = 0 // LIFE_NONE
	c.CurState = 0

	// Clear buffs when not in battle
	if c.ActorBuffModule != nil {
		c.ActorBuffModule.ClearLeaveBattleBuff()
	}

	// Reset battle module
	if c.ActorBattleModule != nil {
		if module, ok := c.ActorBattleModule.(interface{ Reset() bool }); ok {
			module.Reset()
		}
	}

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

// AddBuffTag adds a buff tag to this creature
func (c *Creature) AddBuffTag(tag int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	return c.ActorBuffModule.AddBuffTag(tag)
}

// RemoveBuffTag removes a buff tag from this creature
func (c *Creature) RemoveBuffTag(tag int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	return c.ActorBuffModule.RemoveBuffTag(tag)
}

// HasBuffTag checks if this creature has a specific buff tag
func (c *Creature) HasBuffTag(tag int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	return c.ActorBuffModule.HasBuffTag(tag)
}

// AddBuff adds a buff to this creature
func (c *Creature) AddBuff(templateId int, casterActor *Creature, uid int64, exParam int, bSystem bool) interface{} {
	if c.ActorBuffModule == nil {
		return nil
	}
	return c.ActorBuffModule.AddBuff(templateId, casterActor, c, uid, exParam, bSystem)
}

// RemoveBuffByConfId removes buffs by config ID from this creature
func (c *Creature) RemoveBuffByConfId(cnfId int, reason int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	return c.ActorBuffModule.RemoveBuffByConfId(cnfId, reason)
}

// RemoveBuffByClass removes buffs by class from this creature
func (c *Creature) RemoveBuffByClass(classType int, reason int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	if module, ok := c.ActorBuffModule.(interface{ RemoveBuffByClass(int, int) bool }); ok {
		return module.RemoveBuffByClass(classType, reason)
	}
	return false
}

// RemoveBuffBySubClass removes buffs by subclass from this creature
func (c *Creature) RemoveBuffBySubClass(classID, subClassID int, reason int) bool {
	if c.ActorBuffModule == nil {
		return false
	}
	if module, ok := c.ActorBuffModule.(interface{ RemoveBuffBySubClass(int, int, int) bool }); ok {
		return module.RemoveBuffBySubClass(classID, subClassID, reason)
	}
	return false
}

// ==================== Battle Properties Methods ====================

// GetBattleProps returns the battle properties module
func (c *Creature) GetBattleProps() *battle_props.BattleProps {
	return c.BattleProps
}

// SetBattleProps sets the battle properties module
func (c *Creature) SetBattleProps(props *battle_props.BattleProps) {
	c.BattleProps = props
}

// InitBattleProps initializes battle properties for this creature
func (c *Creature) InitBattleProps() {
	if c.BattleProps == nil {
		c.BattleProps = battle_props.NewBattlePropsWithOwner(c)
	}
}

// GetProperty returns a property value from battle properties
func (c *Creature) GetProperty(prop int32) int {
	if c.BattleProps == nil {
		return 0
	}
	return c.BattleProps.GetProperty(int(prop))
}

// SetProperty sets a property value in battle properties
func (c *Creature) SetProperty(prop, value int32, isNeedRecalculate bool) int {
	if c.BattleProps == nil {
		return 0
	}
	return c.BattleProps.SetPropertyWithRecalculate(int(prop), int(value), isNeedRecalculate)
}

// AddProperty adds value to a property in battle properties
func (c *Creature) AddProperty(prop, value int32) int {
	if c.BattleProps == nil {
		return 0
	}
	oldValue := c.GetProperty(prop)
	newValue := oldValue + int(value)
	return c.SetProperty(prop, int32(newValue), true)
}

// GetBaseProps returns base property value from battle properties
func (c *Creature) GetBaseProps(prop int32) int {
	if c.BattleProps == nil {
		return 0
	}
	return c.BattleProps.GetBasePropsByInt(int(prop))
}

// CanChangeProp checks if a property can be changed
func (c *Creature) CanChangeProp(prop, value int) bool {
	if c.ActorBattleModule == nil {
		return true
	}
	// Use ActorBattleModule.CanChangeProp if available
	if module, ok := c.ActorBattleModule.(interface {
		CanChangeProp(prop, value int) bool
	}); ok {
		return module.CanChangeProp(prop, value)
	}
	return true
}

// GetActorBattleModule returns the actor battle module
// This satisfies the battle_props.BattlePropsOwner interface
func (c *Creature) GetActorBattleModule() battle_props.ActorBattleModule {
	if c.ActorBattleModule == nil {
		return nil
	}
	// Type assertion to the battle_props.ActorBattleModule interface
	if module, ok := c.ActorBattleModule.(interface {
		ReCalculateBattleModules(prop int32)
		OnPropertyChange(prop, oldValue, newValue int)
	}); ok {
		return module
	}
	return nil
}

// ReCalculateBattleModules recalculates battle modules when properties change
func (c *Creature) ReCalculateBattleModules(prop int32) {
	// Forward to ActorBattleModule
	if c.ActorBattleModule == nil {
		return
	}
	if module, ok := c.ActorBattleModule.(interface {
		ReCalculateBattleModules(prop int32)
	}); ok {
		module.ReCalculateBattleModules(prop)
	}
}
