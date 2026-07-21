package battle

import (
	"sort"

	"gameSrv/game/battle/attackdata"
	"gameSrv/pkg/actors"
)

// AttackDataReplaceInfo represents attack data replacement info (protobuf-like)
type AttackDataReplaceInfo struct {
	OldAttackDataId int32
	NewAttackDataId int32
}

// AtkPropByTag represents attack properties by tag (protobuf-like)
type AtkPropByTag struct {
	Tag      int32
	SkillKvs []SkillKV
}

// SkillKV represents a skill key-value pair (protobuf-like)
type SkillKV struct {
	Key int32
	Val int32
}

// PropsKV represents a property key-value pair (protobuf-like)
type PropsKV struct {
	PropKey int32
	PropVal int32
}

// AvatarAttackData represents avatar attack data (protobuf-like)
type AvatarAttackData struct {
	AttackDataId    int32
	Props           []PropsKV
	AbnormalEffects []int32
}

// FormationAtkDataReplaceDO represents formation attack data replacement
type FormationAtkDataReplaceDO struct {
	AtkId        int32
	ReplaceAtkId int32
}

// EntityBattleInfo represents entity battle info (protobuf-like)
type EntityBattleInfo struct {
	AvatarAttackDatas     []*AvatarAttackData
	AttackDataReplaceInfo []*AttackDataReplaceInfo
	AtkDatasPropByTag     []*AtkPropByTag
}

// ActorPropsChangeNtf represents actor props change notification (protobuf-like)
type ActorPropsChangeNtf struct {
	SkillPropsValues []*AvatarAttackData
}

// ActorSkillModule manages skill data for a creature
type ActorSkillModule struct {
	owner               *actors.Creature
	atkDataProps        map[int]*attackdata.FYAttackData
	atkDataPropsChanged map[int]map[int]int32 // Changed skill properties
	atkDataPropsByTag   map[int]map[int]int32 // Properties by skill type tag
	atkDataReplace      map[int]int32         // Attack data replacement map
}

// NewActorSkillModule creates a new ActorSkillModule
func NewActorSkillModule(creature *actors.Creature) *ActorSkillModule {
	return &ActorSkillModule{
		owner:               creature,
		atkDataProps:        make(map[int]*attackdata.FYAttackData),
		atkDataPropsChanged: make(map[int]map[int]int32),
		atkDataPropsByTag:   make(map[int]map[int]int32),
		atkDataReplace:      make(map[int]int32),
	}
}

// Copy copies data from another ActorSkillModule
func (m *ActorSkillModule) Copy(actorSkillModule *ActorSkillModule) {
	// Clear existing data
	m.ClearSkillData()

	// Copy attack data replace
	for k, v := range actorSkillModule.atkDataReplace {
		m.atkDataReplace[k] = v
	}

	// Copy attack data props (deep copy needed)
	for k, v := range actorSkillModule.atkDataProps {
		// Create a new FYAttackData with same ID
		newAtkData := attackdata.NewFYAttackData(v.GetAtkDataId())
		// Copy abnormal effects
		for _, effect := range v.GetAbnormalEffects() {
			newAtkData.AddAbnormalEffectsSingle(effect)
		}
		// Copy skill props
		for prop, val := range v.GetAtkDataSkillPropsMap() {
			newAtkData.SetAtkDataSkillProps(prop, val)
		}
		m.atkDataProps[k] = newAtkData
	}

	// Copy changed props
	for k, v := range actorSkillModule.atkDataPropsChanged {
		newMap := make(map[int]int32)
		for pk, pv := range v {
			newMap[pk] = pv
		}
		m.atkDataPropsChanged[k] = newMap
	}

	// Copy props by tag
	for k, v := range actorSkillModule.atkDataPropsByTag {
		newMap := make(map[int]int32)
		for pk, pv := range v {
			newMap[pk] = pv
		}
		m.atkDataPropsByTag[k] = newMap
	}
}

// SetAtkDataProps sets an attack data property
func (m *ActorSkillModule) SetAtkDataProps(atkDataId int, prop int, value int) {
	attackData := m.getOrCreateAtkData(atkDataId)
	attackData.SetAtkDataSkillProps(prop, int32(value))

	// Track change
	if _, exists := m.atkDataPropsChanged[atkDataId]; !exists {
		m.atkDataPropsChanged[atkDataId] = make(map[int]int32)
	}
	m.atkDataPropsChanged[atkDataId][prop] = int32(value)
}

