package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(roleId int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("router inner login sessionId = %d = %d from  finished", loginRequest.Sid, loginRequest.PlayerId)

	res := &protoGen.InnerLoginResponse{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.PlayerId,
	}
	client.GetInnerClient(global.GAME).SendInnerMsg(protoGen.InnerProtoCode_INNER_LOGIN_RES, loginRequest.PlayerId, res)
}

func innerPlayerDisconnect(roleId int64, request proto.Message) {

	log.Infof("---roleId = %d disconnected", roleId)
}
