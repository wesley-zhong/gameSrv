package handlers

import (
	"gameSrv/pkg/math"
)

// RebornHandler is the reborn handler interface
type RebornHandler interface {
	// Reborn performs rebirth
	Reborn(avatar interface{}, transform *math.Vector3) bool
	// GetRebornReason gets the reborn reason
	GetRebornReason() int32
}