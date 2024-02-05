package discover

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/tcp"
	"gameSrv/pkg/utils"
	"github.com/panjf2000/gnet/v2"
	"github.com/spf13/viper"
)

func InitDiscoverAndRegister(v *viper.Viper, handler tcp.EventHandler, selfType client.GameServerType) error {
	global.SelfServerType = selfType
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

func InitClient(handler tcp.EventHandler) {
	tcp.ClientStart(handler,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay))
}

func connectNode(node *Node) {
	clientConnect := client.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
	node.channelContext = clientConnect
}
