package actors

import (
	"gameSrv/pkg/scene/actors/entity"
)

// RestPoint is a rest point
type RestPoint struct {
	entity.SimpleActor
}

// NewRestPoint creates a new RestPoint
func NewRestPoint() *RestPoint {
	r := &RestPoint{}
	r.SimpleActor = *entity.NewSimpleActor()
	return r
}

// GetActorType returns actor type
func (r *RestPoint) GetActorType() int32 {
	return 17 // EActorType_RestPoint
}