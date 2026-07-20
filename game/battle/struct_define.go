package battle

import (
	"gameSrv/cnfGen/cfg"
)

// Creature represents a creature in the battle system
type Creature interface {
	GetEntityId() int64
	GetConfigId() int64
	// TODO: add more methods as needed
}

// BubbleType represents the bubble type for damage display
type BubbleType int32

const (
	BubbleType_BaseHurt          BubbleType = 0 // 基础攻击
	BubbleType_SkillHurt         BubbleType = 1 // 技能攻击
	BubbleType_CriticalBaseHurt  BubbleType = 2 // 基础攻击暴击
	BubbleType_CriticalSkillHurt BubbleType = 3
	BubbleType_AddHp             BubbleType = 4
	BubbleType_AddSp             BubbleType = 5
	BubbleType_SelfHurt          BubbleType = 6
	BubbleType_Miss              BubbleType = 7
	BubbleType_Immunity          BubbleType = 8
	BubbleType_PickUpMoney       BubbleType = 9
	BubbleType_PickUpItem        BubbleType = 10
	BubbleType_Dodge             BubbleType = 11
	BubbleType_PetHurt           BubbleType = 12 // 宠物攻击
	BubbleType_Max               BubbleType = 13
)

// AttackType represents the type of attack
type AttackType int32

const (
	AttackType_SkillAction AttackType = 0
	AttackType_BuffAction  AttackType = 1
	AttackType_SkillSummon AttackType = 2
	AttackType_ForceKill   AttackType = 3
	AttackType_Unknown     AttackType = 4
)

// ActionType represents the type of action
type ActionType int32

const (
	ActionType_Property              ActionType = 0 // 属性变化
	ActionType_Position              ActionType = 1 // 位置变化
	ActionType_PositionImmediately   ActionType = 2 // 位置立即（闪烁类型）变化
	ActionType_PlayAnim              ActionType = 3 // 播放动画
	ActionType_StopAnim              ActionType = 4 // 停止动画
	ActionType_PlayEffect            ActionType = 5 // 播放特效动画
	ActionType_PlayHitEffect         ActionType = 6 // 播放命中特效动画
	ActionType_PlayBuffEffect        ActionType = 7 // 播放Buff特效动画
	ActionType_PlayAudio             ActionType = 8 // 播放音频
	ActionType_SummonBorn            ActionType = 9 // 召唤物出生
	ActionType_SummonDead            ActionType = 10 // 召唤物死亡
	ActionType_Teleport              ActionType = 11 // 击退 击飞
	ActionType_AddBuff_CastActor     ActionType = 12 // buff出生
	ActionType_RemoveBuff_CastActor  ActionType = 13 // buff移除
	ActionType_ApplyDamage           ActionType = 14
	ActionType_ReceiveDamage         ActionType = 15
	ActionType_UnbeatState           ActionType = 16
)

// DamageResult represents the result of damage calculation
type DamageResult struct {
	Attacker               Creature                 // 攻击者
	Target                 Creature                 // 目标
	TotalDamage            int                      // 伤害值
	BubbleType             BubbleType               // 伤害冒泡类型
	IsSuper                bool                     // 是否暴击
	IsAbnoramlExtra        bool                     // 是否有印记增伤
	IsSermonHighDefence    bool                     // 是否高抗性
	BDodge                 bool                     // 是否躲闪
	EabInnerBlock          bool                     // 是否内角度格挡 + 判断state
	AttackID               int                      // 技能id
	DamageRadio            int                      // 最终伤害倍率
	Attack_Abnormal_Radio  int                      // 异常属性倍率
	IsMaxAbnormal          bool                     //
	HasAbnormalClosing     bool                     //
	BIsPenetration         bool     // 已经穿透护甲
	AttackType             AttackType // 技能 召唤物 buff
	DataType               int       // cfg.DataType
	BuffReason             int       // cfg.BuffTriggerEventEnum
	AbnormalType           int       // cfg.AbnormalElement
	DamageType             int       // cfg.DamageType
	ExSkillEnergyBackstageLists map[int64]int       // 后台充能
	RandomIndex            int                      //
	AbnormalProp           int                      //
	AbnormalStrength       int                      //
	BIgnoreDamageEvent     bool                     // 忽略伤害性事件

	// Property change arrays
	PropWantChangedArray       []int
	PropChangedArray           []int
	AttackPropChangedArray     []int
	TargetPropChangedArray     []int
	TargetPropWantChangedArray []int
	TargetAttackPropChangedArray []int
}

