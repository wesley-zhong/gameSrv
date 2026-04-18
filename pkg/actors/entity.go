package actors

import (
	"errors"
	"fmt"
	"gameSrv/pkg/math"
	"gameSrv/pkg/scene"

	"go.uber.org/zap"
)

// Entity is the base class for all entities in the scene
type Entity struct {
	// entityId is the entity ID
	EntityID int64

	// preAssignedGuid is the GUID exported from the editor, used to record state changes
	PreAssignedGuid int64
	// spawnerPreAssignedId is the preAssignedGuid of the spawner that generated this actors
	SpawnerPreAssignedID int64
	// configId is the configuration ID
	ConfigID int64

	// owner is the owning player (only player-owned actors can have an owner)
	Owner scene.IGamePlayer
	// host is the host client
	Host scene.IGamePlayer
	// scene is the scene where this actors exists
	Scene scene.IScene

	// actorFlags is the bit storage data for ActorFlag, with default NeedSaveToDb as true
	ActorFlags byte

	// phasingId is the scene phase ID, 0 is for public entity list, negative is for multi-player phase, positive is for player-specific phase
	PhasingID int64

	// location is the coordinate information
	Location *math.LocationInfo
	// moveInfo is the movement information
	MoveInfo *math.MoveInfo
	// interactInfo is the common interaction information
	InteractInfo *scene.InteractInfo

	// layerTag is the layer tag
	LayerTag string
	// resetTypes is the list of resetable event types initialized from JSON data
	ResetTypes []int32
	// utilityDropData is the common drop data
	UtilityDropData *UtilityDropData

	// spawnerPropertyCfg is the spawner property configuration
	SpawnerPropertyCfg interface{}

	VisionLevelEnum scene.VisionLevelEnum
	Grid            *scene.Grid

	logger *zap.Logger
}

// NewEntity creates a new base Entity
func NewEntity() *Entity {
	return &Entity{
		Location:     &math.LocationInfo{},
		InteractInfo: &scene.InteractInfo{},
		MoveInfo:     &math.MoveInfo{},
		ActorFlags:   1, // NeedSaveToDb = true by default
	}
}

func (entity *Entity) GetPhasingId() int64 {
	return entity.PhasingID
}

// GetActorType returns the actors type
func (entity *Entity) GetActorType() int32 {
	return 0 // TODO: return actual actor type
}

// GetEntityId returns the entity ID
func (entity *Entity) GetEntityId() int64 {
	return entity.EntityID
}

// GetConfigId returns the configuration ID
func (entity *Entity) GetConfigId() int64 {
	return entity.ConfigID
}

// Init initializes the actors
func (entity *Entity) Init() {
	// To be implemented by subclasses
}

// GetRealEntityId returns the client-used entity ID
func (entity *Entity) GetRealEntityId() int64 {
	return entity.EntityID
}

// GetOnlyVisionUid returns the UID of the client that can see this actors
func (entity *Entity) GetOnlyVisionUid() int64 {
	return 0
}

// CanEnterAoi returns whether this actors needs to enter AOI vision management
func (entity *Entity) CanEnterAoi() bool {
	return true
}

// CanMove returns whether this actors can move
func (entity *Entity) CanMove() bool {
	return true
}

// CanEnterRegion returns whether this actors can enter regions
func (entity *Entity) CanEnterRegion() bool {
	return false
}

// GetActorFlags returns the value of the specified actors flag
func (entity *Entity) GetActorFlags(flag byte) bool {
	return (entity.ActorFlags & flag) != 0
}

// SetActorFlags sets the value of the specified actors flag
func (entity *Entity) SetActorFlags(flag byte, value bool) {
	if value {
		entity.ActorFlags |= flag
	} else {
		entity.ActorFlags &^= flag
	}
}

// IsNeedSaveToDb returns whether to save data to DB
func (entity *Entity) IsNeedSaveToDb() bool {
	return entity.GetActorFlags(1)
}

