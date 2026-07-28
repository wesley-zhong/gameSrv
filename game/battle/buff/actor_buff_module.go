package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/battle"
	"gameSrv/pkg/actors"
)

// FormationBuffDO represents formation buff data (for cross-scene buff persistence)
type FormationBuffDO struct {
	BuffUId       int64
	BuffId        int32
	Layer         int32
	ConsumedTime  int64
	BSystem       bool
}

// EntityBattleInfo represents entity battle info (protobuf-like)
type EntityBattleInfo struct {
	Buffs []*BuffData
}

// ActorBuffModule manages buffs for an actor
type ActorBuffModule struct {
	owner        *actors.Creature
	BuffExec     IFYBuffExecution
	changedBuffs []*BuffData
}

// execRemoveBuffByUid executes RemoveBuffByUid
func (m *ActorBuffModule) execRemoveBuffByUid(buffCnfId int, uid int64, reason int) bool {
	return m.BuffExec.RemoveBuffByUid(buffCnfId, uid, reason)
}

// execAddBuff executes AddBuff
func (m *ActorBuffModule) execAddBuff(templateId int, casterActor, holder *actors.Creature, uid int64, exParam, layer int, life int64, bSystem bool) *FYBuff {
	result := m.BuffExec.AddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
	if result == nil {
		return nil
	}
	return result.(*FYBuff)
}

// execStopAndRemoveBuff executes StopAndRemoveBuff
func (m *ActorBuffModule) execStopAndRemoveBuff(removedBuff *FYBuff, reason int) {
	m.BuffExec.StopAndRemoveBuff(removedBuff, reason)
}

// execTriggerBuffEvent executes TriggerBuffEvent
func (m *ActorBuffModule) execTriggerBuffEvent(eventEnum int, triggerEventParam interface{}) bool {
	return m.BuffExec.TriggerBuffEvent(eventEnum, triggerEventParam)
}

// execHasBuffTag executes HasBuffTag
func (m *ActorBuffModule) execHasBuffTag(tagEnum int) bool {
	return m.BuffExec.HasBuffTag(tagEnum)
}

// execRemoveBuffTag executes RemoveBuffTag
func (m *ActorBuffModule) execRemoveBuffTag(inTag int) bool {
	return m.BuffExec.RemoveBuffTag(inTag)
}

// execAddBuffTag executes AddBuffTag
func (m *ActorBuffModule) execAddBuffTag(inTag int) bool {
	return m.BuffExec.AddBuffTag(inTag)
}

// execClearBuff executes ClearBuff
func (m *ActorBuffModule) execClearBuff(death bool) {
	m.BuffExec.ClearBuff(death)
}

// execClearNotSystemBuff executes ClearNotSystemBuff
func (m *ActorBuffModule) execClearNotSystemBuff() {
	m.BuffExec.ClearNotSystemBuff()
}

// execClearLeaveBattleBuff executes ClearLeaveBattleBuff
func (m *ActorBuffModule) execClearLeaveBattleBuff() {
	m.BuffExec.ClearLeaveBattleBuff()
}

// execRemoveBuffByConfId executes RemoveBuffByConfId
func (m *ActorBuffModule) execRemoveBuffByConfId(cnfId int, reason int) bool {
	return m.BuffExec.RemoveBuffByConfId(cnfId, reason)
}

// execRemoveBuffBySubClass executes RemoveBuffBySubClass
func (m *ActorBuffModule) execRemoveBuffBySubClass(classID, subClassID int, reason int) bool {
	return m.BuffExec.RemoveBuffBySubClass(classID, subClassID, reason)
}

// execRemoveBuffByClass executes RemoveBuffByClass
func (m *ActorBuffModule) execRemoveBuffByClass(classType int, reason int) bool {
	return m.BuffExec.RemoveBuffByClass(classType, reason)
}

// getAllBuffsMap returns the buffs map
func (m *ActorBuffModule) getAllBuffsMap() map[int][]*FYBuff {
	return m.BuffExec.GetAllBuffsMap()
}

