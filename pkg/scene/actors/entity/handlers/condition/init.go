package condition

var (
	// Global handlers list for registration
	handlers []InteractConditionHandler
)

// RegisterHandler registers a new condition handler
func RegisterHandler(handler InteractConditionHandler) {
	handlers = append(handlers, handler)
}

// Init initializes the condition handler manager
func Init() *InteractConditionHandlerManager {
	// Register all handlers
	RegisterHandler(NewCostItemConditionHandler())
	// Can add more handlers here

	return NewInteractConditionHandlerManager(handlers)
}

// GetConditionTypeAlias returns condition type aliases
func GetConditionTypeAlias() map[string]int32 {
	return map[string]int32{
		"CostItem": 1, // CostItem condition type
		// Can add more condition type aliases here
	}
}