package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
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

	res := &protoGen.InnerServerHandShakeRes{
		FromServerId:   0,
		FromServerSid:  discover.MySelfNode.ServiceId,
		FromServerType: int32(discover.MySelfNode.Type),
	}
	validInnerClient.SendInnerMsg(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_RES, 0, res)
}

func handShakeResp(ctx tcp.Channel, request proto.Message) {
	handShake := request.(*protoGen.InnerServerHandShakeRes)
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	log.Infof("handShake finished from serverId=%s  addr =%s handshake finished  ",
		handShake.FromServerSid, validInnerClient.Ctx.RemoteAddr())
}

func heartBeatResponse(ctx tcp.Channel, request proto.Message) {
	log.Infof(" inner  heartBeatResponse context= %s ", ctx.RemoteAddr())
}
