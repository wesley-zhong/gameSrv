package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(roleId int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("world inner login sessionId = %d = %d from  finished", loginRequest.Sid, loginRequest.RoleId)

	res := &protoGen.InnerLoginResponse{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.GetInnerClient(client.GAME).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, loginRequest.RoleId, res)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {
	log.Infof("---roleId = %d disconnected", roleId)
}
