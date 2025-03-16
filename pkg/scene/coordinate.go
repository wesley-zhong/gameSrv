package scene

import "fmt"

// Coordinate 定义二维整数坐标结构
type Coordinate struct {
	X int
	Y int
}

// NewCoordinate 创建新的坐标
func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

// GetX 获取 X 坐标
func (c *Coordinate) GetX() int {
	return c.X
}

// GetY 获取 Y 坐标
func (c *Coordinate) GetY() int {
	return c.Y
}

// String 实现 Stringer 接口
func (c *Coordinate) String() string {
	return fmt.Sprintf("Coordinate{X=%d, Y=%d}", c.X, c.Y)
}

// Equals 判断两个坐标是否相等
func (c *Coordinate) Equals(other *Coordinate) bool {
	if other == nil {
		return false
	}
	return c.X == other.X && c.Y == other.Y
}

// HashCode 计算坐标的哈希值
func (c *Coordinate) HashCode() int {
	return c.X*31 + c.Y
}
