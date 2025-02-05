package main

import (
	"fmt"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/tcp"
	"gameSrv/router/controller"
	"gameSrv/router/dispatcher"
	"gameSrv/router/watcher"
	"github.com/spf13/viper"
	"runtime/debug"
	"sync"
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

	viper.SetConfigName("config")               // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                  // 配置文件格式
	viper.AddConfigPath("/etc/router/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./router/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()              // 查找并读取配置文件
	if err != nil {
		loopWG.Add(-1) // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	controller.Init()
	discover.Init(viper.GetViper(), global.ROUTER)

	handler := &dispatcher.ServerEventHandler{}
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, &tcp.DefaultCodec{})
	err = discover.InitDiscoverAndRegister(viper.GetViper(), watcher.OnDiscoveryServiceChange)
	if err != nil {
		loopWG.Done()
		panic(err)
	}
	loopWG.Wait()
}
