package scene

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/math"
	"log"
)

const UINT16_MAX = 65535

// SceneGridSightModule 网格场景视野模块
type SceneGridSightModule struct {
	*SceneSightModule
	gridMgrs []*GridMgr
}

// NewSceneGridSightModule 创建新的网格场景视野模块
func NewSceneGridSightModule(scn IScene) *SceneGridSightModule {
	gridMgrs := make([]*GridMgr, len(VisionLevelEnums))
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
		rangeType := ForVisionLevelType(VisionLevelType(i))
		if rangeType == nil {
			log.Printf("VisionLevelEnum is null, i = %d", i)
			continue
		}
		// 初始化各级管理器
		s.createGridMgr(rangeType)
	}
	return nil
}

// createGridMgr 创建网格管理器
func (s *SceneGridSightModule) createGridMgr(rangeType *VisionLevelEnum) {
	gridWidth := rangeType.GridWidth
	gridMgr := NewGridMgr(s.GetScene(), rangeType)
	// 把距离大小转化成格子单位
	length := (s.GetSceneSize().X + gridWidth - 1) / gridWidth
	if length <= 0 || length > UINT16_MAX {
		log.Printf("invalid length: %d", length)
	}
	width := (s.GetSceneSize().Y + gridWidth - 1) / gridWidth
	if width <= 0 || width > UINT16_MAX {
		log.Printf("invalid width: %d", width)
	}
	gridMgr.Init(int(length), int(width))
	s.gridMgrs[rangeType.VisionLevelType] = gridMgr

	log.Printf("GridMgr::init length: %d width: %d vision_level: %v", length, width, rangeType)
}

// getGridMgr 获取网格管理器
func (s *SceneGridSightModule) getGridMgr(rangeType *VisionLevelEnum) *GridMgr {
	if int(rangeType.VisionLevelType) >= len(s.gridMgrs) {
		log.Printf("invalid rangeType: %s", rangeType.String())
		return nil
	}
	return s.gridMgrs[rangeType.VisionLevelType]
}

// PlaceEntity entity进入场景
func (s *SceneGridSightModule) PlaceEntity(act IEntity) []IEntity {
	// 如果场景为null,或者非本场景,则Actor的场景管理有问题
	if act.GetScene() == nil || act.GetScene() != s.GetScene() {
		log.Printf("actor scene is null or not same scene, actorScene = %s actorType = %d this.Scene = %s",
			act.GetScene().String(), act.GetActorType(), s.GetScene().String())
	}

	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		log.Printf("gridMgr is null, rangeType:%s", rangeType.String())
		return nil
	}

	// region对象需要特殊处理
	if act.GetActorType() == cfg.ActorType_EActorType_Region {
		// 检查是否是Region类型
		if region, ok := act.(interface {
			GetCoveredCoordinates(ISceneSightModule) []*Coordinate
		}); ok {
			coordSet := region.GetCoveredCoordinates(s)
			gridMgr.PlaceRegionEntity(act, coordSet)
			return nil
		}
		// 如果无法获取覆盖坐标，暂且使用普通实体的处理方式
	}

	pos := act.GetLocation().GetPosition()
	gridX := s.PosToGridX(rangeType, pos.X)
	gridY := s.PosToGridY(rangeType, pos.Y)
	gridMgr.PlaceEntity(act, int(gridX), int(gridY))

	// 如果场景当前没有玩家,可以略过视野计算
	if len(s.GetScene().GetPlayerViewMgrMap()) == 0 {
		return nil
	}

	// 计算附近的entity
	visitVisitor := s.createAroundMeVisitor(act)
	s.VisitGridsInSight(act, visitVisitor)
	return visitVisitor.GetResultList()
}

// createAroundMeVisitor 创建周围的访问者
func (s *SceneGridSightModule) createAroundMeVisitor(act IEntity) IVisitor {
	// 玩家控制角色，能看到所有
	if act.GetActorType() == cfg.ActorType_EActorType_Team {
		return NewVisitEntityVisitor(act)
	}
	// 其他实体只关注玩家控制角色
	return NewVisitAvatarVisitor(act)
}

