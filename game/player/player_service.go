package player

import (
	"gameSrv/game/DO"
	"gameSrv/game/dal"
	"go.mongodb.org/mongo-driver/bson"
)

func FindPlayerAccount(account string) *DO.AccountDO {
	result := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if result == nil {
		return nil
	}
	return result
}

func CreatePlayerAccount(account string) *DO.AccountDO {
	playerAccount := &DO.AccountDO{
		Id:      11,
		Account: account,
	}
	dal.AccountDAO.Insert(playerAccount)
	return nil
}

func AccountLogin(account string) *DO.AccountDO {
	player := FindPlayerAccount(account)
	if player == nil {
		player = CreatePlayerAccount(account)
	}
	return player
}

func LoadRoleFromDB(roleId int64) *DO.RoleDO {
	roleDO := dal.RoleDAO.FindOneById(roleId)
	if roleDO == nil {
		return nil
	}
	return roleDO

}

func UpdateAccount(do *DO.AccountDO) {
	if do == nil {
		return
	}
	do.Pid = 100001
	//dal.AccountDAO.Save(dos.Id, dos)
	dal.AccountDAO.AsynSave(do.Id, do)
}
