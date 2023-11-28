package main

import (
	"fmt"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/network"
	"gameSrv/world/controller"
	"gameSrv/world/networkHandler"
	"github.com/spf13/viper"
	"sync"
)

func main() {
	var loopWG sync.WaitGroup
	loopWG.Add(1)
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
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

	// no need to connect other server  so do not need init
	//clientNetwork := networkHandler.ClientEventHandler{}
	//network.ClientStart(&clientNetwork,
	//	gnet.WithMulticore(true),
	//	gnet.WithReusePort(true),
	//	gnet.WithTCPNoDelay(0))

	controller.Init()

	handler := &networkHandler.ServerEventHandler{}
	go network.ServerStartWithDeCode(viper.GetInt32("port"), handler, network.NewInnerLengthFieldBasedFrameCodecEx())

	err = discover.InitDiscoverAndRegister(viper.GetViper())
	if err != nil {
		loopWG.Done()
		panic(err)
	}
	loopWG.Wait()
}
