package service

import (
	"gameSrv/login/dal"
	"gameSrv/login/dos"
	"gameSrv/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var PlayerService = new(PlayerServiceImpl)

type PlayerServiceImpl struct {
}

func (playerService *PlayerServiceImpl) FindPlayerAccount(account string) *dos.AccountDO {
	result, err := dal.AccountDAO.FindOne(bson.D{{"account", account}})
	if err != nil || result == nil {
		return nil
	}
	return result
}

func (playerService *PlayerServiceImpl) CreatePlayerAccount(account string) *dos.AccountDO {
	pid := utils.NextId()
	playerAccount := &dos.AccountDO{
		Id:      pid,
		Account: account,
		Pid:     pid,
	}
	dal.AccountDAO.Insert(playerAccount)
	return playerAccount
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
