package scene

type VisitorType int32

const (
	VisitorTypeNoneVisitor              VisitorType = iota // 无访问者
	VisitorTypeEntityVisitor                               // 观察所有实体 用于玩家视野
	VisitorTypeAvatarVisitor                               // 只观察avatar 用于怪物身上
	VisitorTypeExcludeSelfAvatarVisitor                    // 排除自身avatar访问者
)

// IVisitor 访问者（Visitor设计模式）
type IVisitor interface {
	GetSelfEntity() IEntity
	GetType() VisitorType
	VisitEntity(actor IEntity)
	GetResultList() []IEntity
}

// Visitor 访问者
type Visitor struct {
	selfActor  IEntity   // 发起访问的实体，获取其周边相关的entity
	resultList []IEntity // 访问结果列表
}

// VisitorResultListInitSize 访问结果列表初始大小
const VisitorResultListInitSize = 10

// NewVisitor 创建新的访问者
func NewVisitor(selfActor IEntity) *Visitor {
	return &Visitor{
		selfActor:  selfActor,
		resultList: make([]IEntity, 0, VisitorResultListInitSize),
	}
}

// GetSelfEntity 获取自身实体
func (v *Visitor) GetSelfEntity() IEntity {
	return v.selfActor
}

// GetType 获取访问者类型（抽象方法）
func (v *Visitor) GetType() VisitorType {
	return VisitorTypeNoneVisitor
}

// VisitEntity 访问一个实体
func (v *Visitor) VisitEntity(actor IEntity) {
	if v.canAddEntity(actor) {
		v.resultList = append(v.resultList, actor)
	}
}

// GetResultList 获取结果实体列表
func (v *Visitor) GetResultList() []IEntity {
	return v.resultList
}

// canAddEntity 判断是否可以添加实体（抽象方法）
func (v *Visitor) canAddEntity(actor IEntity) bool {
	return false
}
