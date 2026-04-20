package actors

import "gameSrv/pkg/scene"

// VolumeBase is the base class for volume actors
type VolumeBase struct {
	SimpleActor
	IsTriggerOneTime bool
}

func (*VolumeBase) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewVolumeBase creates a new VolumeBase
func NewVolumeBase() *VolumeBase {
	v := &VolumeBase{
		IsTriggerOneTime: false,
	}
	v.SimpleActor = *NewSimpleActor()
	return v
}
