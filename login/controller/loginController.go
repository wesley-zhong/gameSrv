package controller

import (
	"gameSrv/login/dal"
	"gameSrv/login/module"
	"gameSrv/login/service"
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
		//log.Warnf(" countId = %s not found", loginDto.AccountId)
		return &LoginRes{
			ErrCode: 1,
		}
	}
	return &LoginRes{
		Pid: 11111,
	}
}

func (login *Login) CreateAccount(loginDto *LoginReq) *LoginRes {
	id := utils.NextId()
	accountDO := &module.AccountDO{
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
