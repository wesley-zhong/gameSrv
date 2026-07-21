package process

import (
	"gameSrv/game/battle/buff"
	"gameSrv/pkg/actors"
)

// Creature placeholder interface
type Creature = actors.Creature

// BuffData is an alias for buff data in buff package
type BuffData = buff.BuffData

// BattleLogData placeholder for protobuf message
type BattleLogData struct {
	EventState     int32
	EventParam     int32
	UID            int64
	RandomIndex    int32
	ExtraTargetIds []int64
	EventReason   int32
}

// IBattleLogProcessFunction defines battle log process function interface
type IBattleLogProcessFunction func(caster, target, holder *Creature, battleLogPush *BattleLogData)

// IBuffAddProcessFunction defines buff add process function interface
type IBuffAddProcessFunction func(caster, target *Creature, buffAddPush *BuffData) bool