// Copyright Epic Games, Inc. All Rights Reserved.

package sight

import (
	"errors"
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

var (
	// ErrUnknownValue 未知值错误
	ErrUnknownValue = errors.New("unknown value")
)

// EntityMoveToRet 实体移动返回结果
type EntityMoveToRet struct {
	MissEntities []scene.IEntity
	MeetEntities []scene.IEntity
}

// NewEntityMoveToRet 创建新的实体移动返回结果
func NewEntityMoveToRet(missEntities, meetEntities []scene.IEntity) *EntityMoveToRet {
	return &EntityMoveToRet{
		MissEntities: missEntities,
		MeetEntities: meetEntities,
	}
}

// SceneSightModule 场景视距管理器基类
type SceneSightModule struct {
	BeginPos  *math.Vector2 // 场景的开始坐标，可能是负坐标，在计算实际坐标时，要减去起始点位置，换算成正数的格子坐标
	EndPos    *math.Vector2 // 场景的结束坐标，场景内所有有效坐标应小于等于此坐标，否则计算出的格子不存在
	SceneSize *math.Vector2 // 场景实际大小
	Scene     scene.IScene  // 所属场景
}

// NewSceneSightModule 创建新的场景视距管理器
func NewSceneSightModule(scn scene.IScene) *SceneSightModule {
	return &SceneSightModule{
		Scene: scn,
	}
}

// Init 初始化组件
// beginPos sceneSize 里面存储的都是实际距离
func (s *SceneSightModule) Init(beginPos, sceneSize *math.Vector2) error {
	s.BeginPos = beginPos
	s.SceneSize = sceneSize
	// 结束坐标需减1，比如场景大小10，起始坐标0，结束坐标应是9
	s.EndPos = math.NewVector2(
		beginPos.X+sceneSize.X-1,
		beginPos.Y+sceneSize.Y-1,
	)
	if sceneSize.X == 0 || sceneSize.Y == 0 {
		return ErrUnknownValue
	}
	return nil
}

// PlaceEntity entity进入场景 (抽象方法，子类实现)
func (s *SceneSightModule) PlaceEntity(act scene.IEntity) []scene.IEntity {
	return nil
}

// RemoveEntity entity离开场景 (抽象方法，子类实现)
func (s *SceneSightModule) RemoveEntity(act scene.IEntity) {
}

// EntityMoveTo entity在场景中移动 (抽象方法，子类实现)
func (s *SceneSightModule) EntityMoveTo(act scene.IEntity, destPos *math.Vector3) *EntityMoveToRet {
	return nil
}

// EntityChangePhasing entity在场景中切换位面 (抽象方法，子类实现)
func (s *SceneSightModule) EntityChangePhasing(act scene.IEntity, newPhasingId int64) *EntityMoveToRet {
	return nil
}

// VisitGridsInSight 查询所有视距级别，entity附近满足visitor条件的entity集合 (抽象方法，子类实现)
func (s *SceneSightModule) VisitGridsInSight(act scene.IEntity) {
}

// IsEntityMoveGrid entity是否移动格子 (抽象方法，子类实现)
func (s *SceneSightModule) IsEntityMoveGrid(act scene.IEntity, pos *math.Vector3) bool {
	return false
}

// FindPossibleRegionSet 寻找pos关联的region列表 (抽象方法，子类实现)
func (s *SceneSightModule) FindPossibleRegionSet(pos *math.Vector3) []scene.IEntity {
	return nil
}

// PosToCoordinate 获取坐标所在视野层的格子位置
func (s *SceneSightModule) PosToCoordinate(rangeType *scene.VisionLevelEnum, posX, posY int32) *scene.Coordinate {
	gridX := (posX - s.BeginPos.X) / rangeType.GridWidth // 坐标要减去起始点的位置
	gridY := (posY - s.BeginPos.Y) / rangeType.GridWidth // 坐标要减去起始点的位置
	return scene.NewCoordinate(gridX, gridY)
}

// PosToGridX 获取坐标所在视野层的X格位置
func (s *SceneSightModule) PosToGridX(rangeType *scene.VisionLevelEnum, posX int32) int32 {
	return (posX - s.BeginPos.X) / rangeType.GridWidth // 坐标要减去起始点的位置
}

// PosToGridY 获取坐标所在视野层的Y格位置
func (s *SceneSightModule) PosToGridY(rangeType *scene.VisionLevelEnum, posY int32) int32 {
	return (posY - s.BeginPos.Y) / rangeType.GridWidth // 坐标要减去起始点的位置
}

// IsSameGrid 判断两个坐标在指定视野层上所在格子是否相等
func (s *SceneSightModule) IsSameGrid(rangeType *scene.VisionLevelEnum, posX1, posY1, posX2, posY2 int32) bool {
	gridWidth := rangeType.GridWidth

	gridX1 := (posX1 - s.BeginPos.X) / gridWidth // 坐标要减去起始点的位置
	gridY1 := (posY1 - s.BeginPos.Y) / gridWidth // 坐标要减去起始点的位置

	gridX2 := (posX2 - s.BeginPos.X) / gridWidth // 坐标要减去起始点的位置
	gridY2 := (posY2 - s.BeginPos.Y) / gridWidth // 坐标要减去起始点的位置

	return gridX1 == gridX2 && gridY1 == gridY2
}

// CoordinateToPos 计算视野格的左上角的起始坐标
func (s *SceneSightModule) CoordinateToPos(rangeType *scene.VisionLevelEnum, gridX, gridY int32) *math.Vector2 {
	// 获取视野格的宽度
	gridWidth := rangeType.GridWidth

	// 计算左上角位置 (格子坐标 * gridWidth) + 起始点位置
	x := gridX*gridWidth + s.BeginPos.X
	y := gridY*gridWidth + s.BeginPos.Y

	// 返回计算得到的实际坐标
	return math.NewVector2(x, y)
}

// InSceneMap 判断坐标是否在地图内
func (s *SceneSightModule) InSceneMap(posX, posY int32) bool {
	return posX >= s.BeginPos.X && posX <= s.EndPos.X &&
		posY >= s.BeginPos.Y && posY <= s.EndPos.Y
}

// GetScene 获取场景
func (s *SceneSightModule) GetScene() scene.IScene {
	return s.Scene
}

// GetBeginPos 获取开始位置
func (s *SceneSightModule) GetBeginPos() *math.Vector2 {
	return s.BeginPos
}

// GetEndPos 获取结束位置
func (s *SceneSightModule) GetEndPos() *math.Vector2 {
	return s.EndPos
}

// GetSceneSize 获取场景大小
func (s *SceneSightModule) GetSceneSize() *math.Vector2 {
	return s.SceneSize
}

// String 返回字符串表示
func (s *SceneSightModule) String() string {
	return "SceneSightModule{" +
		"sceneCfgId" + s.Scene.String() +
		", beginPos=" + s.BeginPos.String() +
		", endPos=" + s.EndPos.String() +
		", sceneSize=" + s.SceneSize.String() +
		"}"
}
