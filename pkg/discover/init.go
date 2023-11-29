package discover

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/network"
	"gameSrv/pkg/utils"
	"github.com/panjf2000/gnet"
	"github.com/spf13/viper"
)

func InitDiscoverAndRegister(v *viper.Viper, handler network.EventHandler, selfType client.GameServerType) error {
	//discover from etcd
	discoverUrls := v.GetStringSlice("discover.url")
	watchServs := v.GetStringSlice("discover.watchServ")
	selfPort := v.GetInt32("port")
	if len(watchServs) > 0 {
		InitClient(handler)
		err := InitDiscovery(discoverUrls, watchServs)
		if err != nil {
			return err
		}
		connect(DiscoverService.serverList, selfType)
	}
	//register  myself to etcd
	serviceName := v.GetString("service.name")
	serviceId := utils.CreateServiceUnitName(serviceName)
	metaData := make(map[string]string)
	err := RegisterService(discoverUrls, serviceName, serviceId, selfPort, selfType, metaData)
	if err != nil {
		return err
	}
	return nil
}

func InitClient(handler network.EventHandler) {
	network.ClientStart(handler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))
}

func connect(serverList map[string]*Node, selfType client.GameServerType) {
	for _, node := range serverList {
		client.InnerClientConnect(node.Type, node.Addr, selfType)
	}
}
