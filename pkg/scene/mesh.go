package scene

import (
	"errors"
	"fmt"
)

// Mesh 定义网格结构
type Mesh struct {
	length    int     // 长度
	width     int     // 宽度
	gridArray []*Grid // 存储网格的数组
}

// NewMesh 创建新的网格
func NewMesh(length, width int) (*Mesh, error) {
	if length <= 0 || width <= 0 {
		return nil, errors.New("length and width must be greater than 0")
	}

	mesh := &Mesh{
		length:    length,
		width:     width,
		gridArray: make([]*Grid, length*width),
	}

	mesh.init()
	return mesh, nil
}

// GetGrid 获取指定坐标的网格
func (m *Mesh) GetGrid(x, y int) *Grid {
	if x >= m.length || x < 0 || y >= m.width || y < 0 {
		fmt.Errorf("coordinates out of bounds: x=%d, y=%d", x, y)
		return nil
	}

	index := y*m.width + x
	if m.gridArray[index] == nil {
		m.gridArray[index] = NewGrid(x, y)
	}

	return m.gridArray[index]
}

// FindGrid 查找指定坐标的网格
func (m *Mesh) FindGrid(x, y int) (*Grid, error) {
	if x >= m.length || x < 0 || y >= m.width || y < 0 {
		return nil, fmt.Errorf("coordinates out of bounds: x=%d, y=%d", x, y)
	}

	return m.gridArray[y*m.width+x], nil
}

// init 初始化网格
func (m *Mesh) init() {
	// 可以在这里添加初始化逻辑
}
