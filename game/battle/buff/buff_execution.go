package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/battle"
	"gameSrv/game/gamedata"
	"gameSrv/pkg/actors"
	"sync"
)

// CheckBuffCondition is a function type for checking buff conditions
type CheckBuffCondition func(buff *FYBuff, param0, param1 int64) bool

// FYBuffExecution is the base class for buff execution
type FYBuffExecution struct {
	owner                 *actors.Creature
	allBuffsMap           map[int][]*FYBuff
	rebornAddBuffIds      []int
	sameBuffMountedLayer  map[int]int
	fyBuffEventTriggerPtr *FYBuffEventTrigger
	mTagSigner            []int
	lastTickTime          int64
	mu                    sync.RWMutex
}

// NewFYBuffExecution creates a new FYBuffExecution
func NewFYBuffExecution(owner *actors.Creature) *FYBuffExecution {
	return &FYBuffExecution{
		owner:                 owner,
		allBuffsMap:           make(map[int][]*FYBuff),
		rebornAddBuffIds:      make([]int, 0),
		sameBuffMountedLayer:  make(map[int]int),
		fyBuffEventTriggerPtr: NewFYBuffEventTrigger(owner),
		mTagSigner:            make([]int, cfg.BuffTagEnum_TAG_MAX),
		lastTickTime:          battle.MilliSeconds(),
	}
}

// TriggerBuffEvent triggers a buff event
func (e *FYBuffExecution) TriggerBuffEvent(
	eventEnum int,
	triggerEventParam interface{},
) bool {
	eventBuffs := e.fyBuffEventTriggerPtr.GetBuffsByEvent(eventEnum)
	if len(eventBuffs) == 0 {
		return false
	}

	changed := false
	triggerProcess := NewFYBuffTriggerProcess()

	// Create copy to avoid concurrent modification
	for _, buffEvent := range eventBuffs {
		buffTemplateId := buffEvent.BuffTempID
		buffProp := e.getBuffProperty(buffTemplateId)
		if buffProp == nil {
			continue
		}

		triggerEvent := e.findBuffTriggerEventCnf(buffProp, eventEnum, buffEvent.EventType)
		if triggerEvent == nil {
			continue
		}

		if e.doTriggerOpt(triggerEventParam, triggerEvent, buffTemplateId, eventEnum, buffEvent.EventType, triggerProcess) {
			changed = true
		}
	}

	if changed {
		e.onBuffTriggerEventFinished()
	}

	return changed
}

// findBuffTriggerEventCnf finds trigger event configuration
func (e *FYBuffExecution) findBuffTriggerEventCnf(
	buffProp *cfg.Buff,
	eventEnum int,
	eventType EBuffEventType,
) *cfg.TriggerEvent {
	if buffProp == nil {
		return nil
	}

	var eventListIf []interface{}
	if eventType == EventType_NormalTrigger {
		eventListIf = buffProp.TriggerEvents
	} else {
		eventListIf = buffProp.RemoveTriggerEvents
	}

	for _, triggerEventIf := range eventListIf {
		triggerEvent, ok := triggerEventIf.(*cfg.TriggerEvent)
		if !ok {
			continue
		}
		if triggerEvent.EventId == int32(eventEnum) {
			return triggerEvent
		}
	}

	return nil
}

// doTriggerOpt performs trigger operation
func (e *FYBuffExecution) doTriggerOpt(
	triggerEventParam interface{},
	triggerEvent *cfg.TriggerEvent,
	buffTemplateId int,
	inReason int,
	eventType EBuffEventType,
	triggerProcess *FYBuffTriggerProcess,
) bool {
	if !triggerProcess.CheckEvent(triggerEventParam, triggerEvent) {
		return false
	}

	e.buffOnOptExeBuffEffect(triggerEventParam, buffTemplateId, triggerEvent, inReason, eventType)
	return true
}

