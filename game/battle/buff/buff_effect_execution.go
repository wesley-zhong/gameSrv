package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/battle"
	"gameSrv/game/gamedata"
	"gameSrv/pkg/actors"
	"sync"
)

// BuffEffectExeFunction executes buff effect
type BuffEffectExeFunction func(buff *FYBuff, target *actors.Creature, buffEffect *cfg.BuffEffect, exLayer int, reason int) bool

// BuffEffectRevertFunction reverts buff effect
type BuffEffectRevertFunction func(buff *FYBuff, target *actors.Creature, buffEffect *cfg.BuffEffect, revertLayer int) bool

// FYBuffEffectExecution handles buff effect execution
type FYBuffEffectExecution struct {
	buffEffectExeFunctions    map[int]BuffEffectExeFunction
	buffEffectRevertFunctions map[int]BuffEffectRevertFunction
}

var (
	// Global instance for effect execution
	effectExecutionInstance *FYBuffEffectExecution
	effectExecutionOnce     sync.Once
)

// GetEffectExecutionInstance returns the singleton instance
func GetEffectExecutionInstance() *FYBuffEffectExecution {
	effectExecutionOnce.Do(func() {
		effectExecutionInstance = &FYBuffEffectExecution{
			buffEffectExeFunctions:    make(map[int]BuffEffectExeFunction),
			buffEffectRevertFunctions: make(map[int]BuffEffectRevertFunction),
		}
		effectExecutionInstance.initBuffEffectFunctions()
	})
	return effectExecutionInstance
}

// initBuffEffectFunctions initializes buff effect functions
func (e *FYBuffEffectExecution) initBuffEffectFunctions() {
	// Property modification effects
	e.registerBuffEffectFunctions(int(cfg.TypeId_ModAvatarProps), e.buffEffectOptModAvatarProps, e.buffEffectOptModAvatarPropsRevert)
	e.registerBuffEffectFunctions(int(cfg.TypeId_ModAvatarPropsForEver), e.buffEffectOptModAvatarPropsForever, nil)

	// Status modification effects
	e.registerBuffEffectFunctions(int(cfg.TypeId_ModAvatarStatus), e.buffEffectOptModAvatarStatus, e.buffEffectOptModAvatarStatusRevert)

	// Buff management effects
	e.registerBuffEffectFunctions(int(cfg.TypeId_StartNewBuff), e.buffEffectOptStartBuff, nil)
	e.registerBuffEffectFunctions(int(cfg.TypeId_RemoveBuff), e.buffEffectOptRemoveBuff, nil)

	// Attack and damage effects
	e.registerBuffEffectFunctions(int(cfg.TypeId_DoAttackData), e.buffEffectOptExecuteAttackData, nil)
	e.registerBuffEffectFunctions(int(cfg.TypeId_BuffDoHurt), e.buffEffectOptBuffDoHurt, nil)
	e.registerBuffEffectFunctions(int(cfg.TypeId_DoAbnormalDamageFormulation), e.buffEffectOptExecuteAbnormalFormulation, nil)

	// Attack data modification effects
	e.registerBuffEffectFunctions(int(cfg.TypeId_ModAttackDataProps), e.buffEffectOptModAttackDataProps, e.buffEffectOptModAttackDataPropsRevert)
	e.registerBuffEffectFunctions(int(cfg.TypeId_ModAttackDataPropsByTag), e.buffEffectOptModAttackDataPropsByTag, e.buffEffectOptModAttackDataPropsByTagRevert)
}

// registerBuffEffectFunctions registers a buff effect function
func (e *FYBuffEffectExecution) registerBuffEffectFunctions(
	effectType int,
	exeFunction BuffEffectExeFunction,
	revertFunction BuffEffectRevertFunction,
) {
	if _, exists := e.buffEffectExeFunctions[effectType]; exists {
		panic("BuffEffectExeFunctions already contains EffectType: " + string(rune(effectType)))
	}
	e.buffEffectExeFunctions[effectType] = exeFunction
	if revertFunction != nil {
		e.buffEffectRevertFunctions[effectType] = revertFunction
	}
}

