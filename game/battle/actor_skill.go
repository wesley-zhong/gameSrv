package battle

// FYAttackData placeholder for attack data
type FYAttackData struct {
	AtkDataId int
	// TODO: add more fields
}

// RoleDO placeholder for role data
type RoleDO struct {
	FormationAttackDataDO     *FormationAttackDataDO
	FormationAtkDataReplaceDO *FormationAtkDataReplaceDO
}

// FormationAttackDataDO placeholder
type FormationAttackDataDO struct {
	AtkId int
	// TODO: add more fields
}

// FormationAtkDataReplaceDO placeholder
type FormationAtkDataReplaceDO struct {
	AtkId       int
	ReplaceAtkId int
}

// ProtoBattle placeholder for protobuf
type ProtoBattle struct {
	EntityBattleInfo *EntityBattleInfo
}

// EntityBattleInfo placeholder
type EntityBattleInfo struct {
	AvatarAttackDatas    []*AvatarAttackData
	AttackDataReplaceInfo []*AttackDataReplaceInfo
	AtkDatasPropByTag    []*AtkPropByTag
}

// AvatarAttackData placeholder
type AvatarAttackData struct {
	AttackDataId int
	Props        []*PropsKV
}

// AttackDataReplaceInfo placeholder
type AttackDataReplaceInfo struct {
	OldAttackDataId int
	NewAttackDataId int
}

// AtkPropByTag placeholder
type AtkPropByTag struct {
	Tag      int
	SkillKvs []*SKillKV
}

// PropsKV placeholder
type PropsKV struct {
	PropKey int
	PropVal int
}

// SKillKV placeholder
type SKillKV struct {
	Key int
	Val int
}

// ActorPropsChangeNtf placeholder
type ActorPropsChangeNtf struct {
	SkillPropsValues []*AvatarAttackData
}

// ActorSkillModule manages skill data for a creature
type ActorSkillModule struct {
	creature *Creature
	atkDataProps            map[int]*FYAttackData
	atkDataPropsChanged     map[int]map[int]int
	atkDataPropsByTag       map[int]map[int]int
	atkDataReplace          map[int]int
}

// NewActorSkillModule creates a new ActorSkillModule
func NewActorSkillModule(creature *Creature) *ActorSkillModule {
	return &ActorSkillModule{
		creature:            creature,
		atkDataProps:        make(map[int]*FYAttackData),
		atkDataPropsChanged: make(map[int]map[int]int),
		atkDataPropsByTag:   make(map[int]map[int]int),
		atkDataReplace:      make(map[int]int),
	}
}

// Copy copies data from another ActorSkillModule
func (m *ActorSkillModule) Copy(other *ActorSkillModule) {
	// Clear existing data
	m.atkDataReplace = make(map[int]int)
	m.atkDataProps = make(map[int]*FYAttackData)
	m.atkDataPropsChanged = make(map[int]map[int]int)
	m.atkDataPropsByTag = make(map[int]map[int]int)

	// Copy from other
	for k, v := range other.atkDataReplace {
		m.atkDataReplace[k] = v
	}
	for k, v := range other.atkDataProps {
		m.atkDataProps[k] = v
	}
	for k, v := range other.atkDataPropsChanged {
		m.atkDataPropsChanged[k] = v
	}
	for k, v := range other.atkDataPropsByTag {
		m.atkDataPropsByTag[k] = v
	}
}

// SetAtkDataProps sets attack data property
func (m *ActorSkillModule) SetAtkDataProps(atkDataId int, prop int, value int) {
	attackData := m.getOrCreateAtkData(atkDataId)
	attackData.setAtkDataSkillProps(prop, value)
}

// setAtkDataSkillProps sets skill props (FYAttackData method)
func (a *FYAttackData) setAtkDataSkillProps(prop int, value int) {
	// TODO: implement
}

// AddAtkAbnormalEffects adds abnormal effects for multiple attack data
func (m *ActorSkillModule) AddAtkAbnormalEffects(atkDataIds []int, abnormalEffects []int) {
	if len(abnormalEffects) == 0 {
		return
	}
	for _, atkDataId := range atkDataIds {
		m.AddAtkAbnormalEffectsSingle(atkDataId, abnormalEffects)
	}
}

// AddAtkAbnormalEffectsSingle adds abnormal effects for single attack data
func (m *ActorSkillModule) AddAtkAbnormalEffectsSingle(atkDataId int, abnormalEffects []int) {
	attackData := m.getOrCreateAtkData(atkDataId)
	attackData.addAbnormalEffects(abnormalEffects)
}

// addAbnormalEffects adds abnormal effects (FYAttackData method)
func (a *FYAttackData) addAbnormalEffects(effects []int) {
	// TODO: implement
}

// AddAtkDataReplace adds attack data replacements
func (m *ActorSkillModule) AddAtkDataReplace(atkDataReplaces map[int]int) {
	for k, v := range atkDataReplaces {
		m.atkDataReplace[k] = v
	}
}

// AddAtkDataReplaceSingle adds single attack data replacement
func (m *ActorSkillModule) AddAtkDataReplaceSingle(oldAtkDataId int, newAtkDataId int) {
	m.atkDataReplace[oldAtkDataId] = newAtkDataId
}

// ClearSkillData clears all skill data
func (m *ActorSkillModule) ClearSkillData() {
	m.atkDataProps = make(map[int]*FYAttackData)
	m.atkDataPropsChanged = make(map[int]map[int]int)
	m.atkDataReplace = make(map[int]int)
}

