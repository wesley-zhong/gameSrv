package actors

import "gameSrv/pkg/scene"

// CommonSwitchActor is a common switch actor
type CommonSwitchActor struct {
	SimpleActor
}

func (c *CommonSwitchActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewCommonSwitchActor creates a new CommonSwitchActor
func NewCommonSwitchActor() *CommonSwitchActor {
	c := &CommonSwitchActor{}
	c.SimpleActor = *NewSimpleActor()
	return c
}

// GetActorType returns actor type
func (c *CommonSwitchActor) GetActorType() int32 {
	return 18 // EActorType_CommonSwitch
}
