package main

import (
	"fmt"
	"gameSrv/gateway/controller"
	"gameSrv/gateway/networkHandler"
	"gameSrv/pkg/network"
	"github.com/spf13/viper"

	"github.com/panjf2000/gnet"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	viper.SetConfigName("config")               // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                  // 配置文件格式
	viper.AddConfigPath("/etc/gateway/config/") // 查找配置文件的路径
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./gateway/config/") // 查找配置文件的路径
	err := viper.ReadInConfig()              // 查找并读取配置文件
	if err != nil {                          // 处理错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	clientNetwork := networkHandler.ClientEventHandler{}
	network.ClientStart(&clientNetwork,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTCPNoDelay(0),
		gnet.WithTicker(true),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))

	controller.Init()

	controller := &networkHandler.ServerEventHandler{}
	network.ServerStart(viper.GetInt32("port"), controller)
}
