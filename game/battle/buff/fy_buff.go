package buff

import (
	"fmt"
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/battle"
	"gameSrv/game/gamedata"
	"gameSrv/pkg/actors"
	"sync"
)

// INT32_MIN is the minimum value for int32
const INT32_MIN = -2147483648

// FYBuff represents a buff instance on a creature
type FYBuff struct {
	Life                      int
	CasterActor               *actors.Creature
	HolderActor               *actors.Creature
	Prop                      *cfg.Buff
	MLayer                    int
	MaxLayer                  int
	PeriodicCount             int
	EndType                   int
	TickTriggerEvent          *cfg.TriggerEvent
	CreationTime              int64
	UID                       int64
	PeriodicTime              int
	StartEffectTime           int64
	ImmediateExeTriggerEvents []*cfg.TriggerEvent
	ExParam                   int
	StartFromSystem           bool
	EffectHurtMap             map[int]int
	CurrentRemoveCounter      int
	TimerTask                 interface{} // Placeholder for timer task
	TriggerFlag               int
	buffEffectRandom          map[int]*BuffEffectRandom
	mu                        sync.RWMutex
}

// init initializes the buff
func (b *FYBuff) init() {
	b.Life = 0
	b.CasterActor = nil
	b.HolderActor = nil
	b.Prop = nil
	b.MLayer = 0
	b.PeriodicCount = 0
	b.EndType = cfg.BuffEndTypeEnum_BUFF_END_NONE
	b.TickTriggerEvent = nil
	b.CreationTime = battle.MilliSeconds()
	b.ImmediateExeTriggerEvents = make([]*cfg.TriggerEvent, 0)
	b.EffectHurtMap = make(map[int]int)
	b.TriggerFlag = 0
	b.ExParam = INT32_MIN
	b.StartFromSystem = false
}

// reset resets the buff state for pooling
func (b *FYBuff) reset() {
	b.Life = 0
	b.CasterActor = nil
	b.HolderActor = nil
	b.Prop = nil
	b.MLayer = 0
	b.MaxLayer = 0
	b.PeriodicCount = 0
	b.EndType = cfg.BuffEndTypeEnum_BUFF_END_NONE
	b.TickTriggerEvent = nil
	b.ImmediateExeTriggerEvents = b.ImmediateExeTriggerEvents[:0]
	b.EffectHurtMap = make(map[int]int)
	b.TriggerFlag = 0
	b.ExParam = INT32_MIN
	b.StartFromSystem = false
	b.CurrentRemoveCounter = 0
}

// SetProperty sets the buff property configuration
func (b *FYBuff) SetProperty(propConfig *cfg.Buff) {
	b.Prop = propConfig
	b.PeriodicTime = b.GetBuffPeriodicTime(propConfig)
}

// GetCnfID returns the config ID of the buff
func (b *FYBuff) GetCnfID() int {
	if b.Prop == nil {
		return 0
	}
	return int(b.Prop.CnfId)
}

// GetClass returns the class of the buff
func (b *FYBuff) GetClass() int {
	if b.Prop == nil {
		return 0
	}
	return int(b.Prop.Class)
}

// GetSubClass returns the subclass of the buff
func (b *FYBuff) GetSubClass() int {
	if b.Prop == nil {
		return 0
	}
	return int(b.Prop.SubClass)
}

// ClientTick handles client-side buff tick
func (b *FYBuff) ClientTick() bool {
	if b.TickTriggerEvent == nil {
		return false
	}
	execution := GetEffectExecutionInstance()
	effectIds := make([]int, len(b.TickTriggerEvent.EffectIds))
	for i, id := range b.TickTriggerEvent.EffectIds {
		effectIds[i] = int(id)
	}
	execution.ExeBuffEffects(b, effectIds, 1, int(cfg.BuffTriggerEventEnum_TICK))
	return false
}

