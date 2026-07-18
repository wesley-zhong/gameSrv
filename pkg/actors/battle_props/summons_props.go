package battle_props

import (
	"gameSrv/cnfGen/cfg"
	"log"
)

const (
	TenThousandthsRate = 10000.0
)

// PropMode defines property mode
type PropMode int32

const (
	PropMode_SELF_PROP_VALUE          PropMode = 1
	PropMode_OWNER_STATIC_PROP_VALUE  PropMode = 2
	PropMode_OWNER_DYNAMIC_PROP_VALUE PropMode = 3
)

// SummonsProps represents battle properties for summons
type SummonsProps struct {
	*LevelBattleProps
	OwnerMaster       BattlePropsOwner
	OwnerDynamicProps map[int]bool
}

// NewSummonsProps creates a new SummonsProps
func NewSummonsProps(maxProps int, owner BattlePropsOwner) *SummonsProps {
	sp := &SummonsProps{
		LevelBattleProps:  NewLevelBattleProps(maxProps, owner),
		OwnerDynamicProps: make(map[int]bool),
	}

	// Get owner master if owner is a SummonActor
	// if summonActor, ok := owner.(*SummonActor); ok {
	//     sp.OwnerMaster = summonActor.GetMaster()
	// }

	if sp.OwnerMaster == nil {
		log.Printf("get null master, actor type: %d, actor config id: %d",
			getActorType(owner), getConfigId(owner))
	}

	return sp
}

func getActorType(owner BattlePropsOwner) int32 {
	// TODO: implement based on actual actor type
	return 0
}

func getConfigId(owner BattlePropsOwner) int64 {
	if owner == nil {
		return 0
	}
	return owner.GetConfigId()
}

// RecalPropsFormula recalculates properties formula
func (sp *SummonsProps) RecalPropsFormula() {
	if sp.BattleProps == nil {
		return
	}

	sp.BattleProps.mu.Lock()
	defer sp.BattleProps.mu.Unlock()

	maxHealthIdx := cfg.HeroProp_MaxHealth
	healthIdx := cfg.HeroProp_Health

	if maxHealthIdx < len(sp.CurProps) && healthIdx < len(sp.CurProps) {
		sp.CurProps[healthIdx] = sp.CurProps[maxHealthIdx]
	}
	if maxHealthIdx < len(sp.BaseProps) && healthIdx < len(sp.BaseProps) {
		sp.BaseProps[healthIdx] = sp.BaseProps[maxHealthIdx]
	}

	// FYDataCharactorPopsUtil.calcAttributes(CurProps, BaseProps) is commented in original
}

// ResetProps resets properties
func (sp *SummonsProps) ResetProps() {
	sp.ResetCustomProps()
}

// ResetCustomProps resets custom properties for summon
func (sp *SummonsProps) ResetCustomProps() {
	// TODO: Get summon entity data from config
	// summonData := ExcelConfigMgr.getTables().getTbSummonEntiy().get(sp.Owner.GetConfigId())
	// if summonData == nil {
	//     return
	// }

	//selfPropParamCount := getPropModeParamCount(int(PropMode_SELF_PROP_VALUE))

	// TODO: Iterate through summonData.Attributes
	// for _, propData := range summonData.Attributes {
	//     propType := HeroProp(propData.PropType)
	//     props := propData.Props
	//     if len(props) == 0 {
	//         log.Printf("get empty props summon actor: %d, propType: %d",
	//             sp.Owner.GetConfigId(), propData.PropType)
	//         continue
	//     }
	//
	//     prop := props[len(props)-1]
	//     if len(props) > selfPropParamCount {
	//         prop = sp.ParseComplexProp(propType, props)
	//     }
	//     sp.SetProperty(propData.PropType, prop)
	// }

	sp.RecalProps()
}

// IsDynamicProp checks if a property type is dynamic
func (sp *SummonsProps) IsDynamicProp(propType int) bool {
	if len(sp.OwnerDynamicProps) == 0 {
		return false
	}
	return sp.OwnerDynamicProps[propType]
}

// ParseComplexProp parses complex property with owner reference
func (sp *SummonsProps) ParseComplexProp(propType int, props []int) int {
	if len(props) == 0 {
		return 0
	}

	propMode := props[0]
	paramCount := getPropModeParamCount(propMode)
	if paramCount != len(props) {
		log.Printf("get error props, summon actor: %d, propType: %d",
			getConfigId(sp.Owner), propType)
		return 0
	}

	ownerProp := 0
	if propMode == int(PropMode_OWNER_STATIC_PROP_VALUE) {
		if sp.OwnerMaster != nil {
			// TODO: Get base prop from owner master
			// ownerProp = sp.OwnerMaster.GetBaseProp(propType)
		}
	} else if propMode == int(PropMode_OWNER_DYNAMIC_PROP_VALUE) {
		if sp.OwnerMaster != nil {
			// TODO: Get battle props from owner master
			// ownerProp = sp.OwnerMaster.GetBattleProps(propType)
		}
		// Record dynamic prop binding
		sp.AddDynamicProp(propType)
	} else {
		log.Printf("get unknown propMode summon actor: %d, propType: %d, propMode: %d",
			getConfigId(sp.Owner), propType, propMode)
		return ownerProp
	}

	if len(props) < 2 {
		return ownerProp
	}

	percentVal := props[1]
	fixVal := props[len(props)-1]
	return int(float64(ownerProp)*(float64(percentVal)/TenThousandthsRate)) + fixVal
}

func getPropModeParamCount(propMode int) int {
	// TODO: Get from config
	// misc := ExcelConfigMgr.getTables().getTbSummonMisc().get(propMode)
	// if misc == nil {
	//     log.Printf("can't found summon misc data, propMode: %d", propMode)
	//     return 0
	// }
	// return misc.ParamCount
	return 0
}

// AddDynamicProp adds a dynamic property binding
func (sp *SummonsProps) AddDynamicProp(propType int) {
	if sp.OwnerDynamicProps == nil {
		sp.OwnerDynamicProps = make(map[int]bool)
	}
	sp.OwnerDynamicProps[propType] = true
}
