package battle_props

import (
	"gameSrv/cnfGen/cfg"
	"sync"
)

// BattleProps represents the base battle properties for creatures
type BattleProps struct {
	// BaseProps are the base property values
	BaseProps []int
	// CurProps are the current property values (Index = 0 is not used)
	CurProps []int
	// AvatarPropsChanged tracks changed properties for client notification
	AvatarPropsChanged map[int]int
	// MinDataCache stores minimum property values cache
	MinDataCache map[int][]int
	// Owner is the creature that owns these properties
	Owner BattlePropsOwner
	// mu protects concurrent access
	mu sync.RWMutex
}

// BattlePropsOwner defines the interface for property owners
type BattlePropsOwner interface {
	GetEntityId() int64
	GetActorBattleModule() ActorBattleModule
	CanChangeProp(prop, value int) bool
	GetConfigId() int64
}

// ActorBattleModule defines the interface for battle module operations
type ActorBattleModule interface {
	ReCalculateBattleModules(prop int32)
	OnPropertyChange(prop, oldValue, newValue int)
}

// NewBattleProps creates a new BattleProps instance
func NewBattleProps() *BattleProps {
	return &BattleProps{
		BaseProps:          make([]int, cfg.HeroProp_Max),
		CurProps:           make([]int, cfg.HeroProp_Max),
		AvatarPropsChanged: make(map[int]int),
		MinDataCache:       make(map[int][]int),
	}
}

// NewBattlePropsWithOwner creates a new BattleProps instance with owner
func NewBattlePropsWithOwner(owner BattlePropsOwner) *BattleProps {
	return &BattleProps{
		BaseProps:          make([]int, cfg.HeroProp_Max),
		CurProps:           make([]int, cfg.HeroProp_Max),
		AvatarPropsChanged: make(map[int]int),
		MinDataCache:       make(map[int][]int),
		Owner:              owner,
	}
}

// InitAddProps adds property value during initialization (not used in battle)
func (bp *BattleProps) InitAddProps(prop, value int) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.CurProps[prop] += value
	bp.BaseProps[prop] += value
}

// InitSetProps sets property value during initialization (not used in battle)
func (bp *BattleProps) InitSetProps(prop, value int) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.CurProps[prop] = value
	bp.BaseProps[prop] = value
}

// SetProperty sets a property value
func (bp *BattleProps) SetProperty(prop, value int) int {
	return bp.SetPropertyWithRecalculate(prop, value, true)
}

// SetPropertyWithRecalculate sets a property value with optional recalculation
func (bp *BattleProps) SetPropertyWithRecalculate(prop, value int, needRecalculateProps bool) int {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if !bp.canChange(prop, value) {
		return bp.GetProperty(prop)
	}

	nowValue := bp.setActorProperty(bp.CurProps, prop, value, needRecalculateProps)
	if bp.Owner == nil {
		return nowValue
	}

	if needRecalculateProps {
		module := bp.Owner.GetActorBattleModule()
		if module != nil {
			module.ReCalculateBattleModules(int32(prop))
		}
	}
	return nowValue
}

func (bp *BattleProps) canChange(prop, value int) bool {
	if bp.Owner == nil {
		return true
	}
	return bp.Owner.CanChangeProp(prop, value)
}

// GetProperty gets a property value
func (bp *BattleProps) GetProperty(prop int) int {
	bp.mu.RLock()
	defer bp.mu.RUnlock()

	if prop < cfg.HeroProp_NONE || prop >= cfg.HeroProp_Max {
		return -1
	}
	if prop < 0 || prop >= len(bp.CurProps) {
		return -1
	}
	return bp.CurProps[prop]
}

// GetBaseProps gets base property value by int
func (bp *BattleProps) GetBaseProps(prop int) int {
	return bp.GetBasePropsByInt(prop)
}