// ExeBuffEffect executes buff effects
func (e *FYBuffEffectExecution) ExeBuffEffect(
	buff *FYBuff,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	if buffEffect == nil || buffEffect.BuffEffectOpt == nil {
		return false
	}

	// Get type ID from interface{}
	var effectType int
	if typeGetter, ok := buffEffect.BuffEffectOpt.(interface{ GetTypeId() int32 }); ok {
		effectType = int(typeGetter.GetTypeId())
	} else {
		return false
	}

	buffEffectFunc := e.buffEffectExeFunctions[effectType]
	if buffEffectFunc == nil {
		return false
	}

	// Check cooldown
	expiredTime := buff.StartEffectTime + int64(buffEffect.CoolDown)
	now := battle.MilliSeconds()
	if buffEffect.CoolDown > 0 && now < expiredTime {
		return false
	}

	// Check random
	if !e.checkBuffEffectRandom(buff, buffEffect) {
		return false
	}

	buff.SetStartEffectTime(now)

	// Select targets and execute effect
	targets := e.selectEffectTargets(buff, buffEffect)
	changed := false

	for _, target := range targets {
		if buffEffectFunc(buff, target, buffEffect, exLayer, reason) {
			changed = true
		}
	}

	return changed
}

// ExeBuffEffects executes multiple buff effects
func (e *FYBuffEffectExecution) ExeBuffEffects(
	buff *FYBuff,
	buffEffectIds []int,
	exLayer int,
	reason int,
) bool {
	executed := false
	for _, effectId := range buffEffectIds {
		if gamedata.Tables == nil || gamedata.Tables.TbBuffEffect == nil {
			continue
		}
		buffEffect := gamedata.Tables.TbBuffEffect.Get(int32(effectId))
		if buffEffect == nil || buffEffect.EffectID == 0 {
			continue
		}

		if e.ExeBuffEffect(buff, buffEffect, exLayer, reason) {
			executed = true
		}
	}
	return executed
}

// checkBuffEffectRandom checks if effect should execute based on random chance
func (e *FYBuffEffectExecution) checkBuffEffectRandom(buff *FYBuff, buffEffect *cfg.BuffEffect) bool {
	// ExePercent field doesn't exist in current cfg structure, default to always execute
	// TODO: Add ExePercent field to BuffEffect config if needed
	return true
}

// RevertBuffAllEffectsWithLayer reverts all buff effects for a given layer
func (e *FYBuffEffectExecution) RevertBuffAllEffectsWithLayer(buff *FYBuff, removeLayer int) bool {
	if len(buff.ImmediateExeTriggerEvents) == 0 {
		return false
	}

	changed := false

	for _, triggerEvent := range buff.ImmediateExeTriggerEvents {
		for _, effectId := range triggerEvent.EffectIds {
			if gamedata.Tables == nil || gamedata.Tables.TbBuffEffect == nil {
				continue
			}
			buffEffect := gamedata.Tables.TbBuffEffect.Get(int32(effectId))
			if buffEffect == nil || buffEffect.EffectID == 0 {
				continue
			}

			// Get type ID from interface{}
			var effectType int
			if typeGetter, ok := buffEffect.BuffEffectOpt.(interface{ GetTypeId() int32 }); ok {
				effectType = int(typeGetter.GetTypeId())
			} else {
				continue
			}

			revertFunc := e.buffEffectRevertFunctions[effectType]
			if revertFunc == nil {
				continue
			}

			targets := e.selectEffectTargets(buff, buffEffect)
			if len(targets) == 0 {
				continue
			}

			for _, target := range targets {
				if revertFunc(buff, target, buffEffect, removeLayer) {
					changed = true
				}
			}
		}
	}

	return changed
}

