package battle_props

// ShipAvatarBattleProps represents battle properties for ship avatars
type ShipAvatarBattleProps struct {
	*BattleProps
}

// NewShipAvatarBattleProps creates a new ShipAvatarBattleProps
func NewShipAvatarBattleProps(maxProps int) *ShipAvatarBattleProps {
	return &ShipAvatarBattleProps{
		BattleProps: NewBattleProps(),
	}
}

// NewShipAvatarBattlePropsWithOwner creates a new ShipAvatarBattleProps with owner
func NewShipAvatarBattlePropsWithOwner(maxProps int, owner BattlePropsOwner) *ShipAvatarBattleProps {
	return &ShipAvatarBattleProps{
		BattleProps: NewBattlePropsWithOwner(owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (sap *ShipAvatarBattleProps) RecalPropsFormula() {
	if sap.BattleProps == nil {
		return
	}
	sap.BattleProps.mu.Lock()
	defer sap.BattleProps.mu.Unlock()

	PropsUtil.CalcAttributes(sap.CurProps, sap.BaseProps)
}

// ResetProps resets properties
func (sap *ShipAvatarBattleProps) ResetProps() {
	if sap.BattleProps == nil || sap.Owner == nil {
		return
	}

	// Check if owner is a ShipAvatarActor
	// For now, we'll clear and let the ship recalculate

	sap.ClearProps()

	// TODO: Call ship.ReCalAttributes() if owner is ShipAvatarActor
	// if shipAvatarActor, ok := sap.Owner.(*ShipAvatarActor); ok {
	//     ship := shipAvatarActor.GetShip()
	//     if ship != nil {
	//         ship.ReCalAttributes()
	//     }
	// }
}
