package scene

import (
	"fmt"
)

// VisionLevelType 视野等级类型
type VisionLevelType int32

const (
	VisionLevelTypeNormal       VisionLevelType = 0
	VisionLevelTypeLittleRemote VisionLevelType = 1
	VisionLevelTypeRemote       VisionLevelType = 2
	VisionLevelTypeSuper        VisionLevelType = 3
)

// VisionLevelEnum 视野等级枚举
type VisionLevelEnum struct {
	VisionLevelType VisionLevelType // 视野等级类型
	GridWidth       int32           // 单个格子宽度
	SightRange      int32           // 视野的范围
	SightRadius     int32           // 视距(以格子为单位)
}

// 预定义的视野等级
var (
	// VisionLevelNormal 正常视距: 80M内可见
	VisionLevelNormal = VisionLevelEnum{
		VisionLevelType: VisionLevelTypeNormal,
		GridWidth:       4000,
		SightRange:      8000,
		SightRadius:     2,
	}

	// VisionLevelLittleRemote 较远视距: 160m内可见
	VisionLevelLittleRemote = VisionLevelEnum{
		VisionLevelType: VisionLevelTypeLittleRemote,
		GridWidth:       8000,
		SightRange:      16000,
		SightRadius:     2,
	}

	// VisionLevelRemote 远视距: 1000m内可见
	VisionLevelRemote = VisionLevelEnum{
		VisionLevelType: VisionLevelTypeRemote,
		GridWidth:       50000,
		SightRange:      100000,
		SightRadius:     2,
	}

	// VisionLevelSuper 超级视距: 4000m内可见
	VisionLevelSuper = VisionLevelEnum{
		VisionLevelType: VisionLevelTypeSuper,
		GridWidth:       200000,
		SightRange:      400000,
		SightRadius:     2,
	}

	// VisionLevelEnums 所有视野等级列表
	VisionLevelEnums = []VisionLevelEnum{
		VisionLevelNormal,
		VisionLevelLittleRemote,
		VisionLevelRemote,
		VisionLevelSuper,
	}
)

// NewVisionLevelEnum 创建新的视野等级枚举
func NewVisionLevelEnum(visionLevelType VisionLevelType, gridWidth, sightRange int32) (VisionLevelEnum, error) {
	// 判断是否可以整除
	if sightRange%gridWidth != 0 {
		return VisionLevelEnum{}, fmt.Errorf("VisionLevelEnum init sightRange error, visionLevelType = %d", visionLevelType)
	}
	// 判断视距是否有效
	sightRadius := sightRange / gridWidth
	if sightRadius <= 0 {
		return VisionLevelEnum{}, fmt.Errorf("VisionLevelEnum init sightRadius error, visionLevelType = %d", visionLevelType)
	}

	return VisionLevelEnum{
		VisionLevelType: visionLevelType,
		GridWidth:       gridWidth,
		SightRange:      sightRange,
		SightRadius:     sightRadius,
	}, nil
}

// ForVisionLevelType 根据视野等级类型获取枚举
func ForVisionLevelType(visionLevelType VisionLevelType) *VisionLevelEnum {
	for _, v := range VisionLevelEnums {
		if v.VisionLevelType == visionLevelType {
			return &v
		}
	}
	return nil
}

// ToVisionLevelType 转换为视野等级类型
func (v *VisionLevelEnum) ToVisionLevelType() VisionLevelType {
	return v.VisionLevelType
}

// GetVisionLevelType 获取视野等级类型
func (v *VisionLevelEnum) GetVisionLevelType() VisionLevelType {
	return v.VisionLevelType
}

// GetGridWidth 获取格子宽度
func (v *VisionLevelEnum) GetGridWidth() int32 {
	return v.GridWidth
}

// GetSightRange 获取视野范围
func (v *VisionLevelEnum) GetSightRange() int32 {
	return v.SightRange
}

// GetSightRadius 获取视距半径(以格子为单位)
func (v *VisionLevelEnum) GetSightRadius() int32 {
	return v.SightRadius
}

// String 返回字符串表示
func (v *VisionLevelEnum) String() string {
	switch v.VisionLevelType {
	case VisionLevelTypeNormal:
		return "VISION_LEVEL_NORMAL"
	case VisionLevelTypeLittleRemote:
		return "VISION_LEVEL_LITTLE_REMOTE"
	case VisionLevelTypeRemote:
		return "VISION_LEVEL_REMOTE"
	case VisionLevelTypeSuper:
		return "VISION_LEVEL_SUPER"
	default:
		return fmt.Sprintf("VisionLevelEnum(%d)", v.VisionLevelType)
	}
}
