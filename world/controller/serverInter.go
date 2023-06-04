package controller

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

func handShake(ctx network.ChannelContext, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	handShake := request.(*protoGen.InnerServerHandShake)
	validInnerClient.ServerId = handShake.FromServerId
	fromServerType := handShake.FromServerType
	client.AddInnerClientConnect(client.GameServerType(fromServerType), validInnerClient)
	log.Infof("client id =%d from serverId=%d  serverType= %d addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, fromServerType, validInnerClient.Ctx.RemoteAddr())
}

func innerHeartBeat(ctx network.ChannelContext, request proto.Message) {
	innerClient := ctx.Context().(*client.ConnInnerClientContext)
	response := &protoGen.InnerHeartBeatResponse{}
	innerClient.SendInnerMsg(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), 0, response)
	log.Infof(" addr =%s  heartbeat time", ctx.RemoteAddr())
}

func heartBeatResponse(ctx network.ChannelContext, request proto.Message) {
	//context := ctx.Context().(*client.ConnInnerClientContext)
	//log.Infof("==== receive sid=%d  addr %s ", context.Sid, ctx.RemoteAddr())
}