// ServerTick handles server-side buff tick
func (b *FYBuff) ServerTick() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Prop == nil {
		return false
	}

	totalTime := int(b.Prop.TotalTime)
	if totalTime < 0 {
		fmt.Printf("buff id =%d total time =%d\n", b.Prop.CnfId, b.Prop.TotalTime)
		return false
	}

	now := battle.MilliSeconds()

	// Calculate life based on TotalTimeType
	if b.Prop.TotalTimeType == cfg.BuffTimeStartEnum_START_ON_EFFECTED {
		b.Life = int(now - b.StartEffectTime)
	} else {
		b.Life = int(now - b.CreationTime)
	}

	mTime := b.Life
	if mTime >= totalTime {
		return true
	}

	if b.TickTriggerEvent == nil || b.PeriodicTime == 0 {
		return false
	}

	iCount := mTime / b.PeriodicTime
	iPeriodic := iCount - b.PeriodicCount

	execution := GetEffectExecutionInstance()
	effectIds := make([]int, len(b.TickTriggerEvent.EffectIds))
	for i, id := range b.TickTriggerEvent.EffectIds {
		effectIds[i] = int(id)
	}
	for i := 0; i < iPeriodic; i++ {
		execution.ExeBuffEffects(b, effectIds, 1, int(cfg.BuffTriggerEventEnum_TICK))
	}

	b.PeriodicCount = iCount

	// Notify battle module of buff update
	if b.HolderActor != nil {
		// TODO: Call onAvatarPropsChangFinish and OnBuffUpdate when ActorBattleModule is implemented
	}

	return false
}

// DoBuffEffect executes buff effects with conditions
func (b *FYBuff) DoBuffEffect(buffEffect *cfg.TriggerEvent, exLayer int, reason int) {
	// First check conditions
	for _, cndId := range buffEffect.CondIds {
		if gamedata.Tables == nil || gamedata.Tables.TbBuffCond == nil {
			continue
		}
		buffCond := gamedata.Tables.TbBuffCond.Get(cndId)
		conditionProcess := &FYBuffConditionProcess{}
		if !conditionProcess.CheckCondType(buffCond, int(cndId), CondTypeHoldCond) {
			continue
		}

		if !conditionProcess.CheckConditionValid(b.HolderActor, int(cndId), CondTypeHoldCond) {
			return
		}
	}

	effectIds := make([]int, len(buffEffect.EffectIds))
	for i, id := range buffEffect.EffectIds {
		effectIds[i] = int(id)
	}
	executed := (&FYBuffEffectExecution{}).ExeBuffEffects(b, effectIds, exLayer, reason)
	if executed && b.StartEffectTime == 0 {
		b.SetStartEffectTime(battle.MilliSeconds())
	}
}

// DoFullLayerEffect executes effects when reaching max layer
func (b *FYBuff) DoFullLayerEffect() {
	if b.Prop == nil {
		return
	}

	for _, buffId := range b.Prop.StartBuffOnMaxLayer {
		if buffId == 0 {
			return
		}
		// TODO: Call AddBuff when ActorBuffModule is implemented
		// b.HolderActor.ActorBattleModule.ActorBuffModule.AddBuff(buffId, b.HolderActor, GenProcessLongId(), 0, false)
		_ = buffId
	}
}

// AddLayer adds layers to the buff
func (b *FYBuff) AddLayer(layer int) {
	b.Refresh()
	if b.MaxLayer > b.MLayer {
		toMax := b.MaxLayer - b.MLayer
		if layer > toMax {
			layer = toMax
		}
		b.MLayer += layer

		for _, triggerEvent := range b.ImmediateExeTriggerEvents {
			b.DoBuffEffect(triggerEvent, layer, int(cfg.BuffTriggerEventEnum_ADD_LAYER))
		}

		if b.MLayer == b.MaxLayer {
			b.DoFullLayerEffect()
		}
		return
	}
	b.DoFullLayerEffect()
}

// InitMaxLayer initializes max layer based on config
func (b *FYBuff) InitMaxLayer() {
	if b.Prop == nil {
		return
	}

	b.MaxLayer = int(b.Prop.MaxLayer)

	if b.Prop.BuffSpecialType == cfg.BuffSpecialType_ParryHitType {
		if b.HolderActor != nil {
			// TODO: Get hero data when HeroAvatarActor is implemented
			// heroData := gamedata.Tables.GetTbHeroData().Get(b.HolderActor.GetResID())
			// if heroData != nil && heroData.BreakWeaponCount != 0 {
			//     b.MaxLayer = heroData.BreakWeaponCount
			// }
		}
	}
}

