package module

type RoleDO struct {
	Id   int64
	Name string
}

type AccountDO struct {
	Id      int64 `bson:"_id"`
	Account string
	Pid     int64
}

type ItemDO struct {
	Id      int64
	Account string
	Pid     int64
}
