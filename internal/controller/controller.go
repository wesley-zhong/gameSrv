package controller

import (
	"gameSvr/internal/client"
	"gameSvr/internal/message"
	"gameSvr/internal/player"
	"gameSvr/pkg/core"
	"gameSvr/pkg/log"
	"gameSvr/pkg/network"
	protoGen "gameSvr/proto"
	"time"

	"google.golang.org/protobuf/proto"
)

var gameInnerClient *client.ConnInnerClientContext

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	core.RegisterMethod(int32(-6), &protoGen.InnerLoginResponse{}, loginResponseFromGameServer)
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	core.RegisterMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)

	context, err := network.Dial("tcp", "127.0.0.1:9007")
	if err != nil {
		log.Error(err)
		return
	}
	gameInnerClient = client.NewInnerClientContext(context)
	//add  msg  to game server to add me
	header := &protoGen.InnerHead{
		FromServerUid:    message.BuildServerUid(message.TypeGateway, 35),
		ToServerUid:      0,
		ReceiveServerUid: 0,
		Id:               0,
		SendType:         0,
		ProtoCode:        message.INNER_PROTO_ADD_SERVER,
		CallbackId:       0,
	}

	innerMessage := &client.InnerMessage{
		InnerHeader: header,
		Body:        nil,
	}
	gameInnerClient.Send(innerMessage)
}

var PlayerMgr = player.NewPlayerMgr() //make(map[int64]network.ChannelContext)

func login(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnClientContext)
	loginRequest := request.(*protoGen.LoginRequest)
	log.Infof("login token = %s id = %d", loginRequest.LoginToken, loginRequest.RoleId)
	innerLoginReq := &protoGen.InnerLoginRequest{
		SessionId: context.Sid,
		AccountId: loginRequest.AccountId,
		RoleId:    loginRequest.RoleId,
	}
	msgHeader := &protoGen.InnerHead{
		FromServerUid:    message.BuildServerUid(message.TypeGateway, 35),
		ToServerUid:      message.BuildServerUid(message.TypeGame, 35),
		ReceiveServerUid: 0,
		Id:               loginRequest.RoleId,
		SendType:         0,
		ProtoCode:        message.INNER_PROTO_LOGIN_REQUEST,
		CallbackId:       0,
	}

	innerMsg := &client.InnerMessage{
		InnerHeader: msgHeader,
		Body:        innerLoginReq,
	}
	gameInnerClient.Send(innerMsg)
	//PlayerContext[loginRequest.RoleId] = ctx
	player := player.NewPlayer(loginRequest.GetRoleId(), context)
	PlayerMgr.Add(player)
	context.Ctx.SetContext(player)
}

func loginResponseFromGameServer(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	innerLoginResponse := request.(*protoGen.InnerLoginResponse)
	log.Infof("login response = %d  sid =%d", innerLoginResponse.RoleId, context.Sid)
	response := &protoGen.LoginResponse{
		ErrorCode:  0,
		ServerTime: time.Now().UnixMilli(),
	}
	//marshal, err := proto.Marshal(response)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//body := make([]byte, len(marshal)+8)
	//writeBuffer := bytes.NewBuffer(body)
	//writeBuffer.Reset()
	//binary.Write(writeBuffer, binary.BigEndian, int32(protoGen.ProtoCode_LOGIN_RESPONSE))
	//binary.Write(writeBuffer, binary.BigEndian, int32(len(marshal)))
	//binary.Write(writeBuffer, binary.BigEndian, marshal)
	//PlayerMgr.GetByRoleId(innerLoginResponse.RoleId).Context.Ctx.AsyncWrite(writeBuffer.Bytes())

	PlayerMgr.GetByRoleId(innerLoginResponse.RoleId).Context.Send(int32(protoGen.ProtoCode_LOGIN_RESPONSE), response)
}

func heartBeat(ctx network.ChannelContext, request proto.Message) {
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

func innerServerKickout(ctx network.ChannelContext, request proto.Message) {
	context := ctx.Context().(*client.ConnInnerClientContext)
	kickOut := request.(*protoGen.KickOutResponse)
	log.Infof("login response = %d  sid =%d", kickOut.Reason, context.Sid)
}