// BuffData represents buff data for client communication
type BuffData struct {
	State           int32
	BuffCnfId       int32
	BuffUid         int64
	CreationTime    int64
	CurLayer        int32
	StartEffectTime int64
}

// EBuffState represents buff state
type EBuffState int32

const (
	EBuffState_BUFF_START EBuffState = 0
	EBuffState_BUFF_END   EBuffState = 1
)

// NewActorBuffModule creates a new ActorBuffModule
func NewActorBuffModule(creature *actors.Creature, isServer bool) *ActorBuffModule {
	var exec IFYBuffExecution
	if isServer {
		exec = NewFYBuffExecutionServer(creature)
	} else {
		exec = NewFYBuffExecutionClient(creature)
	}
	return &ActorBuffModule{
		owner:        creature,
		BuffExec:     exec,
		changedBuffs: make([]*BuffData, 0),
	}
}

// SystemAddBuff adds a buff from system
func (m *ActorBuffModule) SystemAddBuff(bufCnfId int, uid int64) bool {
	ret := m.AddBuff(bufCnfId, m.owner, nil, uid, 0, true)
	if ret != nil {
		buffData := &BuffData{
			State:     int32(EBuffState_BUFF_START),
			BuffCnfId: int32(bufCnfId),
			BuffUid:   uid,
		}
		m.changedBuffs = append(m.changedBuffs, buffData)
	}
	return ret != nil
}

// SystemRemoveBuff removes a buff from system
func (m *ActorBuffModule) SystemRemoveBuff(buffCnfId int, uid int64) bool {
	success := m.execRemoveBuffByUid(buffCnfId, uid, cfg.BuffEndTypeEnum_BUFF_END_NONE)
	if success {
		buffData := &BuffData{
			State:     int32(EBuffState_BUFF_END),
			BuffCnfId: int32(buffCnfId),
			BuffUid:   uid,
		}
		m.changedBuffs = append(m.changedBuffs, buffData)
	}
	return success
}

// AddBuff adds a buff
func (m *ActorBuffModule) AddBuff(
	templateId int,
	casterActor, holder *actors.Creature,
	uid int64,
	exParam int,
	bSystem bool,
) actors.IBuff {
	return m.execAddBuff(templateId, casterActor, holder, uid, exParam, 0, 0, bSystem)
}

// AddBuffWithLayer adds a buff with specific layer
func (m *ActorBuffModule) AddBuffWithLayer(
	templateId int,
	casterActor, holder *actors.Creature,
	uid int64,
	exParam, layer int,
	life int64,
	bSystem bool,
) *FYBuff {
	return m.execAddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
}

// AddBuffWithTriggerFlag adds a buff with trigger flag
func (m *ActorBuffModule) AddBuffWithTriggerFlag(
	templateId int,
	casterActor *actors.Creature,
	uid int64,
	exParam int,
	bSystem bool,
	triggerFlag int,
) *FYBuff {
	fyBuff := m.AddBuff(templateId, casterActor, nil, uid, exParam, bSystem)
	if fyBuff != nil {
		if buff, ok := fyBuff.(*FYBuff); ok {
			buff.TriggerFlag = triggerFlag
			return buff
		}
	}
	return nil
}

// StopAndRemoveBuff stops and removes a buff
func (m *ActorBuffModule) StopAndRemoveBuff(removedBuff *FYBuff, reason int) {
	m.execStopAndRemoveBuff(removedBuff, reason)
}

// TriggerBuffEvent triggers a buff event
func (m *ActorBuffModule) TriggerBuffEvent(
	eventEnum int,
	triggerEventParam interface{},
) bool {
	return m.execTriggerBuffEvent(eventEnum, triggerEventParam)
}

// HasBuffTag checks if creature has a buff tag
func (m *ActorBuffModule) HasBuffTag(tagEnum int) bool {
	return m.execHasBuffTag(tagEnum)
}

// RemoveBuffTag removes a buff tag
func (m *ActorBuffModule) RemoveBuffTag(inTag int) bool {
	return m.execRemoveBuffTag(inTag)
}

