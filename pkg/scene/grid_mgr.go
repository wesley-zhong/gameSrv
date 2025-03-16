package scene

import (
	"errors"
	"log"
)

// VisionLevelType 定义视距类型
type VisionLevelType int

// ProtoScene.ProtEntityType 定义实体类型
type ProtEntityType int

const (
	ProtEntityAvatar ProtEntityType = iota // 假设 PROT_ENTITY_AVATAR 是 0
)

// GridMgr 定义网格管理器
type GridMgr struct {
	scene       *Scene
	rangeType   VisionLevelType
	length      int
	width       int
	sightRadius int
	mesh        *Mesh
}

// NewGridMgr 创建新的网格管理器
func NewGridMgr(scene *Scene, rangeType VisionLevelType) *GridMgr {
	return &GridMgr{
		scene:     scene,
		rangeType: rangeType,
	}
}

// Init 初始化网格管理器
func (g *GridMgr) Init(length, width, sightRadius int) {
	g.length = length
	g.width = width
	g.sightRadius = sightRadius
	g.mesh, _ = NewMesh(length, width)
}

// PlaceEntity 放置实体到指定坐标
func (g *GridMgr) PlaceEntity(actor Actor, coord *Coordinate) error {
	grid := g.GetGrid(coord.X, coord.Y)
	if grid == nil {
		log.Printf("[WARN] getGrid fails. entity: %d @ coord: %v", actor.GetConfigId(), coord)
		return errors.New("grid not found")
	}

	if err := grid.AddEntity(actor); err != nil {
		log.Printf("[WARN] addEntity fails. entity: %d coord: %v", actor.GetConfigId(), coord)
		return err
	}

	actor.SetCoordinate(coord)
	if actor.GetEntityType() == ProtEntityAvatar {
		log.Printf("[GRID] place avatar: %d @ coord: %v", actor.GetConfigId(), coord)
	}

	log.Printf("[FY] place entity_id: %d @ coord: %v", actor.GetActorId(), actor.GetCoordinate())
	return nil
}

// RemoveEntity 移除实体
func (g *GridMgr) RemoveEntity(actor Actor) error {
	grid := actor.GetGrid()
	if grid == nil {
		log.Printf("[ERROR] grid is null: %v", actor)
		return errors.New("grid is null")
	}

	if err := grid.DelEntity(actor); err != nil {
		log.Printf("[ERROR] delEntity fails: %v", actor)
		return err
	}

	if actor.GetEntityType() == ProtEntityAvatar {
		log.Printf("[GRID] remove avatar: %v @ %v", actor, grid)
	}

	log.Printf("[FY] remove entity_id: %d @ %v", actor.GetActorId(), actor.GetCoordinate())
	return nil
}

// EntityMoveTo 移动实体到目标坐标
func (g *GridMgr) EntityMoveTo(actor Actor, destCoord *Coordinate) error {
	curCoord := actor.GetCoordinate()
	curGrid := g.GetGrid(curCoord.X, curCoord.Y)
	if curGrid == nil {
		log.Printf("[WARN] curCoord: %v grid doesn't exist! grid_mgr: %v", curCoord, g)
		return errors.New("current grid not found")
	}

	destGrid := g.GetGrid(destCoord.X, destCoord.Y)
	if destGrid == nil {
		log.Printf("[WARN] destCoord: %v grid doesn't exist! grid_mgr: %v", destCoord, g)
		return errors.New("destination grid not found")
	}

	if curGrid == destGrid {
		return nil
	}

	log.Printf("------------actor id = %d move from %v to %v", actor.GetActorId(), curCoord, destCoord)
	if err := curGrid.DelEntity(actor); err != nil {
		log.Printf("[WARN] delEntity fails: %v", actor)
		return err
	}

	if err := destGrid.AddEntity(actor); err != nil {
		log.Printf("[WARN] addEntity fails: %v", actor)
		return err
	}

	actor.SetCoordinate(destCoord)
	return nil
}

