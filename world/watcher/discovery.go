package watcher

import (
	"gameSrv/pkg/discover"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {

}
