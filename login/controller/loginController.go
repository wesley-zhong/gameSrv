package controller

import (
	"gameSrv/login/dal"
	"gameSrv/login/dos"
	"gameSrv/login/service"
	"gameSrv/pkg/log"
	"gameSrv/pkg/utils"
)

type Login struct {
}

type LoginReq struct {
	AccountId string
	Password  string
}

type LoginRes struct {
	Pid     int64
	ErrCode int32
}

func (login *Login) Login(loginDto *LoginReq) *LoginRes {
	account := service.PlayerService.FindPlayerAccount(loginDto.AccountId)
	if account == nil {
		log.Warnf(" countId = %s not found will create new ", loginDto.AccountId)
		accoutDO := service.PlayerService.CreatePlayerAccount(loginDto.AccountId)
		return &LoginRes{
			Pid: accoutDO.Pid,
		}
	}
	return &LoginRes{
		Pid: account.Pid,
	}
}

func (login *Login) CreateAccount(loginDto *LoginReq) *LoginRes {
	id := utils.NextId()
	accountDO := &dos.AccountDO{
		Id:      id,
		Account: loginDto.AccountId,
		Pid:     id,
	}
	err := dal.AccountDAO.Insert(accountDO)
	if err != nil {
		return nil
	}
	return &LoginRes{Pid: accountDO.Id}
}
