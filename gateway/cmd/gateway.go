package main

import (
	"fmt"
	"gameSrv/gateway/controller"
	"gameSrv/gateway/networkHandler"
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/tcp"
	"github.com/spf13/viper"
	"net/http"
	_ "net/http/pprof"
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

	//for performance
	go func() {
		http.ListenAndServe("localhost:6062", nil)
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

	// msg Register
	controller.Init()
	//package receive handler
	handler := &networkHandler.ServerEventHandler{}
	//start server
	go tcp.ServerStartWithDeCode(viper.GetInt32("port"), handler, tcp.NewLengthFieldBasedFrameCodecEx())

	////register to etcd
	clientNetwork := &networkHandler.ClientEventHandler{}
	err = discover.InitDiscoverAndRegister(viper.GetViper(), clientNetwork, client.GATE_WAY)
	if err != nil {
		panic(err)
		return
	}
	loopWG.Wait()
}
