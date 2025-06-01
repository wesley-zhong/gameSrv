package main

import (
	"fmt"
	"gameSrv/gateway/dispathcer"
	"gameSrv/gateway/watcher"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/tcp"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"sync"

	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")                // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                   // 配置文件格式
	viper.AddConfigPath("/etc/gateway/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./gateway/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()               // 查找并读取配置文件
	if err != nil {
		panic(err)
	}
	global.SelfServerType = global.GATE_WAY
}

func main() {
	var loopWG sync.WaitGroup
	loopWG.Add(1)
	defer func() {
		if x := recover(); x != nil {
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s", x, s)
			loopWG.Done()
		}
	}()

	//for performance
	go func() {
		err := http.ListenAndServe("localhost:6062", nil)
		if err != nil {
			panic(err)
		}
	}()

	//package receive handler
	handler := &dispathcer.ServerEventHandler{}

	//start server
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, &tcp.DefaultCodec{})

	//init tcp aresTcpClient
	err := discover.Init(viper.GetViper(), global.GATE_WAY)
	if err != nil {
		panic(err)
	}
	clientHandler := &dispathcer.ClientEventHandler{}
	err = tcp.ClientStart(clientHandler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay))
	if err != nil {
		panic(err)
	}

	//register to etcd
	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange)
	if err != nil {
		panic(err)
	}
	loopWG.Wait()
}
