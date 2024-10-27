package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func handShakeReq(ctx tcp.Channel, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	handShake := request.(*protoGen.InnerServerHandShakeReq)
	validInnerClient.ServerId = handShake.FromServerId
	fromServerType := handShake.FromServerType
	client.AddInnerClientConnect(client.GameServerType(fromServerType), validInnerClient)
	log.Infof("client id =%d from serverId=%d  serverId= %s addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, handShake.FromServerSid, validInnerClient.Ctx.RemoteAddr())
}
