package interact_handles

// WakeMonsterInteractHandle is the wake monster interaction handler
type WakeMonsterInteractHandle struct{}

// NewWakeMonsterInteractHandle creates a new wake monster interact handler
func NewWakeMonsterInteractHandle() *WakeMonsterInteractHandle {
	return &WakeMonsterInteractHandle{}
}

// Handle handles interaction
func (h *WakeMonsterInteractHandle) Handle(targetActor interface{}, srcActor interface{}, optionType int32, optionParams []int64) int32 {
	// TODO: implement
	return 0 // Success
}