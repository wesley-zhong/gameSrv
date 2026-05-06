package scene

import (
	"fmt"
	"gameSrv/cnfGen/cfg"
	"gameSrv/pkg/math"
	"time"
)

// LevelScene 关卡场景，基于网格的视野系统
// 这种场景类型适用于开放世界关卡，使用网格来管理视野
type LevelScene struct {
	Scene
}

// NewLevelScene 创建新的关卡场景
// gamePlayer: 玩家对象
// sceneCnfId: 场景配置ID
func NewLevelScene(gamePlayer IGamePlayer, sceneCnfId int32) *LevelScene {
	if gamePlayer == nil {
		return nil
	}

	// 获取当前时间戳（毫秒）
	now := time.Now().UnixNano() / 1e6

	// 生成唯一的场景UID
	sceneUid := fmt.Sprintf("level_scene_%d_%d_%d", sceneCnfId, gamePlayer.GetUid(), now)

	// 创建关卡场景
	levelScene := &LevelScene{}

	// 初始化基础场景字段
	levelScene.OwnerUID = gamePlayer.GetUid()
	levelScene.SceneUID = sceneUid
	levelScene.SceneCnfId = sceneCnfId
	levelScene.CreateTime = now
	levelScene.BeginTime = now
	levelScene.AllowOptions = 0 // TODO: 根据场景配置设置允许的选项

	// 初始化各类型映射
	levelScene.PlayerViewMgrMap = make(map[int64]*PlayerViewMgr)
	levelScene.PlayerPreEnterInfoMap = make(map[int64]*ScenePlayerPreSlotInfo)
	levelScene.EntityMap = make(map[int64]IEntity)
	levelScene.PhasingDataMap = make(map[int64]IScenePhasingData)

	// 创建网格视野模块
	levelScene.SceneSightModule = NewSceneGridSightModule(levelScene)
	return levelScene
}

// GetSceneType returns scene type
func (s *LevelScene) GetSceneType() int32 {
	return cfg.SceneType_LEVEL
}

// InitScene 初始化场景（需要在创建后调用）
func (s *LevelScene) InitScene(beginPos, sceneSize *math.Vector2) error {
	if s.SceneSightModule == nil {
		return fmt.Errorf("sight module not initialized")
	}
	return s.SceneSightModule.Init(beginPos, sceneSize)
}

// String returns string representation（覆盖Scene的String方法）
func (s *LevelScene) String() string {
	return fmt.Sprintf("LevelScene{uid=%s, cfgId=%d, owner=%d}",
		s.SceneUID, s.SceneCnfId, s.OwnerUID)
}

// Cleanup 清理场景资源
func (s *LevelScene) Cleanup() {
	// 清理所有玩家
	for uid := range s.PlayerViewMgrMap {
		if viewMgr := s.PlayerViewMgrMap[uid]; viewMgr != nil {
			viewMgr.ResetPlayerViewMgr()
		}
	}

	// 清理所有实体
	s.EntityMap = make(map[int64]IEntity)

	// 清理所有相位数据
	s.PhasingDataMap = make(map[int64]IScenePhasingData)

	// 清理预进入信息
	s.PlayerPreEnterInfoMap = make(map[int64]*ScenePlayerPreSlotInfo)
}

// IsPlayerInScene 检查玩家是否在场景中
func (s *LevelScene) IsPlayerInScene(playerUid int64) bool {
	_, exists := s.PlayerViewMgrMap[playerUid]
	return exists
}

// GetPlayerCount 获取场景中玩家数量
func (s *LevelScene) GetPlayerCount() int {
	return len(s.PlayerViewMgrMap)
}

// GetEntityCount 获取场景中实体数量
func (s *LevelScene) GetEntityCount() int {
	return len(s.EntityMap)
}

// BroadcastToAllPlayers 向场景中所有玩家广播消息
// TODO: 实现广播功能
// func (s *LevelScene) BroadcastToAllPlayers(msg interface{}) error {
// 	// 实现向所有玩家发送消息的逻辑
// 	return nil
// }
