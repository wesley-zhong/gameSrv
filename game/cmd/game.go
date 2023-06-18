package main

import (
	"fmt"
	"gameSrv/game/controller"
	"gameSrv/game/networkHandler"
	"gameSrv/pkg/network"
	"github.com/panjf2000/gnet"
	"github.com/spf13/viper"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()
	//for performance
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	viper.SetConfigName("config")             // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                // 配置文件格式
	viper.AddConfigPath("/etc/game/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./game/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()            // 查找并读取配置文件
	if err != nil {                        // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}
	clientNetwork := networkHandler.ClientEventHandler{}
	network.ClientStart(&clientNetwork,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))

	//dal.Init(viper.GetString("mongo.Addr"), viper.GetString("mongo.userName"), viper.GetString("mongo.password"))
	//account := service.AccountLogin("andy")
	//service.UpdateAccount(account)

	controller.Init()
	networkHandler := &networkHandler.ServerEventHandler{}
	network.ServerStartWithDeCode(viper.GetInt32("port"), networkHandler, network.NewInnerLengthFieldBasedFrameCodecEx())
}
