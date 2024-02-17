package watcher

import (
	"gameSrv/pkg/discover"
	"gameSrv/pkg/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {
	log.Infof("========== watch service %s", node)

}
