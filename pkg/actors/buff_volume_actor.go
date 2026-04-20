package actors

import "gameSrv/pkg/scene"

// BuffVolumeActor is an actor that triggers buff volumes
type BuffVolumeActor struct {
	VolumeBase
}

func (b *BuffVolumeActor) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewBuffVolumeActor creates a new BuffVolumeActor
func NewBuffVolumeActor() *BuffVolumeActor {
	b := &BuffVolumeActor{}
	b.VolumeBase = *NewVolumeBase()
	return b
}

// GetActorType returns actor type
func (b *BuffVolumeActor) GetActorType() int32 {
	return 11 // EActorType_BuffVolume
}
