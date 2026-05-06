package scene

// ScenePhasingData 场景相位数据
type ScenePhasingData struct {
	PhasingId                 int64                          // 相位id
	PlayerMap                 map[int64]IGamePlayer          // 在相位中的玩家 key是playerUid
	EntityMap                 map[int64]IEntity              // 在相位中的实体 key是entityId
	PreAssignedActors         map[int64]IEntity              // 编辑器内预分配了ID的实体列表 key是preAssignedGuid
	SameConfigActors          map[int64]map[IEntity]struct{} // 同ConfigId的Actor记录，key是ConfigID，Set容器内是非战斗单位
	BattleProsChangedEntities map[int64]struct{}             // 战斗过的怪物缓存
	ResetActorHandler         IResetActorHandler             // 重置处理器
	ActorDestroyTimeMap       map[int64]int64                // actor的销毁时间
	AutoVisualActors          map[int64]struct{}             // 可被自动发现的Actor列表
}

// NewScenePhasingData 创建新的ScenePhasingData
func NewScenePhasingData(phasingId int64) *ScenePhasingData {
	return &ScenePhasingData{
		PhasingId:                 phasingId,
		PlayerMap:                 make(map[int64]IGamePlayer),
		EntityMap:                 make(map[int64]IEntity),
		PreAssignedActors:         make(map[int64]IEntity),
		SameConfigActors:          make(map[int64]map[IEntity]struct{}),
		BattleProsChangedEntities: make(map[int64]struct{}),
		ResetActorHandler:         nil, // TODO: implement reset handler
		ActorDestroyTimeMap:       make(map[int64]int64),
	}
}

// EntityEnterPhasing 实体进入相位
func (s *ScenePhasingData) EntityEnterPhasing(act IEntity) error {
	// TODO: 如果是玩家的teamActor实体,额外维护相位中的玩家数据
	// if act.IsAvatarTeamActor() {
	// 	owner := act.GetOwner()
	// 	s.PlayerMap[owner.GetUid()] = owner
	// }

	s.EntityMap[act.GetEntityId()] = act

	// TODO: 实现PreAssignedGuid支持
	// if act.GetPreAssignedGuid() != 0 {
	// 	if _, exists := s.PreAssignedActors[act.GetPreAssignedGuid()]; exists {
	// 		return errors.New("PreAssignedGuid exist, Actor = " + act.String())
	// 	}
	// 	s.PreAssignedActors[act.GetPreAssignedGuid()] = act
	// }

	// TODO: 实现ConfigId支持
	// if act.GetConfigId() != 0 {
	// 	configId := act.GetConfigId()
	// 	if _, exists := s.SameConfigActors[configId]; !exists {
	// 		s.SameConfigActors[configId] = make(map[IEntity]struct{})
	// 	}
	// 	s.SameConfigActors[configId][act] = struct{}{}
	// }

	return nil
}

// EntityLeavePhasing 实体离开相位
func (s *ScenePhasingData) EntityLeavePhasing(act IEntity) {
	// TODO: 如果是玩家的teamActor实体,额外维护相位中的玩家数据
	// if act.IsAvatarTeamActor() {
	// 	delete(s.PlayerMap, act.GetOwnerPlayerUid())
	// }

	delete(s.EntityMap, act.GetEntityId())

	// TODO: 实现PreAssignedGuid支持
	// if act.GetPreAssignedGuid() != 0 {
	// 	delete(s.PreAssignedActors, act.GetPreAssignedGuid())
	// }

	// TODO: 实现ConfigId支持
	// if actors, exists := s.SameConfigActors[act.GetConfigId()]; exists {
	// 	delete(actors, act)
	// }
}

// GetSameConfigActor 获取相同ConfigId的Actor
func (s *ScenePhasingData) GetSameConfigActor(configID int64) []IEntity {
	actors, exists := s.SameConfigActors[configID]
	if !exists {
		return nil
	}
	result := make([]IEntity, 0, len(actors))
	for act := range actors {
		result = append(result, act)
	}
	return result
}

// AddAutoVisualActor 添加自动可视Actor
func (s *ScenePhasingData) AddAutoVisualActor(systemGuid int64) struct{} {
	if s.AutoVisualActors == nil {
		s.AutoVisualActors = make(map[int64]struct{})
	}
	s.AutoVisualActors[systemGuid] = struct{}{}
	return struct{}{}
}

// GetPlayerMap 获取玩家Map
func (s *ScenePhasingData) GetPlayerMap() map[int64]IGamePlayer {
	return s.PlayerMap
}

// GetEntityMap 获取实体Map
func (s *ScenePhasingData) GetEntityMap() map[int64]IEntity {
	return s.EntityMap
}

// GetPreAssignedActors 获取预分配Actors
func (s *ScenePhasingData) GetPreAssignedActors() map[int64]IEntity {
	return s.PreAssignedActors
}

// GetResetActorHandler 获取重置处理器
func (s *ScenePhasingData) GetResetActorHandler() IResetActorHandler {
	return s.ResetActorHandler
}

// GetActorDestroyTimeMap 获取Actor销毁时间Map
func (s *ScenePhasingData) GetActorDestroyTimeMap() map[int64]int64 {
	return s.ActorDestroyTimeMap
}

// GetAutoVisualActors 获取自动可视Actors
func (s *ScenePhasingData) GetAutoVisualActors() map[int64]struct{} {
	return s.AutoVisualActors
}

// GetPhasingId 获取相位ID
func (s *ScenePhasingData) GetPhasingId() int64 {
	return s.PhasingId
}

// GetDeadActorRecords 获取死亡Actor记录
func (s *ScenePhasingData) GetDeadActorRecords() map[int64]IEntity {
	// TODO: 实现死亡记录返回IEntity的map
	return nil
}

