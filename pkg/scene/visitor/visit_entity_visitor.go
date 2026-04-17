// Copyright Epic Games, Inc. All Rights Reserved.

package visitor

import (
	"gameSrv/pkg/scene"
)

// VisitEntityVisitor 实体访问者
// 查找附近的entity
// selfEntity一般是avatar，用于avatar在场景中移动（包括加入，离开场景）
type VisitEntityVisitor struct {
	scene.Visitor
}

// NewVisitEntityVisitor 创建新的实体访问者
func NewVisitEntityVisitor(selfActor scene.IEntity) *VisitEntityVisitor {
	return &VisitEntityVisitor{
		Visitor: *scene.NewVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitEntityVisitor) GetType() scene.VisitorType {
	return scene.VisitorTypeEntityVisitor
}

// canAddEntity 判断是否可以添加实体
func (v *VisitEntityVisitor) canAddEntity(actor scene.IEntity) bool {
	// TODO: 实现基于实体类型的过滤逻辑
	// entityType := actor.GetActorType()

	// switch entityType {
	// case config.ActorTypeEActorType_Monster:
	// 	config.ActorTypeEActorType_NPC:
	// 	...
	// 	return true
	// case config.ActorTypeEActorType_Team:
	// 	// 观察team时要过滤掉观察者自身
	// 	return !actor.Equals(v.selfActor)
	// case config.ActorTypeEActorType_Avatar:
	// 	// 观察eye_point时要过滤掉属于观察者的eye_point
	// 	if v.selfActor.GetActorType() == config.ActorTypeEActorType_Team {
	// 		return actor.GetOwner() != v.selfActor.GetOwner()
	// 	}
	// 	// 观察eye_point时要过滤掉观察者自身
	// 	return !actor.Equals(v.selfActor)
	// default:
	// 	return false
	// }

	return false
}
