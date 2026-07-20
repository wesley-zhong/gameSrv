package buff

import (
	"gameSrv/cnfGen/cfg"
)

// CheckBuffOnTriggerFunction checks conditions on buff trigger
type CheckBuffOnTriggerFunction func(param interface{}, event *cfg.TriggerEvent) bool

// FYBuffTriggerProcess handles buff trigger processing
type FYBuffTriggerProcess struct {
	buffTriggerFunctions map[int]CheckBuffOnTriggerFunction
}

// NewFYBuffTriggerProcess creates a new FYBuffTriggerProcess
func NewFYBuffTriggerProcess() *FYBuffTriggerProcess {
	p := &FYBuffTriggerProcess{
		buffTriggerFunctions: make(map[int]CheckBuffOnTriggerFunction),
	}
	p.initTriggerFunctions()
	return p
}

// CheckEvent checks if event should trigger
func (p *FYBuffTriggerProcess) CheckEvent(triggerEventParam interface{}, triggerEvent *cfg.TriggerEvent) bool {
	checkFunc := p.buffTriggerFunctions[int(triggerEvent.EventId)]
	if checkFunc == nil {
		return false
	}
	return checkFunc(triggerEventParam, triggerEvent)
}

// initTriggerFunctions initializes trigger check functions
func (p *FYBuffTriggerProcess) initTriggerFunctions() {
	// Register trigger event check functions
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_ENTER_TEAM), p.onAvatarEnterTeamTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_LEAVE_TEAM), p.onAvatarLeaveTeamTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_PROPS_CHANGE), p.onAvatarPropertyChangeIntoTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_PROPS_CHANGE_RATE), p.onAvatarPropertyChangeIntoTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_STATE_CHANGE), p.onAvatarStatusChangeIntoTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_BUFF_MOUNTED), p.onOtherBuffMountedTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_MONTAGE_ID), p.onMtgIdAttackTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_MONTAGE_END), p.onMontageEndTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_MONTAGE_TAG), p.onMtgTagAttackTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_HITED), p.onHitedTargetTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_HITED_SKILLTYPE), p.onHitedSkillTypeTargetTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_DO_DAMGE), p.onDoDamageTrigger)
	p.registerCheckFunctions(26, p.onDoDamageBySkillTypeTrigger) // DO_DAMGE_BY_SKILLTYPE - define if needed
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_DO_HEAL), p.onDoDamageTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_DO_CRITICAL_DAMAGE), p.onDoCriticalDamageTrigger)
	p.registerCheckFunctions(27, p.onDoCriticalDamageBySkillTypeTrigger) // DO_CRITICAL_DAMAGE_BYSKILLTYPE - define if needed
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_BE_HITED), p.onBeHitedTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_BE_DAMAGED), p.onBeDamagedTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_BE_HEAL), p.onBeDamagedTrigger)
	p.registerCheckFunctions(28, p.onActionSubAbnormal) // CHARACTER_ACTION_CONSUME_ABNORMAL - define if needed
	p.registerCheckFunctions(29, p.onActionSubAbnormal) // CHARACTER_ACTION_CONSUME_ABNORMAL_ALL - define if needed
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_ABNORMAL_LAYER_CLOSING), p.onAbnormalLayerClosing)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_EXECUTION_END), p.onExecutionEndTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_Kill_MONSTER_TARGET), p.onKilledMonsterTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_PERFECT_DODGE), p.onDefaultTrueTrigger)
	p.registerCheckFunctions(int(cfg.BuffTriggerEventEnum_PARRY_BEGIN), p.onDefaultTrueTrigger)
	p.registerCheckFunctions(30, p.onShipFiredTrigger) // SHIP_FIRE - define if needed
	p.registerCheckFunctions(31, p.onShipFiredBulletTrigger) // SHIP_FIRE_EFFECT - define if needed
	p.registerCheckFunctions(32, p.onClipChangeTrigger) // CLIP_CHANGE - define if needed
	p.registerCheckFunctions(33, p.onClipReloadTrigger) // CLIP_REALOD - define if needed
	p.registerCheckFunctions(34, p.onShipSpeedUpTrigger) // SHIP_SPEED_UP - define if needed
}

// registerCheckFunctions registers a trigger check function
func (p *FYBuffTriggerProcess) registerCheckFunctions(
	eventEnum int,
	checkFunc CheckBuffOnTriggerFunction,
) {
	p.buffTriggerFunctions[eventEnum] = checkFunc
}

// onAvatarEnterTeamTrigger checks enter team condition
func (p *FYBuffTriggerProcess) onAvatarEnterTeamTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	// TODO: Implement when TeamTriggerEvent type is available
	return false
}

// onAvatarLeaveTeamTrigger checks leave team condition
func (p *FYBuffTriggerProcess) onAvatarLeaveTeamTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	// TODO: Implement when TeamTriggerEvent type is available
	return false
}

