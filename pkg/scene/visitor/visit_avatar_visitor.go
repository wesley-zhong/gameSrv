// Copyright Epic Games, Inc. All Rights Reserved.

package visitor

import (
	"gameSrv/pkg/scene"
)

// VisitAvatarVisitor 头像访问者
type VisitAvatarVisitor struct {
	scene.Visitor
}

// NewVisitAvatarVisitor 创建新的头像访问者
func NewVisitAvatarVisitor(selfActor scene.IEntity) *VisitAvatarVisitor {
	return &VisitAvatarVisitor{
		Visitor: *scene.NewVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitAvatarVisitor) GetType() scene.VisitorType {
	return scene.VisitorTypeAvatarVisitor
}

// canAddEntity 判断是否可以添加实体
func (v *VisitAvatarVisitor) canAddEntity(actor scene.IEntity) bool {
	// TODO: 实现基于实体类型的过滤逻辑
	// return actor.GetActorType() == config.ActorTypeEActorType_Team
	return false
}