// Start starts the buff effects
func (b *FYBuff) Start(buffTriggerAddRet *BuffTriggerAddRet) {
	b.TickTriggerEvent = buffTriggerAddRet.TickTriggerEvent
	b.ImmediateExeTriggerEvents = append(b.ImmediateExeTriggerEvents, buffTriggerAddRet.ImmediateExeTriggerEvents...)

	b.LockBuffDamageCalc()
	b.DoBuffTags()

	// Initialize layer
	if b.MLayer == 0 {
		if b.Prop != nil && int(b.Prop.InitLayer) == 0 {
			b.MLayer = 1
		} else if b.Prop != nil {
			b.MLayer = int(b.Prop.InitLayer)
		}
	}

	b.InitMaxLayer()

	for _, triggerEvent := range buffTriggerAddRet.ImmediateExeTriggerEvents {
		b.DoBuffEffect(triggerEvent, b.MLayer, int(cfg.BuffTriggerEventEnum_BUFF_MOUNTED))
	}

	// Trigger buff mounted event
	// TODO: Implement TriggerBuffEvent when ready
	// EventParam := &BuffMountedBuffTriggerEventParam{FromBuffId: b.GetCnfID()}
	// b.HolderActor.TriggerBuffEvent(int(cfg.BuffTriggerEventEnum_BUFF_MOUNTED), EventParam)

	// Handle inactive state
	// TODO: Implement when ActorBattleModule is ready

	// Trigger property change events
	if b.Prop != nil {
		for _, triggerEventIf := range b.Prop.TriggerEvents {
			triggerEvent, ok := triggerEventIf.(*cfg.TriggerEvent)
			if !ok {
				continue
			}
			if triggerEvent.EventId == cfg.BuffTriggerEventEnum_PROPS_CHANGE {
				// TODO: Trigger PROPS_CHANGE event
			} else if triggerEvent.EventId == cfg.BuffTriggerEventEnum_PROPS_CHANGE_RATE {
				// TODO: Trigger PROPS_CHANGE_RATE event
			}
		}
	}
}

// LockBuffDamageCalc locks buff damage calculation
func (b *FYBuff) LockBuffDamageCalc() {
	for _, triggerEvent := range b.ImmediateExeTriggerEvents {
		for _, effectId := range triggerEvent.EffectIds {
			buffEffect := gamedata.Tables.TbBuffEffect.Get(effectId)
			if buffEffect == nil {
				continue
			}
			if buffEffect.BuffEffectOpt == nil {
				continue
			}

			b.lockCalcByBuff(buffEffect.BuffEffectOpt, int(effectId))
		}
	}

	if b.TickTriggerEvent != nil {
		for _, effectId := range b.TickTriggerEvent.EffectIds {
			buffEffect := gamedata.Tables.TbBuffEffect.Get(effectId)
			if buffEffect == nil {
				continue
			}
			if buffEffect.BuffEffectOpt == nil {
				continue
			}

			b.lockCalcByBuff(buffEffect.BuffEffectOpt, int(effectId))
		}
	}
}

// lockCalcByBuff locks damage calculation by buff
func (b *FYBuff) lockCalcByBuff(buffEffectOpt interface{}, effectId int) {
	// Check if it's damage formulation
	// TODO: Implement type checking for DoAbnormalDamageFormulation
	// This requires reflection or type assertions on the protobuf types

	if _, exists := b.EffectHurtMap[effectId]; !exists {
		// Calculate and store damage value
		// b.EffectHurtMap[effectId] = b.CalcDamageByExpression(formulationId, caster, target, b.ExParam)
		b.EffectHurtMap[effectId] = 0
	}
}

// CalcDamageByExpression calculates damage using expression
func (b *FYBuff) CalcDamageByExpression(id int, caster, target *actors.Creature, input int) int {
	// TODO: Implement formula calculation
	return 0
}

// Refresh resets buff life
func (b *FYBuff) Refresh() {
	b.Life = 0
}

