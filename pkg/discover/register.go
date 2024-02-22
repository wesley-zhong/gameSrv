package discover

import (
	"context"
	"fmt"
	"gameSrv/pkg/client"
	"gameSrv/pkg/log"
	"gameSrv/pkg/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var ServiceRegisterInstance *ServiceRegister
var EtcdClient *clientv3.Client

// ServiceRegister create and register quest
type ServiceRegister struct {
	cli     *clientv3.Client //etcd client
	leaseID clientv3.LeaseID //lease ID
	// lease keep-alive chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	serviceId     string //ServiceId
	serviceName   string //value
	MetaData      map[string]string
	localAddr     string
	serverType    client.GameServerType
}

// NewServiceRegister create new quest
func NewServiceRegister(endpoints []string, key, val string, port int32, severType client.GameServerType, lease int64, metaData map[string]string) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("error =", err)
	}

	strIp, err := utils.GetLocalIp()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	serv := &ServiceRegister{
		cli:         cli,
		serviceId:   fmt.Sprintf(key+":%d", port),
		serviceName: val,
		MetaData:    metaData,
		localAddr:   fmt.Sprintf(strIp+":%d", port),
		serverType:  severType,
	}

	if err := serv.putKeyWithLease(lease); err != nil {
		return nil, err
	}
	return serv, nil
}

// set lease
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	// grant lease time
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	// register quest and bind lease
	node := &Node{
		ServiceName:  s.serviceName,
		ServiceId:    s.serviceId,
		RegisterTime: time.Now().UnixMilli(),
		Addr:         s.localAddr,
		MetaData:     s.MetaData,
		Type:         s.serverType,
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
	log.Infof("Register ServiceId:%s  ServiceName:%s  success!", s.serviceId, s.serviceName)
	return nil
}

func (s *ServiceRegister) UpdateNodeValue() error {
	// register quest and bind lease
	node := &Node{
		ServiceName:  s.serviceName,
		ServiceId:    s.serviceId,
		RegisterTime: time.Now().UnixMilli(),
		Addr:         s.localAddr,
		MetaData:     s.MetaData,
	}

	_, err := s.cli.Put(context.Background(), node.getKey(), node.getValue())
	if err != nil {
		return err
	}
	return nil
}

// ListenLeaseRespChan listen  and watch
func (s *ServiceRegister) ListenLeaseRespChan() {
	for {
		select {
		case _, ok := <-s.keepAliveChan:
			if !ok {
				log.Warnf("XXXXdiscover client lose connect")
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

func RegisterService(endpoints []string, serviceName string, serviceId string, port int32, serveType client.GameServerType, metaData map[string]string) error {
	service, err := NewServiceRegister(endpoints, serviceId, serviceName, port, serveType, 6, metaData)
	ServiceRegisterInstance = service
	if err != nil {
		log.Error(err)
		return err
	}
	// listen to keep alive chan
	go service.ListenLeaseRespChan()
	return err
}