// buffOnOptExeBuffEffect executes buff effect on trigger
func (e *FYBuffExecution) buffOnOptExeBuffEffect(
	param interface{},
	buffTemplateId int,
	lubanTriggerEvent *cfg.TriggerEvent,
	inReason int,
	eventType EBuffEventType,
) {
	buffs := e.allBuffsMap[buffTemplateId]
	if len(buffs) == 0 {
		return
	}

	// Create copy to avoid concurrent modification
	buffsCopy := make([]*FYBuff, len(buffs))
	copy(buffsCopy, buffs)

	for _, buff := range buffsCopy {
		// Check trigger flag
		if baseParam, ok := param.(*BuffTriggerEventParamBase); ok {
			if buff.TriggerFlag != 0 && buff.TriggerFlag != baseParam.TriggerFlag {
				continue
			}
		}

		buff.DoBuffEffect(lubanTriggerEvent, 1, inReason)
		e.triggerRemoveBuffCounterEvent(buff, eventType)
	}
}

// triggerRemoveBuffCounterEvent triggers remove buff counter event
func (e *FYBuffExecution) triggerRemoveBuffCounterEvent(buff *FYBuff, buffEvent EBuffEventType) {
	if buffEvent == EventType_RemoveTrigger {
		buff.TriggerBuffDeathCounterEvent()
	}
}

// AddBuff adds a buff to the creature
func (e *FYBuffExecution) AddBuff(
	templateId int,
	casterActor, holder *actors.Creature,
	uid int64,
	exParam, layer int,
	life int64,
	bSystem bool,
) *FYBuff {
	if templateId == 0 {
		return nil
	}

	if holder == nil {
		holder = e.owner
	}

	// Check if can add buff
	// TODO: Check invincible state and health when ActorBattleModule is implemented

	buffProp := e.getBuffProperty(templateId)
	if buffProp == nil || int(buffProp.CnfId) != templateId {
		return nil
	}

	if !e.canAddBuff(buffProp, casterActor) {
		return nil
	}

	// Check buff replace rules
	if e.checkBuffReplace(templateId, buffProp, casterActor) {
		return nil
	}

	bufUid := uid
	if bufUid == 0 {
		bufUid = GenProcessLongId()
	}

	// Create new buff
	buff := GetBuffPoolInstance().NewBuff(bufUid)
	buff.CasterActor = casterActor
	buff.MLayer = layer
	buff.SetProperty(buffProp)
	buff.HolderActor = holder
	buff.CreationTime = battle.MilliSeconds() - life
	buff.ExParam = exParam
	buff.StartFromSystem = bSystem

	// Add to buffs map
	e.allBuffsMap[buff.GetCnfID()] = append(e.allBuffsMap[buff.GetCnfID()], buff)

	// Handle together buffs
	if e.buffIsTogether(buffProp) {
		e.sameBuffMountedLayer[int(buffProp.CnfId)]++
	}

	// Add trigger events
	buffTriggerAddRet := e.fyBuffEventTriggerPtr.AddBuffTriggerEvent(buff, templateId)

	// Start buff
	buff.Start(buffTriggerAddRet)

	// Check if buff should be removed immediately
	if buff.Prop.TotalTime == 0 {
		e.StopAndRemoveBuff(buff, cfg.BuffEndTypeEnum_BUFF_END_TIMEUP)
		return nil
	}

	e.onBuffAdded(buff)
	return buff
}

// StopAndRemoveBuff stops and removes a buff
func (e *FYBuffExecution) StopAndRemoveBuff(removedBuff *FYBuff, reason int) {
	GetBuffPoolInstance().DelBuff(removedBuff)

	if removedBuff.Prop == nil {
		return
	}

	cnfId := int(removedBuff.Prop.CnfId)
	buffs := e.allBuffsMap[cnfId]
	if len(buffs) == 0 {
		return
	}

	for i, buff := range buffs {
		if buff.UID == removedBuff.UID {
			if buff.Stop(reason) {
				// Remove from slice
				e.allBuffsMap[cnfId] = append(
					buffs[:i],
					buffs[i+1:]...,
				)
				e.removeSameBuffRef(buff.Prop)

				if len(e.allBuffsMap[cnfId]) == 0 {
					e.removeBuffTriggers(buff.Prop)
					delete(e.allBuffsMap, cnfId)
				}
			}
			break
		}
	}

	// Handle rebirth buffs
	if reason == int(cfg.BuffEndTypeEnum_BUFF_END_TIMEUP) && removedBuff.Prop != nil {
		for _, startBufId := range removedBuff.Prop.OnNaturalDeadAddBuffIds {
			if startBufId > 0 {
				// TODO: Add buff on rebirth
				_ = startBufId
			}
		}
	} else if removedBuff.Prop != nil {
		for _, startBufId := range removedBuff.Prop.OnUnnaturalDeathAddBuffIds {
			if startBufId > 0 {
				// TODO: Add buff
				_ = startBufId
			}
		}
	}
}

