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
	// TODO: Implement when BuffTagsNeed type is available from cfg
	return true
}

// CheckNoNeedTags checks if buff tags don't exist
func (p *FYBuffConditionProcess) CheckNoNeedTags(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when BuffTagsNoNeed type is available from cfg
	return true
}

// CheckInMontage checks if target is in a specific montage
func (p *FYBuffConditionProcess) CheckInMontage(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when InMontage type is available from cfg
	// Server can't easily determine montage state
	return false
}

// CheckElementType checks if element type matches
func (p *FYBuffConditionProcess) CheckElementType(target *actors.Creature, cond interface{}) bool {
	// TODO: Implement when ElementTypeNeed type is available from cfg
	return false
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
	// TODO: Implement when AbnormalNeed type is available from cfg
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

	// Register condition check functions
	// Type IDs should match those in the cfg package
	// These are placeholder mappings - actual IDs need to match cfg values
	// p.buffConditionFunctions[cfg.BuffTagsNeed.__ID__] = p.CheckNeedTags
	// p.buffConditionFunctions[cfg.BuffTagsNoNeed.__ID__] = p.CheckNoNeedTags
	// p.buffConditionFunctions[cfg.InMontage.__ID__] = p.CheckInMontage
	// p.buffConditionFunctions[cfg.AbnormalNeed.__ID__] = p.CheckAbnormal
	// p.buffConditionFunctions[cfg.ElementTypeNeed.__ID__] = p.CheckElementType
	// p.buffConditionFunctions[cfg.SermonTypeNeed.__ID__] = p.CheckSermonType
	// p.buffConditionFunctions[cfg.HeroHasAllTalentNeed.__ID__] = p.CheckHeroHasAllTalent
	// p.buffConditionFunctions[cfg.MonsterTypeNeed.__ID__] = p.CheckHeroPropMeetTheProp
	// p.buffConditionFunctions[cfg.PropMeetThePropRate.__ID__] = p.CheckHeroPropMeetThePropRate
}
