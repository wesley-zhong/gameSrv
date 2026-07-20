package battle

import (
	"fmt"
	"strings"
	"sync"

	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/actors"
)

// Global configuration tables (to be initialized by the game server)
var (
	ConfigTables *cfg.Tables
)

// BattleLogType represents the type of battle log
type BattleLogType int32

const (
	BattleLogType_NONE              BattleLogType = 0
	BattleLogType_GA                BattleLogType = 1
	BattleLogType_COMBO             BattleLogType = 2
	BattleLogType_MONTAGE           BattleLogType = 3
	BattleLogType_BUFF_CHANGE       BattleLogType = 4
	BattleLogType_BUFF_VALUE_CHANGE BattleLogType = 5
	BattleLogType_COST_ENERGY       BattleLogType = 6
	BattleLogType_DO_ATTACK_DATA    BattleLogType = 7
)

// ElementType represents the element type of a creature
type ElementType int32

const (
	ElementType_None     ElementType = 0
	ElementType_Fire     ElementType = 1
	ElementType_Water    ElementType = 2
	ElementType_Wind     ElementType = 3
	ElementType_Electric ElementType = 4
	ElementType_Grass    ElementType = 5
	ElementType_Ice      ElementType = 6
	ElementType_Rock     ElementType = 7
)

// SermonType represents the sermon type
type SermonType int32

const (
	SermonType_None  SermonType = 0
	SermonType_Fire  SermonType = 1
	SermonType_Water SermonType = 2
)

// BuffTagEnum represents buff tags
type BuffTagEnum int32

const (
	BuffTagEnum_TAG_NONE    BuffTagEnum = 0
	BuffTagEnum_TAG_MAX     BuffTagEnum = 1
	BuffTagEnum_TAG_STUN    BuffTagEnum = 2
	BuffTagEnum_TAG_SILENCE BuffTagEnum = 3
	BuffTagEnum_TAG_FREEZE  BuffTagEnum = 4
	BuffTagEnum_TAG_POISON  BuffTagEnum = 5
)

// FYLogEntity represents entity information in battle log
type FYLogEntity struct {
	EntityID    int64
	ConfigID    int
	ElementType ElementType
	SermonType  SermonType
	BuffTags    []int
	PropData    map[int]int
	DataHash    int64
}

// Reset resets the log entity to default state
func (e *FYLogEntity) Reset() {
	e.EntityID = 0
	e.ConfigID = 0
	e.ElementType = ElementType_None
	e.SermonType = SermonType_None
	e.BuffTags = e.BuffTags[:0]
	e.PropData = make(map[int]int)
	e.DataHash = 0
}

// GetInfoString returns formatted info string
func (e *FYLogEntity) GetInfoString() string {
	if e.EntityID == 0 {
		return ""
	}

	tagStr := ""
	for i, tag := range e.BuffTags {
		if i > 0 {
			tagStr += ","
		}
		tagStr += BuffTagEnum(tag).String()
	}

	return fmt.Sprintf("EntityID:<Yellow_Item>%d</>ConfigID:<Yellow_Item>%d</> <Red_Item>%d</>\n bufftag--%s \n DataHash:%d\n",
		e.EntityID, e.ConfigID, e.ConfigID, tagStr, e.DataHash)
}

// GetHeroNameString returns hero name string
func (e *FYLogEntity) GetHeroNameString() string {
	return fmt.Sprintf("<Yellow_Item>%d</>", e.ConfigID)
}

// String returns string representation of BuffTagEnum
func (b BuffTagEnum) String() string {
	switch b {
	case BuffTagEnum_TAG_NONE:
		return "NONE"
	case BuffTagEnum_TAG_MAX:
		return "MAX"
	case BuffTagEnum_TAG_STUN:
		return "STUN"
	case BuffTagEnum_TAG_SILENCE:
		return "SILENCE"
	case BuffTagEnum_TAG_FREEZE:
		return "FREEZE"
	case BuffTagEnum_TAG_POISON:
		return "POISON"
	default:
		return "UNKNOWN"
	}
}

// FYBattleLog represents a battle log entry
type FYBattleLog struct {
	Index        int
	IsStart      bool
	GAID         int
	ComboID      int
	MontageID    int
	BuffID       int
	State        int
	Type         BattleLogType
	DR           *DamageResult
	SourceEntity FYLogEntity
	TargetEntity FYLogEntity
	CasterHash   int64
	TargetHash   int64
}

