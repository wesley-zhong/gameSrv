package discover

import (
	"gameSrv/pkg/utils"
	"github.com/spf13/viper"
)

func InitDiscoverAndRegister(v *viper.Viper) error {

	//discover from etcd
	discoverUrls := v.GetStringSlice("discover.url")
	watchServs := v.GetStringSlice("discover.watchServ")

	if len(watchServs) > 0 {
		err := InitDiscovery(discoverUrls, watchServs)
		if err != nil {
			return err
		}
	}
	//register  myself to etcd
	serviceName := v.GetString("service.name")
	serviceId := utils.CreateServiceUnitName(serviceName)
	metaData := make(map[string]string)
	err := RegisterService(discoverUrls, serviceName, serviceId, metaData)
	if err != nil {
		return err
	}
	return nil
}
