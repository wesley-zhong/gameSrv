package actors

// KillZVolumeActor is a kill zone volume actor
type KillZVolumeActor struct {
	VolumeBase
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