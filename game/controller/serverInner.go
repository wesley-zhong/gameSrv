package controller

import (
	"gameSrv/pkg/aresTcpClient"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func handShakeReq(ctx tcp.Channel, request proto.Message) {
	validInnerClient := ctx.Context().(*aresTcpClient.ConnInnerClientContext)
	handShake := request.(*protoGen.InnerServerHandShakeReq)
	validInnerClient.ServerId = handShake.FromServerId
	fromServerType := handShake.FromServerType
	aresTcpClient.AddInnerClientConnect(global.GameServerType(fromServerType), validInnerClient)
	log.Infof("=====  handShakeReq ===== %s", request)
	log.Infof("aresTcpClient id =%d from serverId=%d  serverId= %s addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, handShake.FromServerSid, validInnerClient.Ctx.RemoteAddr())

	res := &protoGen.InnerServerHandShakeRes{
		FromServerId:   0,
		FromServerSid:  discover.MySelfNode.ServiceId,
		FromServerType: int32(discover.MySelfNode.Type),
	}
	validInnerClient.SendInnerMsg(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_RES, 0, res)
}

func handShakeRes(ctx tcp.Channel, request proto.Message) {
	handShake := request.(*protoGen.InnerServerHandShakeRes)
	validInnerClient := ctx.Context().(*aresTcpClient.ConnInnerClientContext)
	log.Infof("handShake response finished from serverId=%s  addr =%s handshake finished  ",
		handShake.FromServerSid, validInnerClient.Ctx.RemoteAddr())
}

func heartBeatRes(ctx tcp.Channel, request proto.Message) {
	log.Infof(" inner  heartBeatRes context= %s ", ctx.RemoteAddr())
}
