package main

import (
	"fmt"
	"gameSrv/game/controller"
	"gameSrv/game/dal"
	"gameSrv/game/network"
	"gameSrv/game/watcher"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/pkg/web"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"runtime/debug"
	"sync"

	"net/http"
	_ "net/http/pprof"
)

func init() {
	log.Infof("====")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./game/configs/") // 查找配置文件的路径

	viper.SetConfigName("config")             // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                // 配置文件格式
	viper.AddConfigPath("/etc/game/configs/") // 查找配置文件的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理错误
		panic(fmt.Sprintf("Fatal error configs file: %w \n", err))
	}
}

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
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			panic(err)
		}
	}()

	//mongodb init
	err := dal.InitMongoDB(viper.GetString("mongo.Addr"), viper.GetString("mongo.userName"), viper.GetString("mongo.password"))
	if err != nil {
		panic(err)
	}
	//dal.InitRedisDB(viper.GetString("redis.addr"), viper.GetString("redis.password"))

	//start server
	serverNetworkHandler := &network.ServerEventHandler{}
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), serverNetworkHandler, &tcp.DefaultCodec{})
	//init tcp client
	clientHandler := &network.ClientEventHandler{}
	err = tcp.ClientStart(clientHandler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay))
	if err != nil {
		panic(err)
	}

	// start http server
	httpServer := web.NewHttpServer()
	loginController := &controller.Login{}
	httpServer.HttpMethod.RegisterController(loginController)
	go httpServer.WebAppStart(7788)

	////register to etcd
	// msg Register
	discover.Init(viper.GetViper(), global.GAME)
	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange)
	if err != nil {
		panic(err)
		return
	}

	loopWG.Wait()
}
