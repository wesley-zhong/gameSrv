package scene

import (
	"errors"
	"fmt"
	"gameSrv/pkg/math"
)

var (
	// ErrPlayerNotFound 玩家未找到错误
	ErrPlayerNotFound = errors.New("player not found")
)

// Scene represents scene class implementing IScene
type Scene struct {
	OwnerUID     int64
	SceneUID     string
	SceneCnfId   int32
	CreateTime   int64
	BeginTime    int64
	AllowOptions int32

	// Player view management maps
	PlayerViewMgrMap      map[int64]*PlayerViewMgr
	PlayerPreEnterInfoMap map[int64]*ScenePlayerPreSlotInfo

	// All entities in scene, key is entityId
	EntityMap map[int64]IEntity

	// Scene phasing data, <phasingId, data>
	PhasingDataMap   map[int64]IScenePhasingData
	SceneSightModule ISceneSightModule
}

// NewScene creates a new Scene
func NewScene(sceneUID string, sceneCnfId int32, sceneSightModule ISceneSightModule, threadHashCode int64) *Scene {
	_ = threadHashCode // Reserved for future use
	return &Scene{
		SceneUID:              sceneUID,
		SceneCnfId:            sceneCnfId,
		CreateTime:            0, // CurrentTimeMillis placeholder
		PlayerViewMgrMap:      make(map[int64]*PlayerViewMgr),
		PlayerPreEnterInfoMap: make(map[int64]*ScenePlayerPreSlotInfo),
		EntityMap:             make(map[int64]IEntity),
		PhasingDataMap:        make(map[int64]IScenePhasingData),
		SceneSightModule:      sceneSightModule,
	}
}

// GetSceneCnfId returns the scene configuration ID
func (s *Scene) GetSceneCnfId() int32 {
	return s.SceneCnfId
}

// GetSceneUID returns the scene UID
func (s *Scene) GetSceneUID() int64 {
	if len(s.SceneUID) > 0 {
		// Simple hash conversion
		var hash int64 = 0
		for _, c := range s.SceneUID {
			hash = hash*31 + int64(c)
		}
		return hash
	}
	return 0
}

// GetOwnerUID returns the owner UID
func (s *Scene) GetOwnerUID() int64 {
	return s.OwnerUID
}

// GetSceneType returns the scene type
func (s *Scene) GetSceneType() int32 {
	return 0 // To be implemented by subclasses
}

// GetSightModule returns the vision module
func (s *Scene) GetSightModule() ISceneSightModule {
	return s.SceneSightModule
}

// GetPhasingData gets or creates scene phasing data
func (s *Scene) GetPhasingData(phasingId int64) IScenePhasingData {
	phasingData := s.PhasingDataMap[phasingId]
	if phasingData != nil {
		return phasingData
	}
	newData := NewScenePhasingData(phasingId)
	s.PhasingDataMap[phasingId] = newData
	return newData
}

// GetOrCreatePhasingData gets or creates scene phasing data
func (s *Scene) GetOrCreatePhasingData(phasingId int64) *ScenePhasingData {
	phasingData, exists := s.PhasingDataMap[phasingId]
	if exists {
		if scenePhasingData, ok := phasingData.(*ScenePhasingData); ok {
			return scenePhasingData
		}
	}
	newData := NewScenePhasingData(phasingId)
	s.PhasingDataMap[phasingId] = newData
	return newData
}

// GetEntityMap returns the entity map
func (s *Scene) GetEntityMap() map[int64]IEntity {
	return s.EntityMap
}

// GetPlayerViewMgrMap returns the player view manager map
func (s *Scene) GetPlayerViewMgrMap() map[int64]*PlayerViewMgr {
	return s.PlayerViewMgrMap
}

// GetPlayerPreEnterInfoMap returns the player pre-enter info map
func (s *Scene) GetPlayerPreEnterInfoMap() map[int64]*ScenePlayerPreSlotInfo {
	return s.PlayerPreEnterInfoMap
}

// FindPlayerViewMgr finds a player's view manager
func (s *Scene) FindPlayerViewMgr(uid int64) *PlayerViewMgr {
	return s.PlayerViewMgrMap[uid]
}

// PlayerPreEnter handles player pre-entering scene
func (s *Scene) PlayerPreEnter(player IGamePlayer) error {
	if player == nil {
		return ErrPlayerNotFound
	}
	scenePlayerPreSlotInfo := NewScenePlayerPreSlotInfo()
	scenePlayerPreSlotInfo.UID = player.GetUid()
	s.PlayerPreEnterInfoMap[player.GetUid()] = scenePlayerPreSlotInfo
	return nil
}

// PlayerEnterSceneBegin handles player entering scene start
func (s *Scene) PlayerEnterSceneBegin(player IGamePlayer) {
	if player == nil {
		return
	}
	preEnterPlayer := s.PlayerPreEnterInfoMap[player.GetUid()]
	delete(s.PlayerPreEnterInfoMap, player.GetUid())
	if preEnterPlayer == nil {
		return
	}

	// Add player view
	s.PlayerViewMgrMap[player.GetUid()] = NewPlayerViewMgr(s, player)
}

// PlayerEnterSceneDone handles player entering scene completion
func (s *Scene) PlayerEnterSceneDone(player IGamePlayer) {
	// Complete enter process
	if player == nil {
		return
	}
}

