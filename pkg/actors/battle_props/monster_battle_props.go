package battle_props

import (
	"gameSrv/cnfGen/cfg"
)

// MonsterBattleProps represents battle properties for monsters
type MonsterBattleProps struct {
	*LevelBattleProps
}

// NewMonsterBattleProps creates a new MonsterBattleProps
func NewMonsterBattleProps(maxProps int, owner BattlePropsOwner) *MonsterBattleProps {
	return &MonsterBattleProps{
		LevelBattleProps: NewLevelBattleProps(maxProps, owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (mp *MonsterBattleProps) RecalPropsFormula() {
	if mp.LevelBattleProps != nil {
		mp.LevelBattleProps.RecalPropsFormula()
	}
}

// ResetProps resets properties for monster
func (mp *MonsterBattleProps) ResetProps() {
	if mp.LevelBattleProps != nil {
		mp.LevelBattleProps.ResetProps()
	}
	mp.InitCurHealth()
}

// InitCurHealth initializes current health
func (mp *MonsterBattleProps) InitCurHealth() {
	if mp.BattleProps == nil {
		return
	}

	mp.BattleProps.mu.Lock()
	defer mp.BattleProps.mu.Unlock()

	maxHealthIdx := cfg.HeroProp_MaxHealth
	healthIdx := cfg.HeroProp_Health

	if maxHealthIdx < len(mp.CurProps) && healthIdx < len(mp.CurProps) {
		mp.CurProps[healthIdx] = mp.CurProps[maxHealthIdx]
	}

	// Initialize base props from prop limits
	// TODO: iterate through PropLimit table
	// for _, propLimit := range ExcelConfigMgr.getTables().getTbPropLimit().getDataList() {
	//     if propLimit.ID < len(mp.BaseProps) && propLimit.MaxProp < len(mp.CurProps) {
	//         mp.BaseProps[propLimit.ID] = mp.CurProps[propLimit.MaxProp]
	//     }
	// }
}