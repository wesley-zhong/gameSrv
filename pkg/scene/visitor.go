package scene

// VisitorType 定义访问者类型
type VisitorType int

const (
	VisitNoneVisitor              VisitorType = iota // 无访问者
	IVisitEntityVisitor                              // 实体访问者
	IVisitAvatarVisitor                              // 玩家访问者
	VisitExcludeSelfAvatarVisitor                    // 排除自身的玩家访问者
)

// Visitor 定义访问者接口
type Visitor interface {
	GetSelfActor() Actor
	GetType() VisitorType
	VisitEntity(actor Actor) int
	GetResultList() []Actor
	GetResultListByType(resultList *[]Actor, clazz interface{})
	CanAddEntity(actor Actor) bool
}

// BaseVisitor 定义基础访问者结构
type BaseVisitor struct {
	selfActor  Actor   // 发起访问的实体
	resultList []Actor // 访问结果列表
}

// NewBaseVisitor 创建新的基础访问者
func NewBaseVisitor(selfActor Actor) *BaseVisitor {
	return &BaseVisitor{
		selfActor:  selfActor,
		resultList: make([]Actor, 0, 10), // 初始化容量为 10
	}
}

// GetSelfActor 获取发起访问的实体
func (v *BaseVisitor) GetSelfActor() Actor {
	return v.selfActor
}

// GetResultList 获取结果实体列表
func (v *BaseVisitor) GetResultList() []Actor {
	return v.resultList
}

// GetResultListByType 在结果中获取某一类型的所有实体
func (v *BaseVisitor) GetResultListByType(resultList *[]Actor, clazz interface{}) {
	for _, actor := range v.resultList {
		// 使用类型断言检查 actor 是否属于 clazz 类型
		if _, ok := actor.(interface{}); ok {
			*resultList = append(*resultList, actor)
		}
	}
}

// VisitEntity 访问一个实体
func (v *BaseVisitor) VisitEntity(actor Actor) int {
	if v.CanAddEntity(actor) {
		v.resultList = append(v.resultList, actor)
	}
	return 0
}

// CanAddEntity 判断是否可以添加实体到结果列表
func (v *BaseVisitor) CanAddEntity(actor Actor) bool {
	// 默认实现，子类可以重写
	return true
}