// PlayerLeaveScene handles player leaving scene
func (s *Scene) PlayerLeaveScene(player IGamePlayer, leaveSceneType int32) error {
	_ = leaveSceneType // Reserved for future use
	if player == nil {
		return ErrPlayerNotFound
	}
	// Clear player data in scene
	delete(s.PlayerViewMgrMap, player.GetUid())
	return nil
}

// PlayerInBattleChange handles player battle state change
func (s *Scene) PlayerInBattleChange(player IGamePlayer) {
	_ = player // Reserved for future use
	// TODO: implement battle state change logic
}

// EntityAppear handles entity appearing in scene
func (s *Scene) EntityAppear(act IEntity, context *VisionContext) error {
	_ = context // Reserved for future use
	if act == nil {
		return nil
	}
	// Add entity to entity map
	s.EntityMap[act.GetEntityId()] = act
	// TODO: trigger entity appear event
	return nil
}

// EntityDisappear handles entity disappearing from scene
func (s *Scene) EntityDisappear(act IEntity, context *VisionContext, deadClearTime int64) error {
	_ = context       // Reserved for future use
	_ = deadClearTime // Reserved for future use
	if act == nil {
		return nil
	}
	// Remove entity from entity map
	delete(s.EntityMap, act.GetEntityId())
	// TODO: trigger entity disappear event
	return nil
}

// EntityMoveTo handles entity moving to destination
func (s *Scene) EntityMoveTo(act IEntity, pos, rot *math.Vector3) error {
	_ = rot // Reserved for future use
	if act == nil || s.SceneSightModule == nil {
		return nil
	}
	s.SceneSightModule.EntityMoveTo(act, pos)
	return nil
}

// EntityChangePhasing handles entity changing phasing in scene
func (s *Scene) EntityChangePhasing(act IEntity, newPhasingId int64) error {
	if act == nil {
		return nil
	}
	phasingData := s.GetPhasingData(newPhasingId)
	if phasingData == nil {
		return nil
	}
	phasingData.EntityChangePhasingData(act, 0, newPhasingId)
	return nil
}

// FindMoveEntityIncludeProxy finds entity (returns proxied AvatarTeamActor)
func (s *Scene) FindMoveEntityIncludeProxy(player IGamePlayer, entityId int64) IEntity {
	_ = player // Reserved for future use
	// TODO: implement entity search logic with proxy support
	return s.FindEntity(entityId)
}

// FindRealEntityIncludeProxy finds entity (returns real actor corresponding to EntityId)
func (s *Scene) FindRealEntityIncludeProxy(player IGamePlayer, entityId int64) IEntity {
	_ = player // Reserved for future use
	// TODO: implement entity search logic with proxy support
	return s.FindEntity(entityId)
}

// FindEntity finds entity by entityId
func (s *Scene) FindEntity(entityId int64) IEntity {
	return s.EntityMap[entityId]
}

// ProcessEntityMoveInfo processes entity move info
func (s *Scene) ProcessEntityMoveInfo(player IGamePlayer, moveInfo interface{}) error {
	_ = player   // Reserved for future use
	_ = moveInfo // Reserved for future use
	// TODO: implement entity move info processing
	return nil
}

// GetMonsterLevel gets monster level
func (s *Scene) GetMonsterLevel(monster IEntity) int32 {
	if monster == nil {
		return 1
	}
	// TODO: implement monster level calculation based on scene config
	return 1
}

// GetSceneBuff gets scene buffs
func (s *Scene) GetSceneBuff() []int32 {
	// TODO: implement scene buff retrieval
	return []int32{}
}

// GetMonsterBuff gets monster buffs
func (s *Scene) GetMonsterBuff() []int32 {
	// TODO: implement monster buff retrieval
	return []int32{}
}

// GetSceneProps gets scene properties
func (s *Scene) GetSceneProps() map[int32]int32 {
	// TODO: implement scene properties retrieval
	return make(map[int32]int32)
}

// GetMonsterProps gets monster properties
func (s *Scene) GetMonsterProps() map[int32]int32 {
	// TODO: implement monster properties retrieval
	return make(map[int32]int32)
}

// IsLoadFormationData checks if formation data should be loaded
func (s *Scene) IsLoadFormationData() bool {
	// TODO: implement formation data check
	return false
}

// IsSaveToFormationData checks if should save to formation data
func (s *Scene) IsSaveToFormationData() bool {
	// TODO: implement formation data save check
	return false
}

// AllowOption checks if a scene option is allowed
func (s *Scene) AllowOption(optionType int32) bool {
	return s.AllowOptions&(1<<uint32(optionType)) != 0
}

// ForeachAllPlayer iterates over all players
func (s *Scene) ForeachAllPlayer(consumer func(IGamePlayer)) {
	for _, viewMgr := range s.PlayerViewMgrMap {
		consumer(viewMgr.GetPlayer())
	}
}

// ForeachPhasingPlayer iterates over players in specified phase
func (s *Scene) ForeachPhasingPlayer(phasingId int64, consumer func(IGamePlayer)) {
	for _, viewMgr := range s.PlayerViewMgrMap {
		player := viewMgr.GetPlayer()
		avatarTeam := player.GetAvatarTeam()
		if avatarTeam != nil && avatarTeam.GetPhasingId() == phasingId {
			consumer(player)
		}
	}
}

// String returns string representation
func (s *Scene) String() string {
	return fmt.Sprintf("scene id = %d owner id = %d", s.SceneCnfId, s.OwnerUID)
}
