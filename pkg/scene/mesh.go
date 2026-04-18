package scene

import (
	"fmt"
)

// Mesh 网格容器
type Mesh struct {
	length    int     // 长度  X
	width     int     // 宽度  Y
	gridArray []*Grid // 网格数组
}

// NewMesh 创建新的网格容器
func NewMesh(length, width int) (*Mesh, error) {
	if length <= 0 || width <= 0 {
		return nil, fmt.Errorf("length and width must be greater than 0")
	}
	mesh := &Mesh{
		length:    length,
		width:     width,
		gridArray: make([]*Grid, length*width),
	}
	return mesh, nil
}

// GetGrid 获取指定位置的网格
func (m *Mesh) GetGrid(gridX, gridY int) *Grid {
	if gridX >= m.length || gridX < 0 || gridY >= m.width || gridY < 0 {
		return nil
	}

	index := gridY*m.length + gridX
	val := m.gridArray[index]
	if val == nil {
		val = NewGrid(int32(gridX), int32(gridY))
		m.gridArray[index] = val
	}
	return val
}

// FindGrid 查找指定位置的网格（不存在则返回nil）
func (m *Mesh) FindGrid(gridX, gridY int) *Grid {
	if gridX >= m.length || gridX < 0 || gridY >= m.width || gridY < 0 {
		return nil
	}
	return m.gridArray[gridY*m.length+gridX]
}

// GetLength 获取长度
func (m *Mesh) GetLength() int {
	return m.length
}

// GetWidth 获取宽度
func (m *Mesh) GetWidth() int {
	return m.width
}
