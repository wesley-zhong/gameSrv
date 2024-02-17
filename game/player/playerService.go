package player

import (
	"gameSrv/game/dal"
	"gameSrv/game/do"
	"go.mongodb.org/mongo-driver/bson"
)

func FindPlayerAccount(account string) *do.AccountDO {
	result := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if result == nil {
		return nil
	}
	return result.(*do.AccountDO)
}

func CreatePlayerAccount(account string) *do.AccountDO {
	playerAccount := &do.AccountDO{
		Id:      11,
		Account: account,
	}
	dal.AccountDAO.Insert(playerAccount)
	return nil
}

func AccountLogin(account string) *do.AccountDO {
	player := FindPlayerAccount(account)
	if player == nil {
		player = CreatePlayerAccount(account)
	}
	return player
}

func FindRoleData(roleId int64) *do.RoleDO {
	roleDO := dal.RoleDAO.FindOneById(roleId)
	if roleDO == nil {
		return nil
	}
	return roleDO

}

func UpdateAccount(do *do.AccountDO) {
	if do == nil {
		return
	}
	do.Pid = 100001
	//dal.AccountDAO.Save(do.Id, do)
	dal.AccountDAO.AsynSave(do.Id, do)
}
