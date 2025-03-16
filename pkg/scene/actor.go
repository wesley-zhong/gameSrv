package scene

// Actor 定义实体接口
type Actor interface {
	GetActorID() int64
	Accept(visitor Visitor)
	GetConfigId() int
	GetActorId() int64
	GetEntityType() ProtEntityType
	SetCoordinate(coord *Coordinate)
	GetCoordinate() *Coordinate
	SetGrid(grid *Grid)
	GetGrid() *Grid
}
