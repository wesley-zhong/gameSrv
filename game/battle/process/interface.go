package process

import "gameSrv/pkg/actors"

// Creature placeholder interface
type Creature = actors.Creature

// BattleLogData placeholder for protobuf message
type BattleLogData struct {
	EventState     int32
	EventParam     int32
	UID            int64
	RandomIndex    int32
	ExtraTargetIds []int64
	EventReason   int32
}

// BuffData placeholder for protobuf
type BuffData struct {
	// TODO: define BuffData fields
}

// IBattleLogProcessFunction defines battle log process function interface
type IBattleLogProcessFunction func(caster, target, holder *Creature, battleLogPush *BattleLogData)

// IBuffAddProcessFunction defines buff add process function interface
type IBuffAddProcessFunction func(caster, target *Creature, buffAddPush *BuffData) bool