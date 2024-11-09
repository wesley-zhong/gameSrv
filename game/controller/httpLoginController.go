package controller

type Login struct {
}

type LoginInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (Login) Login(loginInfo *LoginInfo) *LoginInfo {
	loginInfo.Password = "iiii"
	return loginInfo

}
