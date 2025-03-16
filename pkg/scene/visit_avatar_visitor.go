package scene

// VisitAvatarVisitor 定义玩家访问者
type VisitAvatarVisitor struct {
	*BaseVisitor // 嵌入基础访问者
}

// NewVisitAvatarVisitor 创建新的玩家访问者
func NewVisitAvatarVisitor(selfActor Actor) *VisitAvatarVisitor {
	return &VisitAvatarVisitor{
		BaseVisitor: NewBaseVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitAvatarVisitor) GetType() VisitorType {
	return IVisitAvatarVisitor
}

// CanAddEntity 判断是否可以添加实体到结果列表
func (v *VisitAvatarVisitor) CanAddEntity(actor Actor) bool {
	return actor.GetEntityType() == ProtEntityAvatar
}