// RemoveEntity entity离开场景
func (s *SceneGridSightModule) RemoveEntity(act IEntity) {
	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		log.Printf("gridMgr is null, rangeType:%s", rangeType.String())
		return
	}

	// region对象需要特殊处理
	if act.GetActorType() == cfg.ActorType_EActorType_Region {
		// 检查是否是Region类型
		if region, ok := act.(interface {
			GetCoveredCoordinates(ISceneSightModule) []*Coordinate
		}); ok {
			coordSet := region.GetCoveredCoordinates(s)
			gridMgr.RemoveRegionEntity(act, coordSet)
			return
		}
		// 如果无法获取覆盖坐标，暂且使用普通实体的处理方式
	}

	// 从格子中删除
	gridMgr.RemoveEntity(act)
}

// EntityMoveTo entity在场景中移动
func (s *SceneGridSightModule) EntityMoveTo(act IEntity, destPos *math.Vector3) *EntityMoveToRet {
	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		log.Printf("gridMgr is null, rangeType:%s", rangeType.String())
		return nil
	}

	oldGrid := act.GetGrid()
	gridX := s.PosToGridX(rangeType, destPos.X)
	gridY := s.PosToGridY(rangeType, destPos.Y)
	isMoveGrid := gridMgr.EntityMoveTo(act, int(gridX), int(gridY))

	var missVisitor, meetVisitor IVisitor

	if act.GetActorType() == cfg.ActorType_EActorType_Team {
		// 如果所有分层视野中都没有格子变更,则表示视野无需更新,直接return
		if !s.IsEntityMoveGrid(act, destPos) {
			return nil
		}
		// 玩家控制角色，能看到所有
		missVisitor = NewVisitEntityVisitor(act)
		meetVisitor = NewVisitEntityVisitor(act)
		// 查询视野差集
		s.visitDiffGrids(act.GetLocation().GetPosition(), destPos, missVisitor, meetVisitor)
	} else {
		// 如果未移动格子,表示视野无需更新,直接return
		if isMoveGrid {
			return nil
		}
		// 其他实体只关注玩家控制角色
		missVisitor = NewVisitAvatarVisitor(act)
		meetVisitor = NewVisitAvatarVisitor(act)
		// 根据actor的相位信息计算对应相位玩家控制角色的视野
		s.visitPhasingPlayer(act, func(plr IGamePlayer) {
			s.otherEntityMove(plr, rangeType, missVisitor, meetVisitor, oldGrid, act.GetGrid())
		})
	}

	_ = isMoveGrid // 避免未使用变量警告
	return NewEntityMoveToRet(missVisitor.GetResultList(), meetVisitor.GetResultList())
}

