package scene

import (
	"errors"
	"fmt"
	"log"
)

// ForeachPolicy 定义遍历策略
type ForeachPolicy int

const (
	Continue ForeachPolicy = iota // 继续遍历
	Break                         // 中断遍历
)

// Region 定义区域接口
type Region interface{}

// Grid 定义网格结构
type Grid struct {
	entityMap map[int64]Actor     // 存储实体的映射
	regions   map[Region]struct{} // 存储区域的集合
	coordX    int                 // 网格的 X 坐标
	coordY    int                 // 网格的 Y 坐标
}

// NewGrid 创建新的网格
func NewGrid(x, y int) *Grid {
	return &Grid{
		coordX:    x,
		coordY:    y,
		entityMap: make(map[int64]Actor),
		regions:   make(map[Region]struct{}),
	}
}

// Accept 接受访问者
func (g *Grid) Accept(visitor Visitor) {
	if len(g.entityMap) == 0 {
		return
	}

	for _, actor := range g.entityMap {
		actor.Accept(visitor)
	}
}

// Clear 清理网格
func (g *Grid) Clear() {
	g.entityMap = nil
	g.regions = nil
}

// AddEntity 添加实体
func (g *Grid) AddEntity(actor Actor) error {
	if actor.GetGrid() != nil {
		return fmt.Errorf("[GRID] cur_grid not null. entity: %v, cur_grid: %v, this_grid: %v", actor, actor.GetGrid(), g)
	}

	if g.entityMap == nil {
		g.entityMap = make(map[int64]Actor)
	}

	if _, exists := g.entityMap[actor.GetActorID()]; exists {
		return fmt.Errorf("[GRID] entity already exists. this_entity: %v, curGrid: %v", actor, g)
	}

	g.entityMap[actor.GetActorID()] = actor
	actor.SetGrid(g)

	log.Printf("[GRID] add entity: entity_id: %d, grid: %v, curSize = %d", actor.GetActorID(), g, len(g.entityMap))
	return nil
}

// DelEntity 删除实体
func (g *Grid) DelEntity(actor Actor) error {
	if actor.GetGrid() != g {
		return fmt.Errorf("[GRID] grid is different. entity: %v, cur_grid: %v, this_grid: %v", actor, actor.GetGrid(), g)
	}

	if g.entityMap == nil {
		return errors.New("[GRID] entity_info_ptr_ is null")
	}

	if _, exists := g.entityMap[actor.GetActorID()]; !exists {
		return fmt.Errorf("[GRID] cur_entity != this_entity. cur_entity: %v, this_entity: %v, grid: %v", g.entityMap[actor.GetActorID()], actor, g)
	}

	delete(g.entityMap, actor.GetActorID())
	actor.SetGrid(nil)

	log.Printf("[GRID] del entity: entity_id: %d, grid: %v, cur size = %d", actor.GetActorID(), g, len(g.entityMap))
	return nil
}

// AddRegion 添加区域
func (g *Grid) AddRegion(region Region) error {
	if g.regions == nil {
		g.regions = make(map[Region]struct{})
	}

	if _, exists := g.regions[region]; exists {
		return fmt.Errorf("[GRID] region already exists: %v", region)
	}

	g.regions[region] = struct{}{}
	return nil
}

// DelRegion 删除区域
func (g *Grid) DelRegion(region Region) error {
	if g.regions == nil {
		return errors.New("[GRID] region_info_ptr_ is null")
	}

	if _, exists := g.regions[region]; !exists {
		return fmt.Errorf("[GRID] region not exists: %v", region)
	}

	delete(g.regions, region)
	return nil
}

// HasEntity 判断是否有实体
func (g *Grid) HasEntity() bool {
	return len(g.entityMap) > 0
}

// GetAllEntity 获取所有实体
func (g *Grid) GetAllEntity() []Actor {
	entities := make([]Actor, 0, len(g.entityMap))
	for _, actor := range g.entityMap {
		entities = append(entities, actor)
	}
	return entities
}

// GetAllRegion 获取所有区域
func (g *Grid) GetAllRegion() []Region {
	regions := make([]Region, 0, len(g.regions))
	for region := range g.regions {
		regions = append(regions, region)
	}
	return regions
}

// ForeachRegion 遍历所有区域
func (g *Grid) ForeachRegion(f func(region Region) ForeachPolicy) int {
	if len(g.regions) == 0 {
		return 0
	}

	for region := range g.regions {
		policy := f(region)
		if policy == Break {
			return 1
		}
	}
	return 0
}

// String 实现 Stringer 接口
func (g *Grid) String() string {
	if g.entityMap != nil {
		return fmt.Sprintf("[grid: x = %d, y = %d, size = %d]", g.coordX, g.coordY, len(g.entityMap))
	}
	return fmt.Sprintf("[grid: x = %d, y = %d, size = 0]", g.coordX, g.coordY)
}
