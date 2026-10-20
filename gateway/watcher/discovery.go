package watcher

import (
	"gameSrv/pkg/client"
	"gameSrv/pkg/discover"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

var connecting sync.Map // serviceId -> struct{}，防止重复发起连接

func OnDiscoveryServiceChange(node *discover.Node, event mvccpb.Event_EventType) {
	switch event {
	case mvccpb.PUT:
		if _, busy := connecting.LoadOrStore(node.ServiceId, struct{}{}); busy {
			return
		}
		go func() {
			defer connecting.Delete(node.ServiceId)
			clientConnect := client.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
			if clientConnect == nil {
				return
			}
			clientConnect.ServiceId = node.ServiceId
			node.ChannelContext = clientConnect
		}()
	case mvccpb.DELETE:
		if node.ChannelContext != nil {
			node.ChannelContext.Ctx.Close()
		}
		client.DelInnerClientConnect(node.Type, node.ServiceId)
	}
}

// OnInnerClientClosed 由内部连接触发 OnClosed 时调用：
// 清理 InnerClientMap 中的过期连接，并按 etcd 上仍在的节点信息异步重连。
func OnInnerClientClosed(closed *client.ConnInnerClientContext) {
	if closed == nil {
		return
	}

	client.RemoveInnerClient(closed)
	if closed.ServiceId == "" {
		return
	}

	if discover.DiscoverService == nil {
		return
	}
	node := discover.DiscoverService.GetNode(closed.ServiceId)
	if node == nil {
		// etcd 上节点已删除，无需重连（DELETE 事件路径会负责清理）
		return
	}

	if _, busy := connecting.LoadOrStore(node.ServiceId, struct{}{}); busy {
		return
	}
	go func() {
		defer connecting.Delete(node.ServiceId)
		reconnected := client.InnerClientConnect(node.Type, node.Addr, global.SelfServerType)
		if reconnected == nil {
			return
		}
		reconnected.ServiceId = node.ServiceId
		node.ChannelContext = reconnected
		log.Infof("reconnected to service=%s addr=%s", node.ServiceId, node.Addr)
	}()
}