// AddBuffTag adds a buff tag
func (m *ActorBuffModule) AddBuffTag(inTag int) bool {
	return m.execAddBuffTag(inTag)
}

// ClearBuff clears all buffs
func (m *ActorBuffModule) ClearBuff(death bool) {
	m.execClearBuff(death)
}

// ClearNotSystemBuff clears non-system buffs
func (m *ActorBuffModule) ClearNotSystemBuff() {
	m.execClearNotSystemBuff()
}

// ClearLeaveBattleBuff clears buffs on leave battle
func (m *ActorBuffModule) ClearLeaveBattleBuff() {
	m.execClearLeaveBattleBuff()
}

// ClearChangedData clears changed buffs data
func (m *ActorBuffModule) ClearChangedData() {
	m.changedBuffs = make([]*BuffData, 0)
}

// RemoveBuffByConfId removes buffs by config ID
func (m *ActorBuffModule) RemoveBuffByConfId(cnfId int, reason int) bool {
	return m.execRemoveBuffByConfId(cnfId, reason)
}

// RemoveBuffBySubClass removes buffs by subclass
func (m *ActorBuffModule) RemoveBuffBySubClass(classID, subClassID int, reason int) bool {
	return m.execRemoveBuffBySubClass(classID, subClassID, reason)
}

// RemoveBuffByClass removes buffs by class
func (m *ActorBuffModule) RemoveBuffByClass(classType int, reason int) bool {
	return m.execRemoveBuffByClass(classType, reason)
}

// RebornAddBuff adds buffs on rebirth
func (m *ActorBuffModule) RebornAddBuff() {
	// TODO: Implement when reborn system is ready
}

// TickFromClient handles client-side buff tick
func (m *ActorBuffModule) TickFromClient(buffId int) {
	// Client doesn't tick buffs - server handles all buff ticking
	_ = buffId
}

// GetBuffExecution returns the buff execution instance
func (m *ActorBuffModule) GetBuffExecution() IFYBuffExecution {
	return m.BuffExec
}

// GetOwner returns the owner creature
func (m *ActorBuffModule) GetOwner() *actors.Creature {
	return m.owner
}

// OnBuffUpdate is called when buff is updated
// Sends buff update notification to client via network system
func (m *ActorBuffModule) OnBuffUpdate(addedBuff *FYBuff) {
	_ = addedBuff
	// Network notification will be sent when network system is implemented
	// For system buffs, notification is handled by ActorBattleModule.OnAvatarPropsRestFinish
}

// OnBuffRemoved is called when buff is removed
// Sends buff removal notification to client via network system
func (m *ActorBuffModule) OnBuffRemoved(removedBuff *FYBuff, reason int) {
	_ = removedBuff
	_ = reason
	// Network notification will be sent when network system is implemented
	// For system buffs, notification is handled by ActorBattleModule.OnAvatarPropsRestFinish
}

// BuffsToClient converts buffs to client data
func (m *ActorBuffModule) BuffsToClient() []*BuffData {
	buffDataList := make([]*BuffData, 0)

	allBuffs := m.getAllBuffsMap()
	for _, buffs := range allBuffs {
		for _, buff := range buffs {
			buffDataList = append(buffDataList, toClient(buff, EBuffState_BUFF_START))
		}
	}

	return buffDataList
}

// BuffsToClientBuilder converts buffs to client data builder (for protobuf compatibility)
func (m *ActorBuffModule) BuffsToClientBuilder(battleBuilder *EntityBattleInfo) {
	battleBuilder.Buffs = m.BuffsToClient()
}

