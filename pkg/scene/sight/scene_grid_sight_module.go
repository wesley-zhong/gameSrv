// Copyright Epic Games, Inc. All Rights Reserved.

package sight

import (
	"errors"
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

const (
	UINT16_MAX = 65535
)

// SceneGridSightModule 网格场景视野模块
type SceneGridSightModule struct {
	*SceneSightModule
	gridMgrs []*GridMgr
}

// NewSceneGridSightModule 创建新的网格场景视野模块
func NewSceneGridSightModule(scn scene.IScene) *SceneGridSightModule {
	gridMgrs := make([]*GridMgr, len(scene.VisionLevelEnums))
	return &SceneGridSightModule{
		SceneSightModule: NewSceneSightModule(scn),
		gridMgrs:         gridMgrs,
	}
}

// Init 初始化组件
// beginPos sceneSize 里面存储的都是实际距离
func (s *SceneGridSightModule) Init(beginPos, sceneSize *math.Vector2) error {
	if err := s.SceneSightModule.Init(beginPos, sceneSize); err != nil {
		return err
	}

	for i := 0; i < len(s.gridMgrs); i++ {
		rangeType := scene.ForVisionLevelType(scene.VisionLevelType(i))
		if rangeType == nil {
			return errors.New("VisionLevelEnum is null, i = " + string(rune(i)))
		}

		// 初始化各级管理器
		err := s.createGridMgr(rangeType)
		if err != nil {
			return err
		}
	}
	return nil
}

// createGridMgr 创建网格管理器
func (s *SceneGridSightModule) createGridMgr(rangeType *scene.VisionLevelEnum) error {
	gridWidth := rangeType.GridWidth
	gridMgr := NewGridMgr(s.Scene, rangeType)
	// 把距离大小转化成格子单位
	length := (s.SceneSize.X + gridWidth - 1) / gridWidth
	if length <= 0 || length > UINT16_MAX {
		return errors.New("invalid length: " + string(rune(length)))
	}
	width := (s.SceneSize.Y + gridWidth - 1) / gridWidth
	if width <= 0 || width > UINT16_MAX {
		return errors.New("invalid width: " + string(rune(width)))
	}
	err := gridMgr.Init(int(length), int(width))
	if err != nil {
		return err
	}
	s.gridMgrs[rangeType.VisionLevelType] = gridMgr
	return nil
}

// getGridMgr 获取网格管理器
func (s *SceneGridSightModule) getGridMgr(rangeType *scene.VisionLevelEnum) *GridMgr {
	if int(rangeType.VisionLevelType) >= len(s.gridMgrs) {
		panic("invalid rangeType: " + rangeType.String())
	}
	return s.gridMgrs[rangeType.VisionLevelType]
}

// PlaceEntity entity进入场景
func (s *SceneGridSightModule) PlaceEntity(act scene.IEntity) []scene.IEntity {
	// 如果场景为null，或者非本场景，则Actor的场景管理有问题
	if act.GetScene() == nil || act.GetScene() != s.Scene {
		panic("actors scene is null or not same scene, actors = " + act.String())
	}

	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		panic("gridMgr is null, rangeType:" + rangeType.String())
	}

	// TODO: region对象需要特殊处理
	// if region, ok := act.(*abstracts.Region); ok {
	// 	coordSet := region.GetCoveredCoordinates(s)
	// 	gridMgr.PlaceRegionEntity(region, coordSet)
	// 	return nil
	// }

	pos := act.GetLocation().GetPosition()
	gridX := s.PosToGridX(rangeType, pos.X)
	gridY := s.PosToGridY(rangeType, pos.Y)
	gridMgr.PlaceEntity(act, int(gridX), int(gridY))

	// 如果场景当前没有玩家，可以略过视野计算
	// if s.Scene.GetPlayerCount() == 0 {
	// 	return nil
	// }

	// TODO: 计算附近的entity
	// visitVisitor := s.createAroundMeVisitor(act)
	// s.VisitGridsInSight(act)
	// return visitVisitor.GetResultList()
	return nil
}

// RemoveEntity entity离开场景
func (s *SceneGridSightModule) RemoveEntity(act scene.IEntity) {
	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		panic("gridMgr is null, rangeType:" + rangeType.String())
	}

	// TODO: region对象需要特殊处理
	// if region, ok := act.(*abstracts.Region); ok {
	// 	coordSet := region.GetCoveredCoordinates(s)
	// 	gridMgr.RemoveRegionEntity(region, coordSet)
	// 	return
	// }

	// 从格子中删除
	gridMgr.RemoveEntity(act)
}