// RemoveSameBuffRef removes reference to same buff
func (e *FYBuffExecution) removeSameBuffRef(prop *cfg.Buff) {
	if !e.buffIsTogether(prop) {
		return
	}

	cnfId := int(prop.CnfId)
	curLayer := e.sameBuffMountedLayer[cnfId]
	if curLayer > 0 {
		curLayer--
		if curLayer == 0 {
			delete(e.sameBuffMountedLayer, cnfId)
		} else {
			e.sameBuffMountedLayer[cnfId] = curLayer
		}
	}
}

// removeBuffTriggers removes buff triggers
func (e *FYBuffExecution) removeBuffTriggers(prop *cfg.Buff) {
	e.fyBuffEventTriggerPtr.RemoveBuffFromTriggerEvent(int(prop.CnfId))
}

// HasBuffTag checks if creature has a buff tag
func (e *FYBuffExecution) HasBuffTag(tagEnum int) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if tagEnum >= len(e.mTagSigner) {
		return false
	}
	return e.mTagSigner[tagEnum] > 0
}

// RemoveBuffTag removes a buff tag
func (e *FYBuffExecution) RemoveBuffTag(inTag int) bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	if inTag >= len(e.mTagSigner) {
		return false
	}
	if e.mTagSigner[inTag] > 0 {
		e.mTagSigner[inTag]--
	}
	return true
}

// AddBuffTag adds a buff tag
func (e *FYBuffExecution) AddBuffTag(inTag int) bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	if inTag >= len(e.mTagSigner) {
		return false
	}
	e.mTagSigner[inTag]++
	return true
}

// buffIsTogether checks if buff can stack
func (e *FYBuffExecution) buffIsTogether(buffProp *cfg.Buff) bool {
	if buffProp == nil {
		return false
	}
	return buffProp.OverlapByID == cfg.BuffOverLapEnum_BUFF_TOGETHER ||
		buffProp.OverlapBySubClass == cfg.BuffOverLapEnum_BUFF_TOGETHER ||
		buffProp.OverlapByClass == cfg.BuffOverLapEnum_BUFF_TOGETHER
}

// canAddBuff checks if buff can be added
func (e *FYBuffExecution) canAddBuff(buffToAdd *cfg.Buff, casterActor *actors.Creature) bool {
	return buffToAdd != nil
}

// checkBuffReplace checks buff replacement rules
func (e *FYBuffExecution) checkBuffReplace(id int, buffInfo *cfg.Buff, castActor *actors.Creature) bool {
	existBuff := e.findFirstBuffByConfId(id)
	if existBuff != nil {
		return e.checkBuffReplaceByConfId(existBuff, buffInfo, int(existBuff.Prop.OverlapByID))
	}

	existBuff = e.findBuffBySubClass(int(buffInfo.Class), int(buffInfo.SubClass))
	if existBuff != nil {
		return e.checkBuffReplaceBySubClassId(existBuff, buffInfo, int(existBuff.Prop.OverlapBySubClass))
	}

	existBuff = e.findBuffByClass(int(buffInfo.Class))
	if existBuff != nil {
		return e.checkBuffReplaceByClassId(existBuff, buffInfo, int(existBuff.Prop.OverlapByClass))
	}

	return false
}