// SetNeedSaveToDb sets whether to save data to DB
func (entity *Entity) SetNeedSaveToDb(needSaveToDb bool) {
	entity.SetActorFlags(1, needSaveToDb)
}

// IsFromDb returns whether this actors was created from DB
func (entity *Entity) IsFromDb() bool {
	return entity.GetActorFlags(2)
}

// SetFromDb sets whether this actors was created from DB
func (entity *Entity) SetFromDb(fromDb bool) {
	entity.SetActorFlags(2, fromDb)
}

// IsReborn returns whether this actors has been reborn
func (entity *Entity) IsReborn() bool {
	return entity.GetActorFlags(4)
}

// SetReborn sets whether this actors has been reborn
func (entity *Entity) SetReborn(reborn bool) {
	entity.SetActorFlags(4, reborn)
}

// IsNeedReset returns whether this actors needs to be reset
func (entity *Entity) IsNeedReset() bool {
	return entity.GetActorFlags(8)
}

// SetNeedReset sets whether this actors needs to be reset
func (entity *Entity) SetNeedReset(needReset bool) {
	entity.SetActorFlags(8, needReset)
}

// GetOwnerPlayerUid returns the owner player UID
func (entity *Entity) GetOwnerPlayerUid() int64 {
	if entity.Owner == nil {
		return 0
	}
	return entity.Owner.GetUid()
}

// GetControllerPlayer returns the actual controller (host if exists, otherwise owner)
func (entity *Entity) GetControllerPlayer() scene.IGamePlayer {
	if entity.Host != nil {
		return entity.Host
	}
	return entity.Owner
}

// GetControllerUid returns the controller UID
func (entity *Entity) GetControllerUid() int64 {
	controller := entity.GetControllerPlayer()
	if controller == nil {
		return 0
	}
	return controller.GetUid()
}

// CheckPosition checks if a position is valid
func (entity *Entity) CheckPosition(pos *math.Vector3) bool {
	return true
}

// OnReborn is called when the actors is reborn
func (entity *Entity) OnReborn() {
	entity.SetReborn(true)
}

// OnBeforeEnterScene is called before entering the scene
func (entity *Entity) OnBeforeEnterScene(scn scene.IScene, context interface{}) {
	entity.Init()
}

// OnAfterEnterScene is called after entering the scene
func (entity *Entity) OnAfterEnterScene(scn scene.IScene, context interface{}) {
	// TODO: implement
}

// LeaveScene leaves the scene
func (entity *Entity) LeaveScene(context scene.IScene, deadClearTime int64) {
	entity.OnBeforeLeaveScene(context)
	entity.Scene = nil
	entity.OnAfterLeaveScene(context)
}

// OnBeforeLeaveScene is called before leaving the scene
func (entity *Entity) OnBeforeLeaveScene(context interface{}) {
	// To be implemented by subclasses
}

// OnAfterLeaveScene is called after leaving the scene
func (entity *Entity) OnAfterLeaveScene(context interface{}) {
	// To be implemented by subclasses
}

// OnEnterPlayerView is called when entering a player's view
func (entity *Entity) OnEnterPlayerView(p scene.IGamePlayer) {
	// TODO: implement
}

// OnExitPlayerView is called when leaving a player's view
func (entity *Entity) OnExitPlayerView(p scene.IGamePlayer) {
	// TODO: implement
}

// GetCustomInteractType returns the custom interact type
func (entity *Entity) GetCustomInteractType() int32 {
	customType := int32(0) // DEFAULT
	if entity.InteractInfo != nil {
		customType = entity.InteractInfo.CustomInteractType
	}
	return customType
}

// SetCustomInteractType sets the custom interact type
func (entity *Entity) SetCustomInteractType(customType int32) {
	if entity.InteractInfo == nil {
		return
	}
	entity.InteractInfo.CustomInteractType = customType
}

// SetActivated sets the activated state
func (entity *Entity) SetActivated(state int32) {
	if entity.InteractInfo == nil {
		entity.logger.Warn("Entity setActivated fail, not Interact")
		return
	}
	// TODO: implement logic state handling
}

