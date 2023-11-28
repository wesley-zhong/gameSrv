package main

import (
	"fmt"
	"gameSrv/gateway/controller"
	"gameSrv/gateway/networkHandler"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/network"
	"gameSrv/pkg/utils"
	"github.com/spf13/viper"
	"sync"

	"github.com/panjf2000/gnet"
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

	viper.SetConfigName("config")                // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                   // 配置文件格式
	viper.AddConfigPath("/etc/gateway/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./gateway/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()               // 查找并读取配置文件
	if err != nil {
		loopWG.Add(-1) // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	clientNetwork := networkHandler.ClientEventHandler{}
	network.ClientStart(&clientNetwork,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTCPNoDelay(0),
		gnet.WithTicker(true),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))

	// msg Register
	controller.Init()
	//package receive handler
	handler := &networkHandler.ServerEventHandler{}
	//start server
	network.ServerStart(viper.GetInt32("port"), handler)

	//register to etcd
	discoverUrls := viper.GetStringSlice("discover.url")
	discover.InitDiscovery(discoverUrls, []string{"game"})

	serviceName := viper.GetString("service.name")
	serviceId := utils.CreateServiceUnitName(serviceName)

	metaData := make(map[string]string)
	metaData["mk1"] = "hello"
	discover.RegisterService(discoverUrls, serviceName, serviceId, metaData)

	loopWG.Wait()
}
