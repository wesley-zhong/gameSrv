package battle_props

import (
	"gameSrv/cnfGen/cfg"
	"math"
)

// RelaOptEnum defines relation operation enum
type RelaOptEnum int32

const (
	RelaOpt_EQ  RelaOptEnum = 0
	RelaOpt_LT  RelaOptEnum = 1
	RelaOpt_LTE RelaOptEnum = 2
	RelaOpt_BT  RelaOptEnum = 3
	RelaOpt_BTE RelaOptEnum = 4
)

const (
	F_10_000 = 10000.0
	I_10_000 = 10000
)

// FYDataCharactorPopsUtil provides utility functions for character properties calculation
type FYDataCharactorPopsUtil struct{}

// GetTableMgr returns the table manager (to be implemented based on actual config system)
func (u *FYDataCharactorPopsUtil) GetTableMgr() TableManager {
	// TODO: implement based on actual config system
	return nil
}

// CalcAttributes calculates attributes from base attributes
func (u *FYDataCharactorPopsUtil) CalcAttributes(attributes, baseAttributes []int) {
	// Max Health calculation
	a1 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_A1)
	a2 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_A2)
	a3 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_A3)
	baseAttributes[cfg.HeroProp_MaxHealth] = int((float64(a1)*(F_10_000+float64(a2)))/F_10_000 + float64(a3))

	a4 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_A4)
	a5 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_A5)
	attributes[cfg.HeroProp_MaxHealth] = int((float64(baseAttributes[cfg.HeroProp_MaxHealth])*(F_10_000+float64(a4)))/F_10_000 + float64(a5))

	// Attack calculation
	b1 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_B1)
	b2 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_B2)
	b3 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_B3)
	baseAttributes[cfg.HeroProp_Attack] = int((float64(b1)*(F_10_000+float64(b2)))/F_10_000 + float64(b3))

	b4 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_B4)
	b5 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_B5)
	attributes[cfg.HeroProp_Attack] = int((float64(baseAttributes[cfg.HeroProp_Attack])*(F_10_000+float64(b4)))/F_10_000 + float64(b5))

	// Defense calculation
	c1 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_C1)
	c2 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_C2)
	c3 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_C3)
	baseAttributes[cfg.HeroProp_Defense] = int((float64(c1)*(F_10_000+float64(c2)))/F_10_000 + float64(c3))

	c4 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_C4)
	c5 := u.CalcPropGroupVal(attributes, cfg.HeroPropGroup_C5)
	attributes[cfg.HeroProp_Defense] = int((float64(baseAttributes[cfg.HeroProp_Defense])*(F_10_000+float64(c4)))/F_10_000 + float64(c5))
}

// CalcPropGroupVal calculates property group value from attributes array
func (u *FYDataCharactorPopsUtil) CalcPropGroupVal(attributes []int, propGroupId int) int {
	propsGroup := u.GetPropsGroup(propGroupId)
	if propsGroup == nil || propsGroup.Group == int32(cfg.HeroPropGroup_NONE) {
		return 0
	}

	val := 0
	for _, propId := range propsGroup.PropIdList {
		if propId == int32(cfg.HeroProp_NONE) || propId == int32(cfg.HeroProp_Max) {
			continue
		}
		if propId >= 0 && int(propId) < len(attributes) {
			val += attributes[propId]
		}
	}
	return val
}

// GetPropsGroup retrieves property group configuration
func (u *FYDataCharactorPopsUtil) GetPropsGroup(group int) *PropsGroup {
	// TODO: implement based on actual config system
	// This should return from ExcelConfigMgr.getTables().getTbPropsGroup().get(group)
	return nil
}

// GetPropLimit retrieves property limit configuration
func (u *FYDataCharactorPopsUtil) GetPropLimit(heroProp int) *PropLimit {
	// TODO: implement based on actual config system
	// This should return from ExcelConfigMgr.getTables().getTbPropLimit().get(heroProp)
	return nil
}

// GetPropEffect retrieves property effect configuration
func (u *FYDataCharactorPopsUtil) GetPropEffect(prop int) *PropEffect {
	// TODO: implement based on actual config system
	// This should return from ExcelConfigMgr.getTables().getTbPropEffect().get(prop)
	return nil
}

// GetHeroPropMinValue retrieves minimum value for a hero property
func (u *FYDataCharactorPopsUtil) GetHeroPropMinValue(heroProp int) int {
	// TODO: implement based on actual config system
	// This should return from ExcelConfigMgr.getTables().getTbHeroDataInfo().get(heroProp).MinValue
	return 0
}

// RelationValid validates relation between two values
func RelationValid(srcValue, dstValue int, relationShip RelaOptEnum) bool {
	switch relationShip {
	case RelaOpt_EQ:
		return srcValue == dstValue
	case RelaOpt_LT:
		return srcValue < dstValue
	case RelaOpt_LTE:
		return srcValue <= dstValue
	case RelaOpt_BT:
		return srcValue > dstValue
	case RelaOpt_BTE:
		return srcValue >= dstValue
	default:
		return false
	}
}

// ToInter converts float to int with specific rounding logic
func ToInter(value float64) int {
	intValue := int(value)
	intValue2 := intValue + 1
	if math.Abs(float64(intValue2)-value) <= 1e-4 {
		return intValue2
	}
	return intValue
}

// TableManager defines the interface for accessing configuration tables
type TableManager interface {
	GetTbPropsGroup() PropsGroupTable
	GetTbPropLimit() PropLimitTable
	GetTbPropEffect() PropEffectTable
	GetTbHeroDataInfo() HeroDataInfoTable
}

// PropsGroupTable defines property group table interface
type PropsGroupTable interface {
	Get(id int) *PropsGroup
}

// PropLimitTable defines property limit table interface
type PropLimitTable interface {
	Get(id int) *PropLimit
}

// PropEffectTable defines property effect table interface
type PropEffectTable interface {
	Get(id int) *PropEffect
}

// HeroDataInfoTable defines hero data info table interface
type HeroDataInfoTable interface {
	Get(id int) *HeroDataInfo
}

// PropsGroup represents property group configuration
type PropsGroup struct {
	Group      int32
	PropIdList []int32
}

// PropLimit represents property limit configuration
type PropLimit struct {
	ID      int32
	MaxProp int32
}

// PropEffect represents property effect configuration
type PropEffect struct {
	ID     int32
	Effect int32
}

// HeroDataInfo represents hero data information
type HeroDataInfo struct {
	Attribute int32
	MinValue  int32
	Sort      int32
}

// GlobalParamStorage is a placeholder for global parameter storage
var GlobalParamStorage = struct {
	Get func(id int) *GlobalParam
}{
	Get: func(id int) *GlobalParam {
		// TODO: implement based on actual config system
		return nil
	},
}

// GlobalParam represents global parameter configuration
type GlobalParam struct {
	INT int32
}

// Default utility instance
var PropsUtil = &FYDataCharactorPopsUtil{}
