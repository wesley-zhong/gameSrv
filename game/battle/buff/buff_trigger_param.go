package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// BuffTriggerEventParamBase is the base class for buff trigger event parameters
type BuffTriggerEventParamBase struct {
	Target            *actors.Creature
	InteractiveTarget *actors.Creature
	TriggerFlag       int
	ExLayer           int // Extra layer for buff effect stacking
}

// PropsChangeBuffTriggerEventParam handles property change trigger parameters
type PropsChangeBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	Prop        int // cfg.HeroProp constant
	OrigValue   int
	CurValue    int
	CompareRate bool
	StartCheck  bool
}

// ShipFireEventParam handles ship fire trigger parameters
type ShipFireEventParam struct {
	BuffTriggerEventParamBase
	ShipShellTypeEnum int // cfg.ShipShellTypeEnum constant (placeholder)
}

// ShipFireEffectEventParam handles ship fire effect trigger parameters
type ShipFireEffectEventParam struct {
	BuffTriggerEventParamBase
	ShipShellEffectId int64
}

// StateChangeBuffTriggerEventParam handles state change trigger parameters
type StateChangeBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	State int // cfg.ECreatureActionState constant (placeholder)
	IsOut bool
}

// EnterTeamBuffTriggerEventParam handles enter team trigger parameters
type EnterTeamBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AvatarId int64
}

// LeaveTeamBuffTriggerEventParam handles leave team trigger parameters
type LeaveTeamBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AvatarId int64
}

// BuffMountedBuffTriggerEventParam handles buff mounted trigger parameters
type BuffMountedBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	FromBuffId int
}

// HitedBuffTriggerEventParam handles hit trigger parameters
type HitedBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
	DamageType   int
}

// BeHitedBuffTriggerEventParam handles being hit trigger parameters
type BeHitedBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
}

// DodamageBuffTriggerEventParam handles damage dealing trigger parameters
type DodamageBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
	DamageType   int
}

// DodamageBySkillTypeBuffTriggerEventParam handles skill type damage trigger parameters
type DodamageBySkillTypeBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
	DamageType   int
}

// DoCriticalBySkillTypeDamageBuffTriggerEventParam handles critical damage trigger parameters
type DoCriticalBySkillTypeDamageBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int64
	DamageType   int
}

// BeDamagedBuffTriggerEventParam handles being damaged trigger parameters
type BeDamagedBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
}

// MtgIdBuffTriggerEventParam handles montage ID trigger parameters
type MtgIdBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	MtgId int
}

// AttackDataIDBuffTriggerEventParam handles attack data ID trigger parameters
type AttackDataIDBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	AttackDataId int
}

// RebornBuffTriggerEventParam handles reborn trigger parameters
type RebornBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
}

// KilledMonsterEventParam handles killed monster trigger parameters
type KilledMonsterEventParam struct {
	BuffTriggerEventParamBase
	MstType int // cfg.MonsterType constant (placeholder)
}

// MontageEndBuffTriggerEventParam handles montage end trigger parameters
type MontageEndBuffTriggerEventParam struct {
	BuffTriggerEventParamBase
	MtgId int
}

// ExecutionEndParam handles execution end trigger parameters
type ExecutionEndParam struct {
	BuffTriggerEventParamBase
	ExecutionType int
}

// AbnormalLayerClosing handles abnormal layer closing parameters
type AbnormalLayerClosing struct {
	BuffTriggerEventParamBase
	HeroProp        int // cfg.HeroProp constant
	CurrentAttackID int
}

// ActionSubAbnormalEventParam handles action sub abnormal parameters
type ActionSubAbnormalEventParam struct {
	BuffTriggerEventParamBase
	AbnormalLayerType int // cfg.HeroProp constant (placeholder)
}

// ClipCapacityChange handles clip capacity change parameters
type ClipCapacityChange struct {
	BuffTriggerEventParamBase
	Percent int
}

// ClipReloadParam handles clip reload parameters
type ClipReloadParam struct {
	BuffTriggerEventParamBase
}

// ShipSpeedUpParam handles ship speed up parameters
type ShipSpeedUpParam struct {
	BuffTriggerEventParamBase
}

// BuffTriggerAddRet represents the return value for adding buff trigger
type BuffTriggerAddRet struct {
	ImmediateExeTriggerEvents []*cfg.TriggerEvent
	TickTriggerEvent          *cfg.TriggerEvent
}

// FYBuffEffectAddedData represents buff effect added data
type FYBuffEffectAddedData struct {
	BuffEffectAddedDatas map[int]int
	AddNum               int
}