// AddAtkAbnormalEffects adds abnormal effects to multiple attack data IDs
func (m *ActorSkillModule) AddAtkAbnormalEffects(atkDataIds []int, abnormalEffects []int) {
	if len(abnormalEffects) == 0 {
		return
	}

	for _, atkDataId := range atkDataIds {
		m.AddAtkAbnormalEffectsSingle(atkDataId, abnormalEffects)
	}
}

// AddAtkAbnormalEffectsSingle adds abnormal effects to a single attack data
func (m *ActorSkillModule) AddAtkAbnormalEffectsSingle(atkDataId int, abnormalEffects []int) {
	attackData := m.getOrCreateAtkData(atkDataId)
	attackData.AddAbnormalEffects(abnormalEffects)
}

// AddAtkDataReplace adds attack data replacements
func (m *ActorSkillModule) AddAtkDataReplace(atkDataReplaces map[int]int32) {
	for k, v := range atkDataReplaces {
		m.atkDataReplace[k] = v
	}
}

// AddAtkDataReplaceSingle adds a single attack data replacement
func (m *ActorSkillModule) AddAtkDataReplaceSingle(oldAtkDataId int, newAtkDataId int32) {
	m.atkDataReplace[oldAtkDataId] = newAtkDataId
}

// ClearSkillData clears all skill data
func (m *ActorSkillModule) ClearSkillData() {
	m.atkDataProps = make(map[int]*attackdata.FYAttackData)
	m.atkDataPropsChanged = make(map[int]map[int]int32)
	m.atkDataReplace = make(map[int]int32)
	// Note: atkDataPropsByTag is not cleared in original Java code
}

// ClearChangedData clears changed data
func (m *ActorSkillModule) ClearChangedData() {
	m.atkDataPropsChanged = make(map[int]map[int]int32)
}

// GetAtkDataProps gets an attack data property value
// Returns min int value if not found (consistent with Java's Integer.MIN_VALUE)
func (m *ActorSkillModule) GetAtkDataProps(atkDataId int, skillProp int) int {
	attackData := m.GetAtkData(atkDataId)
	if attackData == nil {
		return -1 << 31
	}
	return int(attackData.GetAtkDataSkillProp(skillProp))
}

// GetAtkDataPropsInt32 gets an attack data property value as int32
func (m *ActorSkillModule) GetAtkDataPropsInt32(atkDataId int, skillProp int) int32 {
	attackData := m.GetAtkData(atkDataId)
	if attackData == nil {
		return -1 << 31
	}
	return attackData.GetAtkDataSkillProp(skillProp)
}

// SetAtkDataPropsByTag sets an attack data property by tag
func (m *ActorSkillModule) SetAtkDataPropsByTag(tag int, skillProp int, value int) {
	if _, exists := m.atkDataPropsByTag[tag]; !exists {
		m.atkDataPropsByTag[tag] = make(map[int]int32)
	}
	m.atkDataPropsByTag[tag][skillProp] = int32(value)
}

// GetAtkDataPropsByTag gets an attack data property by tag
func (m *ActorSkillModule) GetAtkDataPropsByTag(tag int, skillProp int) int {
	if skillProps, exists := m.atkDataPropsByTag[tag]; exists {
		return int(skillProps[skillProp])
	}
	return 0
}

// GetAtkData gets an attack data by ID
func (m *ActorSkillModule) GetAtkData(atkDataId int) *attackdata.FYAttackData {
	return m.atkDataProps[atkDataId]
}

// HasAtkData checks if attack data exists
func (m *ActorSkillModule) HasAtkData(atkDataId int) bool {
	_, exists := m.atkDataProps[atkDataId]
	return exists
}

// GetAtkDataReplace gets the replacement attack data ID
func (m *ActorSkillModule) GetAtkDataReplace(atkDataId int) int32 {
	if newId, exists := m.atkDataReplace[atkDataId]; exists {
		return newId
	}
	return 0
}