// selectEffectTargets selects targets for buff effect
func (e *FYBuffEffectExecution) selectEffectTargets(buff *FYBuff, buffEffect *cfg.BuffEffect) []*actors.Creature {
	targets := make([]*actors.Creature, 0)

	if buff.HolderActor == nil {
		return targets
	}

	curCharacter := buff.HolderActor

	// Handle effect target types
	switch buffEffect.EffectTarget {
	case cfg.BuffTargetEnum_MYSELF:
		targets = append(targets, curCharacter)
	case cfg.BuffTargetEnum_ATTACK_TARGET:
		// Get current attack target from ActorBattleModule
		if curCharacter.ActorBattleModule != nil {
			if module, ok := curCharacter.ActorBattleModule.(interface {
				GetCurAttackTarget() *actors.Creature
			}); ok {
				attackTarget := module.GetCurAttackTarget()
				if attackTarget != nil {
					targets = append(targets, attackTarget)
				}
			}
		}
	case cfg.BuffTargetEnum_TEAM_EXCLUDE_ME:
		// TODO: Get team members excluding self
		// This requires team system implementation
	case cfg.BuffTargetEnum_TEAM_INCLUDE_ME:
		// TODO: Get all team members
		// This requires team system implementation
	}

	return targets
}

// calcValue calculates value based on operation type
func (e *FYBuffEffectExecution) calcValue(
	dataType int, // cfg.BuffOpt
	origin, curValue, value, exLayer int,
) int {
	switch dataType {
	case cfg.BuffOpt_ADD:
		return value * exLayer
	case cfg.BuffOpt_ADD_MULTI_10_000:
		return (origin * value * exLayer) / 10000
	case cfg.BuffOpt_ADD_MIN:
		if curValue > value {
			return value - curValue
		}
		return 0
	case cfg.BuffOpt_ADD_MAX:
		if curValue > value {
			return 0
		}
		return value - curValue
	case cfg.BuffOpt_ADD_MULTI_10_000_MIX:
		addVal := (origin * value * exLayer) / 10000
		if curValue > addVal {
			return addVal - curValue
		}
		return 0
	case cfg.BuffOpt_ADD_MULTI_10_000_MAX:
		addVal := (origin * value * exLayer) / 10000
		if curValue > addVal {
			return 0
		}
		return addVal - curValue
	default:
		return 0
	}
}

// buffEffectOptModAvatarProps modifies avatar properties
func (e *FYBuffEffectExecution) buffEffectOptModAvatarProps(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAvatarProps)
	if !ok || target == nil {
		return false
	}

	prop := int(effectOpt.AvatarProp)
	opt := int(effectOpt.OPT)
	value := int(effectOpt.Value)

	if target.BattleProps == nil {
		return false
	}

	// Get current property value
	originValue := target.GetProperty(int32(prop))
	curValue := originValue

	// Calculate the change value using calcValue
	changeValue := e.calcValue(opt, originValue, curValue, value, exLayer)
	if changeValue == 0 {
		return false
	}

	// Apply the change
	newValue := curValue + changeValue
	target.SetProperty(int32(prop), int32(newValue), true)

	// Track the modification for reversion
	effectId := int(buffEffect.EffectID)
	if buff.EffectAddedData == nil {
		buff.EffectAddedData = make(map[int]int)
	}
	buff.EffectAddedData[effectId] = changeValue

	return true
}

// buffEffectOptModAvatarPropsForever modifies avatar properties permanently
func (e *FYBuffEffectExecution) buffEffectOptModAvatarPropsForever(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAvatarPropsForEver)
	if !ok || target == nil {
		return false
	}

	prop := int(effectOpt.AvatarProp)
	opt := int(effectOpt.OPT)
	value := int(effectOpt.Value)

	if target.BattleProps == nil {
		return false
	}

	// Get current base value for permanent modification
	originValue := target.GetBaseProps(int32(prop))
	curValue := target.GetProperty(int32(prop))

	// Calculate the change value using calcValue
	changeValue := e.calcValue(opt, originValue, curValue, value, exLayer)
	if changeValue == 0 {
		return false
	}

	// For permanent modification, we modify both base and current value
	// Using SetAllProps which sets both base and current
	newValue := curValue + changeValue
	target.BattleProps.SetAllProps(prop, newValue)

	return true
}

