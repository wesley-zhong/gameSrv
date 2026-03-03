package service

import (
	"gameSrv/login/dal"
	"gameSrv/login/dos"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var PlayerService = new(PlayerServiceImpl)

type PlayerServiceImpl struct {
}

func (playerService *PlayerServiceImpl) FindPlayerAccount(account string) *dos.AccountDO {
	result := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if result == nil {
		return nil
	}
	return result.(*dos.AccountDO)
}

func (playerService *PlayerServiceImpl) CreatePlayerAccount(account string) *dos.AccountDO {
	playerAccount := &dos.AccountDO{
		Id:      11,
		Account: account,
	}
	dal.AccountDAO.Insert(playerAccount)
	return nil
}

func (playerService *PlayerServiceImpl) AccountLogin(account string) *dos.AccountDO {
	player := playerService.FindPlayerAccount(account)
	if player == nil {
		player = playerService.CreatePlayerAccount(account)
	}
	return player
}

func (playerService *PlayerServiceImpl) UpdateAccount(do *dos.AccountDO) {
	if do == nil {
		return
	}
	do.Pid = 100001
	//dal.AccountDAO.Save(dos.Id, dos)
	dal.AccountDAO.AsynSave(do.Id, do)
}

func createToken(*dos.AccountDO) {
	//discover.DiscoverService.

}
