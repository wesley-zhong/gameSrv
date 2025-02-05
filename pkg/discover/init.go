package discover

import (
	"fmt"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/utils"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

var MySelfNode *Node

func Init(v *viper.Viper, selfType global.GameServerType) error {
	log.Infof("========")

	port := v.GetInt32("port")
	serviceName := v.GetString("service.name")
	addr, err := utils.GetLocalIp()
	if err != nil {
		return err
	}
	MySelfNode = &Node{
		ServiceName:    serviceName,
		ServiceId:      fmt.Sprintf(utils.CreateServiceUnitName(serviceName)+":%d", port),
		RegisterTime:   0,
		Addr:           fmt.Sprintf(addr+":%d", port),
		MetaData:       make(map[string]string),
		Type:           selfType,
		Port:           port,
		ChannelContext: nil,
	}
	global.SelfServiceId = MySelfNode.ServiceId
	return nil
}

type OnWatchServiceChanged func(node *Node, action mvccpb.Event_EventType)

func InitDiscoverAndRegister(v *viper.Viper, onChanged OnWatchServiceChanged) error {
	//discover from etcd
	discoverUrls := v.GetStringSlice("discover.url")
	watchServs := v.GetStringSlice("discover.watchServ")

	if len(watchServs) > 0 {
		err := InitDiscovery(discoverUrls, watchServs, onChanged)
		if err != nil {
			return err
		}
	}
	//register  myself to etcd

	metaData := make(map[string]string)
	err1 := RegisterMySelf(discoverUrls, MySelfNode, metaData)
	if err1 != nil {
		return err1
	}
	return nil
}
