package actors

import "gameSrv/pkg/scene/actors/entity"

// CommonSwitchActor is a common switch actor
type CommonSwitchActor struct {
	entity.SimpleActor
}

// NewCommonSwitchActor creates a new CommonSwitchActor
func NewCommonSwitchActor() *CommonSwitchActor {
	c := &CommonSwitchActor{}
	c.SimpleActor = *entity.NewSimpleActor()
	return c
}

// GetActorType returns actor type
func (c *CommonSwitchActor) GetActorType() int32 {
	return 18 // EActorType_CommonSwitch
}