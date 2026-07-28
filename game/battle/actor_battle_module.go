package battle

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// BuffModule defines the interface for buff module operations
// This interface is implemented by buff.ActorBuffModule to avoid circular imports
type BuffModule interface {
	TriggerBuffEvent(event int, param interface{}) bool
	ClearNotSystemBuff()
	ClearChangedData()
	ClearLeaveBattleBuff()
	ClearBuff(death bool)
	SystemAddBuff(confId int, uid int64) bool
	SystemRemoveBuff(buffCnfId int, uid int64) bool
}

// CreatureActionState represents the creature action state
type CreatureActionState int32

const (
	StateNone             CreatureActionState = 0
	StateUnbeat           CreatureActionState = 1
	StateHarmfulWeak      CreatureActionState = 2
	StateImmunityTypeWeak CreatureActionState = 3
	StateInfighting       CreatureActionState = 4
	StateMAX              CreatureActionState = 5
)

// ECreatureActionState is an alias for CreatureActionState for compatibility
type ECreatureActionState = CreatureActionState

const (
	ECreatureActionState_NONE                    ECreatureActionState = 0
	ECreatureActionState_MAX_VALUE               ECreatureActionState = 5
	ECreatureActionState_STATE_UNBEAT_VALUE       ECreatureActionState = 1
	ECreatureActionState_STATE_HARMFUL_WEAK_VALUE ECreatureActionState = 2
	ECreatureActionState_IMMUNITY_TYPE_WEAK       ECreatureActionState = 3
	ECreatureActionState_STATE_INFIGHTING_VALUE     ECreatureActionState = 4
)

// ShipGunTypeEnum represents the ship gun type
type ShipGunTypeEnum int32

const (
	ShipGunTypeEnumLightGun ShipGunTypeEnum = 0
	ShipGunTypeEnumHeavyGun ShipGunTypeEnum = 1
)

// FAbnormalClosing represents abnormal closing data
type FAbnormalClosing struct {
	AbnormalStrength int
	AbnormalProp     int32
	AbnormalPropMax  int32
	OnceAddValue     int
}

// Reset resets the abnormal closing data
func (f *FAbnormalClosing) Reset() {
	f.AbnormalStrength = 0
	f.OnceAddValue = 0
	f.AbnormalProp = 0
	f.AbnormalPropMax = 0
}

// ActorBattleModule manages battle-related data and operations for a creature
type ActorBattleModule struct {
	owner                 *actors.Creature
	CharacterStateFlag    int64
	States                [int(ECreatureActionState_MAX_VALUE)]int
	Sermon                SermonType
	Element               ElementType
	StateBattleNoSp0      int

	// Sub-modules
	actorBuffModule  BuffModule
	actorSkillModule *ActorSkillModule

	// Combat state
	CurGaList        []int
	CurComboList     []int
	CurMontageList   []int
	curMontageId      int
	newTarget        bool
	shipGunIndex     ShipGunTypeEnum
	randomIndex      int
	doAttackStartIdx int64
	buffUidStart      int64

	// Element terminal strength
	ElementTerminalStrength map[int32]int

	// Current targets
	CurAttackTarget   *actors.Creature
	CurBeAttackTarget *actors.Creature

	// Abnormal closing cache
	AbnormalClosingCachePtr FAbnormalClosing

	// Timer task ID
	timerTaskID interface{}

	pauseTime bool
}

// NewActorBattleModule creates a new ActorBattleModule
func NewActorBattleModule(owner *actors.Creature, isServer bool) *ActorBattleModule {
	m := &ActorBattleModule{
		owner:                 owner,
		Sermon:                SermonType_None,
		Element:               ElementType_None,
		CurGaList:             make([]int, 0),
		CurComboList:          make([]int, 0),
		CurMontageList:        make([]int, 0),
		shipGunIndex:          ShipGunTypeEnumLightGun,
		ElementTerminalStrength: make(map[int32]int),
	}

	// Initialize sub-modules - type assertion will be done when methods are called
	// These are created in the buff package to avoid circular imports
	// m.actorSkillModule = buff.NewActorSkillModule(owner)
	// m.actorBuffModule = buff.NewActorBuffModule(owner, isServer)

	// Initialize element terminal strengths
	m.ElementTerminalStrength[int32(cfg.HeroProp_Ice_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Fire_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Wind_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Thunder_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Light_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Dark_Abnormal)] = 0

	return m
}

