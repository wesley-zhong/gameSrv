package buff

import (
	"gameSrv/cnfGen/cfg"
	"sync"
)

// CheckBuffOnTriggerFunction checks conditions on buff trigger
type CheckBuffOnTriggerFunction func(param interface{}, event *cfg.TriggerEvent) bool

var (
	buffTriggerFunctions map[int]CheckBuffOnTriggerFunction
	initTriggerOnce      sync.Once
)

// CheckEvent checks if event should trigger
func CheckEvent(triggerEventParam interface{}, triggerEvent *cfg.TriggerEvent) bool {
	initTriggerOnce.Do(initTriggerFunctions)
	checkFunc := buffTriggerFunctions[int(triggerEvent.EventId)]
	if checkFunc == nil {
		return false
	}
	return checkFunc(triggerEventParam, triggerEvent)
}

// initTriggerFunctions initializes trigger check functions
func initTriggerFunctions() {
	if buffTriggerFunctions != nil {
		return
	}

	buffTriggerFunctions = make(map[int]CheckBuffOnTriggerFunction)

	// Register trigger event check functions
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_ENTER_TEAM] = onAvatarEnterTeamTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_LEAVE_TEAM] = onAvatarLeaveTeamTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_PROPS_CHANGE] = onAvatarPropertyChangeIntoTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_PROPS_CHANGE_RATE] = onAvatarPropertyChangeIntoTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_STATE_CHANGE] = onAvatarStatusChangeIntoTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_BUFF_MOUNTED] = onOtherBuffMountedTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_MONTAGE_ID] = onMtgIdAttackTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_MONTAGE_END] = onMontageEndTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_MONTAGE_TAG] = onMtgTagAttackTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_HITED] = onHitedTargetTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_HITED_SKILLTYPE] = onHitedSkillTypeTargetTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_DO_DAMGE] = onDoDamageTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_DO_DAMGE_BY_SKILLTYPE] = onDoDamageBySkillTypeTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_DO_HEAL] = onDoDamageTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_DO_CRITICAL_DAMAGE] = onDoCriticalDamageTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_DO_CRITICAL_DAMAGE_BY_SKILLTYPE] = onDoCriticalDamageBySkillTypeTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_BE_HITED] = onBeHitedTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_BE_DAMAGED] = onBeDamagedTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_BE_HEAL] = onBeDamagedTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_CHARACTER_ACTION_CONSUME_ABNORMAL] = onActionSubAbnormal
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_CHARACTER_ACTION_CONSUME_ABNORMAL_ALL] = onActionSubAbnormal
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_ABNORMAL_LAYER_CLOSING] = onAbnormalLayerClosing
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_EXECUTION_END] = onExecutionEndTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_Kill_MONSTER_TARGET] = onKilledMonsterTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_PERFECT_DODGE] = onDefaultTrueTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_PARRY_BEGIN] = onDefaultTrueTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_SHIP_FIRE] = onShipFiredTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_SHIP_FIRE_EFFECT] = onShipFiredBulletTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_CLIP_CHANGE] = onClipChangeTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_CLIP_RELOAD] = onClipReloadTrigger
	buffTriggerFunctions[cfg.BuffTriggerEventEnum_SHIP_SPEED_UP] = onShipSpeedUpTrigger
}

// onAvatarEnterTeamTrigger checks enter team condition
func onAvatarEnterTeamTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	// TODO: Implement when TeamTriggerEvent type is available
	return false
}

// onAvatarLeaveTeamTrigger checks leave team condition
func onAvatarLeaveTeamTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	// TODO: Implement when TeamTriggerEvent type is available
	return false
}

// onAvatarPropertyChangeIntoTrigger checks property change condition
func onAvatarPropertyChangeIntoTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	propsChangeParam, ok := param.(*PropsChangeBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when PropChangeTriggerEvent type is available
	// Check if property matches and condition is met
	_ = propsChangeParam
	return false
}

// onAvatarStatusChangeIntoTrigger checks status change condition
func onAvatarStatusChangeIntoTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	stateChangeParam, ok := param.(*StateChangeBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when StatusChangeTriggerEvent type is available
	_ = stateChangeParam
	return false
}

// onOtherBuffMountedTrigger checks other buff mounted condition
func onOtherBuffMountedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	buffMountedParam, ok := param.(*BuffMountedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when OtherBuffMountedTriggerEvent type is available
	_ = buffMountedParam
	return false
}