// battleLogPool is the sync.Pool for FYBattleLog
var battleLogPool = sync.Pool{
	New: func() interface{} {
		return &FYBattleLog{
			SourceEntity: FYLogEntity{
				BuffTags: make([]int, 0, 8),
				PropData: make(map[int]int),
			},
			TargetEntity: FYLogEntity{
				BuffTags: make([]int, 0, 8),
				PropData: make(map[int]int),
			},
		}
	},
}

// NewBattleLog creates a new FYBattleLog from pool
func NewBattleLog() *FYBattleLog {
	log := battleLogPool.Get().(*FYBattleLog)
	return log
}

// Reset resets the battle log to default state
func (l *FYBattleLog) Reset() {
	l.Index = 0
	l.IsStart = false
	l.GAID = 0
	l.ComboID = 0
	l.MontageID = 0
	l.BuffID = 0
	l.State = 0
	l.Type = BattleLogType_NONE
	l.SourceEntity.Reset()
	l.TargetEntity.Reset()
	l.DR = nil
	l.CasterHash = 0
	l.TargetHash = 0
}

// Recycle returns the battle log to pool
func (l *FYBattleLog) Recycle() {
	l.Reset()
	battleLogPool.Put(l)
}

// String returns string representation
func (l *FYBattleLog) String() string {
	return fmt.Sprintf("Index= %d  type = %d", l.Index, l.Type)
}

// PrintSimpleLog prints simplified battle log
func (l *FYBattleLog) PrintSimpleLog() string {
	var tempStr strings.Builder
	var propStr strings.Builder

	functionHeroProp := func(prop int, value int) string {
		strColor := "<Yellow_Item>"
		plus := "+"
		if value <= 0 {
			strColor = "<Red_Item>"
			plus = ""
		}
		propName := getHeroPropName(prop)
		return fmt.Sprintf("<Blue_Item>%s</>%s%s%d</>", propName, strColor, plus, value)
	}

	functionSkillProp := func(prop int, value int) string {
		strColor := "<Yellow_Item>"
		plus := "+"
		if value <= 0 {
			strColor = "<Red_Item>"
			plus = ""
		}
		propName := getSkillPropName(prop)
		return fmt.Sprintf("<Blue_Item>%s</>%s%s%d</>", propName, strColor, plus, value)
	}

	switch l.Type {
	case BattleLogType_GA:
		return tempStr.String()

	case BattleLogType_COMBO:
		return tempStr.String()

	case BattleLogType_MONTAGE:
		tempStr.Reset()
		if l.IsStart {
			tempStr.WriteString("打出")
		} else {
			return tempStr.String()
		}

		var prop *cfg.SkilldataMontageAttackData
		if ConfigTables != nil {
			prop = ConfigTables.TbMontageAttackData.Get(int32(l.MontageID))
		}
		if prop == nil {
			return ""
		}

		newStr := fmt.Sprintf("%s<Yellow_Item>%s</>", tempStr.String(), prop.GATag)
		tempStr.Reset()
		tempStr.WriteString(newStr)

	case BattleLogType_BUFF_CHANGE:
		tempStr.Reset()
		if l.IsStart {
			tempStr.WriteString("添加")
		} else {
			tempStr.WriteString("移除")
		}

		// TODO: Buff table not available in Go cfg package yet
		// buffProp := cfg.Tables.GetTbBuff().Get(int32(l.BuffID))
		szName := fmt.Sprintf("BuffID_%d", l.BuffID)

		buffStr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("%s Buff<Red_Item>%s</>", buffStr, szName))

	case BattleLogType_COST_ENERGY, BattleLogType_BUFF_VALUE_CHANGE:
		tempStr.Reset()
		propStr.Reset()

		if l.DR != nil {
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionHeroProp(prop, value))
			}, false)
			l.DR.AttackPropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionSkillProp(prop, value))
			}, false)
			tempStr.WriteString(l.SourceEntity.GetHeroNameString() + propStr.String())

			propStr.Reset()
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionHeroProp(prop, value))
			}, true)
			l.DR.AttackPropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionSkillProp(prop, value))
			}, true)

			if tempStr.Len() > 0 {
				tempStr.WriteString(",")
			}
			tempStr.WriteString(l.TargetEntity.GetHeroNameString() + propStr.String())
		}

		// TODO: Buff table not available in Go cfg package yet
		// buffProp := cfg.Tables.GetTbBuff().Get(int32(l.BuffID))
		szName := fmt.Sprintf("BuffID_%d", l.BuffID)

		buffStr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("触发<Red_Item>%s</>效果,%s", szName, buffStr))

	case BattleLogType_DO_ATTACK_DATA:
		tempStr.Reset()
		propStr.Reset()

		if l.DR != nil {
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionHeroProp(prop, value))
			}, false)
			tempStr.WriteString(l.SourceEntity.GetHeroNameString() + propStr.String())

			propStr.Reset()
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				propStr.WriteString(functionHeroProp(prop, value))
			}, true)

			if tempStr.Len() > 0 {
				tempStr.WriteString(",")
			}
			tempStr.WriteString(l.TargetEntity.GetHeroNameString() + propStr.String())

			tempResult := fmt.Sprintf("攻击<Red_Item>%d</>命中 %s", l.DR.AttackID, tempStr.String())
			tempStr.Reset()
			tempStr.WriteString(tempResult)
		}
	}

	ret := fmt.Sprintf("<Red_Item>%d</>-%s %s", l.Index, l.SourceEntity.GetHeroNameString(), tempStr.String())
	return ret
}