// SetBuffModule sets the buff module (called during initialization)
func (m *ActorBattleModule) SetBuffModule(buffModule BuffModule) {
	m.actorBuffModule = buffModule
}

// SetSkillModule sets the skill module (called during initialization)
func (m *ActorBattleModule) SetSkillModule(skillModule *ActorSkillModule) {
	m.actorSkillModule = skillModule
}

// GetBuffModule returns the buff module
func (m *ActorBattleModule) GetBuffModule() BuffModule {
	return m.actorBuffModule
}

// GetSkillModule returns the skill module
func (m *ActorBattleModule) GetSkillModule() *ActorSkillModule {
	return m.actorSkillModule
}

// GetOwner returns the owner creature
func (m *ActorBattleModule) GetOwner() *actors.Creature {
	return m.owner
}

// GetRandomValue generates a random value
// Uses simple counter - in production, this should use scene's random value generator
func (m *ActorBattleModule) GetRandomValue() int {
	m.randomIndex++
	return m.randomIndex
}

// GetBuffUidStart returns the starting buff UID
func (m *ActorBattleModule) GetBuffUidStart() int64 {
	return m.buffUidStart
}

// GetRandomStartIndex returns the random start index
func (m *ActorBattleModule) GetRandomStartIndex() int {
	return m.randomIndex
}

// SetDoAttackStartIndex sets the do attack start index
func (m *ActorBattleModule) SetDoAttackStartIndex(idx int64) {
	m.doAttackStartIdx = idx
}

// GetDoAttackStartIndex gets the do attack start index
func (m *ActorBattleModule) GetDoAttackStartIndex() int64 {
	return m.doAttackStartIdx
}

// ResetBattleInfo resets battle information
func (m *ActorBattleModule) ResetBattleInfo() {
	// TODO: Generate buff UID start from scene's buff ID allocator
	m.buffUidStart = 0
	m.randomIndex = 0
	m.shipGunIndex = ShipGunTypeEnumLightGun

	// Clear non-system buffs via buff module
	if m.actorBuffModule != nil {
		m.actorBuffModule.ClearNotSystemBuff()
	}

	// TODO: Setup timer task when timer system is implemented
}

// ClearAllChanged clears all changed data
func (m *ActorBattleModule) ClearAllChanged() {
	// Clear changed props from BattleProps
	if m.owner != nil && m.owner.BattleProps != nil {
		m.owner.BattleProps.ClearChangedProps()
	}
	// Clear skill module changed data
	if m.actorSkillModule != nil {
		m.actorSkillModule.ClearChangedData()
	}
	// Clear buff module changed data
	if m.actorBuffModule != nil {
		m.actorBuffModule.ClearChangedData()
	}
}

// ClearLeaveBattleBuff clears buffs when leaving battle
func (m *ActorBattleModule) ClearLeaveBattleBuff() {
	if m.actorBuffModule != nil {
		m.actorBuffModule.ClearLeaveBattleBuff()
	}
}

// HasState checks if the creature has a specific state
func (m *ActorBattleModule) HasState(state CreatureActionState) bool {
	if int(state) >= len(m.States) {
		return false
	}
	return (m.CharacterStateFlag & (1 << state)) != 0
}

