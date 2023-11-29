package main

import (
	"fmt"
	"gameSrv/client/controller"
	"gameSrv/client/networkHandler"
	"gameSrv/pkg/network"
	"github.com/panjf2000/gnet"
	"github.com/spf13/viper"
	"sync"
)

func main() {
	wait := sync.WaitGroup{}
	wait.Add(1)
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	//load configs
	viper.SetConfigName("config")               // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                  // 配置文件格式
	viper.AddConfigPath("/etc/client/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./client/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil { // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	clientNetwork := networkHandler.ClientNetwork{}
	network.ClientStart(&clientNetwork,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithCodec(network.NewLengthFieldBasedFrameCodecEx()))

	controller.Init()
	wait.Wait()
}
