// Copyright Epic Games, Inc. All Rights Reserved.

package visitor

import (
	"gameSrv/pkg/scene"
)

// VisitExcludeSelfAvatarVisitor 排除自身头像访问者
type VisitExcludeSelfAvatarVisitor struct {
	scene.Visitor
}

// NewVisitExcludeSelfAvatarVisitor 创建新的排除自身头像访问者
func NewVisitExcludeSelfAvatarVisitor(selfActor scene.IEntity) *VisitExcludeSelfAvatarVisitor {
	return &VisitExcludeSelfAvatarVisitor{
		Visitor: *scene.NewVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitExcludeSelfAvatarVisitor) GetType() scene.VisitorType {
	return scene.VisitorTypeExcludeSelfAvatarVisitor
}

// canAddEntity 判断是否可以添加实体
func (v *VisitExcludeSelfAvatarVisitor) canAddEntity(actor scene.IEntity) bool {
	// TODO: 实现基于实体类型的过滤逻辑
	// return actor.GetActorType() == config.ActorTypeEActorType_Avatar &&
	// 	v.selfActor != nil &&
	// 	!v.selfActor.Equals(actor)
	return false
}
