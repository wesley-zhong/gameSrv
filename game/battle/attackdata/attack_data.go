package attackdata

import (
	"sort"
)

// PropsKV represents a property key-value pair
type PropsKV struct {
	PropKey int32
	PropVal int32
}

// FormationAttackDataPropDO represents attack data property for formation
type FormationAttackDataPropDO struct {
	Prop  int32
	Value int32
}

// FormationAttackDataDO represents attack data for formation
type FormationAttackDataDO struct {
	AtkId           int32
	AtkDataProps    []FormationAttackDataPropDO
	AbnormalEffects []int32
}

// AvatarAttackData represents avatar attack data (protobuf-like structure)
type AvatarAttackData struct {
	AttackDataId    int32
	Props           []PropsKV
	AbnormalEffects []int32
}

// FYAttackData represents attack data with skill properties and abnormal effects
type FYAttackData struct {
	AtkDataId        int
	AbnormalEffects  []int
	AtkDataSkillProps map[int]int32
}

// NewFYAttackData creates a new FYAttackData
func NewFYAttackData(atkDataId int) *FYAttackData {
	return &FYAttackData{
		AtkDataId:        atkDataId,
		AbnormalEffects:  make([]int, 0, 4),
		AtkDataSkillProps: make(map[int]int32, 4),
	}
}

// GetAtkDataId returns the attack data ID
func (a *FYAttackData) GetAtkDataId() int {
	return a.AtkDataId
}

// GetAbnormalEffects returns the abnormal effects
func (a *FYAttackData) GetAbnormalEffects() []int {
	return a.AbnormalEffects
}

// AddAbnormalEffects adds abnormal effects to the attack data
func (a *FYAttackData) AddAbnormalEffects(effects []int) {
	a.AbnormalEffects = append(a.AbnormalEffects, effects...)
}

// AddAbnormalEffectsSingle adds a single abnormal effect
func (a *FYAttackData) AddAbnormalEffectsSingle(effect int) {
	a.AbnormalEffects = append(a.AbnormalEffects, effect)
}

// SetAtkDataSkillProps sets a skill property value
func (a *FYAttackData) SetAtkDataSkillProps(prop int, val int32) {
	a.AtkDataSkillProps[prop] = val
}

// GetAtkDataSkillProp gets a skill property value
func (a *FYAttackData) GetAtkDataSkillProp(prop int) int32 {
	val, ok := a.AtkDataSkillProps[prop]
	if !ok || val == 0 {
		return -1 << 31 // Integer.MIN_VALUE
	}
	return val
}

// GetAtkDataSkillPropsMap returns the skill properties map
func (a *FYAttackData) GetAtkDataSkillPropsMap() map[int]int32 {
	return a.AtkDataSkillProps
}

// HasAbnormalEffect checks if the attack data has a specific abnormal effect
func (a *FYAttackData) HasAbnormalEffect(effect int) bool {
	for _, e := range a.AbnormalEffects {
		if e == effect {
			return true
		}
	}
	return false
}

// ToPB converts to protobuf message (placeholder)
func (a *FYAttackData) ToPB() *AvatarAttackData {
	props := make([]PropsKV, 0, len(a.AtkDataSkillProps))
	for prop, val := range a.AtkDataSkillProps {
		props = append(props, PropsKV{
			PropKey: int32(prop),
			PropVal: val,
		})
	}
	// Sort by PropKey for consistent output
	sort.Slice(props, func(i, j int) bool {
		return props[i].PropKey < props[j].PropKey
	})

	abnormalEffects := make([]int32, len(a.AbnormalEffects))
	for i, effect := range a.AbnormalEffects {
		abnormalEffects[i] = int32(effect)
	}

	return &AvatarAttackData{
		AttackDataId:    int32(a.AtkDataId),
		Props:           props,
		AbnormalEffects: abnormalEffects,
	}
}

// SkillToFormationAttackDataPropDO converts to formation attack data prop DO
func (a *FYAttackData) SkillToFormationAttackDataPropDO(attackDataDO *FormationAttackDataDO) {
	attackDataDO.AtkDataProps = make([]FormationAttackDataPropDO, 0, len(a.AtkDataSkillProps))
	for prop, val := range a.AtkDataSkillProps {
		attackDataDO.AtkDataProps = append(attackDataDO.AtkDataProps, FormationAttackDataPropDO{
			Prop:  int32(prop),
			Value: val,
		})
	}
}

// CreateAtkData creates a new attack data
func CreateAtkData(atkDataId int) *FYAttackData {
	return NewFYAttackData(atkDataId)
}

// Clear clears all data
func (a *FYAttackData) Clear() {
	a.AbnormalEffects = make([]int, 0)
	a.AtkDataSkillProps = make(map[int]int32)
}