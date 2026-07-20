package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// EBuffEventType represents buff event type
type EBuffEventType int

const (
	EventType_NormalTrigger EBuffEventType = 0 // Normal trigger event
	EventType_RemoveTrigger EBuffEventType = 1 // Remove trigger event
)

// FBuffEvent represents a buff event
type FBuffEvent struct {
	EventType  EBuffEventType
	BuffTempID int
}

// FYBuffEventTrigger manages buff event triggers for a creature
type FYBuffEventTrigger struct {
	owner             *actors.Creature
	triggerEventBuffs map[int]map[int]*FBuffEvent
	mountedBuffs      map[int]bool
}

// NewFYBuffEventTrigger creates a new FYBuffEventTrigger
func NewFYBuffEventTrigger(owner *actors.Creature) *FYBuffEventTrigger {
	return &FYBuffEventTrigger{
		owner:             owner,
		triggerEventBuffs: make(map[int]map[int]*FBuffEvent),
		mountedBuffs:      make(map[int]bool),
	}
}

// AddBuffTriggerEvent adds trigger events for a buff
func (t *FYBuffEventTrigger) AddBuffTriggerEvent(buff *FYBuff, buffTemplateId int) *BuffTriggerAddRet {
	buffTriggerAddRet := &BuffTriggerAddRet{
		ImmediateExeTriggerEvents: make([]*cfg.TriggerEvent, 0),
	}

	if buff.Prop == nil {
		return buffTriggerAddRet
	}

	for _, triggerEventIf := range buff.Prop.TriggerEvents {
		triggerEvent, ok := triggerEventIf.(*cfg.TriggerEvent)
		if !ok {
			continue
		}

		if buff.PeriodicTime > 0 {
			buffTriggerAddRet.TickTriggerEvent = triggerEvent
			buffTriggerAddRet.ImmediateExeTriggerEvents = append(
				buffTriggerAddRet.ImmediateExeTriggerEvents,
				triggerEvent,
			)
			continue
		}

		if triggerEvent.EventId == cfg.BuffTriggerEventEnum_BUFF_MOUNTED {
			buffTriggerAddRet.ImmediateExeTriggerEvents = append(
				buffTriggerAddRet.ImmediateExeTriggerEvents,
				triggerEvent,
			)
			continue
		}

		t.addBuffTriggerEvent(
			int(triggerEvent.EventId),
			EventType_NormalTrigger,
			buffTemplateId,
		)
	}

	// Handle remove trigger events
	for _, removeTriggerEventIf := range buff.Prop.RemoveTriggerEvents {
		removeTriggerEvent, ok := removeTriggerEventIf.(*cfg.TriggerEvent)
		if !ok {
			continue
		}

		if removeTriggerEvent.EventId != cfg.BuffTriggerEventEnum_NEVER {
			t.addBuffTriggerEvent(
				int(removeTriggerEvent.EventId),
				EventType_RemoveTrigger,
				buffTemplateId,
			)
		}
	}

	return buffTriggerAddRet
}

// addBuffTriggerEvent adds a single buff trigger event
func (t *FYBuffEventTrigger) addBuffTriggerEvent(
	eventEnum int,
	eventType EBuffEventType,
	buffTemplateId int,
) {
	if _, exists := t.triggerEventBuffs[eventEnum]; !exists {
		t.triggerEventBuffs[eventEnum] = make(map[int]*FBuffEvent)
	}

	t.triggerEventBuffs[eventEnum][buffTemplateId] = &FBuffEvent{
		EventType:  eventType,
		BuffTempID: buffTemplateId,
	}
	t.mountedBuffs[buffTemplateId] = true
}

// RemoveBuffFromTriggerEvent removes buff from trigger events
func (t *FYBuffEventTrigger) RemoveBuffFromTriggerEvent(templateId int) {
	delete(t.mountedBuffs, templateId)

	for eventEnum := range t.triggerEventBuffs {
		if buffs, exists := t.triggerEventBuffs[eventEnum]; exists {
			delete(buffs, templateId)
		}
	}
}

// Clear clears all trigger events
func (t *FYBuffEventTrigger) Clear() {
	t.mountedBuffs = make(map[int]bool)
	t.triggerEventBuffs = make(map[int]map[int]*FBuffEvent)
}

// GetBuffsByEvent returns buffs for a specific event
func (t *FYBuffEventTrigger) GetBuffsByEvent(eventEnum int) map[int]*FBuffEvent {
	return t.triggerEventBuffs[eventEnum]
}

// GetOwner returns the owner creature
func (t *FYBuffEventTrigger) GetOwner() *actors.Creature {
	return t.owner
}