// AddState adds a state to the creature
func (m *ActorBattleModule) AddState(state CreatureActionState, bSyncStateToTag bool, count int) bool {
	if int(state) >= len(m.States) {
		return false
	}

	if state == ECreatureActionState_STATE_HARMFUL_WEAK_VALUE && m.HasImmunityAbility(ECreatureActionState_IMMUNITY_TYPE_WEAK) {
		return false
	}

	index := int(state)
	base := m.States[index]

	m.States[index] += count
	m.CharacterStateFlag |= (1 << state)

	current := m.States[index]
	m.onStateChange(state, base, current, bSyncStateToTag)

	return true
}

// RemoveState removes a state from the creature
func (m *ActorBattleModule) RemoveState(state CreatureActionState, count int, bSyncStateToClient bool) {
	if !m.HasState(state) {
		return
	}

	index := int(state)
	if index >= len(m.States) {
		return
	}

	base := m.States[index]
	if count == -1 {
		count = base
	}

	m.States[index] -= count
	if m.States[index] == 0 {
		m.CharacterStateFlag &^= (1 << state)
	}

	current := m.States[index]
	m.onStateChange(state, base, current, bSyncStateToClient)
}

// ClearStates clears all states
func (m *ActorBattleModule) ClearStates() {
	for i := range m.States {
		m.States[i] = 0
	}
	m.CharacterStateFlag = 0
}

// HasImmunityAbility checks if the creature has immunity ability
func (m *ActorBattleModule) HasImmunityAbility(immunityType CreatureActionState) bool {
	return (m.CharacterStateFlag & (1 << immunityType)) != 0
}

// onStateChange handles state change events
func (m *ActorBattleModule) onStateChange(actionState CreatureActionState, base, current int, bSyncStateToTag bool) {
	bOut := current == 0

	// Trigger buff event for state change via buff module
	if m.actorBuffModule != nil {
		triggerParam := map[string]interface{}{
			"Target": m.owner,
			"State":  int(actionState),
			"IsOut":  bOut,
		}
		m.actorBuffModule.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_STATE_CHANGE), triggerParam)
	}
}

// OnStateTrigger triggers a state change event
func (m *ActorBattleModule) OnStateTrigger(actionState CreatureActionState, isAdd bool) {
	if m.actorBuffModule != nil {
		triggerParam := map[string]interface{}{
			"Target": m.owner,
			"State":  int(actionState),
			"IsOut":  !isAdd,
		}
		m.actorBuffModule.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_STATE_CHANGE), triggerParam)
	}
}

// GetCurHp returns current HP
func (m *ActorBattleModule) GetCurHp() int {
	// Get HP property from owner's BattleProps
	if m.owner == nil {
		return 0
	}
	return m.owner.GetProperty(int32(cfg.HeroProp_Health))
}

// GetMaxHp returns max HP
func (m *ActorBattleModule) GetMaxHp() int {
	// Get max HP property from owner's BattleProps
	if m.owner == nil {
		return 0
	}
	return m.owner.GetProperty(int32(cfg.HeroProp_MaxHealth))
}

// GetProperty returns a property value
func (m *ActorBattleModule) GetProperty(prop int32) int {
	// Get property from owner's BattleProps
	if m.owner == nil {
		return 0
	}
	return m.owner.GetProperty(prop)
}

// SetProperty sets a property value
func (m *ActorBattleModule) SetProperty(prop, value int32, isNeedRecalculate bool) int {
	// Set property in owner's BattleProps
	if m.owner == nil {
		return 0
	}
	return m.owner.SetProperty(prop, value, isNeedRecalculate)
}

// AddProperty adds value to a property
func (m *ActorBattleModule) AddProperty(prop, value int32) int {
	// Add value to property in owner's BattleProps
	if m.owner == nil {
		return 0
	}
	return m.owner.AddProperty(prop, value)
}

// OnAvatarPropsChangFinish handles property change finish event
// Sends property change notification to client via network system
func (m *ActorBattleModule) OnAvatarPropsChangFinish() {
	// Network notification will be sent when network system is implemented
}

// OnAvatarPropsRestFinish handles property reset finish event
// Sends property reset notification to client via network system
func (m *ActorBattleModule) OnAvatarPropsRestFinish(reason int32) {
	_ = reason
	// Network notification will be sent when network system is implemented
}

