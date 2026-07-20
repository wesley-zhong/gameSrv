package battle

import (
	"fmt"
	"testing"
)

// TestBattleLogBasic tests basic battle log functionality
func TestBattleLogBasic(t *testing.T) {
	// Test NewBattleLog
	log := NewBattleLog()
	if log == nil {
		t.Fatal("NewBattleLog returned nil")
	}

	// Test basic fields
	log.Index = 1
	log.Type = BattleLogType_GA
	log.GAID = 100

	if log.Index != 1 {
		t.Errorf("Expected Index 1, got %d", log.Index)
	}
	if log.Type != BattleLogType_GA {
		t.Errorf("Expected Type GA, got %d", log.Type)
	}

	// Test Recycle
	log.Recycle()

	// Test pool reuse
	log2 := NewBattleLog()
	if log2 == nil {
		t.Fatal("NewBattleLog returned nil after recycle")
	}
}

// TestBattleLogCacher tests battle log cacher functionality
func TestBattleLogCacher(t *testing.T) {
	player := &PlayerStub{UID: 12345}
	cacher := NewFYBattleLogCacher(player)

	if cacher.LastIndex != 1 {
		t.Errorf("Expected LastIndex 1, got %d", cacher.LastIndex)
	}

	// Test lock/unlock
	cacher.LockIndex()
	if !cacher.IndexLocked {
		t.Error("Expected IndexLocked to be true")
	}

	cacher.UnLockIndex()
	if cacher.IndexLocked {
		t.Error("Expected IndexLocked to be false")
	}

	// Test attack count
	cacher.SetAttackCount(100)
	if cacher.GetAttackCount(100) != 1 {
		t.Errorf("Expected attack count 1, got %d", cacher.GetAttackCount(100))
	}

	cacher.SetAttackCount(100)
	if cacher.GetAttackCount(100) != 2 {
		t.Errorf("Expected attack count 2, got %d", cacher.GetAttackCount(100))
	}

	cacher.ClearAttackCount()
	if cacher.GetAttackCount(100) != 0 {
		t.Errorf("Expected attack count 0 after clear, got %d", cacher.GetAttackCount(100))
	}
}

// TestBuffTagEnumString tests BuffTagEnum string conversion
func TestBuffTagEnumString(t *testing.T) {
	tests := []struct {
		tag      BuffTagEnum
		expected string
	}{
		{BuffTagEnum_TAG_NONE, "NONE"},
		{BuffTagEnum_TAG_MAX, "MAX"},
		{BuffTagEnum_TAG_STUN, "STUN"},
		{BuffTagEnum_TAG_SILENCE, "SILENCE"},
		{BuffTagEnum_TAG_FREEZE, "FREEZE"},
		{BuffTagEnum_TAG_POISON, "POISON"},
		{BuffTagEnum(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if got := tt.tag.String(); got != tt.expected {
			t.Errorf("BuffTagEnum(%d).String() = %s, want %s", tt.tag, got, tt.expected)
		}
	}
}

// TestLogEntity tests FYLogEntity functionality
func TestLogEntity(t *testing.T) {
	entity := FYLogEntity{
		EntityID: 123,
		ConfigID: 456,
		BuffTags: make([]int, 0, 4),
		PropData: make(map[int]int),
	}

	entity.BuffTags = append(entity.BuffTags, int(BuffTagEnum_TAG_STUN))
	entity.BuffTags = append(entity.BuffTags, int(BuffTagEnum_TAG_POISON))

	infoStr := entity.GetInfoString()
	if infoStr == "" {
		t.Error("Expected non-empty info string")
	}

	heroNameStr := entity.GetHeroNameString()
	if heroNameStr == "" {
		t.Error("Expected non-empty hero name string")
	}

	// Test reset
	entity.Reset()
	if entity.EntityID != 0 {
		t.Errorf("Expected EntityID 0 after reset, got %d", entity.EntityID)
	}
	if len(entity.BuffTags) != 0 {
		t.Errorf("Expected empty BuffTags after reset, got %d", len(entity.BuffTags))
	}
}

// TestGetHeroPropName tests hero property name conversion
func TestGetHeroPropName(t *testing.T) {
	tests := []struct {
		prop     int
		expected string
	}{
		{1, "MaxHealth"},
		{2, "Health"},
		{3, "Attack"},
		{999, "HeroProp_999"},
	}

	for _, tt := range tests {
		if got := getHeroPropName(tt.prop); got != tt.expected {
			t.Errorf("getHeroPropName(%d) = %s, want %s", tt.prop, got, tt.expected)
		}
	}
}

// TestGetSkillPropName tests skill property name conversion
func TestGetSkillPropName(t *testing.T) {
	tests := []struct {
		prop     int
		expected string
	}{
		{2, "Damage_Ratio"},
		{4, "Damage_Type"},
		{49, "Damage_Base"},
		{999, "ESkillProp_999"},
	}

	for _, tt := range tests {
		if got := getSkillPropName(tt.prop); got != tt.expected {
			t.Errorf("getSkillPropName(%d) = %s, want %s", tt.prop, got, tt.expected)
		}
	}
}

// Example usage
func Example() {
	log := NewBattleLog()
	log.Index = 1
	log.Type = BattleLogType_MONTAGE
	log.MontageID = 100
	log.IsStart = true

	fmt.Printf("Battle Log: %s\n", log.String())
	log.Recycle()
}