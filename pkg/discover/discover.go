package discover

import (
	"context"
	"encoding/json"
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var DiscoverService *ServiceDiscovery

type ServiceDiscovery struct {
	cli        *clientv3.Client
	serverList map[string]*Node
	lock       sync.RWMutex
	onChanged  OnWatchServiceChanged
}

type Node struct {
	ServiceName    string
	ServiceId      string
	RegisterTime   int64
	Addr           string
	MetaData       map[string]string
	Type           global.GameServerType
	Port           int32
	ChannelContext *client.ConnInnerClientContext
}

func (node *Node) getKey() string {
	return node.ServiceId
}

func (node *Node) getValue() string {
	marshal, err := json.Marshal(node)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Error(err)
	}

	return &ServiceDiscovery{
		cli:        cli,
		serverList: make(map[string]*Node),
	}
}

func (s *ServiceDiscovery) WatchService(prefixes []string) error {
	for _, prefix := range prefixes {
		resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
		if err != nil {
			return err
		}
		log.Infof("watch service prefix: %s", prefix)

		for _, ev := range resp.Kvs {
			s.SetServiceList(string(ev.Key), string(ev.Value))
		}

		go s.watcher(prefix)
	}
	return nil
}

func (s *ServiceDiscovery) watcher(prefix string) {
	log.Infof("watching prefix: %s", prefix)
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := &Node{}
	if err := json.Unmarshal([]byte(val), node); err != nil {
		log.Error(err)
		return
	}

	existNode := s.serverList[key]
	if existNode != nil && existNode.ChannelContext != nil {
		return
	}

	s.serverList[key] = node
	log.Infof("discover service: id=%s name=%s", key, node.ServiceName)
	s.onChanged(node, mvccpb.PUT)
}

func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.serverList[key]
	if node != nil && node.ChannelContext != nil && node.ChannelContext.Ctx != nil {
		node.ChannelContext.Ctx.Close()
	}

	delete(s.serverList, key)
	s.onChanged(node, mvccpb.DELETE)
	log.Infof("delete service: id=%s", key)
}

func (s *ServiceDiscovery) GetServices() []*Node {
	s.lock.RLock()
	defer s.lock.RUnlock()

	result := make([]*Node, 0, len(s.serverList))
	for _, v := range s.serverList {
		result = append(result, v)
	}
	return result
}

func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}

func (s *ServiceDiscovery) GetKvClient() *clientv3.Client {
	return s.cli
}

func InitDiscovery(endpoints []string, servicePrefixes []string, onChanged OnWatchServiceChanged) error {
	DiscoverService = NewServiceDiscovery(endpoints)
	DiscoverService.onChanged = onChanged
	err := DiscoverService.WatchService(servicePrefixes)
	if err != nil {
		return err
	}
	return nil
}