package controller

type Role struct {
}

type RoleDetailReq struct {
	RoleId int64
}

type RoleDetailRes struct {
	RoleId int64
	Name   string
}

func (roleController *Role) GetRoleDetail(req *RoleDetailReq) *RoleDetailRes {

	return &RoleDetailRes{
		RoleId: 111111,
		Name:   "Wesley",
	}
}
