package battle

// IBuff defines the interface for buff objects
// This avoids circular import with buff package
type IBuff interface {
	GetCnfID() int
	GetUID() int64
}

// CheckBuff is a functional interface for checking buff conditions
type CheckBuff func(buff IBuff) bool

// ConditionCheck is a functional interface for condition checking
type ConditionCheck func(buff IBuff, param0, param1 int64) bool
