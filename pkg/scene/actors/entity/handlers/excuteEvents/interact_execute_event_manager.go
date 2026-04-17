package excuteEvents

import (
	"log"
)

// InteractExecuteEventManager manages interaction post-event handlers
type InteractExecuteEventManager struct {
	handlers map[int]InteractExecuteEventHandler
}

// Instance is the singleton event manager
var Instance *InteractExecuteEventManager

// NewInteractExecuteEventManager creates a new event manager
func NewInteractExecuteEventManager(handlerList []InteractExecuteEventHandler) *InteractExecuteEventManager {
	manager := &InteractExecuteEventManager{
		handlers: make(map[int]InteractExecuteEventHandler),
	}

	Instance = manager

	for _, handler := range handlerList {
		manager.Register(handler)
		log.Printf("Registered InteractExecuteEventHandler: type=%d", handler.GetType())
	}

	return manager
}

// Register registers an event handler
func (m *InteractExecuteEventManager) Register(handler InteractExecuteEventHandler) {
	m.handlers[handler.GetType()] = handler
}

// ExecuteEvent executes a single post-event
func (m *InteractExecuteEventManager) ExecuteEvent(eventConfig interface{}, interactMan interface{}, targetActor interface{}) error {
	// TODO: implement event execution logic
	return nil
}

// ExecuteEvents executes all post-events
func (m *InteractExecuteEventManager) ExecuteEvents(conditionList []interface{}, optionType int32, interactMan interface{}, targetActor interface{}) error {
	// TODO: implement events execution logic
	return nil
}