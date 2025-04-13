package main

import (
	"fmt"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/tcp"
	_ "gameSrv/router/controller"
	"gameSrv/router/dispatcher"
	"gameSrv/router/watcher"
	"github.com/spf13/viper"
	"runtime/debug"
	"sync"
)

func init() {
	viper.SetConfigName("config")               // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                  // 配置文件格式
	viper.AddConfigPath("/etc/router/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./router/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()              // 查找并读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}
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

	// start serer
	handler := &dispatcher.ServerEventHandler{}
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, &tcp.DefaultCodec{})

	//register myself
	err := discover.Init(viper.GetViper(), global.ROUTER)
	if err != nil {
		panic(err)
	}

	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange)
	if err != nil {
		panic(err)
	}
	loopWG.Wait()
}
