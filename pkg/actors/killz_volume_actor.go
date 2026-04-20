package actors

import "gameSrv/pkg/scene"

// KillZVolumeActor is a kill zone volume actor
type KillZVolumeActor struct {
	VolumeBase
}

func (k *KillZVolumeActor) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewKillZVolumeActor creates a new KillZVolumeActor
func NewKillZVolumeActor() *KillZVolumeActor {
	k := &KillZVolumeActor{}
	k.VolumeBase = *NewVolumeBase()
	return k
}

// GetActorType returns actor type
func (k *KillZVolumeActor) GetActorType() int32 {
	return 12 // EActorType_KillZVolume
}
