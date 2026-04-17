package actors

import (
	"gameSrv/pkg/scene/actors/entity"
)

// VolumeBase is the base class for volume actors
type VolumeBase struct {
	entity.SimpleActor
	IsTriggerOneTime bool
}

// NewVolumeBase creates a new VolumeBase
func NewVolumeBase() *VolumeBase {
	v := &VolumeBase{
		IsTriggerOneTime: false,
	}
	v.SimpleActor = *entity.NewSimpleActor()
	return v
}