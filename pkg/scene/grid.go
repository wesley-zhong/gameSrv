package scene

import (
	"fmt"
	"log"
)

// Grid 视野格子
type Grid struct {
	// 相位实体列表<phasingId, <entityId, Entity>> 0是公共相位
	phasingEntityMap map[int64]map[int64]IEntity
	// regions 区域列表 <regionId, Region>
	regions map[int64]IEntity
	// 格子二维坐标
	gridX int32
	gridY int32
}

// NewGrid 创建新的网格
func NewGrid(gridX, gridY int32) *Grid {
	return &Grid{
		gridX:            gridX,
		gridY:            gridY,
		regions:          make(map[int64]IEntity),
		phasingEntityMap: make(map[int64]map[int64]IEntity),
	}
}

// Accept 接受访问者(Visitor设计模式)
func (g *Grid) Accept(visitor interface{}) {
	if len(g.phasingEntityMap) == 0 {
		return
	}

	// 访问公共相位实体
	g.acceptPhasing(visitor, 0)

	// 访问其他相位实体
	if v, ok := visitor.(IVisitor); ok {
		phasingId := v.GetSelfEntity().GetPhasingId()
		if phasingId != 0 {
			// 访问当前所在相位
			g.acceptPhasing(visitor, phasingId)
			// TODO: 如果是avatarTeam，且不在专属相位中，额外访问专属相位中的实体
			// 这里需要判断是否有专属相位ID的方法，并且获取玩家玩家的专属相位ID
			// 这需要 IGamePlayer 接口添加 GetPrivatePhasingId 方法
		}
	}
}

// acceptPhasing 访问相位实体
func (g *Grid) acceptPhasing(visitor interface{}, phasingId int64) {
	actorMap, exists := g.phasingEntityMap[phasingId]
	if !exists || len(actorMap) == 0 {
		return
	}

	if v, ok := visitor.(IVisitor); ok {
		for _, act := range actorMap {
			v.VisitEntity(act)
		}
	}
}

// AddEntity 添加实体
func (g *Grid) AddEntity(act IEntity) {
	// 获取网格信息
	oldGrid := act.GetGrid()
	if oldGrid != nil && oldGrid != g {
		log.Printf("[GRID] cur_grid not null. entity:%v  cur_grid:%v   and this_grid:%v", act, oldGrid, g)
		return
	}

	if g.phasingEntityMap == nil {
		g.phasingEntityMap = make(map[int64]map[int64]IEntity)
	}
	actorMap, exists := g.phasingEntityMap[act.GetPhasingId()]
	if !exists {
		actorMap = make(map[int64]IEntity)
		g.phasingEntityMap[act.GetPhasingId()] = actorMap
	}

	curActor, exists := actorMap[act.GetEntityId()]
	if exists {
		// 这个分支都是异常情况，为了排查bug，日志记录详细一点
		log.Printf("[GRID] entity already exists. this_entity:%v  cur_actor:%v  curGrid:%v", act, curActor, g)
	}

	actorMap[act.GetEntityId()] = act

	// 设置实体的网格
	act.SetGrid(g)
}

// DelEntity 删除实体
func (g *Grid) DelEntity(act IEntity) {
	// 获取网格信息
	oldGrid := act.GetGrid()
	if oldGrid != g {
		log.Printf("[GRID] grid is different. entity: %v,  cur_grid:%v,  and this_grid:%v", act, oldGrid, g)
		return
	}

	actorMap, exists := g.phasingEntityMap[act.GetPhasingId()]
	if !exists {
		return
	}

	curActor, exists := actorMap[act.GetEntityId()]
	if !exists {
		return
	}

	delete(actorMap, act.GetEntityId())

	// 当前相位没有实体，则移除该相位，避免内存问题
	if len(actorMap) == 0 {
		delete(g.phasingEntityMap, act.GetPhasingId())
	}

	if curActor != nil && curActor != act {
		// 当前entity是有效的，并且和进入的entity不相同，不能删除
		log.Printf("[GRID] cur_entity != this_entity. cur_entity:%v, this_entity:%v, grid:%v", curActor, act, g)
	}

	// 删除entity对于调用方肯定成功
	act.SetGrid(nil)
}

// AddRegion 添加区域
func (g *Grid) AddRegion(region IEntity) {
	if region == nil {
		log.Printf("[GRID] AddRegion failed, region is nil")
		return
	}

	if g.regions == nil {
		g.regions = make(map[int64]IEntity)
	}

	g.regions[region.GetEntityId()] = region
}

// DelRegion 删除区域
func (g *Grid) DelRegion(region IEntity) {
	if region == nil {
		log.Printf("[GRID] DelRegion failed, region is nil")
		return
	}

	if g.regions == nil {
		log.Printf("[GRID] DelRegion failed, regions map is nil")
		return
	}

	delete(g.regions, region.GetEntityId())
}

// GetAllRegion 获取所有区域
func (g *Grid) GetAllRegion() []interface{} {
	if g.regions == nil {
		return nil
	}

	result := make([]interface{}, 0, len(g.regions))
	for _, region := range g.regions {
		result = append(result, region)
	}
	return result
}

// Clear 清理
func (g *Grid) Clear() {
	g.phasingEntityMap = nil
	g.regions = nil
}

// GetGridX 获取网格X坐标
func (g *Grid) GetGridX() int32 {
	return g.gridX
}

// GetGridY 获取网格Y坐标
func (g *Grid) GetGridY() int32 {
	return g.gridY
}

// String 返回字符串表示
func (g *Grid) String() string {
	return fmt.Sprintf("[grid: {x=%d,y=%d} phasingSize=%d]", g.gridX, g.gridY, len(g.phasingEntityMap))
}
