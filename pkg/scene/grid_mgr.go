// Copyright Epic Games, Inc. All Rights Reserved.

package sight

import (
	"gameSrv/pkg/scene"
	"log"
)

// GridMgr 网格管理器
type GridMgr struct {
	scene     *scene.Scene
	rangeType *VisionLevelEnum
	length    int // 以格子为单位
	width     int // 以格子为单位
	mesh      *Mesh
}

// NewGridMgr 创建新的网格管理器
func NewGridMgr(scn *scene.Scene, rangeType *VisionLevelEnum) *GridMgr {
	return &GridMgr{
		scene:     scn,
		rangeType: rangeType,
	}
}

// Init 初始化
func (g *GridMgr) Init(length, width int) {
	g.length = length
	g.width = width
	g.mesh = NewMesh(length, width)
}

// PlaceEntity 放置实体
func (g *GridMgr) PlaceEntity(actor *scene.IEntity, gridX, gridY int) {
	grid := g.GetGrid(gridX, gridY)
	grid.AddEntity(actor)
}

// PlaceRegionEntity 放置区域实体
func (g *GridMgr) PlaceRegionEntity(region *abstracts.Region, coords []*scene.Coordinate) {
	for _, coord := range coords {
		grid := g.GetGrid(int(coord.X), int(coord.Y))
		if grid == nil {
			log.Printf("placeRegionEntity getGrid doesn't exist! coord = %v, grid_mgr = %v, sceneInfo = %v, region = %v",
				coord, g, g.scene.GetSightModule(), region)
			continue
		}
		grid.AddRegion(region)
		if region.GetGrid() == nil {
			region.SetGrid(grid)
		}
	}

	// 添加范围内的avatarTeam到region内
	for _, playerViewMgr := range g.scene.GetPlayerViewMgrMap() {
		avatarTeam := playerViewMgr.GetPlayer().GetAvatarTeam()
		if region.IsInRegion(avatarTeam.GetLocation().GetPosition()) {
			region.AddEntity(avatarTeam, true)
		}
	}
}

// RemoveEntity 删除实体
func (g *GridMgr) RemoveEntity(actor *abstracts.Actor) {
	grid := actor.GetGrid()
	if grid == nil {
		log.Printf("grid is null: %v", actor)
		return
	}

	grid.DelEntity(actor)
}

// RemoveRegionEntity 删除区域实体
func (g *GridMgr) RemoveRegionEntity(region *abstracts.Region, coords []*scene.Coordinate) {
	for _, coord := range coords {
		grid := g.GetGrid(int(coord.X), int(coord.Y))
		if grid == nil {
			log.Printf("removeRegionEntity getGrid doesn't exist! coord = %v, grid_mgr = %v, sceneInfo = %v, region = %v",
				coord, g, g.scene.GetSightModule(), region)
			continue
		}
		grid.DelRegion(region)
	}
	region.SetGrid(nil)

	// 从region中移除范围内的所有实体
	region.DelAllEntity(true)
}

// EntityMoveTo 实体移动
func (g *GridMgr) EntityMoveTo(actor *abstracts.Actor, gridX, gridY int) bool {
	curGrid := actor.GetGrid()
	if curGrid == nil {
		exception.NewFyUnknownLogicException("entityMoveTo curGrid doesn't exist! sceneInfo = " + g.scene.GetSightModule().String() + " Entity = " + actor.String())
	}

	destGrid := g.GetGrid(gridX, gridY)
	if destGrid == nil {
		exception.NewFyUnknownLogicException("entityMoveTo getGrid doesn't exist! sceneInfo = " + g.scene.GetSightModule().String() + " Entity = " + actor.String() + " posX = " + string(rune(gridX)) + " posY = " + string(rune(gridY)))
	}

	if curGrid == destGrid {
		return false
	}
	curGrid.DelEntity(actor)
	destGrid.AddEntity(actor)
	return true
}

// VisitGridsInSight 获取整个场景中满足visitor条件的entity集合
func (g *GridMgr) VisitGridsInSight(gridX, gridY int, visitor visitor.Visitor) {
	sightRadius := g.rangeType.GetSightRadius()
	minX := maxInt(0, gridX-int(sightRadius))
	maxX := gridX + int(sightRadius)
	minY := maxInt(0, gridY-int(sightRadius))
	maxY := gridY + int(sightRadius)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			g.visitGrid(i, j, visitor)
		}
	}
}