// GetBasePropsByInt gets base property value by int (kept for compatibility)
func (bp *BattleProps) GetBasePropsByInt(propId int) int {
	bp.mu.RLock()
	defer bp.mu.RUnlock()

	if propId < 0 || propId >= len(bp.BaseProps) {
		return 0
	}
	return bp.BaseProps[propId]
}

// SetAllProps sets both base and current property value
func (bp *BattleProps) SetAllProps(prop, value int) int {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.BaseProps[prop] = value
	oldOwner := bp.Owner
	bp.mu.Unlock()

	result := bp.SetPropertyWithRecalculate(prop, value, true)

	bp.mu.Lock()
	bp.Owner = oldOwner
	bp.mu.Unlock()

	return result
}

// SetActorMinProperty sets minimum property value
func (bp *BattleProps) SetActorMinProperty(prop int, value int) int {
	if cfg.HeroProp_NONE < prop && prop < cfg.HeroProp_Max {
		bp.mu.Lock()
		defer bp.mu.Unlock()

		if _, exists := bp.MinDataCache[prop]; !exists {
			bp.MinDataCache[prop] = make([]int, 0)
		}
		bp.MinDataCache[prop] = append(bp.MinDataCache[prop], value)
		return value
	}
	return 0
}

// GetCurPropMinProperty gets the current minimum property value
func (bp *BattleProps) GetCurPropMinProperty(prop int) int {
	return bp.GetCurPropMinPropertyByInt(prop)
}

// GetCurPropMinPropertyByInt gets the current minimum property value by int (kept for compatibility)
func (bp *BattleProps) GetCurPropMinPropertyByInt(prop int) int {
	bp.mu.RLock()
	defer bp.mu.RUnlock()

	values, exists := bp.MinDataCache[prop]
	if !exists || len(values) == 0 {
		return 0
	}

	// Get the maximum value from all minimum values
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// ClearCurPropMinProperty clears a specific minimum property value
func (bp *BattleProps) ClearCurPropMinProperty(prop int, value int) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	values, exists := bp.MinDataCache[prop]
	if !exists {
		return
	}

	// Remove all occurrences of value
	newValues := make([]int, 0, len(values))
	for _, v := range values {
		if v != value {
			newValues = append(newValues, v)
		}
	}
	bp.MinDataCache[prop] = newValues
}

// ClearProps clears all properties
func (bp *BattleProps) ClearProps() {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	for i := range bp.BaseProps {
		bp.BaseProps[i] = 0
	}
	for i := range bp.CurProps {
		bp.CurProps[i] = 0
	}

	bp.MinDataCache = make(map[int][]int)
}