// SetLockState sets the lock state
func (entity *Entity) SetLockState(lockState int32) {
	if entity.InteractInfo == nil {
		entity.logger.Warn("Entity setLockMode fail, not Interact")
		return
	}
	// TODO: implement
}

// SetInteractMode sets the interact mode
func (entity *Entity) SetInteractMode(mode int32) {
	if entity.InteractInfo == nil {
		entity.logger.Warn("Entity setInteractMode fail, not Interact")
		return
	}
	// TODO: implement
}

// ChangeInteractMode changes the interact mode
func (entity *Entity) ChangeInteractMode(newMode int32) {
	entity.SetInteractMode(newMode)
}

// ChangeLogicState changes the logic state
func (entity *Entity) ChangeLogicState(state int32) {
	if entity.InteractInfo == nil {
		entity.logger.Error("Entity change state failed, not interactable")
		return
	}
	entity.InteractInfo.LogicState = state
}

// SetBornObjectState sets the born object state
func (entity *Entity) SetBornObjectState(state int32) {
	if entity.InteractInfo != nil {
		entity.InteractInfo.BornObjectState = state
	}
}

// InitObjectState initializes the object state
func (entity *Entity) InitObjectState(state int32) {
	if entity.InteractInfo == nil {
		entity.logger.Error("Entity init state failed, not interactable")
		return
	}
	entity.InteractInfo.ObjectState = state
	entity.InteractInfo.BornObjectState = state
}

// ChangeObjectState changes the object state
func (entity *Entity) ChangeObjectState(srcEntity scene.IEntity, state int32) {
	if entity.InteractInfo == nil {
		entity.logger.Error("Entity change state failed, not interactable")
		return
	}
	if entity.InteractInfo.ObjectState == state {
		return
	}
	entity.InteractInfo.ObjectState = state
	// TODO: broadcast change
	// TODO: handle drop
}

// GetLogicState returns the logic state
func (entity *Entity) GetLogicState() int32 {
	if entity.InteractInfo == nil {
		entity.logger.Error("Entity getLogicState fail, not Interact")
		return 0
	}
	return entity.InteractInfo.LogicState
}

// GetObjectState returns the object state
func (entity *Entity) GetObjectState() int32 {
	if entity.InteractInfo == nil {
		entity.logger.Error("Entity getState fail, not Interact")
		return 0
	}
	return entity.InteractInfo.ObjectState
}

// String returns a string representation
func (entity *Entity) String() string {
	return fmt.Sprintf("{PhasingId=%d, ControllerUid=%d, EntityId=%d, PreAssignedGuid=%d}",
		entity.PhasingID, entity.GetControllerUid(), entity.EntityID, entity.PreAssignedGuid)
}

// MotionContext is the movement context
type MotionContext struct {
	// Input fields
	SceneTimeMs int64
	ExcludeUid  int64
	Notify      bool

	// Output fields
	SyncUidList []int64
	DoMove      bool
}

// GetSpeed returns the speed
func (entity *Entity) GetSpeed() *math.Vector3 {
	if entity.MoveInfo == nil {
		return nil
	}
	return entity.MoveInfo.Speed
}

// SetSpeed sets the speed
func (entity *Entity) SetSpeed(speed *math.Vector3) {
	if entity.MoveInfo != nil {
		entity.MoveInfo.Speed = speed
	}
}

// GetLastValidPosition returns the last valid position
func (entity *Entity) GetLastValidPosition() *math.Vector3 {
	if entity.MoveInfo == nil {
		return nil
	}
	return entity.MoveInfo.LastValidPosition
}

// SetLastValidPosition sets the last valid position
func (entity *Entity) SetLastValidPosition(pos *math.Vector3) {
	if entity.MoveInfo != nil {
		entity.MoveInfo.LastValidPosition = pos
	}
}

// GetMotionState returns the motion state
func (entity *Entity) GetMotionState() int32 {
	if entity.MoveInfo == nil {
		return 0
	}
	return entity.MoveInfo.MotionState
}

