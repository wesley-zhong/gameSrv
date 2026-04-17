// Copyright Epic Games, Inc. All Rights Reserved.

package sight

import (
	"gameSrv/pkg/scene"
	"log"
)

// Grid 视野格子
type Grid struct {
	// 相位实体列表<phasingId, <entityId, Entity>> 0是公共相位
	phasingEntityMap map[int64]map[int64]scene.IEntity
	// TODO: 区域 - 需要实现区域功能
	// regions map[interface{}]struct{}
	// 格子二维坐标
	gridX int32
	gridY int32
}

// NewGrid 创建新的网格
func NewGrid(gridX, gridY int32) *Grid {
	return &Grid{
		gridX: gridX,
		gridY: gridY,
	}
}

// Accept 接受访问者(Visitor设计模式)
func (g *Grid) Accept(visitor interface{}) {
	// TODO: 实现Visitor模式
	if len(g.phasingEntityMap) == 0 {
		return
	}

	// 访问公共相位实体
	g.acceptPhasing(visitor, 0)

	// 访问其他相位实体
	// TODO: 获取相位ID
	// phasingId := visitor.GetSelfEntity().GetPhasingId()
	// if phasingId != 0 {
	// 	// 访问当前所在相位
	// 	g.acceptPhasing(visitor, phasingId)
	// 	// 如果是avatarTeam，且不在专属相位中，额外访问专属相位中的实体
	// 	if avatarTeamActor, ok := visitor.GetSelfEntity().(*actor.AvatarTeamActor); ok && !avatarTeamActor.InPrivatePhasing() {
	// 		g.acceptPhasing(visitor, avatarTeamActor.GetPrivatePhasingId())
	// 	}
	// }
}

// acceptPhasing 访问相位实体
func (g *Grid) acceptPhasing(visitor interface{}, phasingId int64) {
	actorMap, exists := g.phasingEntityMap[phasingId]
	if !exists || len(actorMap) == 0 {
		return
	}

	for _, act := range actorMap {
		// TODO: act.VisitorAccept(visitor)
	}
}

// AddEntity 添加实体
func (g *Grid) AddEntity(act scene.IEntity) {
	// TODO: 获取网格信息
	// if act.GetGrid() != nil {
	// 	log.Printf("[GRID] cur_grid not null. entity:%v  cur_grid:%v   and this_grid:%v", act, act.GetGrid(), g)
	// }

	if g.phasingEntityMap == nil {
		g.phasingEntityMap = make(map[int64]map[int64]scene.IEntity)
	}
	actorMap, exists := g.phasingEntityMap[act.GetPhasingId()]
	if !exists {
		actorMap = make(map[int64]scene.IEntity)
		g.phasingEntityMap[act.GetPhasingId()] = actorMap
	}

	curActor, exists := actorMap[act.GetEntityId()]
	if exists {
		// 这个分支都是异常情况，为了排查bug，日志记录详细一点
		log.Printf("[GRID] entity already exists. this_entity:%v  curGrid:%v", act, g)
	}

	actorMap[act.GetEntityId()] = act
	// TODO: act.SetGrid(g)
}

// DelEntity 删除实体
func (g *Grid) DelEntity(act scene.IEntity) {
	// TODO: 获取网格信息
	// if act.GetGrid() != g {
	// 	log.Printf("[GRID] grid is different. entity: %v,  cur_grid:%v,  and this_grid:%v", act, act.GetGrid(), g)
	// }

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

	if curActor != act {
		// 当前entity是有效的，并且和进入的entity不相同，不能删除
		log.Printf("[GRID] cur_entity != this_entity. cur_entity:%v, this_entity:%v, grid:%v", curActor, act, g)
	}

	// 删除entity对于调用方肯定成功
	// TODO: act.SetGrid(nil)
}

// AddRegion 添加区域
func (g *Grid) AddRegion(region interface{}) {
	// TODO: 实现区域功能
	// if g.regions == nil {
	// 	g.regions = make(map[interface{}]struct{})
	// }
	// g.regions[region] = struct{}{}
}

// DelRegion 删除区域
func (g *Grid) DelRegion(region interface{}) {
	// TODO: 实现区域功能
	// if g.regions == nil {
	// 	log.Printf("[GRID] region_info_ptr_ is null")
	// 	return
	// }
	// delete(g.regions, region)
}

// GetAllRegion 获取所有区域
func (g *Grid) GetAllRegion() []interface{} {
	// TODO: 实现区域功能
	// if g.regions == nil {
	// 	return nil
	// }
	// result := make([]interface{}, 0, len(g.regions))
	// for region := range g.regions {
	// 	result = append(result, region)
	// }
	// return result
	return nil
}

// Clear 清理
func (g *Grid) Clear() {
	g.phasingEntityMap = nil
	// TODO: g.regions = nil
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
	return "[grid: {x=" + string(rune(g.gridX)) + ",y=" + string(rune(g.gridY)) + "} phasingSize = " + string(rune(len(g.phasingEntityMap)))
}