// VisitDiffGrids 计算从from_grid 到to_grid，entity集合的减少和新增
func (g *GridMgr) VisitDiffGrids(fromGridX, fromGridY, toGridX, toGridY int, miss, add visitor.Visitor) {
	sightRadius := int(g.rangeType.GetSightRadius())

	leftX1 := fromGridX - sightRadius
	rightX1 := fromGridX + sightRadius
	lowY1 := fromGridY - sightRadius
	upY1 := fromGridY + sightRadius

	leftX2 := toGridX - sightRadius
	rightX2 := toGridX + sightRadius
	lowY2 := toGridY - sightRadius
	upY2 := toGridY + sightRadius

	// 如果两个视野范围完全不重叠
	if rightX1 < leftX2 || rightX2 < leftX1 || upY1 < lowY2 || upY2 < lowY1 {
		// 添加原始视野中的所有实体到miss
		for i := leftX1; i <= rightX1; i++ {
			for j := lowY1; j <= upY1; j++ {
				g.visitGrid(i, j, miss)
			}
		}
		// 添加新视野中的所有实体到add
		for i := leftX2; i <= rightX2; i++ {
			for j := lowY2; j <= upY2; j++ {
				g.visitGrid(i, j, add)
			}
		}
		return
	}

	// 计算差异部分
	lowY := lowY1
	upY := upY1

	if fromGridY < toGridY {
		// 向下移动
		for j := lowY1; j < lowY2; j++ {
			for i := leftX1; i <= rightX1; i++ {
				g.visitGrid(i, j, miss)
			}
		}
		for j := upY2; j > upY1; j-- {
			for i := leftX2; i <= rightX2; i++ {
				g.visitGrid(i, j, add)
			}
		}
		lowY = lowY2
		upY = upY1
	} else if fromGridY > toGridY {
		// 向上移动
		for j := upY1; j > upY2; j-- {
			for i := leftX1; i <= rightX1; i++ {
				g.visitGrid(i, j, miss)
			}
		}
		for j := lowY2; j < lowY1; j++ {
			for i := leftX2; i <= rightX2; i++ {
				g.visitGrid(i, j, add)
			}
		}
		lowY = lowY1
		upY = upY2
	}

	for j := lowY; j <= upY; j++ {
		if fromGridX < toGridX {
			// 向右移动
			for i := leftX1; i < leftX2; i++ {
				g.visitGrid(i, j, miss)
			}
			for i := rightX2; i > rightX1; i-- {
				g.visitGrid(i, j, add)
			}
		} else if fromGridX > toGridX {
			// 向左移动
			for i := leftX2; i < leftX1; i++ {
				g.visitGrid(i, j, add)
			}
			for i := rightX1; i > rightX2; i-- {
				g.visitGrid(i, j, miss)
			}
		}
	}
}

// visitGrid 访问指定的格子
func (g *GridMgr) visitGrid(gridX, gridY int, visitor visitor.Visitor) {
	grid := g.mesh.FindGrid(gridX, gridY)
	if grid != nil {
		grid.Accept(visitor)
	}
}

// GetGrid 获取指定位置的网格
func (g *GridMgr) GetGrid(gridX, gridY int) *Grid {
	return g.mesh.GetGrid(gridX, gridY)
}

// GetRangeType 获取视野等级类型
func (g *GridMgr) GetRangeType() *VisionLevelEnum {
	return g.rangeType
}

// GetLength 获取长度
func (g *GridMgr) GetLength() int {
	return g.length
}

// GetWidth 获取宽度
func (g *GridMgr) GetWidth() int {
	return g.width
}

// maxInt 返回两个整数中的最大值
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// String 返回字符串表示
func (g *GridMgr) String() string {
	return "[scene_id:" + g.scene.String() + ",range_type:" + g.rangeType.String() +
		",length:" + string(rune(g.length)) + ",width:" + string(rune(g.width)) +
		",sightRadius:" + string(rune(g.rangeType.GetSightRadius())) + "]"
}
