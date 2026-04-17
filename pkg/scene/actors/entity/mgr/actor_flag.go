// Copyright Epic Games, Inc. All Rights Reserved.

package mgr

// ActorFlag Actor标志位枚举
type ActorFlag int

const (
	// ActorFlagNeedSaveDB 是否需要落库
	ActorFlagNeedSaveDB ActorFlag = 1 << iota // 0000_0001
	// ActorFlagFromDB 是否是从DB创建
	ActorFlagFromDB // 0000_0010
	// ActorFlagReborn 是否重生
	ActorFlagReborn // 0000_0100
	// ActorFlagNeedReset 是否需要重置
	ActorFlagNeedReset // 0000_1000
)

// Mask 获取标志位的mask
func (f ActorFlag) Mask() int {
	return int(f)
}

// IsSet 检查标志位是否设置
func IsSet(flags byte, flag ActorFlag) bool {
	return (flags & byte(flag.Mask())) != 0
}

// Set 设置或清除标志位
func Set(flags byte, flag ActorFlag, value bool) byte {
	if value {
		return flags | byte(flag.Mask())
	}
	return flags & ^byte(flag.Mask())
}

// SetFlag 设置标志位
func SetFlag(flags byte, flag ActorFlag) byte {
	return flags | byte(flag.Mask())
}

// ClearFlag 清除标志位
func ClearFlag(flags byte, flag ActorFlag) byte {
	return flags & ^byte(flag.Mask())
}
