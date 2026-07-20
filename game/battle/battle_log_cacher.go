package battle

import (
	"fmt"
	"time"

	"gameSrv/pkg/actors"
)

// LogFileConfig handles battle log file operations configuration
type LogFileConfig struct {
	SaveBattleLog bool
}

// LogFile handles battle log file operations
var LogFile = &LogFileConfig{
	SaveBattleLog: false, // Set to true to enable battle log saving
}

// complexLog writes complex log
func (lf *LogFileConfig) complexLog(msg string) {
	if !lf.SaveBattleLog {
		return
	}
	// TODO: Implement actual log file writing
	fmt.Printf("[COMPLEX] %s\n", msg)
}

// simpleLog writes simple log
func (lf *LogFileConfig) simpleLog(msg string) {
	if !lf.SaveBattleLog {
		return
	}
	// TODO: Implement actual log file writing
	fmt.Printf("[SIMPLE] %s\n", msg)
}

// Player represents a game player (placeholder for actual player type)
type Player interface {
	GetUid() int64
}

// PlayerStub is a stub implementation of Player interface
type PlayerStub struct {
	UID int64
}

func (p *PlayerStub) GetUid() int64 {
	return p.UID
}

// FYBattleLogCacher manages battle log caching and recording
type FYBattleLogCacher struct {
	LastIndex        int
	Player           Player
	IndexLocked      bool
	AttackCountMap   map[int]int
	CachedBattleLogs map[int64]*FYBattleLog
}

// NewFYBattleLogCacher creates a new FYBattleLogCacher
func NewFYBattleLogCacher(player Player) *FYBattleLogCacher {
	return &FYBattleLogCacher{
		LastIndex:        1,
		Player:           player,
		IndexLocked:      false,
		AttackCountMap:   make(map[int]int),
		CachedBattleLogs: make(map[int64]*FYBattleLog),
	}
}

// ResetBattleLogIndex resets the battle log index
func (c *FYBattleLogCacher) ResetBattleLogIndex() {
	c.LastIndex = 1
}

// LockIndex locks the index
func (c *FYBattleLogCacher) LockIndex() {
	c.IndexLocked = true
}

// UnLockIndex unlocks the index
func (c *FYBattleLogCacher) UnLockIndex() {
	c.IndexLocked = false
}

// SetAttackCount sets the attack count for an attack data ID
func (c *FYBattleLogCacher) SetAttackCount(attackDataID int) {
	if count, exists := c.AttackCountMap[attackDataID]; exists {
		c.AttackCountMap[attackDataID] = count + 1
	} else {
		c.AttackCountMap[attackDataID] = 1
	}
}

// GetAttackCount gets the attack count for an attack data ID
func (c *FYBattleLogCacher) GetAttackCount(attackDataID int) int {
	if count, exists := c.AttackCountMap[attackDataID]; exists {
		return count
	}
	return 0
}

// ClearAttackCount clears all attack counts
func (c *FYBattleLogCacher) ClearAttackCount() {
	c.AttackCountMap = make(map[int]int)
}

// addBattleLog adds a new battle log
func (c *FYBattleLogCacher) addBattleLog(entity *actors.Creature) *FYBattleLog {
	battleLog := NewBattleLog()

	if entity != nil {
		battleLog.SourceEntity.EntityID = entity.GetEntityId()
		battleLog.SourceEntity.ConfigID = int(entity.GetConfigId())
		// TODO: Set element type, sermon type, buff tags when available
	}

	if c.IndexLocked {
		battleLog.Index = c.LastIndex
	} else {
		battleLog.Index = c.LastIndex
		c.LastIndex++
		c.CachedBattleLogs[int64(battleLog.Index)] = battleLog
	}

	fmt.Printf("create server log=%v\n", battleLog)
	return battleLog
}

// fixBattleLog finalizes and writes the battle log
func (c *FYBattleLogCacher) fixBattleLog(battleLog *FYBattleLog) {
	if !LogFile.SaveBattleLog {
		return
	}

	simple := battleLog.PrintSimpleLog()
	complex := battleLog.PrintComplexLog()

	formattedDate := time.Now().Format("2006-01-02 15:04:05")

	LogFile.complexLog(formattedDate + " " + fmt.Sprint(c.Player.GetUid()) + complex)
	LogFile.complexLog("-------------------------------------------------")

	if simple == "" {
		return
	}
	LogFile.simpleLog(fmt.Sprint(c.Player.GetUid()) + simple)
}

