package attackdata

// FYAttackData represents attack data with skill properties and abnormal effects
type FYAttackData struct {
	AtkDataId        int
	AbnormalEffects  []int
	AtkDataSkillProps map[int]int
}

// NewFYAttackData creates a new FYAttackData
func NewFYAttackData(atkDataId int) *FYAttackData {
	return &FYAttackData{
		AtkDataId:        atkDataId,
		AbnormalEffects:  make([]int, 0, 4),
		AtkDataSkillProps: make(map[int]int),
	}
}

// AddAbnormalEffects adds abnormal effects to the attack data
func (a *FYAttackData) AddAbnormalEffects(effects []int) {
	a.AbnormalEffects = append(a.AbnormalEffects, effects...)
}

// SetAtkDataSkillProps sets a skill property value
func (a *FYAttackData) SetAtkDataSkillProps(prop int, val int) {
	a.AtkDataSkillProps[prop] = val
}

// GetAtkDataSkillProp gets a skill property value
func (a *FYAttackData) GetAtkDataSkillProp(prop int) int {
	val, ok := a.AtkDataSkillProps[prop]
	if !ok {
		return -2147483648 // Integer.MIN_VALUE
	}
	return val
}

// ToPB converts to protobuf message (placeholder)
func (a *FYAttackData) ToPB() interface{} {
	// TODO: implement protobuf conversion
	// ProtoCommon.AvatarAttackData.Builder
	propsKV := make([]interface{}, 0, len(a.AtkDataSkillProps))
	for k, v := range a.AtkDataSkillProps {
		propsKV = append(propsKV, struct {
			PropKey int
			PropVal int
		}{PropKey: k, PropVal: v})
	}
	return struct {
		AttackDataId     int
		Props            []interface{}
		AbnormalEffects  []int
	}{
		AttackDataId:    a.AtkDataId,
		Props:           propsKV,
		AbnormalEffects: a.AbnormalEffects,
	}
}

// SkillToFormationAttackDataPropDO converts to formation attack data prop DO
func (a *FYAttackData) SkillToFormationAttackDataPropDO(attackDataDO interface{}) {
	// TODO: implement formation data conversion
	// This will populate FormationAttackDataDO with skill properties
}

// CreateAtkData creates a new attack data
func CreateAtkData(atkDataId int) *FYAttackData {
	return NewFYAttackData(atkDataId)
}

// Clear clears all data
func (a *FYAttackData) Clear() {
	a.AbnormalEffects = make([]int, 0)
	a.AtkDataSkillProps = make(map[int]int)
}