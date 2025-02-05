package watcher

import (
	"gameSrv/pkg/discover"
	"gameSrv/pkg/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"strconv"
)

var GateWayInfoMap map[string]*GateWayInfo = make(map[string]*GateWayInfo)

func BalanceGateNode() *GateWayInfo {
	for _, v := range GateWayInfoMap {
		return v
	}
	return nil
}

type GateWayInfo struct {
	PlayerCount int
	Node        *discover.Node
}

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {
	log.Infof("========== watch service %s", node)
	info := GateWayInfoMap[node.ServiceId]
	onlicount, err := strconv.Atoi(node.MetaData["cu"])
	if err != nil {
		onlicount = 0
	}
	if info == nil {
		info = &GateWayInfo{
			Node: node,
		}
		GateWayInfoMap[node.ServiceId] = info
	}
	info.PlayerCount = onlicount
}
