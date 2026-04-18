package actors

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
	"log"
)

// RegionShapeType 区域形状类型
const (
	RegionShapeTypeCircle   = 0 // 圆形
	RegionShapeTypeRect    = 1 // 矩形
	RegionShapeTypePolygon = 2 // 多边形
)

// Region is a region entity in scene
// Region 是场景中的区域实体，用于检测实体进入/离开区域，触发相应事件
type Region struct {
	Entity
	// ShapeType 区域形状类型：0=圆形, 1=矩形, 2=多边形
	ShapeType int32
	// RegionSize 区域尺寸（用于圆形半径或矩形宽高）
	RegionSize *math.Vector3
	// PolyPoints 多边形顶点坐标列表
	PolyPoints []*math.Vector2
	// EntitesInRegion 在区域内的实体列表 <entityId, Entity>
	entitiesInRegion map[int64]scene.IEntity
	// TriggerEnter 触发进入事件
	TriggerEnter bool
	// TriggerLeave 触发离开事件
	TriggerLeave bool
}

func (r *Region) EnterScene(scn scene.IScene, context *VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewRegion 创建新的区域
func NewRegion() *Region {
	return &Region{
		RegionSize:      math.NewVector3(0, 0, 0),
		PolyPoints:      make([]*math.Vector2, 0),
		entitiesInRegion: make(map[int64]scene.IEntity),
		TriggerEnter:    true,
		TriggerLeave:    true,
	}
}

// GetOnlyVisionUid 返回视野UID
func (r *Region) GetOnlyVisionUid() int64 {
	return 0
}

// GetPhasingId 返回相位ID
func (r *Region) GetPhasingId() int64 {
	return r.PhasingID
}

// IntersectType 交集类型
type IntersectType int

const (
	IntersectTypeNone    IntersectType = iota // 无交集
	IntersectTypeInward                      // 进入区域
	IntersectTypeOutward                     // 离开区域
	IntersectTypeCross                       // 跨越区域
)

// IsInRegion 检查位置是否在区域内
func (r *Region) IsInRegion(pos *math.Vector3) bool {
	if pos == nil {
		return false
	}

	switch r.ShapeType {
	case RegionShapeTypeCircle:
		return r.isInCircle(pos)
	case RegionShapeTypeRect:
		return r.isInRect(pos)
	case RegionShapeTypePolygon:
		return r.isInPolygon(pos)
	default:
		return false
	}
}

// isInCircle 检查位置是否在圆形区域内
func (r *Region) IsInCircle(pos *math.Vector3) bool {
	if r.Location == nil || r.Location.Position == nil {
		return false
	}
	if r.RegionSize == nil {
		return false
	}

	// 使用X作为半径
	radius := r.RegionSize.X
	centerX := r.Location.Position.X
	centerY := r.Location.Position.Y

	return math.PointInCircle(pos.X, pos.Y, centerX, centerY, radius)
}

// isInCircle 检查位置是否在圆形区域内
func (r *Region) isInCircle(pos *math.Vector3) bool {
	if r.Location == nil || r.Location.Position == nil {
		return false
	}
	if r.RegionSize == nil {
		return false
	}

	// 使用X作为半径
	radius := r.RegionSize.X
	centerX := r.Location.Position.X
	centerY := r.Location.Position.Y

	return math.PointInCircle(pos.X, pos.Y, centerX, centerY, radius)
}

// isInRect 检查位置是否在矩形区域内
func (r *Region) isInRect(pos *math.Vector3) bool {
	if r.Location == nil || r.Location.Position == nil {
		return false
	}
	if r.RegionSize == nil {
		return false
	}

	centerX := r.Location.Position.X
	centerY := r.Location.Position.Y
	width := r.RegionSize.X
	height := r.RegionSize.Y

	// 将中心点转换为左上角点
	leftX := centerX - width/2
	topY := centerY - height/2

	return math.PointInRect(pos.X, pos.Y, leftX, topY, width, height)
}

// isInPolygon 检查位置是否在多边形区域内
func (r *Region) isInPolygon(pos *math.Vector3) bool {
	if len(r.PolyPoints) < 3 {
		return false
	}

	return math.PointInPolygon(pos.X, pos.Y, r.PolyPoints)
}

// AddEntity 添加实体到区域
func (r *Region) AddEntity(entity interface{}, triggerEvt bool) {
	// 支持两种类型：IEntity 和 *actors.Entity（为了向后兼容）
	var entityId int64

	switch e := entity.(type) {
	case scene.IEntity:
		entityId = e.GetEntityId()
		// TODO: 触发实体进入区域事件
		// e.OnEnterRegion(r)
	case *Entity:
		entityId = e.EntityID
	default:
		return
	}

	if _, exists := r.entitiesInRegion[entityId]; exists {
		return // 已经在区域内，不需要重复添加
	}

	r.entitiesInRegion[entityId] = entity

	if triggerEvt && r.TriggerEnter {
		// TODO: 触发实体进入区域事件
		log.Printf("[Region] entity %d entered region %d", entityId, r.EntityID)
	}
}
	if entity == nil {
		return
	}

	entityId := entity.GetEntityId()
	if _, exists := r.entitiesInRegion[entityId]; exists {
		return // 已经在区域内，不需要重复添加
	}

	r.entitiesInRegion[entityId] = entity

	if triggerEvt && r.TriggerEnter {
		// TODO: 触发实体进入区域事件
		// entity.OnEnterRegion(r)
		log.Printf("[Region] entity %d entered region %d", entityId, r.EntityID)
	}
}

