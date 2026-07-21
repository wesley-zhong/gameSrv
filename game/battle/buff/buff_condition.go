package buff

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"gameSrv/pkg/actors"
)

// CondType represents condition type
type CondType int

const (
	CondTypeHoldCond   CondType = 0 // Holder condition
	CondTypeTargetCond CondType = 1 // Target condition
)

// CheckBuffEffectCondition is a function type for checking buff conditions
type CheckBuffEffectCondition func(target *actors.Creature, cond interface{}) bool

// FYBuffConditionProcess handles buff condition processing
type FYBuffConditionProcess struct {
	// buffConditionFunctions maps condition type IDs to check functions
	buffConditionFunctions map[int]CheckBuffEffectCondition
}

// CheckCondType checks if condition type matches
func (p *FYBuffConditionProcess) CheckCondType(condConfig *cfg.BuffCond, condId int, condType CondType) bool {
	if condId == 0 {
		return true
	}
	if condConfig == nil || condConfig.CnfId == 0 {
		return false
	}
	return int(condConfig.CondType) == int(condType)
}

// CheckConditionValid checks if condition is valid for target
func (p *FYBuffConditionProcess) CheckConditionValid(target *actors.Creature, condId int, condType CondType) bool {
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

	// Get check function based on condition type
	checkFunc := p.getBuffConditionFunction(typeId)
	if checkFunc == nil {
		return false
	}

	return checkFunc(target, buffCond.Condtion)
}

// CheckNeedTags checks if required buff tags exist
func (p *FYBuffConditionProcess) CheckNeedTags(target *actors.Creature, cond interface{}) bool {
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
func (p *FYBuffConditionProcess) CheckNoNeedTags(target *actors.Creature, cond interface{}) bool {
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
func (p *FYBuffConditionProcess) CheckInMontage(target *actors.Creature, cond interface{}) bool {
	// InMontage cfg type not available yet - implementation would check:
	// if target.ActorBattleModule != nil {
	//     montageList := target.ActorBattleModule.GetCurMontageList()
	//     // Check if required montage ID is in list
	// }
	return false
}

// CheckElementType checks if element type matches
func (p *FYBuffConditionProcess) CheckElementType(target *actors.Creature, cond interface{}) bool {
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
func (p *FYBuffConditionProcess) CheckSermonType(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when SermonTypeNeed type is available from cfg
	return false
}

// CheckHeroHasAllTalent checks if hero has all required talents
func (p *FYBuffConditionProcess) CheckHeroHasAllTalent(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when HeroHasAllTalentNeed type is available from cfg
	return true
}

// CheckHeroPropMeetTheProp checks if hero property meets condition
func (p *FYBuffConditionProcess) CheckHeroPropMeetTheProp(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when PropMeetTheProp type is available from cfg
	return false
}

// CheckHeroPropMeetThePropRate checks if hero property rate meets condition
func (p *FYBuffConditionProcess) CheckHeroPropMeetThePropRate(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when PropMeetThePropRate type is available from cfg
	return false
}

// CheckAbnormal checks abnormal conditions
func (p *FYBuffConditionProcess) CheckAbnormal(target *actors.Creature, cond interface{}) bool {
	needAbnormal, ok := cond.(*cfg.AbnormalNeed)
	if !ok || target == nil {
		return false
	}

	// TODO: Check creature's abnormal states when abnormal system is available
	_ = needAbnormal.AbnormalCheckIDs

	return true
}

// getBuffConditionFunction returns the condition check function for a given type
func (p *FYBuffConditionProcess) getBuffConditionFunction(typeId int) CheckBuffEffectCondition {
	// Initialize map if not exists
	if p.buffConditionFunctions == nil {
		p.initConditionFunctions()
	}
	return p.buffConditionFunctions[typeId]
}

// initConditionFunctions initializes condition check functions
func (p *FYBuffConditionProcess) initConditionFunctions() {
	if p.buffConditionFunctions != nil {
		return
	}

	p.buffConditionFunctions = make(map[int]CheckBuffEffectCondition)

	// Register condition check functions with their TypeId values
	p.buffConditionFunctions[int(cfg.TypeId_BuffTagsNeed)] = p.CheckNeedTags
	p.buffConditionFunctions[int(cfg.TypeId_BuffTagsNoNeed)] = p.CheckNoNeedTags
	p.buffConditionFunctions[int(cfg.TypeId_ElementTypeNeed)] = p.CheckElementType
	p.buffConditionFunctions[int(cfg.TypeId_AbnormalNeed)] = p.CheckAbnormal
	// InMontage requires a TypeId constant - add when cfg.InMontage type is available
	// SermonType requires a TypeId constant - add when cfg.SermonTypeNeed type is available
	// HeroHasAllTalent requires a TypeId constant - add when cfg.HeroHasAllTalentNeed type is available
}
