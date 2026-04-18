package handlers

// ActorDropType defines actor drop type
type ActorDropType int32

const (
	ActorDropTypeDropgroup ActorDropType = 1
	ActorDropTypeItem      ActorDropType = 2
)

// UtilityDropInterface defines utility drop behavior
type UtilityDropInterface interface {
	// Drop executes drop logic
	Drop(initiator interface{})

	// GetUtilityDropData returns drop data
	GetUtilityDropData() interface{}

	// GetObjectState returns object state
	GetObjectState() int32
}

// DefaultUtilityDropInterface provides default utility drop implementation
type DefaultUtilityDropInterface struct{}

// Drop executes default drop logic
func (d *DefaultUtilityDropInterface) Drop(initiator interface{}) {
	// TODO: implement drop logic
}