package entity

// LevelCreature is a creature with TemplateData\LevelData\Difficulty combat properties
type LevelCreature struct {
	Creature
	ExtData interface{}
}

// NewLevelCreature creates a new LevelCreature
func NewLevelCreature() *LevelCreature {
	l := &LevelCreature{}
	l.Creature = *NewCreature()
	return l
}

func (l *LevelCreature) GetOnlyVisionUid() int64 {
	return 0
}

func (l *LevelCreature) GetPhasingId() int64 {
	return l.PhasingID
}

// InitLevelConfigData initializes level configuration data
func (l *LevelCreature) InitLevelConfigData(configId int64) {
	// TODO: implement
}

// SetupDifficultyData sets up difficulty data
func (l *LevelCreature) SetupDifficultyData(difficultyIndex int32) {
	// TODO: implement
}

// GetExtData returns extended data
func (l *LevelCreature) GetExtData() interface{} {
	return l.ExtData
}

// SetExtData sets extended data
func (l *LevelCreature) SetExtData(object interface{}) {
	l.ExtData = object
}