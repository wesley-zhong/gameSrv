package service

import (
	"gameSrv/login/dal"
	"gameSrv/login/module"
	"go.mongodb.org/mongo-driver/bson"
)

var PlayerService = new(PlayerServiceImpl)

type PlayerServiceImpl struct {
}

func (playerService *PlayerServiceImpl) FindPlayerAccount(account string) *module.AccountDO {
	result := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if result == nil {
		return nil
	}
	return result.(*module.AccountDO)
}

func (playerService *PlayerServiceImpl) CreatePlayerAccount(account string) *module.AccountDO {
	playerAccount := &module.AccountDO{
		Id:      11,
		Account: account,
	}
	dal.AccountDAO.Insert(playerAccount)
	return nil
}

func (playerService *PlayerServiceImpl) AccountLogin(account string) *module.AccountDO {
	player := playerService.FindPlayerAccount(account)
	if player == nil {
		player = playerService.CreatePlayerAccount(account)
	}
	return player
}

func (playerService *PlayerServiceImpl) UpdateAccount(do *module.AccountDO) {
	if do == nil {
		return
	}
	do.Pid = 100001
	//dal.AccountDAO.Save(do.Id, do)
	dal.AccountDAO.AsynSave(do.Id, do)
}
