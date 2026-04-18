package condition

import (
	"log"
)

// InteractConditionHandlerManager manages interaction condition handlers
type InteractConditionHandlerManager struct {
	handlers map[int32]InteractConditionHandler
}

var Instance *InteractConditionHandlerManager

// NewInteractConditionHandlerManager creates a new manager instance
func NewInteractConditionHandlerManager(handlerList []InteractConditionHandler) *InteractConditionHandlerManager {
	mgr := &InteractConditionHandlerManager{
		handlers: make(map[int32]InteractConditionHandler),
	}
	Instance = mgr

	for _, handler := range handlerList {
		mgr.register(handler)
		log.Printf("Registered InteractConditionHandler: type=%d", handler.GetType())
	}
	return mgr
}

// register registers a handler
func (m *InteractConditionHandlerManager) register(handler InteractConditionHandler) {
	m.handlers[handler.GetType()] = handler
}

// GetHandler returns a handler for the given condition
func (m *InteractConditionHandlerManager) GetHandler(condition interface{}) InteractConditionHandler {
	if condition == nil {
		return nil
	}
	// TODO: implement handler lookup based on condition type
	return nil
}

// CheckCondition checks a single condition
func (m *InteractConditionHandlerManager) CheckCondition(condition interface{}, interactMan interface{}) bool {
	handler := m.GetHandler(condition)
	if handler == nil {
		// No registered handler for this condition type, default pass
		return true
	}
	return handler.Check(interactMan, condition)
}

// ExecuteCondition executes a single condition
func (m *InteractConditionHandlerManager) ExecuteCondition(condition interface{}, interactMan interface{}) bool {
	handler := m.GetHandler(condition)
	if handler == nil {
		// No registered handler for this condition type, default pass
		return true
	}
	return handler.Execute(interactMan, condition)
}