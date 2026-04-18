package actors

import "gameSrv/pkg/scene"

// SummonActor is a summon entity actor
type SummonActor struct {
	LevelCreature
}

func (s *SummonActor) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewSummonActor creates a new SummonActor
func NewSummonActor() *SummonActor {
	s := &SummonActor{}
	s.LevelCreature = *NewLevelCreature()
	return s
}

// GetActorType returns actor type
func (s *SummonActor) GetActorType() int32 {
	return 19 // EActorType_Summon
}

// SetMaster sets the master creature
func (s *SummonActor) SetMaster(master *Creature) {
	// TODO: implement
}

// UnlockMaster unlocks the master relationship
func (s *SummonActor) UnlockMaster() {
	// TODO: implement
}
