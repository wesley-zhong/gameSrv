package watcher

import (
	"gameSrv/pkg/aresTcpClient"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {
	switch event {
	case mvccpb.PUT:
		clientConnect := aresTcpClient.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
		clientConnect.ServiceId = node.ServiceId
		node.ChannelContext = clientConnect
	case mvccpb.DELETE:
		if node.ChannelContext != nil {
			node.ChannelContext.Ctx.Close()
		}
		aresTcpClient.DelInnerClientConnect(node.Type, node.ServiceId)
	}
}
