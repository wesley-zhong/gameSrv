package main

import (
	"fmt"
	"gameSrv/login/controller"
	"gameSrv/login/networkHandler"
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/web"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"

	"github.com/spf13/viper"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s", x, s)
		}
	}()
	//for performance
	go func() {
		http.ListenAndServe("localhost:6061", nil)
	}()

	viper.SetConfigName("config")              // 配置文件名，不需要后缀名
	viper.SetConfigType("yml")                 // 配置文件格式
	viper.AddConfigPath("/etc/login/configs/") // 查找配置文件的路径
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./login/configs/") // 查找配置文件的路径
	err := viper.ReadInConfig()             // 查找并读取配置文件
	if err != nil {                         // 处理错误
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	//mongodb init
	//dal.InitMongoDB(viper.GetString("mongo.Addr"), viper.GetString("mongo.userName"), viper.GetString("mongo.password"))
	//dal.InitRedisDB(viper.GetString("redis.addr"), viper.GetString("redis.password"))
	//
	//account := service.AccountLogin("andy")
	//service.UpdateAccount(account)

	//start server
	server := web.NewHttpServer()
	controller.Init(server.HttpMethod)
	server.WebAppStart(viper.GetInt32("port"))

	////register to etcd
	clientNetwork := &networkHandler.ClientEventHandler{}
	err = discover.InitDiscoverAndRegister(viper.GetViper(), clientNetwork, client.LOGIN)
	if err != nil {
		panic(err)
	}
}
