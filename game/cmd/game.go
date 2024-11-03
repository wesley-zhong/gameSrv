package main

import (
	"fmt"
	"gameSrv/game/controller"
	"gameSrv/game/dal"
	"gameSrv/game/dispatcher"
	"gameSrv/game/watcher"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/tcp"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"runtime/debug"
	"sync"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	var loopWG sync.WaitGroup
	loopWG.Add(1)
	defer func() {
		if x := recover(); x != nil {
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s", x, s)
			loopWG.Add(-1)
		}
	}()
	//for performance
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	viper.SetConfigName("config")             // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                // 配置文件格式
	viper.AddConfigPath("/etc/game/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./game/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()            // 查找并读取配置文件
	if err != nil {                        // 处理错误
		panic(fmt.Sprintf("Fatal error configs file: %w \n", err))
	}

	//mongodb init
	dal.InitMongoDB(viper.GetString("mongo.Addr"), viper.GetString("mongo.userName"), viper.GetString("mongo.password"))
	//dal.InitRedisDB(viper.GetString("redis.addr"), viper.GetString("redis.password"))

	// msg Register
	controller.Init()
	discover.Init(viper.GetViper())

	//start server
	serverNetworkHandler := &dispatcher.ServerEventHandler{}
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), serverNetworkHandler, &tcp.DefaultCodec{})

	//init tcp client
	clientHandler := &dispatcher.ClientEventHandler{}
	tcp.ClientStart(clientHandler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay))

	////register to etcd
	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange, global.GAME)
	if err != nil {
		panic(err)
		return
	}
	// start http server
	//httpServer := web.NewHttpServer()
	//httpServer.HttpMethod.RegisterController()
	loopWG.Wait()
}