// EntityMoveTo entity在场景中移动
func (s *SceneGridSightModule) EntityMoveTo(act scene.IEntity, destPos *math.Vector3) *EntityMoveToRet {
	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		panic("gridMgr is null, rangeType:" + rangeType.String())
	}

	// TODO: 实现移动逻辑
	// oldGrid := act.GetGrid()
	// gridX := s.PosToGridX(rangeType, destPos.GetX())
	// gridY := s.PosToGridY(rangeType, destPos.GetY())
	// isMoveGrid := gridMgr.EntityMoveTo(act, int(gridX), int(gridY))
	// ...

	return nil
}

// EntityChangePhasing 实体切换相位，公共视野的实体和相位中的实体，不能互相切换
func (s *SceneGridSightModule) EntityChangePhasing(act scene.IEntity, newPhasingId int64) *EntityMoveToRet {
	// TODO: 实现相位切换逻辑
	return nil
}

// VisitGridsInSight 查询所有视距级别，entity附近满足visitor条件的entity集合
func (s *SceneGridSightModule) VisitGridsInSight(act scene.IEntity) {
	// TODO: 实现视野查询逻辑
	// visitType := s.getVisitorType(act)
	// switch visitType {
	// case visitor.VisitAvatarVisitor, visitor.VisitExcludeSelfAvatarVisitor:
	// 	// 查找周围的角色实体，用直接的方式
	// 	s.visitPhasingPlayer(act, func(player *player.Player) {
	// 		s.visitAvatarDirectly(player, act)
	// 	})
	// 	return
	// ...
	// }
}

// getVisitorType 获取访问者类型
func (s *SceneGridSightModule) getVisitorType(act scene.IEntity) interface{} {
	// TODO: 根据actor类型返回对应的visitor类型
	return nil
}

// IsEntityMoveGrid entity是否移动格子
func (s *SceneGridSightModule) IsEntityMoveGrid(act scene.IEntity, pos *math.Vector3) bool {
	position := act.GetLocation().GetPosition()
	for _, gridMgr := range s.gridMgrs {
		// 有任意一层视野的格子移动，就返回true
		if !s.IsSameGrid(gridMgr.GetRangeType(), position.X, position.Y, pos.X, pos.Y) {
			return true
		}
	}
	return false
}

// FindPossibleRegionSet 寻找pos关联的region列表
func (s *SceneGridSightModule) FindPossibleRegionSet(pos *math.Vector3) []scene.IEntity {
	// TODO: 实现区域查找
	return nil
}

// visitPhasingPlayer 访问对应相位信息的Player
func (s *SceneGridSightModule) visitPhasingPlayer(act scene.IEntity, consumer func(scene.IGamePlayer)) {
	// TODO: 实现相位玩家访问
	if act.GetPhasingId() == 0 {
		// 公共实体访问场景内所有玩家
		// s.Scene.ForeachAllPlayer(consumer)
	} else {
		// 相位实体访问所在相位玩家
		// s.Scene.ForeachPhasingPlayer(act.GetPhasingId(), consumer)

		// 如果实体是在专属相位中,判断相位所属玩家实体是否在专属相位中，不在则需要额外访问所属玩家
		if act.GetPhasingId() < 0 {
			return
		}
		// playerViewMgr := s.Scene.GetPlayerViewMgrMap()[act.GetPhasingId()]
		// if playerViewMgr == nil || playerViewMgr.GetPlayer().GetAvatarTeam().InPrivatePhasing() {
		// 	return
		// }
		// consumer(playerViewMgr.GetPlayer())
	}
}

// isInSightRange 是否在当前层的视野距离内
func (s *SceneGridSightModule) isInSightRange(rangeType *scene.VisionLevelEnum, gridX, gridY, farGridX, farGridY int32) bool {
	sightRadius := rangeType.SightRadius
	return abs(gridX-farGridX) <= sightRadius && abs(gridY-farGridY) <= sightRadius
}

// abs 绝对值
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// visitDiffGrids 查询两个position entity集合的差集：减少以及新增
func (s *SceneGridSightModule) visitDiffGrids(fromPos, toPos *math.Vector3, missVisitor, meetVisitor interface{}) {
	for _, gridMgr := range s.gridMgrs {
		rangeType := gridMgr.GetRangeType()
		fromGridX := s.PosToGridX(rangeType, fromPos.X)
		fromGridY := s.PosToGridY(rangeType, fromPos.Y)
		toGridX := s.PosToGridX(rangeType, toPos.X)
		toGridY := s.PosToGridY(rangeType, toPos.Y)

		if fromGridX == toGridX && fromGridY == toGridY {
			continue
		}

		// TODO: gridMgr.VisitDiffGrids(int(fromGridX), int(fromGridY), int(toGridX), int(toGridY), missVisitor, meetVisitor)
	}
}