// NewDamageResult creates a new DamageResult
func NewDamageResult() *DamageResult {
	return &DamageResult{
		PropWantChangedArray:         make([]int, cfg.HeroProp_Max),
		PropChangedArray:             make([]int, cfg.HeroProp_Max),
		AttackPropChangedArray:       make([]int, cfg.ESkillProp_MAX),
		TargetPropChangedArray:       make([]int, cfg.HeroProp_Max),
		TargetPropWantChangedArray:   make([]int, cfg.HeroProp_Max),
		TargetAttackPropChangedArray: make([]int, cfg.ESkillProp_MAX),
		ExSkillEnergyBackstageLists: make(map[int64]int),
	}
}

// RecordPropertyWantChange records property change that will happen
func (dr *DamageResult) RecordPropertyWantChange(propType int, valueDelta int, isTarget bool) {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetPropWantChangedArray) {
			dr.TargetPropWantChangedArray[propType] = valueDelta
		}
	} else {
		if propType >= 0 && propType < len(dr.PropWantChangedArray) {
			dr.PropWantChangedArray[propType] = valueDelta
		}
	}
}

// GetPropertyWantChangeValue gets the property change value
func (dr *DamageResult) GetPropertyWantChangeValue(propType int, isTarget bool) int {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetPropWantChangedArray) {
			return dr.TargetPropWantChangedArray[propType]
		}
	} else {
		if propType >= 0 && propType < len(dr.PropWantChangedArray) {
			return dr.PropWantChangedArray[propType]
		}
	}
	return 0
}

// RecordPropertyChange records actual property change
func (dr *DamageResult) RecordPropertyChange(propType int, valueDelta int, isTarget bool) {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetPropChangedArray) {
			dr.TargetPropChangedArray[propType] = valueDelta
		}
	} else {
		if propType >= 0 && propType < len(dr.PropChangedArray) {
			dr.PropChangedArray[propType] = valueDelta
		}
	}
}

// GetPropertyChangeValue gets the actual property change value
func (dr *DamageResult) GetPropertyChangeValue(propType int, isTarget bool) int {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetPropChangedArray) {
			return dr.TargetPropChangedArray[propType]
		}
	} else {
		if propType >= 0 && propType < len(dr.PropChangedArray) {
			return dr.PropChangedArray[propType]
		}
	}
	return 0
}

// RecordAttackPropertyChange records attack property change
func (dr *DamageResult) RecordAttackPropertyChange(propType int, valueDelta int, isTarget bool) {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetAttackPropChangedArray) {
			dr.TargetAttackPropChangedArray[propType] = valueDelta
		}
	} else {
		if propType >= 0 && propType < len(dr.AttackPropChangedArray) {
			dr.AttackPropChangedArray[propType] = valueDelta
		}
	}
}

// GetAttackPropertyChangeValue gets the attack property change value
func (dr *DamageResult) GetAttackPropertyChangeValue(propType int, isTarget bool) int {
	if isTarget {
		if propType >= 0 && propType < len(dr.TargetAttackPropChangedArray) {
			return dr.TargetAttackPropChangedArray[propType]
		}
	} else {
		if propType >= 0 && propType < len(dr.AttackPropChangedArray) {
			return dr.AttackPropChangedArray[propType]
		}
	}
	return 0
}

// PropertyParallelForEach iterates over property changes
func (dr *DamageResult) PropertyParallelForEach(fn func(prop int, value int), isTarget bool) {
	var arr []int
	if isTarget {
		arr = dr.TargetPropChangedArray
	} else {
		arr = dr.PropChangedArray
	}
	for i, value := range arr {
		if value != 0 {
			fn(i, value)
		}
	}
}

// AttackPropertyParallelForEach iterates over attack property changes
func (dr *DamageResult) AttackPropertyParallelForEach(fn func(prop int, value int), isTarget bool) {
	var arr []int
	if isTarget {
		arr = dr.TargetAttackPropChangedArray
	} else {
		arr = dr.AttackPropChangedArray
	}
	for i, value := range arr {
		if value != 0 {
			fn(i, value)
		}
	}
}