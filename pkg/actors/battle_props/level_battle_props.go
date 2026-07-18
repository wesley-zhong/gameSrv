package battle_props

// LevelBattleProps represents battle properties for level-based creatures
type LevelBattleProps struct {
	*BattleProps
}

// NewLevelBattleProps creates a new LevelBattleProps
func NewLevelBattleProps(maxProps int, owner BattlePropsOwner) *LevelBattleProps {
	return &LevelBattleProps{
		BattleProps: NewBattlePropsWithOwner(owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (lp *LevelBattleProps) RecalPropsFormula() {
	if lp.BattleProps == nil {
		return
	}
	lp.BattleProps.mu.Lock()
	defer lp.BattleProps.mu.Unlock()

	PropsUtil.CalcAttributes(lp.CurProps, lp.BaseProps)
}

// ResetProps resets properties for level creature
func (lp *LevelBattleProps) ResetProps() {
	if lp.BattleProps == nil {
		return
	}

	// Check if owner is a LevelCreature
	// This requires type assertion to LevelCreature
	// For now, we'll implement the basic logic

	lp.ClearProps()

	// TODO: Get monster template data from owner
	// monsterTmplData := creature.GetMonsterTmplData()
	// if monsterTmplData != nil {
	//     for prop, value := range monsterTmplData.HeroAttributes {
	//         lp.InitSetProps(prop, value)
	//     }
	// }

	// TODO: Apply scene monster properties
	// scene := owner.GetScene()
	// if scene != nil {
	//     for prop, value := range scene.GetMonsterProps() {
	//         lp.InitAddProps(prop, value)
	//     }
	// }

	// TODO: Apply monster level properties (growth)
	// monsterLevelProps := creature.GetMonsterLevelProps()
	// if monsterLevelProps != nil {
	//     for prop, growthValue := range monsterLevelProps.GrowthAttributes {
	//         propsVal := lp.GetProperty(prop)
	//         newValue := int(float64(propsVal) + float64(propsVal*int64(growthValue))/F_10_000)
	//         lp.InitSetProps(prop, newValue)
	//     }
	// }

	// TODO: Apply monster difficulty data
	// monsterDifficultyData := creature.GetMonsterDifficultyData()
	// if monsterDifficultyData != nil {
	//     for prop, difficultyValue := range monsterDifficultyData.HeroAttributes {
	//         propsVal := lp.GetProperty(prop)
	//         newValue := int(float64(propsVal) + float64(propsVal*int64(difficultyValue))/F_10_000)
	//         lp.InitSetProps(prop, newValue)
	//     }
	// }

	// Set level property
	// level := creature.GetLevel()
	// lp.InitSetProps(cfg.HeroProp_Level, level)

	lp.RecalProps()
}
