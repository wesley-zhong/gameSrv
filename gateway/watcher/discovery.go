package watcher

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {
	switch event {
	case mvccpb.PUT:
		clientConnect := client.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
		node.ChannelContext = clientConnect
		break
	case mvccpb.DELETE:
		if node.ChannelContext != nil {
			node.ChannelContext.Ctx.Close()
		}
		break
	}
}
