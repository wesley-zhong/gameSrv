package controller

type RoleController struct {
}

type RoleDetailReq struct {
	RoleId int64
}

type RoleDetailRes struct {
	RoleId int64
	Name   string
}

func (roleController *RoleController) getRoleDetail(req *RoleDetailReq) *RoleDetailRes {

	return &RoleDetailRes{
		RoleId: 111111,
		Name:   "Wesley",
	}
}