// EntityChangePhasing 实体切换相位，公共视野的实体和相位中的实体，不能互相切换
func (s *SceneGridSightModule) EntityChangePhasing(act IEntity, newPhasingId int64) *EntityMoveToRet {
	if act.GetPhasingId() == newPhasingId {
		return nil
	}

	rangeType := act.GetVisionLevelEnum()
	gridMgr := s.getGridMgr(rangeType)
	if gridMgr == nil {
		log.Printf("gridMgr is null, rangeType:%s", rangeType.String())
		return nil
	}

	// TODO 优化方向,Grid中新增ChangePhasing的方法，根据相位精细化访问actor,不需要双向过滤
	// 先从旧相位中离开
	s.RemoveEntity(act)
	// 进入新相位并计算新相位的视野
	// act.SetPhasingId(newPhasingId)
	meetEntities := s.PlaceEntity(act)

	var missEntities []IEntity

	// 玩家控制实体和其他实体分开处理
	if act.GetActorType() == cfg.ActorType_EActorType_Team {
		playerViewMgr := s.GetScene().FindPlayerViewMgr(act.GetOwnerPlayerUid())
		if playerViewMgr != nil {
			// 先暂时将自身从视野管理中移除
			playerViewMgr.DelEntityInView(act)
			// 创建一份当前视野的实体set，用于过滤
			currActorSet := make(map[IEntity]struct{})
			entitiesInView := playerViewMgr.GetEntitiesInView()
			for _, entity := range entitiesInView {
				if entity, ok := entity.(IEntity); ok {
					currActorSet[entity] = struct{}{}
				}
			}
			// 恢复自身加入视野管理中
			playerViewMgr.AddEntityInView(act)
			// 双向过滤meet和当前视野中同时存在的实体，则meet中剩下的就是真正出现的实体，set剩下的就是真正消失的实体
			filteredMeet := make([]IEntity, 0)
			for _, entity := range meetEntities {
				if _, exists := currActorSet[entity]; exists {
					delete(currActorSet, entity)
				} else {
					filteredMeet = append(filteredMeet, entity)
				}
			}
			meetEntities = filteredMeet
			// set剩下的就是真正消失的实体
			missEntities = make([]IEntity, 0, len(currActorSet))
			for entity := range currActorSet {
				missEntities = append(missEntities, entity)
			}
		}
	} else {
		// TODO: 需要实现act.GetAoiViewMgr().GetViewingPlayers()方法
		// missEntities = make([]IEntity, 0)
		// for _, viewPlayer := range act.GetAoiViewMgr().GetViewingPlayers() {
		// 	// 如果其他实体的新相位，是玩家实体的专属相位，则过滤该玩家实体
		// 	if viewPlayer.GetAvatarTeam().GetPrivatePhasingId() == newPhasingId {
		// 		continue
		// 	}
		// 	missEntities = append(missEntities, viewPlayer.GetAvatarTeam())
		// }
	}

	return NewEntityMoveToRet(missEntities, meetEntities)
}

// VisitGridsInSight 查询所有视距级别，entity附近满足visitor条件的entity集合
func (s *SceneGridSightModule) VisitGridsInSight(act IEntity, visitor IVisitor) {
	visitType := visitor.GetType()

	switch visitType {
	case VisitorTypeEntityVisitor, VisitorTypeExcludeSelfAvatarVisitor:
		// 非主角或eye_point的实体，只需要查找周围的player_eye_entity即可
		if act.GetActorType() != cfg.ActorType_EActorType_Avatar &&
			act.GetActorType() != cfg.ActorType_EActorType_EYE_POINT {
			s.visitPlayerEyeEntityDirectly(act)
			return
		}
		// 主角或eye_point的实体，走通用的流程
		position := act.GetLocation().GetPosition()
		for _, gridMgr := range s.gridMgrs {
			if gridMgr == nil {
				continue
			}
			gridX := s.PosToGridX(gridMgr.GetRangeType(), position.X)
			gridY := s.PosToGridY(gridMgr.GetRangeType(), position.Y)
			gridMgr.VisitGridsInSight(int(gridX), int(gridY), visitor)
		}

	case VisitorTypeAvatarVisitor:
		// 查找周围的角色实体，用直接的方式
		s.visitPhasingPlayer(act, func(plr IGamePlayer) {
			s.visitAvatarDirectly(plr, act)
		})

	default:
		// 剩余的情况，走通用的流程
		position := act.GetLocation().GetPosition()
		for _, gridMgr := range s.gridMgrs {
			if gridMgr == nil {
				continue
			}
			gridX := s.PosToGridX(gridMgr.GetRangeType(), position.X)
			gridY := s.PosToGridY(gridMgr.GetRangeType(), position.Y)
			gridMgr.VisitGridsInSight(int(gridX), int(gridY), visitor)
		}
	}
}

// IsEntityMoveGrid entity是否移动格子
func (s *SceneGridSightModule) IsEntityMoveGrid(act IEntity, pos *math.Vector3) bool {
	position := act.GetLocation().GetPosition()
	for _, gridMgr := range s.gridMgrs {
		if gridMgr == nil {
			continue
		}
		// 有任意一层视野的格子移动,就返回true
		if !s.IsSameGrid(gridMgr.GetRangeType(), position.X, position.Y, pos.X, pos.Y) {
			return true
		}
	}
	return false
}

