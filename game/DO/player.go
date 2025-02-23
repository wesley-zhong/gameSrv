package DO

type RoleDO struct {
	Id   int64
	Name string
}

type Role struct {
	Id    int64
	CnfId int32
}

type AccountDO struct {
	Id      int64 `bson:"_id"`
	Account string
	Pid     int64
}