// buffEffectOptModAvatarPropsRevert reverts avatar property modifications
func (e *FYBuffEffectExecution) buffEffectOptModAvatarPropsRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAvatarProps)
	if !ok || target == nil {
		return false
	}

	prop := int(effectOpt.AvatarProp)
	effectId := int(buffEffect.EffectID)

	if target.BattleProps == nil {
		return false
	}

	// Get the previously stored change value
	_, exists := buff.EffectAddedData[effectId]
	if !exists {
		return false
	}

	// Calculate the reversion value based on the revert layer
	opt := int(effectOpt.OPT)
	value := int(effectOpt.Value)

	originValue := target.GetProperty(int32(prop))
	curValue := originValue

	// Calculate the change to revert (negative of original change)
	changeToRevert := -e.calcValue(opt, originValue, curValue, value, revertLayer)

	// Apply the reversion
	newValue := curValue + changeToRevert
	target.SetProperty(int32(prop), int32(newValue), true)

	// Remove the tracking entry if fully reverted
	delete(buff.EffectAddedData, effectId)

	return true
}

// buffEffectOptStartBuff starts a new buff
func (e *FYBuffEffectExecution) buffEffectOptStartBuff(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.StartNewBuff)
	if !ok || target == nil {
		return false
	}

	addCount := int(effectOpt.AddCount)
	if addCount <= 0 {
		addCount = 1
	}

	for _, buffId := range effectOpt.BuffIds {
		for i := 0; i < addCount; i++ {
			target.AddBuff(int(buffId), buff.CasterActor, 0, buff.ExParam, buff.StartFromSystem)
		}
	}

	return true
}

// buffEffectOptRemoveBuff removes buffs
func (e *FYBuffEffectExecution) buffEffectOptRemoveBuff(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.RemoveBuff)
	if !ok || target == nil {
		return false
	}

	// ByClass determines the removal method:
	// 0 = by config ID (Values contains buff IDs to remove)
	// 1 = by class (Values contains class IDs, first value is classID)
	// 2 = by subclass (Values contains [classID, subClassID])
	switch effectOpt.ByClass {
	case 0: // By config ID
		for _, buffId := range effectOpt.Values {
			target.RemoveBuffByConfId(int(buffId), reason)
		}
	case 1: // By class
		// Use ActorBuffModule to remove by class
		if module, ok := target.ActorBuffModule.(*ActorBuffModule); ok {
			for _, classId := range effectOpt.Values {
				module.RemoveBuffByClass(int(classId), reason)
			}
		}
	case 2: // By subclass
		// Use ActorBuffModule to remove by subclass
		if module, ok := target.ActorBuffModule.(*ActorBuffModule); ok && len(effectOpt.Values) >= 2 {
			module.RemoveBuffBySubClass(int(effectOpt.Values[0]), int(effectOpt.Values[1]), reason)
		}
	}

	return true
}

// buffEffectOptModAvatarStatus modifies avatar status
func (e *FYBuffEffectExecution) buffEffectOptModAvatarStatus(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAvatarStatus)
	if !ok || target == nil {
		return false
	}

	state := int32(effectOpt.Status)
	opt := effectOpt.Opt

	if target.ActorBattleModule == nil {
		return false
	}

	// Apply state modification based on operation type
	switch opt {
	case cfg.StatusOpt_SET:
		// Set state (remove then add)
		target.ActorBattleModule.RemoveState(state, 999, true)
		target.ActorBattleModule.AddState(state, true, 1)
	case cfg.StatusOpt_REMOVE:
		// Remove state
		target.ActorBattleModule.RemoveState(state, exLayer, true)
	}

	// Track the modification for reversion
	effectId := int(buffEffect.EffectID)
	if buff.EffectAddedData == nil {
		buff.EffectAddedData = make(map[int]int)
	}
	buff.EffectAddedData[effectId] = exLayer

	return true
}

