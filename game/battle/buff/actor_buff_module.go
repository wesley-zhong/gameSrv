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
	owner           *actors.Creature
	serverBuffExec  *FYBuffExecutionServer
	clientBuffExec  *FYBuffExecutionClient
	changedBuffs    []*BuffData
}

// getBuffExec returns the appropriate buff execution instance
func (m *ActorBuffModule) getBuffExec() *FYBuffExecution {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.FYBuffExecution
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.FYBuffExecution
	}
	return nil
}

// execRemoveBuffByUid executes RemoveBuffByUid on the appropriate executor
func (m *ActorBuffModule) execRemoveBuffByUid(buffCnfId int, uid int64, reason int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.RemoveBuffByUid(buffCnfId, uid, reason)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.RemoveBuffByUid(buffCnfId, uid, reason)
	}
	return false
}

// execAddBuff executes AddBuff on the appropriate executor
func (m *ActorBuffModule) execAddBuff(templateId int, casterActor, holder *actors.Creature, uid int64, exParam, layer int, life int64, bSystem bool) *FYBuff {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.AddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.AddBuff(templateId, casterActor, holder, uid, exParam, layer, life, bSystem)
	}
	return nil
}

// execStopAndRemoveBuff executes StopAndRemoveBuff on the appropriate executor
func (m *ActorBuffModule) execStopAndRemoveBuff(removedBuff *FYBuff, reason int) {
	if m.serverBuffExec != nil {
		m.serverBuffExec.StopAndRemoveBuff(removedBuff, reason)
	} else if m.clientBuffExec != nil {
		m.clientBuffExec.StopAndRemoveBuff(removedBuff, reason)
	}
}

// execTriggerBuffEvent executes TriggerBuffEvent on the appropriate executor
func (m *ActorBuffModule) execTriggerBuffEvent(eventEnum int, triggerEventParam interface{}) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.TriggerBuffEvent(eventEnum, triggerEventParam)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.TriggerBuffEvent(eventEnum, triggerEventParam)
	}
	return false
}

// execHasBuffTag executes HasBuffTag on the appropriate executor
func (m *ActorBuffModule) execHasBuffTag(tagEnum int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.HasBuffTag(tagEnum)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.HasBuffTag(tagEnum)
	}
	return false
}

// execRemoveBuffTag executes RemoveBuffTag on the appropriate executor
func (m *ActorBuffModule) execRemoveBuffTag(inTag int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.RemoveBuffTag(inTag)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.RemoveBuffTag(inTag)
	}
	return false
}

// execAddBuffTag executes AddBuffTag on the appropriate executor
func (m *ActorBuffModule) execAddBuffTag(inTag int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.AddBuffTag(inTag)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.AddBuffTag(inTag)
	}
	return false
}

// execClearBuff executes ClearBuff on the appropriate executor
func (m *ActorBuffModule) execClearBuff(death bool) {
	if m.serverBuffExec != nil {
		m.serverBuffExec.ClearBuff(death)
	} else if m.clientBuffExec != nil {
		m.clientBuffExec.ClearBuff(death)
	}
}

// execClearNotSystemBuff executes ClearNotSystemBuff on the appropriate executor
func (m *ActorBuffModule) execClearNotSystemBuff() {
	if m.serverBuffExec != nil {
		m.serverBuffExec.ClearNotSystemBuff()
	} else if m.clientBuffExec != nil {
		m.clientBuffExec.ClearNotSystemBuff()
	}
}

// execClearLeaveBattleBuff executes ClearLeaveBattleBuff on the appropriate executor
func (m *ActorBuffModule) execClearLeaveBattleBuff() {
	if m.serverBuffExec != nil {
		m.serverBuffExec.ClearLeaveBattleBuff()
	} else if m.clientBuffExec != nil {
		m.clientBuffExec.ClearLeaveBattleBuff()
	}
}

// execRemoveBuffByConfId executes RemoveBuffByConfId on the appropriate executor
func (m *ActorBuffModule) execRemoveBuffByConfId(cnfId int, reason int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.RemoveBuffByConfId(cnfId, reason)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.RemoveBuffByConfId(cnfId, reason)
	}
	return false
}

// execRemoveBuffBySubClass executes RemoveBuffBySubClass on the appropriate executor
func (m *ActorBuffModule) execRemoveBuffBySubClass(classID, subClassID int, reason int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.RemoveBuffBySubClass(classID, subClassID, reason)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.RemoveBuffBySubClass(classID, subClassID, reason)
	}
	return false
}

// execRemoveBuffByClass executes RemoveBuffByClass on the appropriate executor
func (m *ActorBuffModule) execRemoveBuffByClass(classType int, reason int) bool {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.RemoveBuffByClass(classType, reason)
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.RemoveBuffByClass(classType, reason)
	}
	return false
}

// getAllBuffsMap returns the buffs map from the appropriate executor
func (m *ActorBuffModule) getAllBuffsMap() map[int][]*FYBuff {
	if m.serverBuffExec != nil {
		return m.serverBuffExec.allBuffsMap
	}
	if m.clientBuffExec != nil {
		return m.clientBuffExec.allBuffsMap
	}
	return nil
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
	if isServer {
		exec := NewFYBuffExecutionServer(creature)
		return &ActorBuffModule{
			owner:           creature,
			serverBuffExec:  exec,
			changedBuffs:    make([]*BuffData, 0),
		}
	}

	exec := NewFYBuffExecutionClient(creature)
	return &ActorBuffModule{
		owner:           creature,
		clientBuffExec:  exec,
		changedBuffs:    make([]*BuffData, 0),
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
// Returns interface{} to satisfy IActorBuffModule (actually returns *FYBuff)
func (m *ActorBuffModule) AddBuff(
	templateId int,
	casterActor, holder *actors.Creature,
	uid int64,
	exParam int,
	bSystem bool,
) interface{} {
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
	if m.serverBuffExec != nil {
		for _, buffId := range m.serverBuffExec.rebornAddBuffIds {
			m.AddBuff(buffId, m.owner, m.owner, GenProcessLongId(), 0, false)
		}
		m.serverBuffExec.rebornAddBuffIds = make([]int, 0)
	}
}

// TickFromClient handles client-side buff tick
func (m *ActorBuffModule) TickFromClient(buffId int) {
	if m.clientBuffExec != nil {
		// TODO: Implement client tick when needed
		_ = buffId
	}
}

// GetBuffExecution returns the buff execution instance
func (m *ActorBuffModule) GetBuffExecution() *FYBuffExecution {
	return m.getBuffExec()
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
	exec := m.getBuffExec()
	if exec == nil {
		return nil
	}
	return exec.findFirstBuffByConfId(cnfId)
}

// FindBuffByClass finds buff by class
func (m *ActorBuffModule) FindBuffByClass(classID int) *FYBuff {
	exec := m.getBuffExec()
	if exec == nil {
		return nil
	}
	return exec.findBuffByClass(classID)
}

// FindBuffBySubClass finds buff by subclass
func (m *ActorBuffModule) FindBuffBySubClass(classId, subClassID int) *FYBuff {
	exec := m.getBuffExec()
	if exec == nil {
		return nil
	}
	return exec.findBuffBySubClass(classId, subClassID)
}

// FindBuffByCastActor finds buff by caster actor
func (m *ActorBuffModule) FindBuffByCastActor(id int, castActor *actors.Creature) *FYBuff {
	exec := m.getBuffExec()
	if exec == nil {
		return nil
	}
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
