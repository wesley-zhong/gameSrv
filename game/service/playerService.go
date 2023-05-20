package service

import (
	"gameSrv/game/dal"
	"gameSrv/game/module"
	"go.mongodb.org/mongo-driver/bson"
)

func FindPlayerAccount(account string) *module.AccountDO {
	result := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if result == nil {
		return nil
	}
	return result.(*module.AccountDO)
}

func CreatePlayerAccount(account string) *module.AccountDO {
	playerAccount := &module.AccountDO{
		Id:      11,
		Account: account,
	}
	dal.AccountDAO.Insert(playerAccount)
	return nil
}

func AccountLogin(account string) *module.AccountDO {
	player := FindPlayerAccount(account)
	if player == nil {
		player = CreatePlayerAccount(account)
	}
	return player
}

func UpdateAccount(do *module.AccountDO) {
	if do == nil {
		return
	}
	do.Pid = 100001
	//dal.AccountDAO.Save(do.Id, do)
	dal.AccountDAO.AsynSave(do.Id, do)
}