// OnBuffUpdate is called when a buff is updated
// Sends buff update notification to client via network system
func (m *ActorBattleModule) OnBuffUpdate(buff interface{}) {
	_ = buff
	// Network notification will be sent when network system is implemented
}

// OnBuffRemoved is called when a buff is removed
// Sends buff remove notification to client via network system
func (m *ActorBattleModule) OnBuffRemoved(removedBuff interface{}, reason int) {
	_ = removedBuff
	_ = reason
	// Network notification will be sent when network system is implemented
}

// OnPlayerLeaveScene is called when the player leaves the scene
func (m *ActorBattleModule) OnPlayerLeaveScene() {
	// Clear non-system buffs when leaving scene
	m.ClearLeaveBattleBuff()
}

// HandleDead handles creature death
func (m *ActorBattleModule) HandleDead(killerActor *actors.Entity) bool {
	// Clear all buffs on death
	if m.actorBuffModule != nil {
		m.actorBuffModule.ClearBuff(true)
	}

	// TODO: Cancel timer task when timer system is implemented
	// m.owner.getScene().getTimerTickContainer().cancelTask(m.timerTaskID)

	return true
}

// ReCalculateBattleProps recalculates battle properties
// Called when a property changes that affects other properties
func (m *ActorBattleModule) ReCalculateBattleProps(inProperty int32) {
	// This is called by BattleProps when a property changes
	// The actual recalculation logic would be here based on game rules
	// For now, this is a placeholder
	_ = inProperty
}

// ReCalculateBattleModules recalculates battle modules (alias for compatibility)
func (m *ActorBattleModule) ReCalculateBattleModules(prop int32) {
	m.ReCalculateBattleProps(prop)
}

// OnPropertyChange is called when a property changes
// This allows the battle module to react to property changes
func (m *ActorBattleModule) OnPropertyChange(prop, oldValue, newValue int) {
	// This is called by BattleProps when a property changes
	// Used for triggering effects based on property thresholds
	_ = prop
	_ = oldValue
	_ = newValue
}

// SystemAddBuff adds a system buff
func (m *ActorBattleModule) SystemAddBuff(confId int, uid int64) bool {
	if m.actorBuffModule != nil {
		return m.actorBuffModule.SystemAddBuff(confId, uid)
	}
	return false
}

// SystemDelBuff removes a system buff
func (m *ActorBattleModule) SystemDelBuff(buffCnfId int, uid int64) bool {
	if m.actorBuffModule != nil {
		return m.actorBuffModule.SystemRemoveBuff(buffCnfId, uid)
	}
	return false
}

// AddGAId adds a GA ID to current GA list
func (m *ActorBattleModule) AddGAId(gaId int) {
	m.CurGaList = append(m.CurGaList, gaId)
}

// RemoveGAId removes a GA ID from current GA list
func (m *ActorBattleModule) RemoveGAId(gaId int) {
	for i, id := range m.CurGaList {
		if id == gaId {
			m.CurGaList = append(m.CurGaList[:i], m.CurGaList[i+1:]...)
			break
		}
	}
}

// AddComboId adds a combo ID to current combo list
func (m *ActorBattleModule) AddComboId(comboId int) {
	m.CurComboList = append(m.CurComboList, comboId)
}

// RemoveComboId removes a combo ID from current combo list
func (m *ActorBattleModule) RemoveComboId(comboId int) {
	for i, id := range m.CurComboList {
		if id == comboId {
			m.CurComboList = append(m.CurComboList[:i], m.CurComboList[i+1:]...)
			break
		}
	}
}

// AddMontageID adds a montage ID to current montage list
func (m *ActorBattleModule) AddMontageID(montageId int) {
	m.CurMontageList = append(m.CurMontageList, montageId)
	m.curMontageId = montageId
	m.initSkillPropGainOnce(montageId)
}