// onHitedTargetTrigger checks hit target condition
func onHitedTargetTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	hitedParam, ok := param.(*HitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when HitedTriggerEvent type is available
	_ = hitedParam
	return false
}

// onHitedSkillTypeTargetTrigger checks skill type hit condition
func onHitedSkillTypeTargetTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	hitedParam, ok := param.(*HitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when HitedSkillTypeTriggerEvent type is available
	_ = hitedParam
	return false
}

// onExecutionEndTrigger checks execution end condition
func onExecutionEndTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	executionParam, ok := param.(*ExecutionEndParam)
	if !ok {
		return false
	}

	// TODO: Implement when ExecutionEnd type is available
	_ = executionParam
	return false
}

// onAbnormalLayerClosing checks abnormal layer closing condition
func onAbnormalLayerClosing(param interface{}, event *cfg.TriggerEvent) bool {
	abnormalParam, ok := param.(*AbnormalLayerClosing)
	if !ok {
		return false
	}

	// TODO: Implement when DoAbnormalClosing type is available
	_ = abnormalParam
	return false
}

// onActionSubAbnormal checks action sub abnormal condition
func onActionSubAbnormal(param interface{}, event *cfg.TriggerEvent) bool {
	actionParam, ok := param.(*ActionSubAbnormalEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ActionSubAbnormalEvent type is available
	_ = actionParam
	return false
}

// onDoDamageTrigger checks damage dealing condition
func onDoDamageTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	damageParam, ok := param.(*DodamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoDamageTriggerEvent type is available
	_ = damageParam
	return false
}

// onDoDamageBySkillTypeTrigger checks skill type damage condition
func onDoDamageBySkillTypeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	damageParam, ok := param.(*DodamageBySkillTypeBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoDamageBySkillTypeTriggerEvent type is available
	_ = damageParam
	return false
}

// onDoCriticalDamageTrigger checks critical damage condition
func onDoCriticalDamageTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	criticalParam, ok := param.(*DoCriticalBySkillTypeDamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoCriticalDamageTriggerEvent type is available
	_ = criticalParam
	return false
}

// onDoCriticalDamageBySkillTypeTrigger checks skill type critical damage condition
func onDoCriticalDamageBySkillTypeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	criticalParam, ok := param.(*DoCriticalBySkillTypeDamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoCriticalDamageBySkillTypeTriggerEvent type is available
	_ = criticalParam
	return false
}

// onBeHitedTrigger checks being hit condition
func onBeHitedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	beHitedParam, ok := param.(*BeHitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when BeHitedTriggerEvent type is available
	_ = beHitedParam
	return false
}

// onBeDamagedTrigger checks being damaged condition
func onBeDamagedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	beDamagedParam, ok := param.(*BeDamagedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when BeDamageTriggerEvent type is available
	_ = beDamagedParam
	return false
}

// onShipFiredTrigger checks ship fire condition
func onShipFiredTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	shipFireParam, ok := param.(*ShipFireEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ShipFireEvent type is available
	_ = shipFireParam
	return false
}

// onShipFiredBulletTrigger checks ship fire bullet condition
func onShipFiredBulletTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	shipEffectParam, ok := param.(*ShipFireEffectEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ShipFireEffectEvent type is available
	_ = shipEffectParam
	return false
}

// onClipChangeTrigger checks clip change condition
func onClipChangeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	clipParam, ok := param.(*ClipCapacityChange)
	if !ok {
		return false
	}

	// TODO: Implement when ClipCapcityChangeEvent type is available
	_ = clipParam
	return false
}

// onClipReloadTrigger checks clip reload condition
func onClipReloadTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}

// onShipSpeedUpTrigger checks ship speed up condition
func onShipSpeedUpTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}

// onMtgIdAttackTrigger checks montage ID condition
func onMtgIdAttackTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MtgIdBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MtgIdTriggerEvent type is available
	_ = mtgParam
	return false
}

// onKilledMonsterTrigger checks killed monster condition
func onKilledMonsterTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	killedParam, ok := param.(*KilledMonsterEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when KilledMonsterEvent type is available
	_ = killedParam
	return false
}

// onMontageEndTrigger checks montage end condition
func onMontageEndTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MontageEndBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MontageEndTriggerEvent type is available
	_ = mtgParam
	return false
}

// onMtgTagAttackTrigger checks montage tag condition
func onMtgTagAttackTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MtgIdBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MtgTagTriggerEvent type is available
	_ = mtgParam
	return false
}

// onDefaultTrueTrigger always returns true
func onDefaultTrueTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}
