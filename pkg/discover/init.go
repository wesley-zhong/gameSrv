package discover

import (
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/utils"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

var MySelfNode *Node

func Init(v *viper.Viper) {
	log.Infof("========")
	serviceName := v.GetString("service.name")
	MySelfNode = &Node{
		ServiceName:    serviceName,
		ServiceId:      utils.CreateServiceUnitName(serviceName),
		RegisterTime:   0,
		Addr:           "",
		MetaData:       nil,
		Type:           0,
		ChannelContext: nil,
	}
}

type OnWatchServiceChanged func(node *Node, action mvccpb.Event_EventType)

func InitDiscoverAndRegister(v *viper.Viper, onChanged OnWatchServiceChanged, selfType global.GameServerType) error {
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
