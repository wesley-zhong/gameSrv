package interact_handles

// CustomInteractable is a custom interact interface
type CustomInteractable interface {
	// Handle processes interaction
	Handle(targetActor interface{}, srcActor interface{}, optionType int32, optionParams []int64) int32
	// GetCustomHandleType gets the custom handler type
	GetCustomHandleType() int32
}