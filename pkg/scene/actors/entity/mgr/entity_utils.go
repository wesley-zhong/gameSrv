// Copyright Epic Games, Inc. All Rights Reserved.

package mgr

// ActorType represents actor type
type ActorType int32

const (
	ActorTypeNone           ActorType = 0
	ActorTypeGadget         ActorType = 1
	ActorTypeMonster        ActorType = 2
	ActorTypeChest          ActorType = 3
	ActorTypeSummon         ActorType = 4
	ActorTypeBuffVolume     ActorType = 5
	ActorTypeKillZVolume    ActorType = 6
	ActorTypeLadder         ActorType = 7
	ActorTypeInvisibleWall ActorType = 8
	ActorTypeInteractive    ActorType = 9
	ActorTypeTeleporter     ActorType = 10
	ActorTypeRestPoint      ActorType = 11
	ActorTypeCommonSwitch  ActorType = 12
	ActorTypeHeroAvatar    ActorType = 13
	ActorTypeChestDestruction ActorType = 14
	ActorTypeGameFlow      ActorType = 15
	ActorTypeSummonSpawner ActorType = 16
	ActorTypeSummonMachine ActorType = 17
)

const (
	PeerBits      = 3
	CategoryBits  = 5
	IsSyncedBits  = 1
	SequenceBits  = 32 - (PeerBits + IsSyncedBits + CategoryBits)
	PeerShift     = CategoryBits + IsSyncedBits + SequenceBits
	CategoryShift = IsSyncedBits + SequenceBits
	IsSyncedShift = SequenceBits
	EffectCate      = 16
	AttackUnitCate   = 17
	CameraCate       = 18
	ManagerCate      = 19
	LocalGadgetCate  = 20
	LocalMassiveCate  = 21
	LevelRuntimeID = 0<<PeerShift | ManagerCate<<CategoryShift | 1<<IsSyncedShift | 1
)

// GetEntityId 获取实体ID
func GetEntityId(actorType ActorType, index int32) int32 {
	return int32(actorType)<<24 | index
}

// GetEntityType 获取实体类型
func GetEntityType(entityId int32) ActorType {
	return ActorType(entityId >> 24)
}

// IsGadget 判断是否是Gadget
func IsGadget(actor interface{}) bool {
	if act, ok := actor.(ActorTypeGetter); ok {
		return act.GetActorType() == ActorTypeGadget
	}
	return false
}

// ActorTypeGetter 获取Actor类型的接口
type ActorTypeGetter interface {
	GetActorType() ActorType
}