// checkBuffReplaceByConfId checks buff replacement by config ID
func (e *FYBuffExecution) checkBuffReplaceByConfId(existBuff *FYBuff, buffToAdd *cfg.Buff, overlapEnum int) bool {
	buffOverlap := e.checkFullLayerReplace(existBuff, overlapEnum)

	switch buffOverlap {
	case int(cfg.BuffOverLapEnum_BUFF_CONFLICT):
		return true
	case int(cfg.BuffOverLapEnum_BUFF_TOGETHER):
		nowSameLayer := e.sameBuffMountedLayer[int(existBuff.Prop.CnfId)]
		if nowSameLayer >= existBuff.MaxLayer {
			e.StopAndRemoveBuff(existBuff, int(cfg.BuffEndTypeEnum_BUFF_END_REPLACE))
		}
		return false
	case int(cfg.BuffOverLapEnum_BUFF_FRESH):
		exLayer := 1
		if existBuff.Prop.AddLayerOnRepeatMounted != 0 {
			exLayer = int(existBuff.Prop.AddLayerOnRepeatMounted)
		}
		existBuff.AddLayer(exLayer)
		return true
	default:
		return true
	}
}

// checkBuffReplaceBySubClassId checks buff replacement by subclass
func (e *FYBuffExecution) checkBuffReplaceBySubClassId(existBuff *FYBuff, buffToAdd *cfg.Buff, overlapEnum int) bool {
	switch overlapEnum {
	case int(cfg.BuffOverLapEnum_BUFF_CONFLICT):
		return true
	case int(cfg.BuffOverLapEnum_BUFF_TOGETHER):
		return false
	default:
		if existBuff.Prop.Level < buffToAdd.Level {
			e.StopAndRemoveBuff(existBuff, int(cfg.BuffEndTypeEnum_BUFF_END_REPLACE))
			return false
		}
		return true
	}
}

// checkBuffReplaceByClassId checks buff replacement by class
func (e *FYBuffExecution) checkBuffReplaceByClassId(existBuff *FYBuff, buffToAdd *cfg.Buff, overlapEnum int) bool {
	switch overlapEnum {
	case int(cfg.BuffOverLapEnum_BUFF_CONFLICT):
		return true
	case int(cfg.BuffOverLapEnum_BUFF_TOGETHER), int(cfg.BuffOverLapEnum_BUFF_FRESH):
		return false
	default:
		if existBuff.Prop.Level < buffToAdd.Level {
			e.StopAndRemoveBuff(existBuff, int(cfg.BuffEndTypeEnum_BUFF_END_REPLACE))
			return false
		}
		return true
	}
}

// checkFullLayerReplace checks full layer replace condition
func (e *FYBuffExecution) checkFullLayerReplace(existBuff *FYBuff, overlapEnum int) int {
	buffOverlap := overlapEnum
	if existBuff.IsFullLayer() && existBuff.Prop.IsShareOnMaxLayer {
		buffOverlap = int(cfg.BuffOverLapEnum_BUFF_CONFLICT)
	}
	return buffOverlap
}

// findFirstBuffByConfId finds first buff by config ID
func (e *FYBuffExecution) findFirstBuffByConfId(cnfId int) *FYBuff {
	buffs := e.allBuffsMap[cnfId]
	if len(buffs) == 0 {
		return nil
	}
	return buffs[0]
}

// findBuffByClass finds buff by class
func (e *FYBuffExecution) findBuffByClass(classID int) *FYBuff {
	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			if buff.Prop != nil && int(buff.Prop.Class) == classID {
				return buff
			}
		}
	}
	return nil
}

// findBuffBySubClass finds buff by subclass
func (e *FYBuffExecution) findBuffBySubClass(classId, subClassID int) *FYBuff {
	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			if buff.Prop != nil && int(buff.Prop.Class) == classId && int(buff.Prop.SubClass) == subClassID {
				return buff
			}
		}
	}
	return nil
}

// getBuffProperty gets buff property by ID
func (e *FYBuffExecution) getBuffProperty(buffId int) *cfg.Buff {
	if gamedata.Tables == nil || gamedata.Tables.TbBuff == nil {
		return nil
	}
	return gamedata.Tables.TbBuff.Get(int32(buffId))
}

