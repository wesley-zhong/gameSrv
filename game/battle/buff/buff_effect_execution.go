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
	// Register buff effect execution and revert functions
	// Type IDs should match those in the cfg package
	// e.registerBuffEffectFunctions(cfg.ModAvatarProps.__ID__, e.buffEffectOptModAvatarProps, e.buffEffectOptModAvatarPropsRevert)
	// e.registerBuffEffectFunctions(cfg.ModAvatarPropsForEver.__ID__, e.buffEffectOptModAvatarPropsForever, nil)
	// ... more registrations
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
		// TODO: Get current attack target when ActorBattleModule is implemented
		// if curCharacter.ActorBattleModule.CurAttackTarget != nil {
		//     targets = append(targets, curCharacter.ActorBattleModule.CurAttackTarget)
		// }
	case cfg.BuffTargetEnum_TEAM_EXCLUDE_ME:
		// TODO: Get team members excluding self
	case cfg.BuffTargetEnum_TEAM_INCLUDE_ME:
		// TODO: Get all team members
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
	// TODO: Implement when ModAvatarProps type is available
	return false
}

// buffEffectOptModAvatarPropsForever modifies avatar properties permanently
func (e *FYBuffEffectExecution) buffEffectOptModAvatarPropsForever(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when ModAvatarPropsForEver type is available
	return false
}

// buffEffectOptModAvatarPropsRevert reverts avatar property modifications
func (e *FYBuffEffectExecution) buffEffectOptModAvatarPropsRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	// TODO: Implement when ModAvatarProps type is available
	return false
}

// buffEffectOptStartBuff starts a new buff
func (e *FYBuffEffectExecution) buffEffectOptStartBuff(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when StartNewBuff type is available
	return false
}

// buffEffectOptRemoveBuff removes buffs
func (e *FYBuffEffectExecution) buffEffectOptRemoveBuff(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when RemoveBuff type is available
	return false
}

// buffEffectOptModAvatarStatus modifies avatar status
func (e *FYBuffEffectExecution) buffEffectOptModAvatarStatus(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when ModAvatarStatus type is available
	return false
}

// buffEffectOptModAvatarStatusRevert reverts avatar status modifications
func (e *FYBuffEffectExecution) buffEffectOptModAvatarStatusRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	// TODO: Implement when ModAvatarStatus type is available
	return false
}

// buffEffectOptExecuteAttackData executes attack data
func (e *FYBuffEffectExecution) buffEffectOptExecuteAttackData(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when DoAttackData type is available
	return false
}

// buffEffectOptBuffDoHurt deals damage
func (e *FYBuffEffectExecution) buffEffectOptBuffDoHurt(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when BuffDoHurt type is available
	return false
}

// buffEffectOptExecuteAbnormalFormulation executes abnormal damage formulation
func (e *FYBuffEffectExecution) buffEffectOptExecuteAbnormalFormulation(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when DoAbnormalDamageFormulation type is available
	return false
}

// buffEffectOptModAttackDataProps modifies attack data properties
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataProps(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when ModAttackDataProps type is available
	return false
}

// buffEffectOptModAttackDataPropsRevert reverts attack data property modifications
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	// TODO: Implement when ModAttackDataProps type is available
	return false
}

// buffEffectOptModAttackDataPropsByTag modifies attack data properties by tag
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsByTag(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	exLayer int,
	reason int,
) bool {
	// TODO: Implement when ModAttackDataPropsByTag type is available
	return false
}

// buffEffectOptModAttackDataPropsByTagRevert reverts attack data property modifications by tag
func (e *FYBuffEffectExecution) buffEffectOptModAttackDataPropsByTagRevert(
	buff *FYBuff,
	target *actors.Creature,
	buffEffect *cfg.BuffEffect,
	revertLayer int,
) bool {
	// TODO: Implement when ModAttackDataPropsByTag type is available
	return false
}