// DelEntity 从区域删除实体
func (r *Region) DelEntity(entity scene.IEntity, triggerEvt bool) {
	if entity == nil {
		return
	}

	entityId := entity.GetEntityId()
	if _, exists := r.entitiesInRegion[entityId]; !exists {
		return // 不在区域内，不需要删除
	}

	delete(r.entitiesInRegion, entityId)

	if triggerEvt && r.TriggerLeave {
		// TODO: 触发实体离开区域事件
		// entity.OnLeaveRegion(r)
		log.Printf("[Region] entity %d left region %d", entityId, r.EntityID)
	}
}

// DelAllEntity 删除区域内所有实体
func (r *Region) DelAllEntity(triggerEvt bool) {
	if triggerEvt && r.TriggerLeave {
		// 触发所有实体离开区域事件
		for entityId := range r.entitiesInRegion {
			// TODO: 触发实体离开区域事件
			log.Printf("[Region] entity %d left region %d (del all)", entityId, r.EntityID)
		}
	}

	// 清空实体列表
	r.entitiesInRegion = make(map[int64]scene.IEntity)
}

// GetEntityCount 获取区域内实体数量
func (r *Region) GetEntityCount() int {
	return len(r.entitiesInRegion)
}

// HasEntity 检查实体是否在区域内
func (r *Region) HasEntity(entityId int64) bool {
	_, exists := r.entitiesInRegion[entityId]
	return exists
}

// GetAllEntities 获取区域内所有实体
func (r *Region) GetAllEntities() []scene.IEntity {
	result := make([]scene.IEntity, 0, len(r.entitiesInRegion))
	for _, entity := range r.entitiesInRegion {
		result = append(result, entity)
	}
	return result
}

// SetTriggerEnter 设置是否触发进入事件
func (r *Region) SetTriggerEnter(trigger bool) {
	r.TriggerEnter = trigger
}

// SetTriggerLeave 设置是否触发离开事件
func (r *Region) SetTriggerLeave(trigger bool) {
	r.TriggerLeave = trigger
}

// SetGrid 设置区域所在的网格（ IEntity 接口实现）
func (r *Region) SetGrid(grid *scene.Grid) {
	r.Grid = grid
}

// GetGrid 获取区域所在的网格
func (r *Region) GetGrid() *scene.Grid {
	return r.Grid
}

