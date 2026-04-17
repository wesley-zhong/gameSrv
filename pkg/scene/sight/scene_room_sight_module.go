// Copyright Epic Games, Inc. All Rights Reserved.

package sight

import (
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"
)

// SceneRoomSightModule 以房间作为视野，只能挂在房间场景上
type SceneRoomSightModule struct {
	*SceneSightModule
	grid *Grid
}

// NewSceneRoomSightModule 创建新的房间视野模块
func NewSceneRoomSightModule(scn scene.IScene) *SceneRoomSightModule {
	return &SceneRoomSightModule{
		SceneSightModule: NewSceneSightModule(scn),
		grid:             NewGrid(0, 0),
	}
}

// PlaceEntity entity进入场景
func (s *SceneRoomSightModule) PlaceEntity(act scene.IEntity) []scene.IEntity {
	// 如果场景为null，或者非本场景，则Actor的场景管理有问题
	if act.GetScene() == nil || act.GetScene() != s.Scene {
		panic("actors scene is null or not same scene, actorId = " + string(rune(act.GetEntityId())) +
			" this.Scene = " + s.Scene.String())
	}

	// TODO: region对象单独维护
	// if region, ok := act.(scene.IRegion); ok {
	// 	s.grid.AddRegion(region)
	// 	region.SetGrid(s.grid)
	// 	// 添加范围内的avatarTeam到region内
	// 	for _, playerViewMgr := range s.Scene.GetPlayerViewMgrMap() {
	// 		teamActor := playerViewMgr.GetPlayer().GetAvatarTeam()
	// 		if region.IsInRegion(teamActor.GetLocation().GetPosition()) {
	// 			region.AddEntity(teamActor, true)
	// 		}
	// 	}
	// 	return nil
	// }

	s.grid.AddEntity(act)
	// TODO: act.SetGrid(s.grid)

	// 如果场景当前没有玩家，可以略过视野计算
	// if s.Scene.GetPlayerCount() == 0 {
	// 	return nil
	// }

	// TODO: 如果是AvatarTeamActor，则返回房间里所有可见的实体
	// if _, ok := act.(*actor.AvatarTeamActor); ok {
	// 	return s.getRoomEntityList(act)
	// }

	// 其他实体直接返回房间里的玩家Avatar
	return s.getRoomAvatarList()
}

// getRoomEntityList 获取房间里所有可见的实体
func (s *SceneRoomSightModule) getRoomEntityList(act scene.IEntity) []scene.IEntity {
	// TODO: visitor := visitor.NewVisitEntityVisitor(act)
	// s.VisitGridsInSight(act)
	// return visitVisitor.GetResultList()
	return nil
}

// getRoomAvatarList 获取房间中玩家的实体
func (s *SceneRoomSightModule) getRoomAvatarList() []scene.IEntity {
	result := make([]scene.IEntity, 0)

	for _, playerViewMgr := range s.Scene.GetPlayerViewMgrMap() {
		// TODO: result = append(result, playerViewMgr.GetPlayer().GetAvatarTeam())
		_ = playerViewMgr
	}
	return result
}

// RemoveEntity entity离开场景
func (s *SceneRoomSightModule) RemoveEntity(act scene.IEntity) {
	// TODO: region对象单独维护
	// if region, ok := act.(scene.IRegion); ok {
	// 	s.grid.DelRegion(region)
	// 	region.SetGrid(nil)
	// 	return
	// }

	s.grid.DelEntity(act)
}

// EntityMoveTo 以整个房间作为视野范围时，不会因移动导致视野变更
func (s *SceneRoomSightModule) EntityMoveTo(act scene.IEntity, destPos *math.Vector3) *EntityMoveToRet {
	return nil
}

// EntityChangePhasing 切换相位
func (s *SceneRoomSightModule) EntityChangePhasing(act scene.IEntity, newPhasingId int64) *EntityMoveToRet {
	return nil
}

// VisitGridsInSight 查询所有视距级别，entity附近满足visitor条件的entity集合
func (s *SceneRoomSightModule) VisitGridsInSight(act scene.IEntity) {
	s.grid.Accept(act)
}

// IsEntityMoveGrid entity是否移动格子
func (s *SceneRoomSightModule) IsEntityMoveGrid(act scene.IEntity, prevPos *math.Vector3) bool {
	return false
}

// FindPossibleRegionSet 寻找pos关联的region列表
func (s *SceneRoomSightModule) FindPossibleRegionSet(pos *math.Vector3) []scene.IEntity {
	// TODO: return s.grid.GetAllRegion()
	return nil
}
