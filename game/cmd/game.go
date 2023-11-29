package main

import (
	"fmt"
	"gameSrv/game/controller"
	"gameSrv/game/networkHandler"
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/network"
	"github.com/spf13/viper"
	"sync"

	"net/http"
	_ "net/http/pprof"
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

	//mongodb init
	//dal.InitMongoDB(viper.GetString("mongo.Addr"), viper.GetString("mongo.userName"), viper.GetString("mongo.password"))
	//dal.InitRedisDB(viper.GetString("redis.addr"), viper.GetString("redis.password"))
	//
	//account := service.AccountLogin("andy")
	//service.UpdateAccount(account)

	// msg Register
	controller.Init()

	//start server
	serverNetworkHandler := &networkHandler.ServerEventHandler{}
	go network.ServerStartWithDeCode(viper.GetInt32("port"), serverNetworkHandler, network.NewInnerLengthFieldBasedFrameCodecEx())

	////register to etcd
	clientNetwork := &networkHandler.ClientEventHandler{}
	err = discover.InitDiscoverAndRegister(viper.GetViper(), clientNetwork, client.GAME)
	if err != nil {
		panic(err)
		return
	}
	loopWG.Wait()
}
