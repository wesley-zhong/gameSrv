package scene

// VisionType 视野类型
type VisionType int32

const (
	VisionTypeNone          VisionType = 0
	VisionTypeMeet          VisionType = 1
	VisionTypeReborn        VisionType = 2
	VisionTypeMiss          VisionType = 3
	VisionTypeDie           VisionType = 4
	VisionTypeReplace       VisionType = 5
	VisionTypeReplaceServer VisionType = 6
	VisionTypePickup        VisionType = 7
)

// VisionContext 视野上下文
type VisionContext struct {
	Type               VisionType         // 视野类型
	Param              int64              // 参数
	ExcludeEnterAoiUid map[int64]struct{} // 此次AOI视野更新将指定uid排除在外
	ExcludeNotifyUid   map[int64]struct{} // 此次AOI通知将指定uid排除在外
}

// NewVisionContext 创建新的VisionContext
func NewVisionContext(typ VisionType) *VisionContext {
	return &VisionContext{
		Type:               typ,
		Param:              0,
		ExcludeEnterAoiUid: make(map[int64]struct{}),
		ExcludeNotifyUid:   make(map[int64]struct{}),
	}
}

// NewVisionContextWithUids 创建带有排除UID的VisionContext
func NewVisionContextWithUids(typ VisionType, param, excludeEnterAoiUid, excludeNotifyUid int64) *VisionContext {
	vc := &VisionContext{
		Type:               typ,
		Param:              param,
		ExcludeEnterAoiUid: make(map[int64]struct{}),
		ExcludeNotifyUid:   make(map[int64]struct{}),
	}
	if excludeEnterAoiUid > 0 {
		vc.ExcludeEnterAoiUid[excludeEnterAoiUid] = struct{}{}
	}
	if excludeNotifyUid > 0 {
		vc.ExcludeNotifyUid[excludeNotifyUid] = struct{}{}
	}
	return vc
}

// NewVisionContextWithSets 创建带有排除UID集合的VisionContext
func NewVisionContextWithSets(typ VisionType, param int64, excludeEnterAoiUid, excludeNotifyUid map[int64]struct{}) *VisionContext {
	vc := &VisionContext{
		Type:               typ,
		Param:              param,
		ExcludeEnterAoiUid: make(map[int64]struct{}),
		ExcludeNotifyUid:   make(map[int64]struct{}),
	}
	if excludeEnterAoiUid != nil {
		vc.ExcludeEnterAoiUid = excludeEnterAoiUid
	}
	if excludeNotifyUid != nil {
		vc.ExcludeNotifyUid = excludeNotifyUid
	}
	return vc
}

// NewVisionContextFromContext 从现有VisionContext创建
func NewVisionContextFromContext(typ VisionType, context *VisionContext) *VisionContext {
	return &VisionContext{
		Type:               typ,
		Param:              context.Param,
		ExcludeEnterAoiUid: context.ExcludeEnterAoiUid,
		ExcludeNotifyUid:   context.ExcludeNotifyUid,
	}
}

// BuildContextOfAddUid 添加UID到排除集合
func BuildContextOfAddUid(context *VisionContext, excludeEnterAoiUid, excludeNotifyUid int64) *VisionContext {
	newExcludeEnterAoi := addUidToSet(context.ExcludeEnterAoiUid, excludeEnterAoiUid)
	newExcludeNotify := addUidToSet(context.ExcludeNotifyUid, excludeNotifyUid)
	return NewVisionContextWithSets(context.Type, context.Param, newExcludeEnterAoi, newExcludeNotify)
}

// addUidToSet 添加UID到集合
func addUidToSet(originalSet map[int64]struct{}, uid int64) map[int64]struct{} {
	if uid <= 0 {
		return originalSet
	}

	if len(originalSet) == 0 {
		newSet := make(map[int64]struct{})
		newSet[uid] = struct{}{}
		return newSet
	}

	newSet := make(map[int64]struct{}, len(originalSet)+1)
	for k, v := range originalSet {
		newSet[k] = v
	}
	newSet[uid] = struct{}{}
	return newSet
}

// GetExcludeEnterAoiUid 获取排除进入AOI的UID列表
func (v *VisionContext) GetExcludeEnterAoiUid() map[int64]struct{} {
	return v.ExcludeEnterAoiUid
}

// GetExcludeNotifyUid 获取排除通知的UID列表
func (v *VisionContext) GetExcludeNotifyUid() map[int64]struct{} {
	return v.ExcludeNotifyUid
}

// GetType 获取视野类型
func (v *VisionContext) GetType() VisionType {
	return v.Type
}

// GetParam 获取参数
func (v *VisionContext) GetParam() int64 {
	return v.Param
}

// 预定义的VisionContext常量
var (
	VisionContextNone          = NewVisionContext(VisionTypeNone)
	VisionContextMeet          = NewVisionContext(VisionTypeMeet)
	VisionContextReborn        = NewVisionContext(VisionTypeReborn)
	VisionContextMiss          = NewVisionContext(VisionTypeMiss)
	VisionContextDie           = NewVisionContext(VisionTypeDie)
	VisionContextReplace       = NewVisionContext(VisionTypeReplace)
	VisionContextReplaceServer = NewVisionContext(VisionTypeReplaceServer)
	VisionContextPickup        = NewVisionContext(VisionTypePickup)
)
