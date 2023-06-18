
GameSrv is a lightweight online game framework  written with Golang.Inlcude 3 types server： **Gateway**， **GameSrv**  and  **World** .It can be used in many types of online game.Tcp network use [gnet](https://github.com/panjf2000/gnet)

**Build**
---
**Linux**

```
git clone https://github.com/wesley-zhong/gameSrv.git
cd gameSrv
make  clean/build
```
**GoLand**

Open directly

**Run And  Test**
---
```
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

controller.Init()

```
//register client msg process handler
func Init() {
	core.RegisterMethod(int32(protoGen.ProtoCode_LOGIN_REQUEST), &protoGen.LoginRequest{}, login)
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
		
//connect server
client.InnerClientConnect(client.GAME, viper.GetString("gameServerAddr"), client.GATE_WAY)
```