// RemoveBuffByConfId removes buffs by config ID
func (e *FYBuffExecution) RemoveBuffByConfId(cnfId int, reason int) bool {
	buffs := e.allBuffsMap[cnfId]
	if len(buffs) == 0 {
		return false
	}

	for _, buff := range buffs {
		e.StopAndRemoveBuff(buff, reason)
	}
	return true
}

// RemoveBuffBySubClass removes buffs by subclass
func (e *FYBuffExecution) RemoveBuffBySubClass(classID, subClassID int, reason int) bool {
	removed := false

	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			condMatch := false
			if classID == 0 {
				condMatch = buff.GetSubClass() == subClassID
			} else {
				condMatch = buff.GetClass() == classID && buff.GetSubClass() == subClassID
			}

			if condMatch {
				e.StopAndRemoveBuff(buff, reason)
				removed = true
			}
		}
	}

	return removed
}

// RemoveBuffByClass removes buffs by class
func (e *FYBuffExecution) RemoveBuffByClass(classType int, reason int) bool {
	removed := false

	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			if buff.Prop != nil && int(buff.Prop.Class) == classType {
				e.StopAndRemoveBuff(buff, reason)
				removed = true
			}
		}
	}

	return removed
}

// RemoveBuffByUid removes buff by UID
func (e *FYBuffExecution) RemoveBuffByUid(cnfId int, bufUid int64, reason int) bool {
	if bufUid == 0 {
		return e.RemoveBuffByConfId(cnfId, reason)
	}

	buffs := e.allBuffsMap[cnfId]
	for _, buff := range buffs {
		if buff.UID == bufUid {
			e.StopAndRemoveBuff(buff, reason)
			return true
		}
	}
	return false
}

// ClearBuff clears all buffs
func (e *FYBuffExecution) ClearBuff(death bool) {
	if death {
		for _, buffs := range e.allBuffsMap {
			for _, buff := range buffs {
				if !buff.Prop.RemoveOnHolderDead {
					e.rebornAddBuffIds = append(e.rebornAddBuffIds, buff.GetCnfID())
				}
			}
		}
	}

	for _, buffs := range e.allBuffsMap {
		var prop *cfg.Buff
		for _, buff := range buffs {
			prop = buff.Prop
			buff.Stop(cfg.BuffEndTypeEnum_BUFF_END_HAND)
			GetBuffPoolInstance().DelBuff(buff)
		}
		if prop != nil {
			e.removeBuffTriggers(prop)
		}
	}

	e.allBuffsMap = make(map[int][]*FYBuff)
}

// ClearNotSystemBuff clears non-system buffs
func (e *FYBuffExecution) ClearNotSystemBuff() {
	var buffsToRemove []*FYBuff

	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			if !buff.StartFromSystem {
				buffsToRemove = append(buffsToRemove, buff)
			}
		}
	}

	for _, buff := range buffsToRemove {
		e.StopAndRemoveBuff(buff, cfg.BuffEndTypeEnum_BUFF_END_HAND)
	}
}

// ClearLeaveBattleBuff clears buffs on leave battle
func (e *FYBuffExecution) ClearLeaveBattleBuff() {
	var buffsToRemove []*FYBuff

	for _, buffs := range e.allBuffsMap {
		for _, buff := range buffs {
			if buff.Prop.RemoveOnOutOfCombat {
				buffsToRemove = append(buffsToRemove, buff)
			}
		}
	}

	for _, buff := range buffsToRemove {
		e.StopAndRemoveBuff(buff, cfg.BuffEndTypeEnum_BUFF_END_HAND)
	}
}

// onBuffAdded is called when buff is added
func (e *FYBuffExecution) onBuffAdded(buff *FYBuff) {
	// To be implemented by subclasses
}

// onBuffTriggerEventFinished is called when buff trigger event finishes
func (e *FYBuffExecution) onBuffTriggerEventFinished() {
	// To be implemented by subclasses
}

// onPlayerLeaveScene is called when player leaves scene
func (e *FYBuffExecution) onPlayerLeaveScene() {
	// To be implemented by subclasses
}

// synBuffToClient syncs buff to client
func (e *FYBuffExecution) synBuffToClient() {
	// To be implemented by subclasses
}
