package controller

type LoginController struct {
}

type LoginReq struct {
	AccountId string
	Password  string
}

type LoginRes struct {
	Pid int64
}

func (login *LoginController) Login(loginDto *LoginReq) *LoginRes {
	return &LoginRes{
		Pid: 11111,
	}
}