// BuffsToFormationDatas converts buffs to formation data format
// Used for saving buff data when leaving scene (for non-system buffs or buffs with RemoveOnLeaveScene=false)
func (m *ActorBuffModule) BuffsToFormationDatas() []*FormationBuffDO {
	buffDatas := make([]*FormationBuffDO, 0)

	allBuffs := m.getAllBuffsMap()
	currTime := battle.MilliSeconds()

	for _, buffs := range allBuffs {
		for _, buff := range buffs {
			// Only save system buffs or buffs that don't remove on leave scene
			// This is the inverse of the Java condition which was:
			// if (!buff.isStartFromSystem() || buff.getProp().RemoveOnLeaveScene)
			// So we include when: buff.StartFromSystem && !buff.Prop.RemoveOnLeaveScene
			if !buff.StartFromSystem {
				continue
			}
			if buff.Prop.RemoveOnLeaveScene {
				continue
			}

			formationBuffDO := &FormationBuffDO{
				BuffUId:      buff.UID,
				BuffId:       int32(buff.GetCnfID()),
				Layer:        int32(buff.MLayer),
				ConsumedTime: currTime - buff.CreationTime,
				BSystem:      buff.StartFromSystem,
			}
			buffDatas = append(buffDatas, formationBuffDO)
		}
	}

	return buffDatas
}

// AddBuffFromFormation adds a buff from formation data
// Called when entering scene to restore saved buffs
func (m *ActorBuffModule) AddBuffFromFormation(buffDO *FormationBuffDO) {
	m.AddBuffWithLayer(
		int(buffDO.BuffId),
		m.owner,
		m.owner,
		buffDO.BuffUId,
		0,
		int(buffDO.Layer),
		buffDO.ConsumedTime,
		buffDO.BSystem,
	)
}

// toClient converts FYBuff to BuffData
func toClient(buff *FYBuff, buffState EBuffState) *BuffData {
	return &BuffData{
		State:           int32(buffState),
		BuffCnfId:       int32(buff.GetCnfID()),
		BuffUid:         buff.UID,
		CreationTime:    buff.CreationTime,
		CurLayer:        int32(buff.MLayer),
		StartEffectTime: buff.StartEffectTime,
	}
}

// BuffsChangedToPB converts changed buffs to protobuf format
func (m *ActorBuffModule) BuffsChangedToPB() []*BuffData {
	if len(m.changedBuffs) == 0 {
		return nil
	}

	result := make([]*BuffData, len(m.changedBuffs))
	copy(result, m.changedBuffs)
	return result
}

// BuffsRemovedToPB converts removed buffs to protobuf format
func (m *ActorBuffModule) BuffsRemovedToPB() []*BuffData {
	if len(m.changedBuffs) == 0 {
		return nil
	}

	var removedBuffs []*BuffData
	remainingBuffs := make([]*BuffData, 0)

	for _, buffData := range m.changedBuffs {
		if buffData.State == int32(EBuffState_BUFF_END) {
			removedBuffs = append(removedBuffs, buffData)
		} else {
			remainingBuffs = append(remainingBuffs, buffData)
		}
	}

	m.changedBuffs = remainingBuffs
	return removedBuffs
}

// GetBuffByConfId gets first buff by config ID
func (m *ActorBuffModule) GetBuffByConfId(cnfId int) *FYBuff {
	buff := m.BuffExec.GetBuffByConfId(cnfId)
	if buff == nil {
		return nil
	}
	return buff.(*FYBuff)
}

// FindBuffByClass finds buff by class
func (m *ActorBuffModule) FindBuffByClass(classID int) *FYBuff {
	buff := m.BuffExec.GetBuffByClass(classID)
	if buff == nil {
		return nil
	}
	return buff.(*FYBuff)
}

// FindBuffBySubClass finds buff by subclass
func (m *ActorBuffModule) FindBuffBySubClass(classId, subClassID int) *FYBuff {
	buff := m.BuffExec.GetBuffBySubClass(classId, subClassID)
	if buff == nil {
		return nil
	}
	return buff.(*FYBuff)
}

// FindBuffByCastActor finds buff by caster actor
func (m *ActorBuffModule) FindBuffByCastActor(id int, castActor *actors.Creature) *FYBuff {
	allBuffs := m.getAllBuffsMap()
	for _, buffs := range allBuffs {
		for _, buff := range buffs {
			if buff.CasterActor == castActor && buff.GetCnfID() == id {
				return buff
			}
		}
	}
	return nil
}
