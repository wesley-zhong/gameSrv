package controller

import "gameSrv/login/dal"

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
	roleDO := dal.RoleDAO.FindOneById(req.RoleId)
	return &RoleDetailRes{
		RoleId: roleDO.Id,
		Name:   roleDO.Name,
	}
}
