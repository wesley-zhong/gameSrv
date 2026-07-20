package battle

// ShipClipModule represents ship clip module
type ShipClipModule struct {
	Position        int // 炮位
	ShellCfgId      int // 实际的炮弹id
	ClipShellCount  int // 弹夹里的炮弹数量
	BagShellCount   int // 背包里的炮弹数量
	Capacity        int // 弹夹容量
	FireShellCfgId  int // 发射时炮弹id
}

// NewShipClipModule creates a new ShipClipModule
func NewShipClipModule() *ShipClipModule {
	return &ShipClipModule{}
}

// AddClipShellCount adds shells to clip
func (m *ShipClipModule) AddClipShellCount(count int) {
	m.ClipShellCount += count
}

// DelClipShellCount removes shells from clip
func (m *ShipClipModule) DelClipShellCount(count int) {
	m.ClipShellCount -= count
}

// ToPack converts to protobuf message (placeholder)
func (m *ShipClipModule) ToPack() interface{} {
	// TODO: implement protobuf conversion
	// ProtoBattle.OceanClipInfo.Builder
	return struct {
		Position        int
		ShellCfgId      int
		ClipShellCount  int
		BagShellCount   int
	}{
		Position:        m.Position,
		ShellCfgId:      m.FireShellCfgId,
		ClipShellCount:  m.ClipShellCount,
		BagShellCount:   m.BagShellCount,
	}
}