// GetCoveredCoordinates 获取区域覆盖的所有格子坐标
// 根据区域形状计算覆盖的所有格子坐标
func (r *Region) GetCoveredCoordinates(sightModule scene.ISceneSightModule) []*scene.Coordinate {
	if sightModule == nil || r.Location == nil || r.Location.Position == nil {
		return nil
	}

	visionLevelEnum := r.GetVisionLevelEnum()
	if visionLevelEnum == nil {
		return nil
	}

	gridWidth := visionLevelEnum.GetGridWidth()
	if gridWidth <= 0 {
		return nil
	}

	switch r.ShapeType {
	case RegionShapeTypeCircle:
		return r.getCircleCoveredCoordinates(sightModule, gridWidth)
	case RegionShapeTypeRect:
		return r.getRectCoveredCoordinates(sightModule, gridWidth)
	case RegionShapeTypePolygon:
		return r.getPolygonCoveredCoordinates(sightModule, gridWidth)
	default:
		// 默认返回区域中心点所在的格子
		coord := sightModule.PosToCoordinate(visionLevelEnum, r.Location.Position.X, r.Location.Position.Y)
		return []*scene.Coordinate{coord}
	}
}

// getCircleCoveredCoordinates 获取圆形区域覆盖的格子坐标
func (r *Region) getCircleCoveredCoordinates(sightModule scene.ISceneSightModule, gridWidth int32) []*scene.Coordinate {
	if r.RegionSize == nil {
		return nil
	}

	radius := r.RegionSize.X
	centerX := r.Location.Position.X
	centerY := r.Location.Position.Y

	// 计算圆形的包围盒
	minX := centerX - radius
	maxX := centerX + radius
	minY := centerY - radius
	maxY := centerY + radius

	radiusSq := int64(radius) * int64(radius)
	coords := make([]*scene.Coordinate, 0)

	// 遍历包围盒内的格子
	for x := minX; x <= maxX; x += gridWidth {
		for y := minY; y <= maxY; y += gridWidth {
			// 检查格子中心是否在圆内
			// 简化：检查格子左上角是否在圆内
			if math.DistanceSquared2(x, y, centerX, centerY) <= radiusSq {
				coord := sightModule.PosToCoordinate(r.GetVisionLevelEnum(), x, y)
				coords = append(coords, coord)
			}
		}
	}

	return coords
}

// getRectCoveredCoordinates 获取矩形区域覆盖的格子坐标
func (r *Region) getRectCoveredCoordinates(sightModule scene.ISceneSightModule, gridWidth int32) []*scene.Coordinate {
	if r.RegionSize == nil {
		return nil
	}

	centerX := r.Location.Position.X
	centerY := r.Location.Position.Y
	width := r.RegionSize.X
	height := r.RegionSize.Y

	// 将中心点转换为左上角点
	leftX := centerX - width/2
	topY := centerY - height/2
	rightX := centerX + width/2
	bottomY := centerY + height/2

	coords := make([]*scene.Coordinate, 0)

	// 遍历矩形内的格子
	for x := leftX; x < rightX; x += gridWidth {
		for y := topY; y < bottomY; y += gridWidth {
			coord := sightModule.PosToCoordinate(r.GetVisionLevelEnum(), x, y)
			coords = append(coords, coord)
		}
	}

	return coords
}

// getPolygonCoveredCoordinates 获取多边形区域覆盖的格子坐标
// 使用扫描线算法计算多边形覆盖的格子
func (r *Region) getPolygonCoveredCoordinates(sightModule scene.ISceneSightModule, gridWidth int32) []*scene.Coordinate {
	if len(r.PolyPoints) < 3 {
		return nil
	}

	// 计算多边形的包围盒
	minX := r.PolyPoints[0].X
	maxX := r.PolyPoints[0].X
	minY := r.PolyPoints[0].Y
	maxY := r.PolyPoints[0].Y

	for _, point := range r.PolyPoints {
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	coords := make([]*scene.Coordinate, 0)

	// 遍历包围盒内的格子
	for x := minX; x <= maxX; x += gridWidth {
		for y := minY; y <= maxY; y += gridWidth {
			// 检查格子左上角是否在多边形内
			if math.PointInPolygon(x, y, r.PolyPoints) {
				coord := sightModule.PosToCoordinate(r.GetVisionLevelEnum(), x, y)
				coords = append(coords, coord)
			}
		}
	}

	return coords
}