package rebornhandler

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

// JumpCliffReborn handles cliff jump reborn
type JumpCliffReborn struct{}

// NewJumpCliffReborn creates a new JumpCliffReborn
func NewJumpCliffReborn() *JumpCliffReborn {
	return &JumpCliffReborn{}
}

// Reborn handles reborn process
func (j *JumpCliffReborn) Reborn(entity scene.IEntity, position *math.Vector3, rotation *math.Vector3) bool {
	// TODO: implement reborn logic
	return true
}

// GetRebornReason returns reborn reason
func (j *JumpCliffReborn) GetRebornReason() int32 {
	return 0
}