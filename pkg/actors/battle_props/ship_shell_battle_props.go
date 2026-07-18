package battle_props

// ShipShellBattleProps represents battle properties for ship shells
type ShipShellBattleProps struct {
	*BattleProps
}

// NewShipShellBattleProps creates a new ShipShellBattleProps
func NewShipShellBattleProps(maxProps int) *ShipShellBattleProps {
	return &ShipShellBattleProps{
		BattleProps: NewBattleProps(),
	}
}

// NewShipShellBattlePropsWithOwner creates a new ShipShellBattleProps with owner
func NewShipShellBattlePropsWithOwner(maxProps int, owner BattlePropsOwner) *ShipShellBattleProps {
	return &ShipShellBattleProps{
		BattleProps: NewBattlePropsWithOwner(owner),
	}
}

// RecalPropsFormula recalculates properties formula
func (ssp *ShipShellBattleProps) RecalPropsFormula() {
	// Ship shell may not need property calculation
	// Implementation can be added if needed
}

// ResetProps resets properties
func (ssp *ShipShellBattleProps) ResetProps() {
	// Ship shell may not need property reset
	// Implementation can be added if needed
}