// FindPossibleRegionSet 查找pos关联的region列表
func (s *SceneGridSightModule) FindPossibleRegionSet(pos *math.Vector3) []IEntity {
	regionSet := make(map[IEntity]struct{})
	for _, gridMgr := range s.gridMgrs {
		if gridMgr == nil {
			continue
		}
		gridX := s.PosToGridX(gridMgr.GetRangeType(), pos.X)
		gridY := s.PosToGridY(gridMgr.GetRangeType(), pos.Y)
		grid := gridMgr.GetGrid(int(gridX), int(gridY))
		if grid == nil {
			log.Printf("findPossibleRegionSet getGrid doesn't exist! gridX = %d, gridY=%d, grid_mgr = %v, sceneInfo = %v",
				gridX, gridY, gridMgr, s)
			continue
		}
		regions := grid.GetAllRegion()
		for _, region := range regions {
			if entity, ok := region.(IEntity); ok {
				regionSet[entity] = struct{}{}
			}
		}
	}

	result := make([]IEntity, 0, len(regionSet))
	for region := range regionSet {
		result = append(result, region)
	}
	return result
}

// visitPhasingPlayer 访问对应相位信息的Player,执行consumer
func (s *SceneGridSightModule) visitPhasingPlayer(act IEntity, consumer func(IGamePlayer)) {
	if act.GetPhasingId() == 0 {
		// 公共实体访问场景内所有玩家
		s.GetScene().ForeachAllPlayer(consumer)
	} else {
		// 相位实体访问所在相位玩家
		s.GetScene().ForeachPhasingPlayer(act.GetPhasingId(), consumer)
		// 如果实体是在专属相位中,判断相位所属玩家实体是否在专属相位中，不在则需要额外访问所属玩家
		if act.GetPhasingId() < 0 {
			return
		}
		playerViewMgr := s.GetScene().FindPlayerViewMgr(act.GetPhasingId())
		if playerViewMgr == nil || playerViewMgr.GetPlayer().GetAvatarTeam() == nil {
			return
		}
		if playerViewMgr.GetPlayer().InPrivatePhasing() {
			return
		}
		consumer(playerViewMgr.GetPlayer())
	}
}

// visitPlayerEyeEntityDirectly 直接查找周围的player_eye_entity
func (s *SceneGridSightModule) visitPlayerEyeEntityDirectly(act IEntity) {
	// 非主角或eye_point的实体，只需要查找周围的player_eye_entity即可
	// 通过遍历场景中的玩家，查找其avatarTeam是否在实体的视野范围内
	s.visitPhasingPlayer(act, func(plr IGamePlayer) {
		avatarTeam := plr.GetAvatarTeam()
		if avatarTeam == nil {
			return
		}

		rangeType := act.GetVisionLevelEnum()
		if rangeType == nil {
			return
		}

		pos := avatarTeam.GetLocation().GetPosition()
		avatarGridX := s.PosToGridX(rangeType, pos.X)
		avatarGridY := s.PosToGridY(rangeType, pos.Y)

		oldGrid := act.GetGrid()
		if oldGrid == nil {
			return
		}

		if s.isInSightRange(rangeType, oldGrid.GetGridX(), oldGrid.GetGridY(), avatarGridX, avatarGridY) {
			// avatarTeam在视野范围内，需要通知（这里TODO需要后续实现）
			_ = act
			_ = avatarTeam
		}
	})
}

// visitAvatarDirectly 直接访问AvatarTeam(不通过格子)
func (s *SceneGridSightModule) visitAvatarDirectly(player IGamePlayer, act IEntity) {
	// 如果entity处于更高的视野，直接通过视野距离获取avatar
	rangeType := act.GetVisionLevelEnum()
	if rangeType == nil {
		return
	}
	avatarTeam := player.GetAvatarTeam()
	if avatarTeam == nil {
		return
	}
	pos := avatarTeam.GetLocation().GetPosition()
	avatarGridX := s.PosToGridX(rangeType, pos.X)
	avatarGridY := s.PosToGridY(rangeType, pos.Y)

	oldGrid := act.GetGrid()
	if oldGrid == nil {
		return
	}

	if s.isInSightRange(rangeType, oldGrid.GetGridX(), oldGrid.GetGridY(), avatarGridX, avatarGridY) {
		// TODO: s.notifyAvatarInView(act, avatarTeam)
		_ = act
		_ = avatarTeam
	}
}

