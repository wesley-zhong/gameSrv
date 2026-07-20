package process

// BuffAddCheckProcess provides buff addition check functions
type BuffAddCheckProcess struct{}

// CheckBuffAddFromBuff checks if buff can be added from buff
func (b *BuffAddCheckProcess) CheckBuffAddFromBuff(caster, target *Creature, buffAddPush *BuffData) bool {
	return true
}

// CheckBuffAddFromAttackData checks if buff can be added from attack data
func (b *BuffAddCheckProcess) CheckBuffAddFromAttackData(caster, target *Creature, buffAddPush *BuffData) bool {
	return true
}

// CheckBuffAddFromMontageCreature checks if buff can be added from montage creature
func (b *BuffAddCheckProcess) CheckBuffAddFromMontageCreature(caster, target *Creature, buffAddPush *BuffData) bool {
	return true
}

// Global instance
var BuffAddCheck = &BuffAddCheckProcess{}