// IsRecordActorDead 检查是否记录了Actor死亡
func (s *ScenePhasingData) IsRecordActorDead(systemGuid int64) bool {
	if s.ResetActorHandler == nil {
		return false
	}
	return s.ResetActorHandler.IsRecordActorDead(systemGuid)
}

// AddDeadRecord 添加死亡记录
func (s *ScenePhasingData) AddDeadRecord(systemGuid int64) {
	// TODO: 实现死亡记录添加
}

// RemoveResetActor 移除重置Actor
func (s *ScenePhasingData) RemoveResetActor(phasingId int64, entityID int64) {
	// TODO: 实现移除重置Actor
}

// ClearDeadActorRecords 清除死亡Actor记录
func (s *ScenePhasingData) ClearDeadActorRecords() {
	// TODO: 实现清除死亡记录
}

// EntityChangePhasingData 实体改变相位数据
func (s *ScenePhasingData) EntityChangePhasingData(act IEntity, oldPhasingId int64, newPhasingId int64) {
	// TODO: 实现相位数据改变
}

// GetEntityChangePhasing 获取实体改变相位
func (s *ScenePhasingData) GetEntityChangePhasing(act IEntity, newPhasingId int64) {
	// TODO: 实现获取实体改变相位
}

// GetPreAssignedActor 获取预分配的Actor
func (s *ScenePhasingData) GetPreAssignedActor(systemGuid int64, preAssignedId int64) IEntity {
	return s.PreAssignedActors[preAssignedId]
}

// RemoveAutoVisualActor 移除自动可视Actor
func (s *ScenePhasingData) RemoveAutoVisualActor(systemGuid int64) bool {
	if s.AutoVisualActors == nil {
		return false
	}
	_, exists := s.AutoVisualActors[systemGuid]
	delete(s.AutoVisualActors, systemGuid)
	return exists
}

// FilterAlwaysActiveTags 过滤始终激活的标签
func (s *ScenePhasingData) FilterAlwaysActiveTags(player IGamePlayer) {
	// TODO: 实现过滤始终激活标签
}

// FilterConditionLayerTags 过滤条件层级标签
func (s *ScenePhasingData) FilterConditionLayerTags(player IGamePlayer) {
	// TODO: 实现过滤条件层级标签
}

// LoadActorFromLevelTag 从层级标签加载Actor
func (s *ScenePhasingData) LoadActorFromLevelTag(player IGamePlayer, layerTag string) bool {
	// TODO: 实现从层级标签加载Actor
	return false
}

// UnLoadActorWithTag 卸载带标签的Actor
func (s *ScenePhasingData) UnLoadActorWithTag(player IGamePlayer, layerTag string) bool {
	// TODO: 实现卸载带标签的Actor
	return false
}

// UnLoadLevelActors 卸载层级Actor
func (s *ScenePhasingData) UnLoadLevelActors(player IGamePlayer, levelActorCfgs interface{}) bool {
	// TODO: 实现卸载层级Actor
	return false
}

// LoadLevelInstanceActors 加载层级实例Actor
func (s *ScenePhasingData) LoadLevelInstanceActors(player IGamePlayer, layerTagActors interface{}) bool {
	// TODO: 实现加载层级实例Actor
	return false
}

// LoadActorFromDO 从数据对象加载Actor
func (s *ScenePhasingData) LoadActorFromDO(player IGamePlayer, sceneActorData interface{}) error {
	// TODO: 实现从数据对象加载Actor
	return nil
}

// CheckWaitLoadTags 检查等待加载标签
func (s *ScenePhasingData) CheckWaitLoadTags(player IGamePlayer) {
	// TODO: 实现检查等待加载标签
}

// InitActorFromDO 从数据对象初始化Actor
func (s *ScenePhasingData) InitActorFromDO(player IGamePlayer, sceneDataDO interface{}) error {
	// TODO: 实现从数据对象初始化Actor
	return nil
}

// FlushPlayerSceneDataDO 刷新玩家场景数据对象
func (s *ScenePhasingData) FlushPlayerSceneDataDO(player IGamePlayer, sceneDataDO interface{}) {
	// TODO: 实现刷新玩家场景数据对象
}

// GetMonsterLevel 获取怪物等级
func (s *ScenePhasingData) GetMonsterLevel(monster interface{}) int32 {
	// TODO: 实现获取怪物等级
	return 1
}

// GetSceneBuff 获取场景Buff
func (s *ScenePhasingData) GetSceneBuff() []int32 {
	// TODO: 实现获取场景Buff
	return []int32{}
}

// GetMonsterBuff 获取怪物Buff
func (s *ScenePhasingData) GetMonsterBuff() []int32 {
	// TODO: 实现获取怪物Buff
	return []int32{}
}

// GetSceneProps 获取场景属性
func (s *ScenePhasingData) GetSceneProps() map[int32]int32 {
	// TODO: 实现获取场景属性
	return map[int32]int32{}
}

// IsLoadFormationData 检查是否加载编队数据
func (s *ScenePhasingData) IsLoadFormationData() bool {
	// TODO: 实现检查是否加载编队数据
	return false
}

// IsSaveToFormationData 检查是否保存到编队数据
func (s *ScenePhasingData) IsSaveToFormationData() bool {
	// TODO: 实现检查是否保存到编队数据
	return false
}

// AllowOption 检查是否允许选项
func (s *ScenePhasingData) AllowOption(optionType int32) bool {
	// TODO: 实现检查是否允许选项
	return false
}

// GetNumber 获取数值
func (s *ScenePhasingData) GetNumber() int32 {
	// TODO: 实现获取数值
	return 0
}
