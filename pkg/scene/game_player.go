package scene

import (
	"gameSrv/pkg/math"
)

// IGamePlayer 玩家接口
type IGamePlayer interface {
	GetUid() int64

	// AvatarTeam 相关方法
	GetAvatarTeam() IEntity
	GetCurAvatarConfId() int64
	InPrivatePhasing() bool

	GetLevel() int32
	GetLifeState() int32

	GetExp() int64
	GetExceedID() int64
	GetCampType() int32

	// 位置设置
	SetCachePosRot(pos, rot *math.Vector3)

	// 事件处理
	OnTeamAvatarDead(actor IEntity)
}