// IsContains checks if a seq no exists in cached logs
func (c *FYBattleLogCacher) IsContains(seqNo int64) bool {
	_, exists := c.CachedBattleLogs[seqNo]
	return exists
}

// GetBattleLog gets a battle log by seq no
func (c *FYBattleLogCacher) GetBattleLog(seqNo int64) *FYBattleLog {
	return c.CachedBattleLogs[seqNo]
}

// Count returns the number of cached battle logs
func (c *FYBattleLogCacher) Count() int {
	return len(c.CachedBattleLogs)
}

// Reset clears all cached battle logs
func (c *FYBattleLogCacher) Reset() {
	if len(c.CachedBattleLogs) == 0 {
		return
	}

	for _, log := range c.CachedBattleLogs {
		if log != nil {
			log.Recycle()
		}
	}
	c.CachedBattleLogs = make(map[int64]*FYBattleLog)
}

// RecordGA records a GA (Game Ability) event
func (c *FYBattleLogCacher) RecordGA(entity *actors.Creature, isStart bool, gaid int) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_GA
	battleLog.GAID = gaid
	battleLog.IsStart = isStart
	// TODO: battleLog.State = entity.ActorBattleModule.CharacterStateFlag
	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordCombo records a combo event
func (c *FYBattleLogCacher) RecordCombo(entity *actors.Creature, isStart bool, comboID int) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_COMBO
	battleLog.ComboID = comboID
	battleLog.IsStart = isStart
	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordMontage records a montage event
func (c *FYBattleLogCacher) RecordMontage(entity *actors.Creature, isStart bool, montageID int) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_MONTAGE
	battleLog.MontageID = montageID
	battleLog.IsStart = isStart
	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordBuff records a buff change event
func (c *FYBattleLogCacher) RecordBuff(entity *actors.Creature, isStart bool, buffID int) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_BUFF_CHANGE
	battleLog.IsStart = isStart
	battleLog.BuffID = buffID
	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordBattleAction records a battle action event
func (c *FYBattleLogCacher) RecordBattleAction(entity *actors.Creature, logType BattleLogType, dr *DamageResult) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = logType
	battleLog.DR = dr

	if dr != nil {
		battleLog.BuffID = dr.AttackID

		// Set target entity info
		if dr.Target != nil {
			// Type assertion for Creature interface
			if target, ok := dr.Target.(interface{ GetEntityId() int64 }); ok {
				battleLog.TargetEntity.EntityID = target.GetEntityId()
			}
			if target, ok := dr.Target.(interface{ GetConfigId() int64 }); ok {
				battleLog.TargetEntity.ConfigID = int(target.GetConfigId())
			}
			// TODO: battleLog.TargetHash = dr.Target.CalculateBattleHash()
		}
	}

	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordBuffValueChange records a buff value change event
func (c *FYBattleLogCacher) RecordBuffValueChange(entity *actors.Creature, isStart bool, buffID int, dr *DamageResult) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_BUFF_VALUE_CHANGE
	battleLog.IsStart = isStart
	battleLog.BuffID = buffID
	battleLog.DR = dr

	if dr != nil && dr.Target != nil {
		// Type assertion for Creature interface
		if target, ok := dr.Target.(interface{ GetEntityId() int64 }); ok {
			battleLog.TargetEntity.EntityID = target.GetEntityId()
		}
		if target, ok := dr.Target.(interface{ GetConfigId() int64 }); ok {
			battleLog.TargetEntity.ConfigID = int(target.GetConfigId())
		}
	}

	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}

// RecordCostEnergy records a cost energy event
func (c *FYBattleLogCacher) RecordCostEnergy(entity *actors.Creature, dr *DamageResult) {
	battleLog := c.addBattleLog(entity)
	battleLog.Type = BattleLogType_COST_ENERGY
	battleLog.DR = dr

	if dr != nil && dr.Target != nil {
		// Type assertion for Creature interface
		if target, ok := dr.Target.(interface{ GetEntityId() int64 }); ok {
			battleLog.TargetEntity.EntityID = target.GetEntityId()
		}
		if target, ok := dr.Target.(interface{ GetConfigId() int64 }); ok {
			battleLog.TargetEntity.ConfigID = int(target.GetConfigId())
		}
	}

	// TODO: battleLog.CasterHash = entity.CalculateBattleHash()
	c.fixBattleLog(battleLog)
}
