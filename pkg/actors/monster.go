package actors

import "gameSrv/pkg/scene"

// Monster is a monster class
type Monster struct {
	LevelCreature
}

func (m *Monster) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewMonster creates a new Monster
func NewMonster() *Monster {
	m := &Monster{}
	m.LevelCreature = *NewLevelCreature()
	return m
}

// GetActorType returns actor type
func (m *Monster) GetActorType() int32 {
	return 1 // EActorType_Monster
}

// GetSubType returns monster subtype
func (m *Monster) GetSubType() int32 {
	return 0
}

// Reset resets the monster
func (m *Monster) Reset() bool {
	// TODO: implement
	return true
}

// HandleDead handles death
func (m *Monster) HandleDead(killerActor *Entity) bool {
	// TODO: implement
	return true
}

// HandleInteract handles interaction
func (m *Monster) HandleInteract(interactMan *Entity, optionType int32, optionParams []int64) int32 {
	// TODO: implement
	return 0 // SUCCESS
}

// GetResID returns resource ID
func (m *Monster) GetResID() int64 {
	return m.ConfigID
}

// Drop handles drop logic
func (m *Monster) Drop(player interface{}) bool {
	// TODO: implement
	return false
}
