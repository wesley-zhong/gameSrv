package entity

import (
	"gameSrv/pkg/math"
)

// Region is a region entity in scene
type Region struct {
	Entity
}

func (r *Region) GetOnlyVisionUid() int64 {
	return 0
}

func (r *Region) GetPhasingId() int64 {
	return r.PhasingID
}

// IntersectType represents intersection type with region
type IntersectType int

const (
	IntersectTypeNone    IntersectType = iota
	IntersectTypeInward                // Enter region
	IntersectTypeOutward               // Leave region
	IntersectTypeCross                 // Cross region
)

// IsInRegion checks if a position is inside region
func (r *Region) IsInRegion(pos *math.Vector3) bool {
	// TODO: implement polygon intersection check
	return false
}

// AddEntity adds an entity to region
func (r *Region) AddEntity(entity *Entity, triggerEvt bool) {
	// TODO: implement
}

// DelEntity removes an entity from region
func (r *Region) DelEntity(entity *Entity, triggerEvt bool) {
	// TODO: implement
}

// DelAllEntity removes all entities from region
func (r *Region) DelAllEntity(triggerEvt bool) {
	// TODO: implement
}