// PrintComplexLog prints detailed battle log
func (l *FYBattleLog) PrintComplexLog() string {
	var tempStr strings.Builder

	switch l.Type {
	case BattleLogType_GA:
		if l.IsStart {
			tempStr.WriteString("<Element_QuickFrozen_22>开启</>")
		} else {
			tempStr.WriteString("<Red_Item>结束</>")
		}

		var prop *cfg.SkilldataGACombo
		if ConfigTables != nil {
			prop = ConfigTables.TbGACombox.Get(int32(l.GAID))
		}
		if prop == nil {
			return ""
		}
		AssetName := prop.AbilityPath
		gaString := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("%s GA:<Yellow_Item>%d-%s</>\n当前状态Flag:%d", gaString, l.GAID, AssetName, l.State))

	case BattleLogType_COMBO:
		if l.IsStart {
			tempStr.WriteString("<Element_QuickFrozen_22>开启</>")
		} else {
			tempStr.WriteString("<Red_Item>结束</>")
		}

		comstr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("%s COMBO:<Yellow_Item>%d</>", comstr, l.ComboID))

	case BattleLogType_MONTAGE:
		if l.IsStart {
			tempStr.WriteString("<Element_QuickFrozen_22>开启</>")
		} else {
			tempStr.WriteString("<Red_Item>结束</>")
		}

		var prop *cfg.SkilldataMontageAttackData
		if ConfigTables != nil {
			prop = ConfigTables.TbMontageAttackData.Get(int32(l.MontageID))
		}
		if prop == nil {
			return ""
		}
		mtgStr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("%s Montage:<Yellow_Item>%d</> %s", mtgStr, l.MontageID, prop.GATag))

	case BattleLogType_BUFF_CHANGE:
		if l.IsStart {
			tempStr.WriteString("添加")
		} else {
			tempStr.WriteString("移除")
		}
		buffStr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("%s Buff:%d", buffStr, l.BuffID))

	case BattleLogType_BUFF_VALUE_CHANGE:
		if l.DR != nil {
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(攻击方)<Blue_Item>%s</>(属性)增减<Red_Item>%d</> \n", getHeroPropName(prop), value))
			}, false)
			l.DR.AttackPropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(攻击方)<Blue_Item>%s</>(AttackData)增减<Red_Item>%d</> \n", getSkillPropName(prop), value))
			}, false)
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(受击方)<Blue_Item>%s</>(属性)增减<Red_Item>%d</> \n", getHeroPropName(prop), value))
			}, true)
			l.DR.AttackPropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(受击方)<Blue_Item>%s</>(AttackData)增减<Red_Item>%d</> \n", getSkillPropName(prop), value))
			}, true)
		}

		buffValStr := tempStr.String()
		tempStr.Reset()
		attackID := 0
		if l.DR != nil {
			attackID = l.DR.AttackID
		}
		tempStr.WriteString(fmt.Sprintf("触发Buff:<Yellow_Item>%d</>\n %s", attackID, buffValStr))

	case BattleLogType_DO_ATTACK_DATA:
		attackID := 0
		if l.DR != nil {
			attackID = l.DR.AttackID
		}
		tempStr.WriteString(fmt.Sprintf("基于%d（AttackID）结算伤害\n ", attackID))

		if l.DR != nil {
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(攻击方)<Blue_Item>%s</>(属性)增减<Red_Item>%d</> \n", getHeroPropName(prop), value))
			}, false)
			l.DR.PropertyParallelForEach(func(prop int, value int) {
				tempStr.WriteString(fmt.Sprintf("(受击方)<Blue_Item>%s</>(属性)增减<Red_Item>%d</> \n", getHeroPropName(prop), value))
			}, true)

			for k, v := range l.DR.ExSkillEnergyBackstageLists {
				tempStr.WriteString(fmt.Sprintf("(后台角色)<Yellow_Item>%d</> 大招能量增减<Red_Item>%d</> \n", k, v))
			}
		}

		attactStr := tempStr.String()
		tempStr.Reset()
		tempStr.WriteString(fmt.Sprintf("基于%d（AttackID）结算伤害\n %s", attackID, attactStr))
	}

	targetEntityString := l.TargetEntity.GetInfoString()
	if targetEntityString != "" {
		targetEntityString = fmt.Sprintf("TargetEntity: %s", targetEntityString)
	}

	ret := tempStr.String()
	tempStr.Reset()

	tempStr.WriteString(fmt.Sprintf("BattleLog<Yellow_Item>%d</>\nSourceEntity:%s %s %s",
		l.Index, targetEntityString, targetEntityString, ret))

	return tempStr.String()
}