// onAvatarPropertyChangeIntoTrigger checks property change condition
func (p *FYBuffTriggerProcess) onAvatarPropertyChangeIntoTrigger(param interface{}, event *cfg.TriggerEvent) bool {
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
func (p *FYBuffTriggerProcess) onAvatarStatusChangeIntoTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	stateChangeParam, ok := param.(*StateChangeBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when StatusChangeTriggerEvent type is available
	_ = stateChangeParam
	return false
}

// onOtherBuffMountedTrigger checks other buff mounted condition
func (p *FYBuffTriggerProcess) onOtherBuffMountedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	buffMountedParam, ok := param.(*BuffMountedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when OtherBuffMountedTriggerEvent type is available
	_ = buffMountedParam
	return false
}

// onHitedTargetTrigger checks hit target condition
func (p *FYBuffTriggerProcess) onHitedTargetTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	hitedParam, ok := param.(*HitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when HitedTriggerEvent type is available
	_ = hitedParam
	return false
}

// onHitedSkillTypeTargetTrigger checks skill type hit condition
func (p *FYBuffTriggerProcess) onHitedSkillTypeTargetTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	hitedParam, ok := param.(*HitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when HitedSkillTypeTriggerEvent type is available
	_ = hitedParam
	return false
}

// onExecutionEndTrigger checks execution end condition
func (p *FYBuffTriggerProcess) onExecutionEndTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	executionParam, ok := param.(*ExecutionEndParam)
	if !ok {
		return false
	}

	// TODO: Implement when ExecutionEnd type is available
	_ = executionParam
	return false
}

// onAbnormalLayerClosing checks abnormal layer closing condition
func (p *FYBuffTriggerProcess) onAbnormalLayerClosing(param interface{}, event *cfg.TriggerEvent) bool {
	abnormalParam, ok := param.(*AbnormalLayerClosing)
	if !ok {
		return false
	}

	// TODO: Implement when DoAbnormalClosing type is available
	_ = abnormalParam
	return false
}

// onActionSubAbnormal checks action sub abnormal condition
func (p *FYBuffTriggerProcess) onActionSubAbnormal(param interface{}, event *cfg.TriggerEvent) bool {
	actionParam, ok := param.(*ActionSubAbnormalEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ActionSubAbnormalEvent type is available
	_ = actionParam
	return false
}

// onDoDamageTrigger checks damage dealing condition
func (p *FYBuffTriggerProcess) onDoDamageTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	damageParam, ok := param.(*DodamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoDamageTriggerEvent type is available
	_ = damageParam
	return false
}

// onDoDamageBySkillTypeTrigger checks skill type damage condition
func (p *FYBuffTriggerProcess) onDoDamageBySkillTypeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	damageParam, ok := param.(*DodamageBySkillTypeBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoDamageBySkillTypeTriggerEvent type is available
	_ = damageParam
	return false
}

// onDoCriticalDamageTrigger checks critical damage condition
func (p *FYBuffTriggerProcess) onDoCriticalDamageTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	criticalParam, ok := param.(*DoCriticalBySkillTypeDamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoCriticalDamageTriggerEvent type is available
	_ = criticalParam
	return false
}

// onDoCriticalDamageBySkillTypeTrigger checks skill type critical damage condition
func (p *FYBuffTriggerProcess) onDoCriticalDamageBySkillTypeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	criticalParam, ok := param.(*DoCriticalBySkillTypeDamageBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when DoCriticalDamageBySkillTypeTriggerEvent type is available
	_ = criticalParam
	return false
}

// onBeHitedTrigger checks being hit condition
func (p *FYBuffTriggerProcess) onBeHitedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	beHitedParam, ok := param.(*BeHitedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when BeHitedTriggerEvent type is available
	_ = beHitedParam
	return false
}

// onBeDamagedTrigger checks being damaged condition
func (p *FYBuffTriggerProcess) onBeDamagedTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	beDamagedParam, ok := param.(*BeDamagedBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when BeDamageTriggerEvent type is available
	_ = beDamagedParam
	return false
}

// onShipFiredTrigger checks ship fire condition
func (p *FYBuffTriggerProcess) onShipFiredTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	shipFireParam, ok := param.(*ShipFireEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ShipFireEvent type is available
	_ = shipFireParam
	return false
}

// onShipFiredBulletTrigger checks ship fire bullet condition
func (p *FYBuffTriggerProcess) onShipFiredBulletTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	shipEffectParam, ok := param.(*ShipFireEffectEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when ShipFireEffectEvent type is available
	_ = shipEffectParam
	return false
}

// onClipChangeTrigger checks clip change condition
func (p *FYBuffTriggerProcess) onClipChangeTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	clipParam, ok := param.(*ClipCapacityChange)
	if !ok {
		return false
	}

	// TODO: Implement when ClipCapcityChangeEvent type is available
	_ = clipParam
	return false
}

// onClipReloadTrigger checks clip reload condition
func (p *FYBuffTriggerProcess) onClipReloadTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}

// onShipSpeedUpTrigger checks ship speed up condition
func (p *FYBuffTriggerProcess) onShipSpeedUpTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}

// onMtgIdAttackTrigger checks montage ID condition
func (p *FYBuffTriggerProcess) onMtgIdAttackTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MtgIdBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MtgIdTriggerEvent type is available
	_ = mtgParam
	return false
}

// onKilledMonsterTrigger checks killed monster condition
func (p *FYBuffTriggerProcess) onKilledMonsterTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	killedParam, ok := param.(*KilledMonsterEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when KilledMonsterEvent type is available
	_ = killedParam
	return false
}

// onMontageEndTrigger checks montage end condition
func (p *FYBuffTriggerProcess) onMontageEndTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MontageEndBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MontageEndTriggerEvent type is available
	_ = mtgParam
	return false
}

// onMtgTagAttackTrigger checks montage tag condition
func (p *FYBuffTriggerProcess) onMtgTagAttackTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	mtgParam, ok := param.(*MtgIdBuffTriggerEventParam)
	if !ok {
		return false
	}

	// TODO: Implement when MtgTagTriggerEvent type is available
	_ = mtgParam
	return false
}

// onDefaultTrueTrigger always returns true
func (p *FYBuffTriggerProcess) onDefaultTrueTrigger(param interface{}, event *cfg.TriggerEvent) bool {
	return true
}