// DoBuffTags adds buff tags to holder
func (b *FYBuff) DoBuffTags() {
	if b.Prop == nil || b.HolderActor == nil {
		return
	}

	for _, buffTag := range b.Prop.BuffTags {
		// TODO: Call AddBuffTag when implemented
		_ = buffTag
	}
}

// RemoveBuffTags removes buff tags from holder
func (b *FYBuff) RemoveBuffTags() {
	if b.Prop == nil || b.HolderActor == nil {
		return
	}

	for _, buffTag := range b.Prop.BuffTags {
		// TODO: Call RemoveBuffTag when implemented
		_ = buffTag
	}
}

// RemoveLayer removes layers from buff
func (b *FYBuff) RemoveLayer(revertLayer int) bool {
	if revertLayer == 0 || revertLayer >= b.MLayer {
		return true // Go to stop flow
	}

	b.MLayer -= revertLayer
	(&FYBuffEffectExecution{}).RevertBuffAllEffectsWithLayer(b, revertLayer)
	b.Refresh()
	return false
}

// Stop stops the buff
func (b *FYBuff) Stop(endType int) bool {
	revertLayer := 0
	if endType == cfg.BuffEndTypeEnum_BUFF_END_TIMEUP && b.Prop != nil {
		revertLayer = int(b.Prop.RemoveLayerOnTimeOut)
		if revertLayer > b.MLayer {
			revertLayer = b.MLayer
		}
	}

	b.MLayer -= revertLayer
	(&FYBuffEffectExecution{}).RevertBuffAllEffectsWithLayer(b, revertLayer)
	b.Refresh()

	if revertLayer == 0 || b.MLayer == 0 {
		b.EndType = endType
		b.Destroy()
		return true
	}
	return false
}

// Destroy cleans up buff resources
func (b *FYBuff) Destroy() {
	b.RemoveBuffTags()
	if b.TimerTask != nil {
		// TODO: Cancel timer task when timer system is implemented
		b.TimerTask = nil
	}
}

// GetBuffPeriodicTime gets periodic time from buff config
func (b *FYBuff) GetBuffPeriodicTime(buff *cfg.Buff) int {
	if buff == nil {
		return 0
	}

	for _, triggerEventIf := range buff.TriggerEvents {
		triggerEvent, ok := triggerEventIf.(*cfg.TriggerEvent)
		if !ok {
			continue
		}
		if triggerEvent.EventId == cfg.BuffTriggerEventEnum_BUFF_MOUNTED {
			// Try to cast to CommonTriggerEvent to get PeriodicTime
			// TODO: Implement proper type checking when CommonTriggerEvent type is available
			_ = triggerEvent
			return 0
		}
	}
	return 0
}

// SetStartEffectTime sets the start effect time
func (b *FYBuff) SetStartEffectTime(time int64) {
	b.StartEffectTime = time
}

// IsFullLayer checks if buff is at max layer
func (b *FYBuff) IsFullLayer() bool {
	return b.MLayer == b.MaxLayer
}

// TriggerBuffDeathCounterEvent triggers death counter event
func (b *FYBuff) TriggerBuffDeathCounterEvent() {
	b.CurrentRemoveCounter++
	if b.CheckRemoveCounterOver() {
		// TODO: Call StopAndRemoveBuff when ActorBuffModule is implemented
		_ = b.HolderActor
	}
}

// CheckRemoveCounterOver checks if remove counter is exceeded
func (b *FYBuff) CheckRemoveCounterOver() bool {
	if b.Prop == nil {
		return false
	}
	return b.CurrentRemoveCounter >= int(b.Prop.BuffDeathCounter)
}

// GetBuffEffectRandom returns buff effect random map
func (b *FYBuff) GetBuffEffectRandom() map[int]*BuffEffectRandom {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.buffEffectRandom == nil {
		b.buffEffectRandom = make(map[int]*BuffEffectRandom)
	}
	return b.buffEffectRandom
}

// SetBuffEffectRandom sets buff effect random map
func (b *FYBuff) SetBuffEffectRandom(m map[int]*BuffEffectRandom) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.buffEffectRandom = m
}