func (bp *BattleProps) setActorProperty(propData []int, prop, value int, needChangeEvent bool) int {
	if prop >= cfg.HeroProp_Max || prop >= len(propData) {
		return 0
	}

	oldValue := propData[prop]
	finalValue := value

	if needChangeEvent {
		propLimit := PropsUtil.GetPropLimit(prop)
		if propLimit != nil {
			maxPropIdx := int(propLimit.MaxProp)
			if maxPropIdx >= 0 && maxPropIdx < len(propData) {
				maxProp := propData[maxPropIdx]
				if propLimit.MaxProp == cfg.HeroProp_MaxPoise {
					maxPoiseCurrentRateIdx := cfg.HeroProp_MaxPoiseCurrentRate
					if maxPoiseCurrentRateIdx < len(propData) {
						maxProp = int(float64(maxProp) * (float64(propData[maxPoiseCurrentRateIdx]) * 0.0001))
					}
				}
				if finalValue > maxProp {
					finalValue = maxProp
				}
			}
		} else {
			propEffect := PropsUtil.GetPropEffect(prop)
			if propEffect != nil {
				effectPropIdx := int(propEffect.Effect)
				if effectPropIdx >= 0 && effectPropIdx < len(propData) {
					effectProp := propData[effectPropIdx]
					calcValue := oldValue
					nowValue := finalValue

					if prop == cfg.HeroProp_MaxPoiseCurrentRate || prop == cfg.HeroProp_MaxPoise {
						maxPoiseCurrentRateIdx := cfg.HeroProp_MaxPoiseCurrentRate
						maxPoiseIdx := cfg.HeroProp_MaxPoise
						if maxPoiseCurrentRateIdx < len(propData) && maxPoiseIdx < len(propData) {
							calcValue = int(float64(propData[maxPoiseCurrentRateIdx]) * 0.0001 * float64(propData[maxPoiseIdx]))

							if prop == cfg.HeroProp_MaxPoise {
								nowValue = int(float64(finalValue) * 0.0001 * float64(propData[maxPoiseCurrentRateIdx]))
							} else {
								nowValue = int(float64(finalValue) * 0.0001 * float64(propData[maxPoiseIdx]))
							}
						}
					}

					delta := calcValue - effectProp
					k := 0.0
					if calcValue != 0 {
						k = 1 - ((float64(delta) * 0.1) / (float64(calcValue) * 0.1))
					}

					bp.mu.Unlock()
					bp.SetPropertyWithRecalculate(int(propEffect.Effect), ToInter(float64(nowValue)*k), false)
					bp.mu.Lock()
				}
			}
		}
	}

	// Get minimum value
	minValue := PropsUtil.GetHeroPropMinValue(prop)

	minProperty := bp.GetCurPropMinPropertyByInt(prop)
	if minProperty != 0 {
		minValue = minProperty
	}

	if finalValue < minValue {
		finalValue = minValue
	}

	propData[prop] = finalValue

	if oldValue != finalValue && bp.Owner != nil {
		bp.AvatarPropsChanged[prop] = finalValue
		module := bp.Owner.GetActorBattleModule()
		if module != nil {
			module.OnPropertyChange(prop, oldValue, finalValue)
		}
	}
	return finalValue
}

// RecalProps recalculates properties wrapper
func (bp *BattleProps) RecalProps() {
	bp.RecalPropsFormula()
}

// RecalPropsFormula recalculates properties formula (to be implemented by subclasses)
func (bp *BattleProps) RecalPropsFormula() {
	// To be implemented by subclasses
}

// ResetProps resets properties (to be implemented by subclasses)
func (bp *BattleProps) ResetProps() {
	// To be implemented by subclasses
}

// Reborn handles reborn
func (bp *BattleProps) Reborn() {
	// Default implementation does nothing, can be overridden
}

// GetHash returns hash value
func (bp *BattleProps) GetHash() int64 {
	// TODO: implement hash calculation if needed
	return 0
}

// ClearChangedProps clears the changed properties tracking
func (bp *BattleProps) ClearChangedProps() {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.AvatarPropsChanged = make(map[int]int)
}

// ToClient fills entity builder with property data for client
func (bp *BattleProps) ToClient(builder EntityBattleInfoBuilder) {
	bp.mu.RLock()
	defer bp.mu.RUnlock()

	// Add current properties
	for _, attr := range bp.CurProps {
		builder.AddCurAttr(attr)
	}
	// Add base properties
	for _, attr := range bp.BaseProps {
		builder.AddAttr(attr)
	}
}

// EntityBattleInfoBuilder defines the interface for building entity battle info
type EntityBattleInfoBuilder interface {
	AddCurAttr(value int)
	AddAttr(value int)
}

// ChangedDataToClient generates property change notification for client
func (bp *BattleProps) ChangedDataToClient() interface{} {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if len(bp.AvatarPropsChanged) == 0 {
		return nil
	}

	// TODO: implement based on actual protobuf definition
	// This should return ProtoBattle.OceanBattlePropChangeNtf

	result := make(map[int]int)
	for prop, value := range bp.AvatarPropsChanged {
		result[prop] = value
	}

	bp.AvatarPropsChanged = make(map[int]int)
	return result
}
