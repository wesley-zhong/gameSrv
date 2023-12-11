package controller

import (
	"gameSrv/login/service"
	"gameSrv/pkg/log"
)

type Login struct {
}

type LoginReq struct {
	AccountId string
	Password  string
}

type LoginRes struct {
	Pid int64
}

func (login *Login) Login(loginDto *LoginReq) *LoginRes {
	account := service.FindPlayerAccount(loginDto.AccountId)
	if account == nil {
		log.Warnf(" countId = %s not found", loginDto.AccountId)
		return nil
	}
	return &LoginRes{
		Pid: 11111,
	}
}
