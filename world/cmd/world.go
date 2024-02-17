package main

import (
	"fmt"
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/tcp"
	"gameSrv/world/controller"
	"gameSrv/world/dispatcher"
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

	viper.SetConfigName("config")              // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                 // 配置文件格式
	viper.AddConfigPath("/etc/world/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./world/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()             // 查找并读取配置文件
	if err != nil {
		loopWG.Add(-1) // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	clientNetwork := &dispatcher.ClientEventHandler{}
	controller.Init()

	handler := &dispatcher.ServerEventHandler{}
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, &tcp.DefaultCodec{})

	err = discover.InitDiscoverAndRegister(viper.GetViper(), clientNetwork, client.WORLD)
	if err != nil {
		loopWG.Done()
		panic(err)
	}
	loopWG.Wait()
}
