package controller

import (
	"gameSrv/gateway/player"
	"gameSrv/pkg/client"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"time"

	"google.golang.org/protobuf/proto"
)

var playerConn = make(map[int64]*client.ConnClientContext)

func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
	core.RegisterMethod(int32(protoGen.ProtoCode_KICK_OUT_RESPONSE), &protoGen.KickOutResponse{}, innerServerKickout)
	core.RegisterMethod(int32(protoGen.ProtoCode_PERFORMANCE_TEST_RES), &protoGen.PerformanceTestRes{}, performanceRes)

	go startConnection(1000)

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

func performanceRes(ctx network.ChannelContext, res proto.Message) {
	performanceRes := res.(*protoGen.PerformanceTestRes)
	log.Infof("------response id =%d", performanceRes.SomeIdAdd)
}

func startConnection(count int) {
	for i := 0; i < count; i++ {
		client := client.ClientConnect("124.222.26.216:9001")
		//add  msg  to game server to add me
		playerConn[client.Sid] = client

	}

	req := &protoGen.PerformanceTestReq{
		SomeId:   2,
		SomeBody: "hello",
	}
	timer := time.NewTimer(3 * time.Second)

	go func() {
		for i := 0; i < 10000; i++ {
			for {
				timer.Reset(100 * time.Millisecond)
				select {
				case <-timer.C:
					for _, v := range playerConn {
						v.Send(int32(protoGen.ProtoCode_PERFORMANCE_TEST_REQ), req)
					}
				}
			}
		}
	}()

}