// buffEffectOptModAvatarStatusRevert reverts avatar status modifications
func (e *FYBuffEffectExecution) buffEffectOptModAvatarStatusRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAvatarStatus)
	if !ok || target == nil {
		return false
	}

	state := int32(effectOpt.Status)
	opt := effectOpt.Opt

	if target.ActorBattleModule == nil {
		return false
	}

	// Revert state modification based on operation type
	switch opt {
	case cfg.StatusOpt_SET:
		// To revert set, remove the state
		target.ActorBattleModule.RemoveState(state, revertLayer, true)
	case cfg.StatusOpt_REMOVE:
		// To revert remove, add the state back
		target.ActorBattleModule.AddState(state, true, revertLayer)
	}

	// Remove the tracking entry
	effectId := int(buffEffect.EffectID)
	delete(buff.EffectAddedData, effectId)

	return true
}

// buffEffectOptExecuteAttackData executes attack data
func (e *FYBuffEffectExecution) buffEffectOptExecuteAttackData(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.DoAttackData)
	if !ok || target == nil {
		return false
	}

	// TODO: Implement attack data execution when ActorBattleModule is available
	_ = effectOpt.AttackDataId
	_ = exLayer

	return true
}

// buffEffectOptBuffDoHurt deals damage
func (e *FYBuffEffectExecution) buffEffectOptBuffDoHurt(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.BuffDoHurt)
	if !ok || target == nil {
		return false
	}

	// TODO: Implement damage dealing when damage system is available
	_ = effectOpt.DamageSource
	_ = effectOpt.DamageType
	_ = effectOpt.HideDmgShow
	_ = exLayer

	return true
}

// buffEffectOptExecuteAbnormalFormulation executes abnormal damage formulation
func (e *FYBuffEffectExecution) buffEffectOptExecuteAbnormalFormulation(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.DoAbnormalDamageFormulation)
	if !ok || target == nil {
		return false
	}

	// TODO: Implement abnormal damage formulation when damage system is available
	// Store the formulation ID for later use in buff.EffectHurtMap
	_ = effectOpt.FormulationId
	_ = effectOpt.DamgeType
	_ = exLayer

	return true
}

// buffEffectOptModAttackDataProps modifies attack data properties
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataProps(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAttackDataProps)
	if !ok || target == nil {
		return false
	}

	// Use ActorSkillModule to modify attack data properties
	// This would need to be called when buff is applied to holder
	for _, atkDataId := range effectOpt.AttatckDataIds {
		value := e.calcValue(int(effectOpt.OPT), 0, 0, int(effectOpt.Value), exLayer)
		// Would call: target.ActorSkillModule.SetAtkDataProps(atkDataId, effectOpt.SkillProp, value)
		_ = atkDataId
		_ = effectOpt.SkillProp
		_ = value
	}

	return true
}

// buffEffectOptModAttackDataPropsRevert reverts attack data property modifications
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAttackDataProps)
	if !ok || target == nil {
		return false
	}

	// Revert attack data property modifications
	// This would reset properties to original values
	for _, atkDataId := range effectOpt.AttatckDataIds {
		// Would call: target.ActorSkillModule.SetAtkDataProps(atkDataId, effectOpt.SkillProp, 0)
		_ = atkDataId
		_ = effectOpt.SkillProp
	}
	_ = revertLayer

	return true
}

// buffEffectOptModAttackDataPropsByTag modifies attack data properties by tag
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsByTag(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAttackDataPropsByTag)
	if !ok || target == nil {
		return false
	}

	// TODO: Implement attack data property modification by tag when ActorBattleModule is available
	_ = effectOpt.SkillProp
	_ = effectOpt.OPT
	_ = effectOpt.Value
	_ = effectOpt.Tag
	_ = exLayer

	return true
}

// buffEffectOptModAttackDataPropsByTagRevert reverts attack data property modifications by tag
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsByTagRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	effectOpt, ok := buffEffect.BuffEffectOpt.(*cfg.ModAttackDataPropsByTag)
	if !ok || target == nil {
		return false
	}

	// TODO: Implement attack data property reversion by tag when ActorBattleModule is available
	_ = effectOpt.SkillProp
	_ = effectOpt.OPT
	_ = effectOpt.Value
	_ = effectOpt.Tag
	_ = revertLayer

	return true
}
