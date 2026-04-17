// Copyright Epic Games, Inc. All Rights Reserved.

package scene

import "gameSrv/pkg/math"

// IGamePlayer 玩家接口
type IGamePlayer interface {
	GetUid() int64

	// AvatarTeam 相关方法
	GetAvatarTeam() IEntity
	GetCurAvatarConfId() int64
	InPrivatePhasing() bool

	// Avatar 属性获取方法
	GetConfigID() int64
	GetLevel() int32
	GetLifeState() int32
	GetActorID() int64
	GetExp() int64
	GetExceedID() int64
	GetCampType() int32

	// 战斗属性
	GetBattleProps() map[int32]int32

	// 位置设置
	SetCachePosRot(pos, rot *math.Vector3)

	// 事件处理
	OnTeamAvatarDead(actor IEntity)
	OnAvatarStateChange(sync bool)

	// 状态相关
	MakeAvatarState() interface{}
	ReCalAttributes()
	ReCalAllAttackDataEffect()
	GetBuffManager() interface{}
}

// AvatarStateInfo 头像状态信息接口
type AvatarStateInfo interface {
	GetCampType() int32
}
