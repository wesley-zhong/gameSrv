package discover

import (
	"context"
	"gameSrv/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var ServiceRegisterInstance *ServiceRegister
var EtcdClient *clientv3.Client

// ServiceRegister create and register quest
type ServiceRegister struct {
	cli     *clientv3.Client //etcd aresTcpClient
	leaseID clientv3.LeaseID //lease ID
	// lease keep-alive chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	MySelfNode    *Node
}

func Register(endpoints []string, node *Node, lease int64, metaData map[string]string) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("error =", err)
	}

	serv := &ServiceRegister{
		cli:        cli,
		MySelfNode: node,
	}

	if err := serv.registerMeWithLease(node, lease); err != nil {
		return nil, err
	}
	return serv, nil

}
func (s *ServiceRegister) registerMeWithLease(node *Node, ttl int64) error {
	// grant lease time
	resp, err := s.cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}
	_, err = s.cli.Put(context.Background(), node.getKey(), node.getValue(), clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	// set keep-alive logic
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	log.Infof("leaseId = %d", s.leaseID)
	s.keepAliveChan = leaseRespChan
	log.Infof("Register ServiceId:%s  ServiceName:%s  success!", node.ServiceId, node.ServiceName)
	return nil
}

//func (s *ServiceRegister) UpdateNodeValue() error {
//	// register quest and bind lease
//	node := &Node{
//		ServiceName:  s.MySelfNode.ServiceName,
//		ServiceId:    s.MySelfNode.ServiceId,
//		RegisterTime: time.Now().UnixMilli(),
//		Addr:         s.MySelfNode.Addr,
//		MetaData:     s.MySelfNode.MetaData,
//	}
//
//	_, err := s.cli.Put(context.Background(), node.getKey(), node.getValue())
//	if err != nil {
//		return err
//	}
//	return nil
//}

// ListenLeaseRespChan listen  and watch
func (s *ServiceRegister) ListenLeaseRespChan() {
	for {
		select {
		case _, ok := <-s.keepAliveChan:
			if !ok {
				log.Warnf("XXXXdiscover aresTcpClient lose connect")
				return
			}
		}
	}
}

// Close
func (s *ServiceRegister) Close() error {
	// revoke lease
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Infof("close finshed")
	return s.cli.Close()
}

func RegisterMySelf(endpoints []string, node *Node, metaData map[string]string) error {
	service, err := Register(endpoints, node, 6, metaData)
	ServiceRegisterInstance = service
	if err != nil {
		log.Error(err)
		return err
	}
	// listen to keep alive chan
	go service.ListenLeaseRespChan()
	return err
}
