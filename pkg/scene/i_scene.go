package scene

import "gameSrv/pkg/math"

// IScene 抽象场景接口
type IScene interface {
	// GetSceneCnfId 获取场景配置ID
	GetSceneCnfId() int32

	// GetSceneUID 获取场景唯一ID
	GetSceneUID() int64

	// GetOwnerUID 获取拥有者UID
	GetOwnerUID() int64

	// GetSceneType 获取场景类型
	GetSceneType() int32

	// GetSightModule 获取视野模块
	GetSightModule() ISceneSightModule

	// GetPhasingData 获取或创建相位数据
	GetPhasingData(phasingId int64) IScenePhasingData

	// GetEntityMap 获取实体Map
	GetEntityMap() map[int64]IEntity

	// GetPlayerViewMgrMap 获取玩家视野管理器Map
	GetPlayerViewMgrMap() map[int64]*PlayerViewMgr

	// GetPlayerPreEnterInfoMap 获取玩家预进入信息Map
	GetPlayerPreEnterInfoMap() map[int64]*ScenePlayerPreSlotInfo

	// PlayerPreEnter 玩家预进入场景
	PlayerPreEnter(player IGamePlayer) error

	// PlayerEnterSceneBegin 玩家进入场景开始
	PlayerEnterSceneBegin(player IGamePlayer)

	// PlayerEnterSceneDone 玩家进入场景完成
	PlayerEnterSceneDone(player IGamePlayer)

	// PlayerLeaveScene 玩家离开场景
	PlayerLeaveScene(player IGamePlayer, leaveSceneType int32) error

	// PlayerInBattleChange 玩家战斗状态变化
	PlayerInBattleChange(player IGamePlayer)

	// EntityAppear 实体出现
	EntityAppear(act IEntity, context *VisionContext) error

	// EntityDisappear 实体消失
	EntityDisappear(act IEntity, context *VisionContext, deadClearTime int64) error

	// EntityMoveTo 实体移动
	EntityMoveTo(act IEntity, pos, rot *math.Vector3) error

	// EntityChangePhasing 实体改变相位
	EntityChangePhasing(act IEntity, newPhasingId int64) error

	// FindMoveEntityIncludeProxy 查找实体（返回代理AvatarTeamActor）
	FindMoveEntityIncludeProxy(player IGamePlayer, entityId int64) IEntity

	// FindRealEntityIncludeProxy 查找实体（返回真实Actor）
	FindRealEntityIncludeProxy(player IGamePlayer, entityId int64) IEntity

	// FindEntity 查找实体
	FindEntity(entityId int64) IEntity

	// ProcessEntityMoveInfo 处理实体移动信息
	ProcessEntityMoveInfo(player IGamePlayer, moveInfo interface{}) error

	// GetMonsterLevel 获取怪物等级
	GetMonsterLevel(monster IEntity) int32

	// GetSceneBuff 获取场景Buff
	GetSceneBuff() []int32

	// GetMonsterBuff 获取怪物Buff
	GetMonsterBuff() []int32

	// GetSceneProps 获取场景属性
	GetSceneProps() map[int32]int32

	// GetMonsterProps 获取怪物属性
	GetMonsterProps() map[int32]int32

	// IsLoadFormationData 检查是否加载编队数据
	IsLoadFormationData() bool

	// IsSaveToFormationData 检查是否保存到编队数据
	IsSaveToFormationData() bool

	// String 返回字符串表示
	String() string

	// AllowOption 检查场景选项是否允许
	AllowOption(optionType int32) bool

	// ForeachAllPlayer 遍历所有玩家
	ForeachAllPlayer(consumer func(IGamePlayer))

	// ForeachPhasingPlayer 遍历指定相位的玩家
	ForeachPhasingPlayer(phasingId int64, consumer func(IGamePlayer))

	// FindPlayerViewMgr 查找玩家的视野管理器
	FindPlayerViewMgr(uid int64) *PlayerViewMgr
}

// IScenePhasingData 场景相位数据接口
type IScenePhasingData interface {
	GetPhasingId() int64

	GetEntityMap() map[int64]IEntity

	GetPreAssignedActors() map[int64]IEntity

	GetActorDestroyTimeMap() map[int64]int64

	GetResetActorHandler() IResetActorHandler

	GetDeadActorRecords() map[int64]IEntity

	IsRecordActorDead(systemGuid int64) bool

	AddDeadRecord(systemGuid int64)

	RemoveResetActor(phasingId int64, entityID int64)

	ClearDeadActorRecords()

	EntityChangePhasingData(act IEntity, oldPhasingId int64, newPhasingId int64)

	EntityLeavePhasing(act IEntity)

	GetEntityChangePhasing(act IEntity, newPhasingId int64)

	GetPreAssignedActor(systemGuid int64, preAssignedId int64) IEntity

	GetSameConfigActor(configID int64) []IEntity

	GetAutoVisualActors() map[int64]struct{}

	AddAutoVisualActor(systemGuid int64) struct{}

	RemoveAutoVisualActor(systemGuid int64) bool

	FilterAlwaysActiveTags(player IGamePlayer)

	FilterConditionLayerTags(player IGamePlayer)

	LoadActorFromLevelTag(player IGamePlayer, layerTag string) bool

	UnLoadActorWithTag(player IGamePlayer, layerTag string) bool

	UnLoadLevelActors(player IGamePlayer, levelActorCfgs interface{}) bool

	LoadLevelInstanceActors(player IGamePlayer, layerTagActors interface{}) bool

	LoadActorFromDO(player IGamePlayer, sceneActorData interface{}) error

	CheckWaitLoadTags(player IGamePlayer)

	InitActorFromDO(player IGamePlayer, sceneDataDO interface{}) error

	FlushPlayerSceneDataDO(player IGamePlayer, sceneDataDO interface{})

	GetMonsterLevel(monster interface{}) int32

	GetSceneBuff() []int32

	GetMonsterBuff() []int32

	GetSceneProps() map[int32]int32

	IsLoadFormationData() bool

	IsSaveToFormationData() bool

	AllowOption(optionType int32) bool

	GetNumber() int32
}

// IResetActorHandler 重置Actor处理器接口
type IResetActorHandler interface {
	IsRecordActorDead(systemGuid int64) bool
	GetDeadActorRecords() map[int64]struct{}
}
