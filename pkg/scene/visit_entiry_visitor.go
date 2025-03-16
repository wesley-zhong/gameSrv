package scene

// VisitEntityVisitor 定义实体访问者
type VisitEntityVisitor struct {
	*BaseVisitor // 嵌入基础访问者
}

// NewVisitEntityVisitor 创建新的实体访问者
func NewVisitEntityVisitor(selfActor Actor) *VisitEntityVisitor {
	return &VisitEntityVisitor{
		BaseVisitor: NewBaseVisitor(selfActor),
	}
}

// GetType 获取访问者类型
func (v *VisitEntityVisitor) GetType() VisitorType {
	return IVisitEntityVisitor
}

// CanAddEntity 判断是否可以添加实体到结果列表
func (v *VisitEntityVisitor) CanAddEntity(actor Actor) bool {
	//entityType := actor.GetEntityType()
	//switch entityType {
	//case ProtEntityMonster, ProtEntityNPC, ProtEntityGadget, ProtEntityWeather:
	//	return true
	//case ProtEntityAvatar:
	//	if v.selfActor == nil {
	//		return false
	//	}
	//	selfEntityType := v.selfActor.GetEntityType()
	//	if selfEntityType == ProtEntityAvatar {
	//		return !actor.Equals(v.selfActor)
	//	}
	//	return true
	//case ProtEntityEyePoint:
	//	if v.selfActor == nil {
	//		return false
	//	}
	//	selfEntityType := v.selfActor.GetEntityType()
	//	if selfEntityType == ProtEntityAvatar {
	//		return actor.GetPlayer() != v.selfActor.GetPlayer()
	//	} else if selfEntityType == ProtEntityEyePoint {
	//		return !actor.Equals(v.selfActor)
	//	}
	//	return true
	//default:
	//	return false
	//}
	return false
}