// otherEntityMove 其他实体移动时计算视野差集
func (s *SceneGridSightModule) otherEntityMove(player IGamePlayer, rangeType *VisionLevelEnum, missVisitor, meetVisitor IVisitor, oldGrid, newGrid *Grid) {
	if oldGrid == nil || newGrid == nil || rangeType == nil {
		return
	}

	// 获取avatarTeam的位置
	avatarTeam := player.GetAvatarTeam()
	if avatarTeam == nil {
		return
	}
	pos := avatarTeam.GetLocation().GetPosition()
	avatarGridX := s.PosToGridX(rangeType, pos.X)
	avatarGridY := s.PosToGridY(rangeType, pos.Y)

	// 计算avatarTeam是否在旧格子视野中
	oldInSight := s.isInSightRange(rangeType, oldGrid.GetGridX(), oldGrid.GetGridY(), avatarGridX, avatarGridY)
	// 计算avatarTeam是否在新格子视野中
	newInSight := s.isInSightRange(rangeType, newGrid.GetGridX(), newGrid.GetGridY(), avatarGridX, avatarGridY)

	if oldInSight && !newInSight {
		// avatarTeam从视野中消失
		missVisitor.VisitEntity(avatarTeam)
	} else if !oldInSight && newInSight {
		// avatarTeam进入视野
		meetVisitor.VisitEntity(avatarTeam)
	}
}

// isInSightRange 是否在当前层的视野距离内
func (s *SceneGridSightModule) isInSightRange(rangeType *VisionLevelEnum, gridX, gridY, farGridX, farGridY int32) bool {
	sightRadius := rangeType.SightRadius
	return absInt32(gridX-farGridX) <= sightRadius && absInt32(gridY-farGridY) <= sightRadius
}

// visitDiffGrids 查询两个position entity集合的差集：减少以及新增
func (s *SceneGridSightModule) visitDiffGrids(fromPos, toPos *math.Vector3, missVisitor, addVisitor IVisitor) {
	for _, gridMgr := range s.gridMgrs {
		if gridMgr == nil {
			continue
		}
		rangeType := gridMgr.GetRangeType()
		fromGridX := s.PosToGridX(rangeType, fromPos.X)
		fromGridY := s.PosToGridY(rangeType, fromPos.Y)
		toGridX := s.PosToGridX(rangeType, toPos.X)
		toGridY := s.PosToGridY(rangeType, toPos.Y)

		if fromGridX == toGridX && fromGridY == toGridY {
			continue
		}
		gridMgr.VisitDiffGrids(int(fromGridX), int(fromGridY), int(toGridX), int(toGridY), missVisitor, addVisitor)
	}
}

// absInt32 绝对值
func absInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// PosToGridX 将位置X转换为格子X
func (s *SceneGridSightModule) PosToGridX(rangeType *VisionLevelEnum, x int32) int32 {
	gridWidth := rangeType.GridWidth
	return x / gridWidth
}

// PosToGridY 将位置Y转换为格子Y
func (s *SceneGridSightModule) PosToGridY(rangeType *VisionLevelEnum, y int32) int32 {
	gridWidth := rangeType.GridWidth
	return y / gridWidth
}

// IsSameGrid 判断是否是同一个格子
func (s *SceneGridSightModule) IsSameGrid(rangeType *VisionLevelEnum, x1, y1, x2, y2 int32) bool {
	_ = rangeType // 避免未使用参数警告
	return s.PosToGridX(rangeType, x1) == s.PosToGridX(rangeType, x2) &&
		s.PosToGridY(rangeType, y1) == s.PosToGridY(rangeType, y2)
}
