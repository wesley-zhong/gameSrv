package actors

import "gameSrv/pkg/scene"

// SceneLoadingType scene loading type
type SceneLoadingType int32

const (
	SceneLoadingTypeELoadingDefault SceneLoadingType = 0
	SceneLoadingTypeELoadingBigMap  SceneLoadingType = 1
)

// TeleporterActor is a teleporter actor
type TeleporterActor struct {
	SimpleActor
	LoadingType            SceneLoadingType
	TeleportDataMap        map[int32]map[int64]struct{} // <SceneId, 传送目的地列表>
	LinkedResourceLevelIDs []int64                      // 关联的资源关卡ID列表
}

func (t *TeleporterActor) EnterScene(scn scene.IScene, context *scene.VisionContext) error {
	//TODO implement me
	panic("implement me")
}

// NewTeleporterActor creates a new TeleporterActor
func NewTeleporterActor() *TeleporterActor {
	actor := &TeleporterActor{
		TeleportDataMap: make(map[int32]map[int64]struct{}),
	}
	actor.SimpleActor = *NewSimpleActor()
	return actor
}

// AddTeleportData 添加传送数据
func (t *TeleporterActor) AddTeleportData(sceneId int32, teleportTarget int64) {
	teleportTargets, exists := t.TeleportDataMap[sceneId]
	if !exists {
		teleportTargets = make(map[int64]struct{})
		t.TeleportDataMap[sceneId] = teleportTargets
	}
	teleportTargets[teleportTarget] = struct{}{}
}

// GetActorType 获取Actor类型
func (t *TeleporterActor) GetActorType() int32 {
	return 2 // ActorTypeEActorTypeTeleporter
}

// IsCanAutoVisual 是否可自动发现
func (t *TeleporterActor) IsCanAutoVisual() bool {
	// TODO: implement
	return false
}

// GetTeleportDataMap 获取传送数据map
func (t *TeleporterActor) GetTeleportDataMap() map[int32]map[int64]struct{} {
	return t.TeleportDataMap
}

// GetLinkedResourceLevelIDs 获取关联的资源关卡ID列表
func (t *TeleporterActor) GetLinkedResourceLevelIDs() []int64 {
	return t.LinkedResourceLevelIDs
}

// SetLinkedResourceLevelIDs 设置关联的资源关卡ID列表
func (t *TeleporterActor) SetLinkedResourceLevelIDs(ids []int64) {
	t.LinkedResourceLevelIDs = ids
}