// RemoveMontageID removes a montage ID from current montage list
func (m *ActorBattleModule) RemoveMontageID(montageId int) {
	for i, id := range m.CurMontageList {
		if id == montageId {
			m.CurMontageList = append(m.CurMontageList[:i], m.CurMontageList[i+1:]...)
			break
		}
	}

	m.resetSkillPropGainOnce()
}

// InitSkillPropGainOnce initializes skill prop gain once for a montage
func (m *ActorBattleModule) InitSkillPropGainOnce(montageId int) {
	m.initSkillPropGainOnce(montageId)
}

// ResetSkillPropGainOnce resets skill prop gain once
func (m *ActorBattleModule) ResetSkillPropGainOnce() {
	m.resetSkillPropGainOnce()
}

// initSkillPropGainOnce initializes skill prop gain once for a montage
// Tracks which montage is active for one-time skill properties
func (m *ActorBattleModule) initSkillPropGainOnce(montageId int) {
	m.curMontageId = montageId
	m.newTarget = false
}

// resetSkillPropGainOnce resets skill prop gain once tracking
// Called when montage ends
func (m *ActorBattleModule) resetSkillPropGainOnce() {
	m.curMontageId = 0
	m.newTarget = false
}

// GetCurMontageId returns the current montage ID
func (m *ActorBattleModule) GetCurMontageId() int {
	return m.curMontageId
}

// SetNewTarget sets the new target flag
func (m *ActorBattleModule) SetNewTarget(newTarget bool) {
	m.newTarget = newTarget
}

// GetNewTarget gets the new target flag
func (m *ActorBattleModule) GetNewTarget() bool {
	return m.newTarget
}

// GetCurAttackTarget returns the current attack target
func (m *ActorBattleModule) GetCurAttackTarget() *actors.Creature {
	return m.CurAttackTarget
}

// SetCurAttackTarget sets the current attack target
func (m *ActorBattleModule) SetCurAttackTarget(target *actors.Creature) {
	m.CurAttackTarget = target
}

// GetCurBeAttackTarget returns the current being attacked target
func (m *ActorBattleModule) GetCurBeAttackTarget() *actors.Creature {
	return m.CurBeAttackTarget
}

// SetCurBeAttackTarget sets the current being attacked target
func (m *ActorBattleModule) SetCurBeAttackTarget(target *actors.Creature) {
	m.CurBeAttackTarget = target
}

// CheckBuffChange checks if a buff change is valid based on current state
func (m *ActorBattleModule) CheckBuffChange(buffId, reason int) bool {
	// TODO: Implement buff change validation based on GA/Combo/Montage/Move state
	_ = buffId
	_ = reason
	return true
}

// CheckAttackData checks if attack data is valid for current state
func (m *ActorBattleModule) CheckAttackData(attackDataId int) bool {
	// TODO: Implement attack data validation based on current montage
	_ = attackDataId
	return true
}

// OnGA handles GA (General Action) events
func (m *ActorBattleModule) OnGA(entity *actors.Creature, isStart bool, gaId int) bool {
	if gaId == 0 {
		return false
	}

	if isStart {
		m.CurGaList = append(m.CurGaList, gaId)
		// TODO: Apply skill costs from GACombo config
	} else {
		for i, id := range m.CurGaList {
			if id == gaId {
				m.CurGaList = append(m.CurGaList[:i], m.CurGaList[i+1:]...)
				break
			}
		}
	}

	return true
}

// OnCombo handles combo events
func (m *ActorBattleModule) OnCombo(entity *actors.Creature, isStart bool, comboId int) bool {
	if comboId == 0 {
		return false
	}

	if isStart {
		// Clear combo abnormal effects when starting a new combo
		// TODO: Implement combo abnormal effects

		// Check if combo is valid for any current GA
		isAdd := false
		for _, gaId := range m.CurGaList {
			// TODO: Check GACombo config for this combo
			_ = gaId
		}

		if isAdd {
			m.CurComboList = append(m.CurComboList, comboId)
		}
		return isAdd
	} else {
		for i, id := range m.CurComboList {
			if id == comboId {
				m.CurComboList = append(m.CurComboList[:i], m.CurComboList[i+1:]...)
				break
			}
		}
	}

	return true
}

