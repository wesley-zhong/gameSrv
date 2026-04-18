// Copyright Epic Games, Inc. All Rights Reserved.

package scene

// VisitExcludeSelfAvatarVisitor 排除自身头像访问者
type VisitExcludeSelfAvatarVisitor struct {
	*Visitor
}

// NewVisitExcludeSelfAvatarVisitor 创建新的排除自身头像访问者
func NewVisitExcludeSelfAvatarVisitor(selfActor IEntity) *VisitExcludeSelfAvatarVisitor {
	return &VisitExcludeSelfAvatarVisitor{
		Visitor: NewVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitExcludeSelfAvatarVisitor) GetType() VisitorType {
	return VisitorTypeExcludeSelfAvatarVisitor
}

// canAddEntity 判断是否可以添加实体
func (v *VisitExcludeSelfAvatarVisitor) canAddEntity(actor IEntity) bool {
	// TODO: 实现基于实体类型的过滤逻辑
	// return actor.GetActorType() == config.ActorTypeEActorType_Avatar &&
	// 	v.selfActor != nil &&
	// 	!v.selfActor.Equals(actor)
	return false
}
