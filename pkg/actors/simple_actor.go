package actors

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

// SimpleActor is an immovable simple entity
type SimpleActor struct {
	Entity
}

func (s *SimpleActor) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleActor) GetOnlyVisionUid() int64 {
	return 0
}

func (s *SimpleActor) GetPhasingId() int64 {
	return s.PhasingID
}

// NewSimpleActor creates a new SimpleActor
func NewSimpleActor() *SimpleActor {
	s := &SimpleActor{}
	s.Entity = *NewEntity()
	s.MoveInfo = nil
	return s
}

// CanMove returns whether this actors can move
func (s *SimpleActor) CanMove() bool {
	return false
}

// GetSpeed returns the speed (always zero for SimpleActor)
func (s *SimpleActor) GetSpeed() *math.Vector3 {
	return math.ZeroVector3()
}

// SetSpeed sets the speed (no-op for SimpleActor)
func (s *SimpleActor) SetSpeed(speed *math.Vector3) {
	// No-op for SimpleActor
}

// GetLastValidPosition returns the last valid position
func (s *SimpleActor) GetLastValidPosition() *math.Vector3 {
	return s.Location.Position
}

// SetLastValidPosition sets the last valid position (no-op for SimpleActor)
func (s *SimpleActor) SetLastValidPosition(lastValidPosition *math.Vector3) {
	// No-op for SimpleActor
}

// GetLastValidRotation returns the last valid rotation
func (s *SimpleActor) GetLastValidRotation() *math.Vector3 {
	return s.Location.Rotation
}

// SetLastValidRotation sets the last valid rotation (no-op for SimpleActor)
func (s *SimpleActor) SetLastValidRotation(lastValidRotation *math.Vector3) {
	// No-op for SimpleActor
}

// GetTeleportCachePos returns the teleport cache position
func (s *SimpleActor) GetTeleportCachePos() *math.Vector3 {
	return s.Location.Position
}

// SetTeleportCachePos sets the teleport cache position (no-op for SimpleActor)
func (s *SimpleActor) SetTeleportCachePos(teleportCachePos *math.Vector3) {
	// No-op for SimpleActor
}

// GetTeleportCacheRot returns the teleport cache rotation
func (s *SimpleActor) GetTeleportCacheRot() *math.Vector3 {
	return s.Location.Rotation
}

// SetTeleportCacheRot sets the teleport cache rotation (no-op for SimpleActor)
func (s *SimpleActor) SetTeleportCacheRot(teleportCacheRot *math.Vector3) {
	// No-op for SimpleActor
}

// GetMotionState returns the motion state
func (s *SimpleActor) GetMotionState() int32 {
	return 0 // MOTION_NONE
}

// SetMotionState sets the motion state (no-op for SimpleActor)
func (s *SimpleActor) SetMotionState(motionState int32) {
	// No-op for SimpleActor
}

// GetLastMoveSceneTimeMs returns the last move scene time
func (s *SimpleActor) GetLastMoveSceneTimeMs() int64 {
	return 0
}

// SetLastMoveSceneTimeMs sets the last move scene time (no-op for SimpleActor)
func (s *SimpleActor) SetLastMoveSceneTimeMs(lastMoveSceneTimeMs int64) {
	// No-op for SimpleActor
}

// GetLastPosValidTimeMs returns the last position valid time
func (s *SimpleActor) GetLastPosValidTimeMs() int64 {
	return 0
}

// SetLastPosValidTimeMs sets the last position valid time (no-op for SimpleActor)
func (s *SimpleActor) SetLastPosValidTimeMs(lastPosValidTimeMs int64) {
	// No-op for SimpleActor
}

// SetCachePosRot sets the cache position and rotation (no-op for SimpleActor)
func (s *SimpleActor) SetCachePosRot(pos, rot *math.Vector3) {
	// No-op for SimpleActor
}

// GetCachePosOrCurPos returns the cache position or current position
func (s *SimpleActor) GetCachePosOrCurPos() *math.Vector3 {
	return s.Location.Position
}

// GetCacheRotOrCurRot returns the cache rotation or current rotation
func (s *SimpleActor) GetCacheRotOrCurRot() *math.Vector3 {
	return s.Location.Rotation
}

// Drop implements drop functionality (no-op for SimpleActor)
func (s *SimpleActor) Drop(initiator *Entity) {
	// TODO: implement drop functionality
}