// OnMontage handles montage events
func (m *ActorBattleModule) OnMontage(entity *actors.Creature, isStart bool, montageId int) bool {
	if isStart {
		m.initSkillPropGainOnce(montageId)

		// Check if montage is valid for any current combo
		isAdd := false
		for _, comboId := range m.CurComboList {
			// TODO: Check ComboMontage config for this montage
			_ = comboId
		}

		// Check if montage is valid for any current GA
		if !isAdd {
			for _, gaId := range m.CurGaList {
				// TODO: Check GACombo config for this montage
				_ = gaId
			}
		}

		if isAdd || true { // TODO: Remove '|| true' when validation is implemented
			m.CurMontageList = append(m.CurMontageList, montageId)
			m.curMontageId = montageId

			// Trigger montage events via buff module
			if m.actorBuffModule != nil {
				mtgEventParam := map[string]interface{}{
					"Target": entity,
					"MtgId":  montageId,
				}
				m.actorBuffModule.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_MONTAGE_ID), mtgEventParam)
				m.actorBuffModule.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_MONTAGE_TAG), mtgEventParam)
			}
		}

		return isAdd
	} else {
		m.resetSkillPropGainOnce()

		// Remove montage from list
		for i, id := range m.CurMontageList {
			if id == montageId {
				m.CurMontageList = append(m.CurMontageList[:i], m.CurMontageList[i+1:]...)
				break
			}
		}

		// Trigger montage end event via buff module
		if m.actorBuffModule != nil {
			mtgEventParam := map[string]interface{}{
				"Target": entity,
				"MtgId":  montageId,
			}
			m.actorBuffModule.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_MONTAGE_END), mtgEventParam)
		}

		// TODO: Handle time dilation recovery
	}

	return true
}

// SetPauseTime sets the pause time flag
func (m *ActorBattleModule) SetPauseTime() {
	m.pauseTime = true
	// TODO: Pause scene countdown if applicable
}

// RecoverTime recovers from time dilation
func (m *ActorBattleModule) RecoverTime(dilationK float32) {
	m.pauseTime = false
	// TODO: Resume scene countdown if applicable
	_ = dilationK
}

// IsPauseTime returns whether time is paused
func (m *ActorBattleModule) IsPauseTime() bool {
	return m.pauseTime
}

// SetAttackDataProp sets attack data property
func (m *ActorBattleModule) SetAttackDataProp(atkDataId, propId, value int) {
	if m.actorSkillModule != nil {
		m.actorSkillModule.SetAtkDataProps(atkDataId, propId, value)
	}
}

// AddAttackDataAbnormalEffects adds abnormal effects to attack data
func (m *ActorBattleModule) AddAttackDataAbnormalEffects(atkDataId int, abnormalEffects []int) {
	if m.actorSkillModule != nil {
		m.actorSkillModule.AddAtkAbnormalEffectsSingle(atkDataId, abnormalEffects)
	}
}

// GetAttackDataProp gets attack data property
func (m *ActorBattleModule) GetAttackDataProp(atkDataId, propId int) int {
	if m.actorSkillModule != nil {
		return m.actorSkillModule.GetAtkDataProps(atkDataId, propId)
	}
	return 0
}

// Init initializes the battle module
func (m *ActorBattleModule) Init() {
	// Initialize element terminal strengths
	m.ElementTerminalStrength[int32(cfg.HeroProp_Ice_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Fire_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Wind_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Thunder_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Light_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Dark_Abnormal)] = 0
}

