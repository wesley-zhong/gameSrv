package battle

import (
	"gameSrv/cnfGen/cfg"
	"gameSrv/game/gamedata"
	"time"
)

// MilliSeconds returns current time in milliseconds
func MilliSeconds() int64 {
	return time.Now().UnixMilli()
}

// GetBuffProperty gets buff configuration by id
func GetBuffProperty(id int) *cfg.Buff {
	if gamedata.Tables == nil || gamedata.Tables.TbBuff == nil {
		return nil
	}
	return gamedata.Tables.TbBuff.Get(int32(id))
}

// GetAttackActionProperty gets attack data configuration by id
func GetAttackActionProperty(id int) *cfg.SkilldataMontageAttackData {
	if gamedata.Tables == nil || gamedata.Tables.TbMontageAttackData == nil {
		return nil
	}
	return gamedata.Tables.TbMontageAttackData.Get(int32(id))
}