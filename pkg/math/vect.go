package math

import (
	"fmt"
)

type Vector2 struct {
	X int32
	Y int32
}

// String 返回字符串表示
func (v *Vector2) String() string {
	return fmt.Sprintf("Vector2{%d,%d}", v.X, v.Y)
}

type Vector3 struct {
	X int32
	Y int32
	Z int32
}

// String 返回字符串表示
func (v *Vector3) String() string {
	return fmt.Sprintf("Vector3{%d,%d,%d}", v.X, v.Y, v.Z)
}

type Vector4 struct {
	X int32
	Y int32
	Z int32
	R int32
}

// String 返回字符串表示
func (v *Vector4) String() string {
	return fmt.Sprintf("Vector4{%d,%d,%d,%d}", v.X, v.Y, v.Z, v.R)
}

func NewVector2(x, y int32) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

// NewVector3 creates a new Vector3
func NewVector3(x, y, z int32) *Vector3 {
	return &Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

// ZeroVector3 returns a zero vector
func ZeroVector3() *Vector3 {
	return &Vector3{X: 0, Y: 0, Z: 0}
}

// LocationInfo is the location information
type LocationInfo struct {
	Position *Vector3
	Rotation *Vector3
	BornPos  *Vector3
	BornRot  *Vector3
}

// GetPosition 获取位置
func (l *LocationInfo) GetPosition() *Vector3 {
	return l.Position
}

// GetRotation 获取旋转
func (l *LocationInfo) GetRotation() *Vector3 {
	return l.Rotation
}

// MoveInfo is the movement information
type MoveInfo struct {
	LastValidPosition   *Vector3
	LastValidRotation   *Vector3
	TeleportCachePos    *Vector3
	TeleportCacheRot    *Vector3
	Speed               *Vector3
	MotionState         int32
	LastMoveSceneTimeMs int64
	LastPosValidTimeMs  int64
}
