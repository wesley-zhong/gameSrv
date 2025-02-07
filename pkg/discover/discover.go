package discover

import (
	"context"
	"encoding/json"
	"gameSrv/pkg/client"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var DiscoverService *ServiceDiscovery

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli        *clientv3.Client //etcd client
	serverList map[string]*Node //服务列表
	lock       sync.Mutex
	onChanged  OnWatchServiceChanged
}
type Node struct {
	ServiceName    string                `json:"serviceName"`
	ServiceId      string                `json:"serviceId"`
	RegisterTime   int64                 `json:"registerTime"`
	Addr           string                `json:"addr"`
	MetaData       map[string]string     `json:"metaData"`
	Type           global.GameServerType `json:"type"`
	Port           int32                 `json:"port"`
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

// NewServiceDiscovery  新建发现服务
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

// WatchService 初始化服务列表和监视
func (s *ServiceDiscovery) WatchService(prefixes []string) error {
	//根据前缀获取现有的key
	for _, prefix := range prefixes {
		resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
		if err != nil {
			return err
		}
		log.Infof(" ###  WatchService get  prefix:%s now...", prefix)

		for _, ev := range resp.Kvs {
			s.SetServiceList(string(ev.Key), string(ev.Value))
		}

		//监视前缀，修改变更的server
		go s.watcher(prefix)
	}
	return nil
}

// watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	log.Infof("watching prefix:%s now...", prefix)
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: //删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	node := &Node{}
	err := json.Unmarshal([]byte(val), node)
	if err != nil {
		log.Error(err)
		return
	}
	existNode := s.serverList[key]
	if existNode != nil && existNode.ChannelContext != nil {
		return
	}
	s.serverList[key] = node
	log.Infof("### discover quest :ServiceId  %s:  ServiceName: %s", key, val)
	s.onChanged(node, mvccpb.PUT)
	//go connectNode(node)
}

// DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	node := s.serverList[key]
	//关闭
	if node != nil && node.ChannelContext != nil && node.ChannelContext.Ctx != nil {
		node.ChannelContext.Ctx.Close()
	}
	delete(s.serverList, key)
	s.onChanged(node, mvccpb.DELETE)
	log.Infof("-------del ServiceId: %s", key)
}

// GetServices 获取服务地址
func (s *ServiceDiscovery) GetServices() []*Node {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]*Node, 0)

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

// Close 关闭服务
func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}

func (s *ServiceDiscovery) GetKvClient() {

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
