package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/pkg/actors"
	"sync"
)

// CondType represents condition type
type CondType int

const (
	CondTypeHoldCond   CondType = 0 // Holder condition
	CondTypeTargetCond CondType = 1 // Target condition
)

// CheckBuffEffectCondition is a function type for checking buff conditions
type CheckBuffEffectCondition func(target *actors.Creature, cond interface{}) bool

var (
	buffConditionFunctions map[int]CheckBuffEffectCondition
	initConditionOnce      sync.Once
)

// CheckCondType checks if condition type matches
func CheckCondType(condConfig *cfg.BuffCond, condId int, condType CondType) bool {
	if condId == 0 {
		return true
	}
	if condConfig == nil || condConfig.CnfId == 0 {
		return false
	}
	return int(condConfig.CondType) == int(condType)
}

// CheckConditionValid checks if condition is valid for target
func CheckConditionValid(target *actors.Creature, condId int, condType CondType) bool {
	if condId == 0 {
		return true
	}

	if gamedata.Tables == nil || gamedata.Tables.TbBuffCond == nil {
		return false
	}

	buffCond := gamedata.Tables.TbBuffCond.Get(int32(condId))
	if buffCond == nil {
		return false
	}

	// Get type ID from interface{}
	var typeId int
	if typeGetter, ok := buffCond.Condtion.(interface{ GetTypeId() int32 }); ok {
		typeId = int(typeGetter.GetTypeId())
	} else {
		return false
	}

	// Initialize and get check function based on condition type
	initConditionOnce.Do(initConditionFunctions)
	checkFunc := buffConditionFunctions[typeId]
	if checkFunc == nil {
		return false
	}

	return checkFunc(target, buffCond.Condtion)
}

// CheckNeedTags checks if required buff tags exist
func CheckNeedTags(target *actors.Creature, cond interface{}) bool {
	needTags, ok := cond.(*cfg.BuffTagsNeed)
	if !ok || target == nil {
		return false
	}

	// Check if creature has all required tags
	for _, tag := range needTags.BuffTags {
		if !target.HasBuffTag(int(tag)) {
			return false
		}
	}
	return true
}

// CheckNoNeedTags checks if buff tags don't exist
func CheckNoNeedTags(target *actors.Creature, cond interface{}) bool {
	noNeedTags, ok := cond.(*cfg.BuffTagsNoNeed)
	if !ok || target == nil {
		return false
	}

	// Check if creature doesn't have any of the forbidden tags
	for _, tag := range noNeedTags.BuffTags {
		if target.HasBuffTag(int(tag)) {
			return false
		}
	}
	return true
}

// CheckInMontage checks if target is in a specific montage
func CheckInMontage(target *actors.Creature, cond interface{}) bool {
	// InMontage cfg type not available yet - implementation would check:
	// if target.ActorBattleModule != nil {
	//     montageList := target.ActorBattleModule.GetCurMontageList()
	//     // Check if required montage ID is in list
	// }
	return false
}

// CheckElementType checks if element type matches
func CheckElementType(target *actors.Creature, cond interface{}) bool {
	needElement, ok := cond.(*cfg.ElementTypeNeed)
	if !ok || target == nil {
		return false
	}

	// Check if creature's element type matches any required type
	// Element system will need to be implemented when element types are defined
	_ = needElement.ElementTypeIDs

	// Placeholder: return true if no specific element requirement
	return len(needElement.ElementTypeIDs) == 0
}

// CheckSermonType checks if sermon type matches
func CheckSermonType(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when SermonTypeNeed type is available from cfg
	return false
}

// CheckHeroHasAllTalent checks if hero has all required talents
func CheckHeroHasAllTalent(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when HeroHasAllTalentNeed type is available from cfg
	return true
}

// CheckHeroPropMeetTheProp checks if hero property meets condition
func CheckHeroPropMeetTheProp(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when PropMeetTheProp type is available from cfg
	return false
}

// CheckHeroPropMeetThePropRate checks if hero property rate meets condition
func CheckHeroPropMeetThePropRate(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when PropMeetThePropRate type is available from cfg
	return false
}

// CheckAbnormal checks abnormal conditions
func CheckAbnormal(target *actors.Creature, cond interface{}) bool {
	needAbnormal, ok := cond.(*cfg.AbnormalNeed)
	if !ok || target == nil {
		return false
	}

	// TODO: Check creature's abnormal states when abnormal system is available
	_ = needAbnormal.AbnormalCheckIDs

	return true
}

// initConditionFunctions initializes condition check functions
func initConditionFunctions() {
	if buffConditionFunctions != nil {
		return
	}

	buffConditionFunctions = make(map[int]CheckBuffEffectCondition)

	// Register condition check functions with their TypeId values
	buffConditionFunctions[int(cfg.TypeId_BuffTagsNeed)] = CheckNeedTags
	buffConditionFunctions[int(cfg.TypeId_BuffTagsNoNeed)] = CheckNoNeedTags
	buffConditionFunctions[int(cfg.TypeId_ElementTypeNeed)] = CheckElementType
	buffConditionFunctions[int(cfg.TypeId_AbnormalNeed)] = CheckAbnormal
	// InMontage requires a TypeId constant - add when cfg.InMontage type is available
	// SermonType requires a TypeId constant - add when cfg.SermonTypeNeed type is available
	// HeroHasAllTalent requires a TypeId constant - add when cfg.HeroHasAllTalentNeed type is available
}