// ClearChangedData clears changed data
func (m *ActorSkillModule) ClearChangedData() {
	m.atkDataPropsChanged = make(map[int]map[int]int)
}

// GetAtkDataProps gets attack data property
func (m *ActorSkillModule) GetAtkDataProps(atkDataId int, skillProp int) int {
	attackData, ok := m.atkDataProps[atkDataId]
	if !ok {
		return -2147483648 // Integer.MIN_VALUE
	}
	return attackData.getAtkDataSkillProp(skillProp)
}

// getAtkDataSkillProp gets skill prop (FYAttackData method)
func (a *FYAttackData) getAtkDataSkillProp(skillProp int) int {
	// TODO: implement
	return 0
}

// SetAtkDataPropsByTag sets attack data props by tag
func (m *ActorSkillModule) SetAtkDataPropsByTag(tag int, skillProp int, value int) {
	atkProps := m.getOrCreateSkillProps(m.atkDataPropsByTag, tag)
	atkProps[skillProp] = value
}

// GetAtkDataPropsByTag gets attack data props by tag
func (m *ActorSkillModule) GetAtkDataPropsByTag(tag int, skillProp int) int {
	intCursors, ok := m.atkDataPropsByTag[tag]
	if !ok {
		return 0
	}
	return intCursors[skillProp]
}

// SkillsToClient converts skills to client data
func (m *ActorSkillModule) SkillsToClient(battleBuilder *ProtoBattle) {
	for _, attackData := range m.atkDataProps {
		battleBuilder.EntityBattleInfo.AvatarAttackDatas = append(
			battleBuilder.EntityBattleInfo.AvatarAttackDatas,
			attackData.toPB(),
		)
	}
	for oldId, newId := range m.atkDataReplace {
		battleBuilder.EntityBattleInfo.AttackDataReplaceInfo = append(
			battleBuilder.EntityBattleInfo.AttackDataReplaceInfo,
			&AttackDataReplaceInfo{
				OldAttackDataId: oldId,
				NewAttackDataId: newId,
			},
		)
	}
	for tag, skillProps := range m.atkDataPropsByTag {
		builder := &AtkPropByTag{Tag: tag}
		for k, v := range skillProps {
			builder.SkillKvs = append(builder.SkillKvs, &SKillKV{Key: k, Val: v})
		}
		battleBuilder.EntityBattleInfo.AtkDatasPropByTag = append(
			battleBuilder.EntityBattleInfo.AtkDatasPropByTag,
			builder,
		)
	}
}

// toPB converts to protobuf (FYAttackData method)
func (a *FYAttackData) toPB() *AvatarAttackData {
	// TODO: implement
	return &AvatarAttackData{AttackDataId: a.AtkDataId}
}

// SkillsToFormationDatas converts skills to formation data
func (m *ActorSkillModule) SkillsToFormationDatas(attackDatas []*FormationAttackDataDO, atkDataReplaces []*FormationAtkDataReplaceDO) {
	for _, attackData := range m.atkDataProps {
		attackDataDO := &FormationAttackDataDO{
			AtkId: attackData.AtkDataId,
		}
		attackData.skillToFormationAttackDataPropDO(attackDataDO)
		// TODO: set abnormal effects
		attackDatas = append(attackDatas, attackDataDO)
	}

	for oldId, newId := range m.atkDataReplace {
		attackDataReplaceDO := &FormationAtkDataReplaceDO{
			AtkId:       oldId,
			ReplaceAtkId: newId,
		}
		atkDataReplaces = append(atkDataReplaces, attackDataReplaceDO)
	}
}

// skillToFormationAttackDataPropDO converts to formation attack data prop DO
func (a *FYAttackData) skillToFormationAttackDataPropDO(do *FormationAttackDataDO) {
	// TODO: implement
}

// SkillsChangedToPB converts changed skills to protobuf
func (m *ActorSkillModule) SkillsChangedToPB(entityBuilder *ActorPropsChangeNtf) {
	for atkDataId, props := range m.atkDataPropsChanged {
		avatarAttackDataBuilder := &AvatarAttackData{
			AttackDataId: atkDataId,
		}
		for k, v := range props {
			propsKV := &PropsKV{PropKey: k, PropVal: v}
			avatarAttackDataBuilder.Props = append(avatarAttackDataBuilder.Props, propsKV)
		}
		entityBuilder.SkillPropsValues = append(entityBuilder.SkillPropsValues, avatarAttackDataBuilder)
	}
	m.ClearChangedData()
}

// getOrCreateSkillProps gets or creates skill props map
func (m *ActorSkillModule) getOrCreateSkillProps(collection map[int]map[int]int, key int) map[int]int {
	result, ok := collection[key]
	if !ok {
		result = make(map[int]int)
		collection[key] = result
	}
	return result
}

// getOrCreateAtkData gets or creates attack data
func (m *ActorSkillModule) getOrCreateAtkData(atkDataId int) *FYAttackData {
	attackData, ok := m.atkDataProps[atkDataId]
	if !ok {
		attackData = &FYAttackData{AtkDataId: atkDataId}
		m.atkDataProps[atkDataId] = attackData
	}
	return attackData
}

// GetAtkData gets attack data by id
func (m *ActorSkillModule) GetAtkData(atkDataId int) *FYAttackData {
	return m.atkDataProps[atkDataId]
}