// VisitGridsInSight 访问视野范围内的网格
func (g *GridMgr) VisitGridsInSight(center Coordinate, visitor Visitor, sightRadius int) {
	if sightRadius == 0 {
		sightRadius = g.sightRadius
	}

	for i := max(0, center.GetX()-sightRadius); i <= center.GetX()+sightRadius; i++ {
		for j := max(0, center.GetY()-sightRadius); j <= center.GetY()+sightRadius; j++ {
			g.VisitGrid(i, j, visitor)
		}
	}
}

// VisitDiffGrids 计算从 fromCoord 到 toCoord 的实体集合变化
func (g *GridMgr) VisitDiffGrids(center1, center2 Coordinate, t1, t2 Visitor) {
	leftX1 := center1.GetX() - g.sightRadius
	rightX1 := center1.GetX() + g.sightRadius
	lowY1 := center1.GetY() - g.sightRadius
	upY1 := center1.GetY() + g.sightRadius

	leftX2 := center2.GetX() - g.sightRadius
	rightX2 := center2.GetX() + g.sightRadius
	lowY2 := center2.GetY() - g.sightRadius
	upY2 := center2.GetY() + g.sightRadius

	if rightX1 < leftX2 || rightX2 < leftX1 || upY1 < lowY2 || upY2 < lowY1 {
		for i := leftX1; i <= rightX1; i++ {
			for j := lowY1; j <= upY1; j++ {
				g.VisitGrid(i, j, t1)
			}
		}
		for i := leftX2; i <= rightX2; i++ {
			for j := lowY2; j <= upY2; j++ {
				g.VisitGrid(i, j, t2)
			}
		}
		return
	}

	lowY := lowY1
	upY := upY1
	if center1.GetY() < center2.GetY() {
		for j := lowY1; j < lowY2; j++ {
			for i := leftX1; i <= rightX1; i++ {
				g.VisitGrid(i, j, t1)
			}
		}
		for j := upY2; j > upY1; j-- {
			for i := leftX2; i <= rightX2; i++ {
				g.VisitGrid(i, j, t2)
			}
		}
		lowY = lowY2
		upY = upY1
	} else if center1.GetY() > center2.GetY() {
		for j := upY1; j > upY2; j-- {
			for i := leftX1; i <= rightX1; i++ {
				g.VisitGrid(i, j, t1)
			}
		}
		for j := lowY2; j < lowY1; j++ {
			for i := leftX2; i <= rightX2; i++ {
				g.VisitGrid(i, j, t2)
			}
		}
		lowY = lowY1
		upY = upY2
	}

	for j := lowY; j <= upY; j++ {
		if center1.GetX() < center2.GetX() {
			for i := leftX1; i < leftX2; i++ {
				g.VisitGrid(i, j, t1)
			}
			for i := rightX2; i > rightX1; i-- {
				g.VisitGrid(i, j, t2)
			}
		} else if center1.GetX() > center2.GetX() {
			for i := leftX2; i < leftX1; i++ {
				g.VisitGrid(i, j, t2)
			}
			for i := rightX1; i > rightX2; i-- {
				g.VisitGrid(i, j, t1)
			}
		}
	}
}

// VisitGrid 访问指定坐标的网格
func (g *GridMgr) VisitGrid(x, y int, visitor Visitor) {
	grid := g.GetGrid(x, y)
	if grid != nil {
		grid.Accept(visitor)
	}
}

// GetGrid 获取指定坐标的网格
func (g *GridMgr) GetGrid(x, y int) *Grid {
	if g.mesh == nil {
		return nil
	}
	return g.mesh.GetGrid(x, y)
}

// GetGridByCoord 通过坐标获取网格
func (g *GridMgr) GetGridByCoord(coord Coordinate) *Grid {
	return g.GetGrid(coord.GetX(), coord.GetY())
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