// SkillsToClient converts skills to client data (protobuf-like structure)
func (m *ActorSkillModule) SkillsToClient() *EntityBattleInfo {
	battleInfo := &EntityBattleInfo{
		AvatarAttackDatas:      make([]*AvatarAttackData, 0),
		AttackDataReplaceInfo:  make([]*AttackDataReplaceInfo, 0),
		AtkDatasPropByTag:      make([]*AtkPropByTag, 0),
	}

	// Add all attack datas
	for _, attackData := range m.atkDataProps {
		pbData := attackData.ToPB()
		props := make([]PropsKV, len(pbData.Props))
		for i, p := range pbData.Props {
			props[i] = PropsKV{
				PropKey: p.PropKey,
				PropVal: p.PropVal,
			}
		}
		battleInfo.AvatarAttackDatas = append(battleInfo.AvatarAttackDatas, &AvatarAttackData{
			AttackDataId:    pbData.AttackDataId,
			Props:           props,
			AbnormalEffects: pbData.AbnormalEffects,
		})
	}

	// Add attack data replace info
	for oldId, newId := range m.atkDataReplace {
		battleInfo.AttackDataReplaceInfo = append(battleInfo.AttackDataReplaceInfo, &AttackDataReplaceInfo{
			OldAttackDataId: int32(oldId),
			NewAttackDataId: newId,
		})
	}

	// Add attack data props by tag
	for tag, skillProps := range m.atkDataPropsByTag {
		atkPropByTag := &AtkPropByTag{
			Tag:      int32(tag),
			SkillKvs: make([]SkillKV, 0),
		}
		for prop, val := range skillProps {
			atkPropByTag.SkillKvs = append(atkPropByTag.SkillKvs, SkillKV{
				Key: int32(prop),
				Val: val,
			})
		}
		// Sort by Key for consistent output
		sort.Slice(atkPropByTag.SkillKvs, func(i, j int) bool {
			return atkPropByTag.SkillKvs[i].Key < atkPropByTag.SkillKvs[j].Key
		})
		battleInfo.AtkDatasPropByTag = append(battleInfo.AtkDatasPropByTag, atkPropByTag)
	}

	return battleInfo
}

// SkillsToFormationDatas converts skills to formation data
func (m *ActorSkillModule) SkillsToFormationDatas() ([]*attackdata.FormationAttackDataDO, []*FormationAtkDataReplaceDO) {
	attackDatas := make([]*attackdata.FormationAttackDataDO, 0)
	atkDataReplaces := make([]*FormationAtkDataReplaceDO, 0)

	for _, attackData := range m.atkDataProps {
		attackDataDO := &attackdata.FormationAttackDataDO{
			AtkId: int32(attackData.GetAtkDataId()),
		}
		attackData.SkillToFormationAttackDataPropDO(attackDataDO)
		abnormalEffects := make([]int32, len(attackData.GetAbnormalEffects()))
		for i, effect := range attackData.GetAbnormalEffects() {
			abnormalEffects[i] = int32(effect)
		}
		attackDataDO.AbnormalEffects = abnormalEffects
		attackDatas = append(attackDatas, attackDataDO)
	}

	for oldId, newId := range m.atkDataReplace {
		atkDataReplaceDO := &FormationAtkDataReplaceDO{
			AtkId:        int32(oldId),
			ReplaceAtkId: newId,
		}
		atkDataReplaces = append(atkDataReplaces, atkDataReplaceDO)
	}

	return attackDatas, atkDataReplaces
}

// SkillsChangedToPB converts changed skills to protobuf-like structure
func (m *ActorSkillModule) SkillsChangedToPB() *ActorPropsChangeNtf {
	notify := &ActorPropsChangeNtf{
		SkillPropsValues: make([]*AvatarAttackData, 0),
	}

	for atkDataId, props := range m.atkDataPropsChanged {
		avatarAttackData := &AvatarAttackData{
			AttackDataId: int32(atkDataId),
			Props:        make([]PropsKV, 0),
		}

		for prop, val := range props {
			avatarAttackData.Props = append(avatarAttackData.Props, PropsKV{
				PropKey: int32(prop),
				PropVal: val,
			})
		}

		notify.SkillPropsValues = append(notify.SkillPropsValues, avatarAttackData)
	}

	// Clear after generating
	m.ClearChangedData()

	return notify
}

// getOrCreateAtkData gets or creates an attack data
func (m *ActorSkillModule) getOrCreateAtkData(atkDataId int) *attackdata.FYAttackData {
	if attackData, exists := m.atkDataProps[atkDataId]; exists {
		return attackData
	}
	attackData := attackdata.CreateAtkData(atkDataId)
	m.atkDataProps[atkDataId] = attackData
	return attackData
}

// getOrCreateSkillProps gets or creates skill props map
func (m *ActorSkillModule) getOrCreateSkillProps(collection map[int]map[int]int32, key int) map[int]int32 {
	if skillProps, exists := collection[key]; exists {
		return skillProps
	}
	skillProps := make(map[int]int32)
	collection[key] = skillProps
	return skillProps
}

// GetOwner returns the owner creature
func (m *ActorSkillModule) GetOwner() *actors.Creature {
	return m.owner
}
