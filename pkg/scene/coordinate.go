package scene

// Coordinate 二维整数坐标
type Coordinate struct {
	X int32
	Y int32
}

// NewCoordinate 创建新的Coordinate
func NewCoordinate(x, y int32) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

// Equals 判断是否相等
func (c *Coordinate) Equals(other *Coordinate) bool {
	if c == other {
		return true
	}
	if other == nil {
		return false
	}
	return c.X == other.X && c.Y == other.Y
}

// String 返回字符串表示
func (c *Coordinate) String() string {
	return "Coordinate{" + "x=" + string(rune(c.X)) + ", y=" + string(rune(c.Y)) + "}"
}
