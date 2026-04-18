package scene

import "gameSrv/pkg/math"

// IEntity 实体接口，定义所有实体共有的行为
type IEntity interface {
	GetEntityId() int64
	GetConfigId() int64
	GetActorType() int32
	GetRealEntityId() int64
	GetPhasingId() int64
	GetOnlyVisionUid() int64
	CanEnterAoi() bool
	CanMove() bool
	CanEnterRegion() bool
	GetActorFlags(flag byte) bool
	SetActorFlags(flag byte, value bool)
	IsNeedSaveToDb() bool
	SetNeedSaveToDb(needSaveToDb bool)
	IsFromDb() bool
	SetFromDb(fromDb bool)
	IsReborn() bool
	SetReborn(reborn bool)
	IsNeedReset() bool
	SetNeedReset(needReset bool)
	GetOwnerPlayerUid() int64
	GetControllerPlayer() IGamePlayer
	GetControllerUid() int64
	GetOwner() IGamePlayer
	SetOwner(owner IGamePlayer)
	GetCustomInteractType() int32
	SetCustomInteractType(customType int32)
	GetLogicState() int32
	GetObjectState() int32
	InitObjectState(state int32)
	ChangeObjectState(srcEntity IEntity, state int32)
	Init()
	OnReborn()
	OnBeforeEnterScene(scene IScene, context interface{})
	OnAfterEnterScene(scene IScene, context interface{})
	LeaveScene(context IScene, deadClearTime int64)
	OnBeforeLeaveScene(context interface{})
	OnAfterLeaveScene(context interface{})
	OnEnterPlayerView(p IGamePlayer)
	OnExitPlayerView(p IGamePlayer)
	CheckPosition(pos *math.Vector3) bool
	GetSpeed() *math.Vector3
	SetSpeed(speed *math.Vector3)
	GetLastValidPosition() *math.Vector3
	SetLastValidPosition(pos *math.Vector3)
	GetMotionState() int32
	SetMotionState(state int32)
	GetLastMoveSceneTimeMs() int64
	SetLastMoveSceneTimeMs(ms int64)
	GetCachePosOrCurPos() *math.Vector3
	GetCacheRotOrCurRot() *math.Vector3
	SetBornPosRot(pos, rot *math.Vector3)
	SetCurPosRot(pos, rot *math.Vector3)
	GetInteractInfo() *InteractInfo
	SetInteractInfo(interactInfo *InteractInfo)
	GetLocation() *math.LocationInfo
	GetVisionLevelEnum() *VisionLevelEnum
	GetScene() IScene
	GetGrid() *Grid
	SetGrid(grid *Grid)
	EnterScene(scn IScene, context *VisionContext) error
	String() string
}

// InteractInfo is the interaction information
type InteractInfo struct {
	// CustomInteractType is the custom interaction handler type
	CustomInteractType int32
	// TriggeredGuids is the list of triggers
	TriggeredGuids []int64

	// LogicState is the bit flag state value IsActivated|CanInteract|IsLocked
	LogicState int32
	// BornLogicState is the born logic state
	BornLogicState int32

	// ObjectState is the corresponding Entity's enum in scene_object.xlsx
	ObjectState int32
	// BornObjectState is the born object state
	BornObjectState int32
}
