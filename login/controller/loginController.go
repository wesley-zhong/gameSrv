package controller

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
	return &LoginRes{
		Pid: 11111,
	}
}
