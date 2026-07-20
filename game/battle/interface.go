package battle

// Buff represents a buff in the battle system
// This is a placeholder - the actual definition is in the buff package
type Buff struct {
	// TODO: define Buff fields or use buff.FYBuff
}

// CheckBuff is a functional interface for checking buff conditions
type CheckBuff func(buff *Buff) bool

// ConditionCheck is a functional interface for condition checking
type ConditionCheck func(buff *Buff, param0, param1 int64) bool
