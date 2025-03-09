package DO

type ItemDO struct {
	Id      int64 `bson:"_id"`
	Account string
}

type Item struct {
	Id    int64 `bson:"_id"`
	CnfId int32
	//properties
}
