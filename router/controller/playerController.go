package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

func innerPlayerLogin(pid int64, request proto.Message) {
	loginRequest := request.(*protoGen.InnerLoginRequest)
	log.Infof("router inner login sessionId = %d = %d from  finished", loginRequest.Sid, loginRequest.RoleId)

	res := &protoGen.InnerLoginResponse{
		Sid:    loginRequest.Sid,
		RoleId: loginRequest.RoleId,
	}
	client.SendInnerToGameServer(loginRequest.RoleId, protoGen.InnerProtoCode_INNER_LOGIN_RES, res)
}

func innerPlayerDisconnect(pid int64, request proto.Message) {

	log.Infof("---pid = %d disconnected", pid)
}
