package actors

import (
	"gameSrv/pkg/scene/actors/entity"
)

// SummonActor is a summon entity actor
type SummonActor struct {
	entity.LevelCreature
}

// NewSummonActor creates a new SummonActor
func NewSummonActor() *SummonActor {
	s := &SummonActor{}
	s.LevelCreature = *entity.NewLevelCreature()
	return s
}

// GetActorType returns actor type
func (s *SummonActor) GetActorType() int32 {
	return 19 // EActorType_Summon
}

// SetMaster sets the master creature
func (s *SummonActor) SetMaster(master *entity.Creature) {
	// TODO: implement
}

// UnlockMaster unlocks the master relationship
func (s *SummonActor) UnlockMaster() {
	// TODO: implement
}