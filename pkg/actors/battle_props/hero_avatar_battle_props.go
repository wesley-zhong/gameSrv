package battle_props

import (
	"gameSrv/cnfGen/cfg"
)

// HeroAvatarBattleProps represents battle properties for hero avatars
type HeroAvatarBattleProps struct {
	*BattleProps
}

// NewHeroAvatarBattleProps creates a new HeroAvatarBattleProps
func NewHeroAvatarBattleProps(maxProps int) *HeroAvatarBattleProps {
	return &HeroAvatarBattleProps{
		BattleProps: NewBattleProps(),
	}
}

// NewHeroAvatarBattlePropsWithOwner creates a new HeroAvatarBattleProps with owner
func NewHeroAvatarBattlePropsWithOwner(maxProps int, owner BattlePropsOwner) *HeroAvatarBattleProps {
	return &HeroAvatarBattleProps{
		BattleProps: NewBattlePropsWithOwner(owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (hap *HeroAvatarBattleProps) RecalPropsFormula() {
	if hap.BattleProps == nil {
		return
	}
	hap.BattleProps.mu.Lock()
	defer hap.BattleProps.mu.Unlock()

	PropsUtil.CalcAttributes(hap.CurProps, hap.BaseProps)
}

// ResetProps resets properties
func (hap *HeroAvatarBattleProps) ResetProps() {
	if hap.BattleProps == nil || hap.Owner == nil {
		return
	}

	// Check if owner is a hero avatar actor
	// This requires type assertion to HeroAvatarActor
	// For now, we'll clear and let the avatar recalculate

	hap.ClearProps()

	// TODO: Call avatar.reCalAttributes() if owner is HeroAvatarActor
	// if heroAvatarActor, ok := hap.Owner.(*HeroAvatarActor); ok {
	//     avatar := heroAvatarActor.GetAvatar()
	//     if avatar != nil {
	//         avatar.ReCalAttributes()
	//     }
	// }
}

// CallReservedAttributes calls reserved attributes calculation
func (hap *HeroAvatarBattleProps) CallReservedAttributes(prop int, enums int) {
	globalParams := GlobalParamStorage.Get(enums)
	if globalParams != nil {
		currentValue := hap.GetProperty(prop)
		newValue := int(float64(currentValue) * (float64(globalParams.INT) / F_10_000))
		hap.SetPropertyWithRecalculate(prop, newValue, false)
	}
}

// Reborn handles reborn for hero avatar
func (hap *HeroAvatarBattleProps) Reborn() {
	// Set stamina to max stamina
	hap.SetPropertyWithRecalculate(cfg.HeroProp_Stamina, hap.GetProperty(cfg.HeroProp_MaxStamina), false)

	// Call reserved attributes for SP0, SP1 and ExSkillEnergy
	// TODO: define these constants from ParamKeyEnum
	hap.CallReservedAttributes(cfg.HeroProp_SP0, 0)           // ParamKeyEnum.DEATH_RETAIN_SP0_RATIO_VALUE
	hap.CallReservedAttributes(cfg.HeroProp_SP1, 0)           // ParamKeyEnum.DEATH_RETAIN_SP1_RATIO_VALUE
	hap.CallReservedAttributes(cfg.HeroProp_ExSkillEnergy, 0) // ParamKeyEnum.DEATH_RETAIN_EX_RATIO_VALUE

	// Special case: for config ID 1024 (Liliya), clear SP0 and SP1
	if hap.Owner != nil && hap.Owner.GetConfigId() == 1024 {
		hap.SetPropertyWithRecalculate(cfg.HeroProp_SP0, 0, false)
		hap.SetPropertyWithRecalculate(cfg.HeroProp_SP1, 0, false)
	}
}
