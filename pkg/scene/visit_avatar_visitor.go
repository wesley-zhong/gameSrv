package scene

// VisitAvatarVisitor 头像访问者
type VisitAvatarVisitor struct {
	*Visitor
}

// NewVisitAvatarVisitor 创建新的头像访问者
func NewVisitAvatarVisitor(selfActor IEntity) *VisitAvatarVisitor {
	return &VisitAvatarVisitor{
		Visitor: NewVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitAvatarVisitor) GetType() VisitorType {
	return VisitorTypeAvatarVisitor
}

// canAddEntity 判断是否可以添加实体
func (v *VisitAvatarVisitor) canAddEntity(actor IEntity) bool {
	// TODO: 实现基于实体类型的过滤逻辑
	// return actor.GetActorType() == config.ActorTypeEActorType_Team
	return false
}