// getHeroPropName returns the name of a hero property
func getHeroPropName(prop int) string {
	switch prop {
	case cfg.HeroProp_MaxHealth:
		return "MaxHealth"
	case cfg.HeroProp_Health:
		return "Health"
	case cfg.HeroProp_Attack:
		return "Attack"
	case cfg.HeroProp_Defense:
		return "Defense"
	case cfg.HeroProp_Critical_Rate:
		return "Critical_Rate"
	case cfg.HeroProp_Critical_DamageRate:
		return "Critical_DamageRate"
	case cfg.HeroProp_MaxPoise:
		return "MaxPoise"
	case cfg.HeroProp_Poise:
		return "Poise"
	case cfg.HeroProp_DamageReduction:
		return "DamageReduction"
	case cfg.HeroProp_AmplifyDamage:
		return "AmplifyDamage"
	case cfg.HeroProp_MaxStamina:
		return "MaxStamina"
	case cfg.HeroProp_Stamina:
		return "Stamina"
	case cfg.HeroProp_MaxExSkillEnergy:
		return "MaxExSkillEnergy"
	case cfg.HeroProp_ExSkillEnergy:
		return "ExSkillEnergy"
	default:
		return fmt.Sprintf("HeroProp_%d", prop)
	}
}

// getSkillPropName returns the name of a skill property
func getSkillPropName(prop int) string {
	switch prop {
	case cfg.ESkillProp_Damage_Ratio:
		return "Damage_Ratio"
	case cfg.ESkillProp_Damage_Type:
		return "Damage_Type"
	case cfg.ESkillProp_HitPowerType:
		return "HitPowerType"
	case cfg.ESkillProp_Hit_Type:
		return "Hit_Type"
	case cfg.ESkillProp_Buff_ID:
		return "Buff_ID"
	case cfg.ESkillProp_Poise_Attack:
		return "Poise_Attack"
	case cfg.ESkillProp_ExSkillEnergy_Ratio:
		return "ExSkillEnergy_Ratio"
	case cfg.ESkillProp_SPGain:
		return "SPGain"
	case cfg.ESkillProp_Abnormal_Ratio:
		return "Abnormal_Ratio"
	case cfg.ESkillProp_Damage_Base:
		return "Damage_Base"
	case cfg.ESkillProp_Damage_Critical_Rate:
		return "Damage_Critical_Rate"
	case cfg.ESkillProp_Damage_Critical_DamageRate:
		return "Damage_Critical_DamageRate"
	default:
		return fmt.Sprintf("ESkillProp_%d", prop)
	}
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
	Attacker               *actors.Creature                 // 攻击者
	Target                 *actors.Creature                 // 目标
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
