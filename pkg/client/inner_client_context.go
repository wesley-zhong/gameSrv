package client

import (
	"errors"
	"gameSrv/pkg/common"
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"
)

var sId int64

func genSid() int64 {
	return atomic.AddInt64(&sId, 1)
}

var ICodec = &tcp.DefaultCodec{}

var (
	innerClientLock sync.RWMutex
	InnerClientMap  = make(map[global.GameServerType]*ConnInnerClientContext)
)

func Connect(addr string) (tcp.Channel, error) {
	context, err := tcp.Dial("tcp", addr)
	if err != nil {
		//log.Infof("----- connect failed 3 s after reconnect %v", err.Error())
		return nil, err
	}
	return context, nil
}

func InnerClientConnect(serverType global.GameServerType, addr string, myServerType global.GameServerType) *ConnInnerClientContext {
	if !tcp.ClientInited() {
		log.Error(errors.New(" XXXXXXXX  net work client not init ，pleaser init first！"))
		return nil
	}

connect:
	context, err := Connect(addr)
	if err != nil {
		log.Infof("----- connect failed. 3 s later  will  reconnect %v", err.Error())
		time.Sleep(3 * time.Second)
		goto connect
	}

	log.Infof("----- connect  success  %s ", addr)
	gameInnerClient := NewInnerClientContext(context)

	innerClientLock.Lock()
	InnerClientMap[serverType] = gameInnerClient
	innerClientLock.Unlock()

	handShake := &protoGen.InnerServerHandShakeReq{
		FromServerId:   common.BuildServerUid(int(serverType), 35),
		FromServerType: int32(myServerType),
		FromServerSid:  global.SelfServiceId,
	}

	header := &protoGen.InnerHead{
		Id: 0,
	}

	packet := &tcp.MsgPacket{MsgId: int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE_REQ), Header: header, Body: handShake}
	gameInnerClient.Send(packet)
	log.Infof("----send handShak req = %s", addr)
	return gameInnerClient
}

func AddInnerClientConnect(key global.GameServerType, ctx *ConnInnerClientContext) {
	innerClientLock.Lock()
	InnerClientMap[key] = ctx
	innerClientLock.Unlock()
}

func DelInnerClientConnect(key global.GameServerType, serviceId string) {
	innerClientLock.Lock()
	defer innerClientLock.Unlock()
	ctx, ok := InnerClientMap[key]
	if ok && ctx.ServiceId == serviceId {
		delete(InnerClientMap, key)
	}
}

// RemoveInnerClient 清理 map 中指向同一连接的过期条目（断线时调用）。
func RemoveInnerClient(stale *ConnInnerClientContext) {
	innerClientLock.Lock()
	defer innerClientLock.Unlock()
	for k, v := range InnerClientMap {
		if v == stale {
			delete(InnerClientMap, k)
		}
	}
}

func getInnerClient(clientType global.GameServerType) *ConnInnerClientContext {
	innerClientLock.RLock()
	defer innerClientLock.RUnlock()
	return InnerClientMap[clientType]
}

type ConnInnerClientContext struct {
	Ctx       tcp.Channel
	Sid       int64
	ServerId  int64 //this client from which server
	ServiceId string
}

func NewInnerClientContext(context tcp.Channel) *ConnInnerClientContext {
	c := &ConnInnerClientContext{Ctx: context, Sid: genSid()}
	context.SetContext(c)
	return c
}

func (client *ConnInnerClientContext) Send(packet *tcp.MsgPacket) {
	encode, err := ICodec.InnerEncode(packet)
	if err != nil {
		log.Error(err)
		return
	}
	if err := client.Ctx.AsyncWrite(encode); err != nil {
		log.Error(err)
	}
}

//func (client *ConnInnerClientContext) SendInnerMsg(innerCode protoGen.InnerProtoCode, pid int64, body proto.Message) {
//	header := &protoGen.InnerHead{Id: pid}
//	packet := &tcp.MsgPacket{
//		MsgId:  int16(innerCode),
//		Header: header,
//		Body:   body,
//	}
//	client.Send(packet)
//}

func (client *ConnInnerClientContext) SendMsg(protoCode int16, pid int64, body proto.Message) {
	header := &protoGen.InnerHead{Id: pid}
	packet := &tcp.MsgPacket{
		MsgId:  protoCode,
		Header: header,
		Body:   body,
	}
	client.Send(packet)
}

func (client *ConnInnerClientContext) sendBytesMsg(body []byte) {
	if err := client.Ctx.AsyncWrite(body); err != nil {
		log.Error(err)
	}
}

func (client *ConnInnerClientContext) SendInnerMsg(innerCode protoGen.InnerProtoCode, pid int64, body proto.Message) {
	header := &protoGen.InnerHead{Id: pid}
	packet := &tcp.MsgPacket{
		MsgId:  int16(innerCode),
		Header: header,
		Body:   body,
	}
	client.Send(packet)
}
