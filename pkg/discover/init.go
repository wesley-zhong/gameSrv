package discover

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/utils"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

type OnWatchServiceChanged func(node *Node, action mvccpb.Event_EventType)

func InitDiscoverAndRegister(v *viper.Viper, onChanged OnWatchServiceChanged, selfType client.GameServerType) error {
	global.SelfServerType = selfType
	//discover from etcd
	discoverUrls := v.GetStringSlice("discover.url")
	watchServs := v.GetStringSlice("discover.watchServ")
	selfPort := v.GetInt32("port")
	if len(watchServs) > 0 {
		//InitClient(onChanged)
		err := InitDiscovery(discoverUrls, watchServs, onChanged)
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

//func InitClient(handler tcp.EventHandler) {
//	tcp.ClientStart(handler,
//		gnet.WithMulticore(true),
//		gnet.WithReusePort(true),
//		gnet.WithTicker(true),
//		gnet.WithTCPNoDelay(gnet.TCPNoDelay))
//}

func connectNode(node *Node) {
	clientConnect := client.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
	node.ChannelContext = clientConnect
}
