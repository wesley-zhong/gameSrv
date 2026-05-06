package scene

import (
	"gameSrv/pkg/math"
	"log"
)

// SceneRoomSightModule 以房间作为视野，只能挂在房间场景上
type SceneRoomSightModule struct {
	*SceneSightModule
	grid *Grid
}

// NewSceneRoomSightModule 创建新的房间视野模块
func NewSceneRoomSightModule(scn IScene) *SceneRoomSightModule {
	return &SceneRoomSightModule{
		SceneSightModule: NewSceneSightModule(scn),
		grid:             NewGrid(0, 0),
	}
}

// PlaceEntity entity进入场景
func (s *SceneRoomSightModule) PlaceEntity(act IEntity) []IEntity {
	// 如果场景为null,或者非本场景,则Actor的场景管理有问题
	if act.GetScene() == nil || act.GetScene() != s.GetScene() {
		log.Printf("actor scene is null or not same scene, actorScene = %s actorType = %d this.Scene = %s",
			act.GetScene().String(), act.GetActorType(), s.GetScene().String())
	}

	// region对象单独维护
	if act.GetActorType() == 0 { // TODO: 使用 cfg.ActorType_EActorType_Region
		s.grid.AddRegion(act)
		act.SetGrid(s.grid)

		// 添加范围内的avatarTeam到region内
		for _, playerViewMgr := range s.GetScene().GetPlayerViewMgrMap() {
			teamActor := playerViewMgr.GetPlayer().GetAvatarTeam()
			if teamActor == nil {
				continue
			}
			// 检查AvatarTeam是否在区域内
			if regionEntity, ok := act.(interface {
				IsInRegion(*math.Vector3) bool
				AddEntity(IEntity, bool)
			}); ok {
				if regionEntity.IsInRegion(teamActor.GetLocation().GetPosition()) {
					regionEntity.AddEntity(teamActor, true)
				}
			}
		}
		return nil
	}

	s.grid.AddEntity(act)
	// act.SetGrid(s.grid)

	// 如果场景当前没有玩家,可以略过视野计算
	if len(s.GetScene().GetPlayerViewMgrMap()) == 0 {
		return nil
	}

	// TODO: 如果是AvatarTeamActor,则返回房间里所有可见的实体
	// if _, ok := act.(*AvatarTeamActor); ok {
	// 	return s.getRoomEntityList(act)
	// }

	// 其他实体直接返回房间里的玩家Avatar
	return s.getRoomAvatarList()
}

// RemoveEntity entity离开场景
func (s *SceneRoomSightModule) RemoveEntity(act IEntity) {
	// region对象单独维护
	if act.GetActorType() == 0 { // TODO: 使用 cfg.ActorType_EActorType_Region
		s.grid.DelRegion(act)
		act.SetGrid(nil)
		return
	}

	s.grid.DelEntity(act)
}

// EntityMoveTo 以整个房间作为视野范围时，不会因移动导致视野变更
func (s *SceneRoomSightModule) EntityMoveTo(act IEntity, destPos *math.Vector3) *EntityMoveToRet {
	return nil
}

// EntityChangePhasing 切换相位
func (s *SceneRoomSightModule) EntityChangePhasing(act IEntity, newPhasingId int64) *EntityMoveToRet {
	return nil
}

// VisitGridsInSight 查询所有视距级别，entity附近满足visitor条件的entity集合
func (s *SceneRoomSightModule) VisitGridsInSight(act IEntity) {
	visitVisitor := NewVisitor(act)
	s.grid.Accept(visitVisitor)
}

// IsEntityMoveGrid entity是否移动格子
func (s *SceneRoomSightModule) IsEntityMoveGrid(act IEntity, prevPos *math.Vector3) bool {
	return false
}

// FindPossibleRegionSet 查找pos关联的region列表
func (s *SceneRoomSightModule) FindPossibleRegionSet(pos *math.Vector3) []IEntity {
	regions := s.grid.GetAllRegion()
	if regions == nil {
		return nil
	}

	result := make([]IEntity, 0)
	for _, region := range regions {
		if entity, ok := region.(IEntity); ok {
			result = append(result, entity)
		}
	}
	return result
}

// getRoomEntityList 获取房间里所有可见的实体
func (s *SceneRoomSightModule) getRoomEntityList(act IEntity) []IEntity {
	visitVisitor := NewVisitor(act)
	s.VisitGridsInSight(act)
	return visitVisitor.GetResultList()
}

// getRoomAvatarList 获取房间中玩家的实体
func (s *SceneRoomSightModule) getRoomAvatarList() []IEntity {
	result := make([]IEntity, 0)

	for _, playerViewMgr := range s.GetScene().GetPlayerViewMgrMap() {
		result = append(result, playerViewMgr.GetPlayer().GetAvatarTeam())
	}
	return result
}