// SetMotionState sets the motion state
func (entity *Entity) SetMotionState(state int32) {
	if entity.MoveInfo != nil {
		entity.MoveInfo.MotionState = state
	}
}

// GetLastMoveSceneTimeMs returns the last move scene time
func (entity *Entity) GetLastMoveSceneTimeMs() int64 {
	if entity.MoveInfo == nil {
		return 0
	}
	return entity.MoveInfo.LastMoveSceneTimeMs
}

// SetLastMoveSceneTimeMs sets the last move scene time
func (entity *Entity) SetLastMoveSceneTimeMs(ms int64) {
	if entity.MoveInfo != nil {
		entity.MoveInfo.LastMoveSceneTimeMs = ms
	}
}

// GetCachePosOrCurPos returns the cache position or current position
func (entity *Entity) GetCachePosOrCurPos() *math.Vector3 {
	if entity.MoveInfo != nil {
		// TODO: implement logic for cache pos
	}
	return entity.Location.Position
}

// GetCacheRotOrCurRot returns the cache rotation or current rotation
func (entity *Entity) GetCacheRotOrCurRot() *math.Vector3 {
	if entity.MoveInfo != nil {
		// TODO: implement logic for cache rot
	}
	return entity.Location.Rotation
}

// SetBornPosRot sets the birth position and rotation
func (entity *Entity) SetBornPosRot(pos, rot *math.Vector3) {
	entity.Location.BornPos = pos
	entity.Location.BornRot = rot
	entity.SetCurPosRot(pos, rot)
}

// SetCurPosRot sets the current position and rotation
func (entity *Entity) SetCurPosRot(pos, rot *math.Vector3) {
	entity.Location.Position = pos
	entity.Location.Rotation = rot
}

// GetOwner returns the owner player
func (entity *Entity) GetOwner() scene.IGamePlayer {
	return entity.Owner
}

// SetOwner sets the owner player
func (entity *Entity) SetOwner(owner scene.IGamePlayer) {
	entity.Owner = owner
}

// GetInteractInfo returns the interact info
func (entity *Entity) GetInteractInfo() *scene.InteractInfo {
	return entity.InteractInfo
}

// SetInteractInfo sets the interact info
func (entity *Entity) SetInteractInfo(interactInfo *scene.InteractInfo) {
	entity.InteractInfo = interactInfo
}

func (entity *Entity) GetLocation() *math.LocationInfo {
	return entity.Location
}

// GetVisionLevelEnum returns the vision level enum
func (entity *Entity) GetVisionLevelEnum() *scene.VisionLevelEnum {
	return &entity.VisionLevelEnum
}

// GetScene returns the scene
func (entity *Entity) GetScene() scene.IScene {
	if entity.Scene == nil {
		return nil
	}
	// Type assertion from interface{} to scene.IScene
	if scn, ok := entity.Scene.(scene.IScene); ok {
		return scn
	}
	return nil
}

func (entity *Entity) GetGrid() *scene.Grid {
	return entity.Grid
}

func (entity *Entity) SetGrid(grid *scene.Grid) {
	entity.Grid = grid
}

// EnterScene enters the scene
func (entity *Entity) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	defer func() {
		if r := recover(); r != nil {
			entity.logger.Error("Actor enterScene panic", zap.Any("error", r))
		}
	}()

	entity.Scene = scn
	if entity.PhasingID == 0 {
		return errors.New("Actor enter 0 phasing")
	}

	entity.OnBeforeEnterScene(scn, context)
	scn.EntityAppear(entity, context)
	entity.OnAfterEnterScene(scn, context)
	return nil
}

// UtilityDropData is the utility drop data
type UtilityDropData struct {
	DropType  int32
	DropID    int32
	DropCount int32
	DropState int32
}

// recordMotionStateSet is the set of motion states that need to record position
var recordMotionStateSet = []int32{
	1, // MOTION_WALK
	2, // MOTION_NAV_WALK
	3, // MOTION_SWIM_MOVE
	4, // MOTION_RESET
}
