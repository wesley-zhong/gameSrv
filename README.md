
GameSrv is a lightweight online game framework  written with Golang.Inlcude 3 types server： **Gateway**， **GameSrv**  and  **World** .It can be used in many types of online game.Tcp network use [gnet](https://github.com/panjf2000/gnet)

**Build**
---
**Linux**

```
git clone https://github.com/wesley-zhong/gameSrv.git
cd gameSrv
make  clean/build
```
**Windows**

Open directly with Goland  directly or VsCode

**Run And  Test**
---
```
//only run linux or mac os
make run  //start the server

//start client
cd build/client
./client  
```

**How To Use**
---

start server
```
// msg Register
controller.Init()

//package receive handler
handler := &networkHandler.ServerEventHandler{}
	
//start server
network.ServerStart(viper.GetInt32("port"), handler)
```

register client msg process handler

```
func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
	core.RegisterMethod(int32(protoGen.ProtoCode_HEART_BEAT_REQUEST), &protoGen.HeartBeatRequest{}, heartBeat)
}
	
```
function implement
```
func login(ctx network.ChannelContext, request proto.Message) {
// logic code here

// wrap inner msg and inner msgId to  send to game server 
log.Infof("====== loginAddr=%s now loginCount =%d", ctx.RemoteAddr(), PlayerMgr.GetSize())
client.GetInnerClient(client.GAME).SendInnerMsgProtoCode(protoGen.InnerProtoCode_INNER_LOGIN_REQ, existPlayer.Pid, innerRequest)
}

func heartBeat(ctx network.ChannelContext, request proto.Message) {
	player := ctx.Context().(*player.Player)
	heartBeat := request.(*protoGen.HeartBeatRequest)
	log.Infof(" context= %d  heartbeat time = %d", player.Context.Sid, heartBeat.ClientTime)

	response := &protoGen.HeartBeatResponse{
		ClientTime: heartBeat.ClientTime,
		ServerTime: time.Now().UnixMilli(),
	}
	//send to client
	player.Context.Send(int32(protoGen.ProtoCode_HEART_BEAT_RESPONSE), response)
}

```

InnerClient : server's tcp connection
```
//init InnerClient
clientNetwork := networkHandler.ClientEventHandler{}
network.ClientStart(&clientNetwork,
	gnet.WithMulticore(true),
	gnet.WithReusePort(true),
	gnet.WithTCPNoDelay(0),
	gnet.WithTicker(true),
	gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))
		
//build a connection from GateWay(client.GATE_WAY) to Game(client.GAME)
client.InnerClientConnect(client.GAME, viper.GetString("gameServerAddr"), client.GATE_WAY)
```
**Deploy**

*Docker*
```
  cd build
  make  docker_build
  make docker_run
```