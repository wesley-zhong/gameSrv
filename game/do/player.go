package do

type RoleDO struct {
	Id   int64
	Name string
}

type AccountDO struct {
	Id      int64 `bson:"_id"`
	Account string
	Pid     int64
}
