package battle_props

// DropPodBattleProps represents battle properties for drop pods
type DropPodBattleProps struct {
	*BattleProps
}

// NewDropPodBattleProps creates a new DropPodBattleProps
func NewDropPodBattleProps() *DropPodBattleProps {
	return &DropPodBattleProps{
		BattleProps: NewBattleProps(),
	}
}

// NewDropPodBattlePropsWithOwner creates a new DropPodBattleProps with owner
func NewDropPodBattlePropsWithOwner(owner BattlePropsOwner) *DropPodBattleProps {
	return &DropPodBattleProps{
		BattleProps: NewBattlePropsWithOwner(owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (dpp *DropPodBattleProps) RecalPropsFormula() {
	if dpp.BattleProps == nil {
		return
	}
	dpp.BattleProps.mu.Lock()
	defer dpp.BattleProps.mu.Unlock()

	PropsUtil.CalcAttributes(dpp.CurProps, dpp.BaseProps)
}

// ResetProps resets properties
func (dpp *DropPodBattleProps) ResetProps() {
	if dpp.BattleProps == nil || dpp.Owner == nil {
		return
	}

	// Check if owner is a DropPodAvatarActor
	// For now, we'll clear and let the drop pod recalculate

	dpp.ClearProps()

	// TODO: Call dropPodAvatar.ReCalAttributes() if owner is DropPodAvatarActor
	// if dropPodAvatarActor, ok := dpp.Owner.(*DropPodAvatarActor); ok {
	//     dropPodAvatar := dropPodAvatarActor.GetDropPodAvatar()
	//     if dropPodAvatar != nil {
	//         dropPodAvatar.ReCalAttributes()
	//     }
	// }
}
