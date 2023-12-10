package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"time"
)

func handShake(ctx tcp.ChannelContext, request proto.Message) {
	validInnerClient := ctx.Context().(*client.ConnInnerClientContext)
	handShake := request.(*protoGen.InnerServerHandShake)
	validInnerClient.ServerId = handShake.FromServerId
	fromServerType := handShake.FromServerType
	client.AddInnerClientConnect(client.GameServerType(fromServerType), validInnerClient)
	log.Infof("client id =%d from serverId=%d  serverType= %d addr =%s handshake finished",
		validInnerClient.Sid, validInnerClient.ServerId, fromServerType, validInnerClient.Ctx.RemoteAddr())
}

func handShake1(ctx tcp.ChannelContext, request **protoGen.InnerServerHandShake) {
	log.Infof("hjahah")

}

func heartBeat(ctx tcp.ChannelContext, request proto.Message) {
	player := ctx.Context().(*player.Player)
	//context := ctx.Context().(*client.ConnClientContext)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" contex= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	//	PlayerMgr.Get()
	//PlayerMgr.GetByContext(context).Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
	player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

func heartBeatResponse(ctx tcp.ChannelContext, request proto.Message) {
	//context := ctx.Context().(*client.ConnInnerClientContext)
	//log.Infof("==== receive sid=%d  addr %s ", context.Sid, ctx.RemoteAddr())
}
