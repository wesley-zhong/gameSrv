package condition

// InteractConditionHandler defines the interface for interaction condition handlers
type InteractConditionHandler interface {
	// Check checks if the condition is met (pre-check, no actual operation)
	Check(interactMan interface{}, condition interface{}) bool

	// Execute executes the condition (such as deducting items)
	Execute(interactMan interface{}, condition interface{}) bool

	// GetType returns the condition type
	GetType() int32
}