// Reset resets the battle module
func (m *ActorBattleModule) Reset() bool {
	// Reset states
	m.ClearStates()
	m.CharacterStateFlag = 0

	// Reset element terminal strength
	m.ElementTerminalStrength[int32(cfg.HeroProp_Ice_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Fire_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Wind_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Thunder_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Light_Abnormal)] = 0
	m.ElementTerminalStrength[int32(cfg.HeroProp_Dark_Abnormal)] = 0

	// Clear combat lists
	m.CurGaList = make([]int, 0)
	m.CurComboList = make([]int, 0)
	m.CurMontageList = make([]int, 0)

	// Reset targets
	m.CurAttackTarget = nil
	m.CurBeAttackTarget = nil

	return true
}

// Destroy cleans up battle module resources
func (m *ActorBattleModule) Destroy() {
	// TODO: Cancel timer task when timer system is implemented
	// if m.timerTaskID != nil {
	//     m.owner.getScene().getTimerTickContainer().cancelTask(m.timerTaskID)
	// }

	if m.actorBuffModule != nil {
		m.actorBuffModule.ClearBuff(true)
	}
}

// OnCostEnergyChange handles energy cost change event
func (m *ActorBattleModule) OnCostEnergyChange(skillId, eventReason int) bool {
	// TODO: Implement skill cost validation and application
	_ = skillId
	_ = eventReason
	return true
}

// ApplySkillCost applies skill cost
func (m *ActorBattleModule) ApplySkillCost(skillId int) bool {
	// TODO: Implement skill cost application
	_ = skillId
	return true
}

// EnergyRealCost performs real energy cost
func (m *ActorBattleModule) EnergyRealCost(property, costValue int) bool {
	// TODO: Implement energy real cost
	_ = property
	_ = costValue
	return true
}

// OnEntityDeathAbnormalDiffusion handles entity death abnormal diffusion
func (m *ActorBattleModule) OnEntityDeathAbnormalDiffusion(caster, target *actors.Creature) {
	// TODO: Implement abnormal diffusion on entity death
	_ = caster
	_ = target
}

// RecordEntityDeadAbnormalDiffusion records entity dead abnormal diffusion
func (m *ActorBattleModule) RecordEntityDeadAbnormalDiffusion(caster, target *actors.Creature, abnormalElements []int32, abnormalValues []int) {
	// TODO: Implement abnormal diffusion recording
	_ = caster
	_ = target
	_ = abnormalElements
	_ = abnormalValues
}

// CanChangeProp checks if property can be changed
func (m *ActorBattleModule) CanChangeProp(prop, value int) bool {
	return m.StateBattleNoSp0 == 0
}

// GetElementTerminalStrength returns element terminal strength
func (m *ActorBattleModule) GetElementTerminalStrength(prop int32) int {
	return m.ElementTerminalStrength[int32(prop)]
}

// SetElementTerminalStrength sets element terminal strength
func (m *ActorBattleModule) SetElementTerminalStrength(prop, value int) {
	m.ElementTerminalStrength[int32(prop)] = value
}

// GetAbnormalClosingCache returns the abnormal closing cache
func (m *ActorBattleModule) GetAbnormalClosingCache() *FAbnormalClosing {
	return &m.AbnormalClosingCachePtr
}

// GetShipGunIndex returns the ship gun index
func (m *ActorBattleModule) GetShipGunIndex() ShipGunTypeEnum {
	return m.shipGunIndex
}

// SetShipGunIndex sets the ship gun index
func (m *ActorBattleModule) SetShipGunIndex(index ShipGunTypeEnum) {
	m.shipGunIndex = index
}

// GetCurGaList returns the current GA list
func (m *ActorBattleModule) GetCurGaList() []int {
	result := make([]int, len(m.CurGaList))
	copy(result, m.CurGaList)
	return result
}

// GetCurComboList returns the current combo list
func (m *ActorBattleModule) GetCurComboList() []int {
	result := make([]int, len(m.CurComboList))
	copy(result, m.CurComboList)
	return result
}

// GetCurMontageList returns the current montage list
func (m *ActorBattleModule) GetCurMontageList() []int {
	result := make([]int, len(m.CurMontageList))
	copy(result, m.CurMontageList)
	return result
}
