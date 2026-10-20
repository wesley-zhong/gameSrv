package main

import (
	"fmt"
	_ "gameSrv/gateway/constants"
	_ "gameSrv/gateway/controller"
	"gameSrv/gateway/network"
	"gameSrv/gateway/watcher"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/pkg/utils"
	"net/http"
	_ "net/http/pprof"

	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func main() {
	utils.IdGenInit(11, 23)

	//for performance
	go func() {
		http.ListenAndServe("localhost:6062", nil)
	}()

	viper.SetConfigName("config")                // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                   // 配置文件格式
	viper.AddConfigPath("/etc/gateway/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./gateway/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()               // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	log.Init("./log/gateway.log", zapcore.InfoLevel, true)

	//package receive handler
	handler := &network.ServerEventHandler{}
	discover.Init(viper.GetViper(), global.GATE_WAY)

	//start server
	go func() {
		tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, &tcp.DefaultCodec{})
	}()

	//init tcp client
	clientHandler := &network.ClientEventHandler{}
	tcp.ClientStart(clientHandler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay))

	////register to etcd
	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange)
	if err != nil {
		panic(err)
	}

	// block forever; servers run in background goroutines
	select {}
}
