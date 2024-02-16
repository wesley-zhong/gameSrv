package client

import (
	"errors"
	"fmt"
	"gameSrv/pkg/common"
	"gameSrv/pkg/log"
	"gameSrv/pkg/tcp"
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
	"sync/atomic"
	"time"
)

var sId int64

func genSid() int64 {
	return atomic.AddInt64(&sId, 1)
}

type GameServerType int32

// ConnInnerClientContext -------------server inner client ---------------
const (
	GATE_WAY GameServerType = 1
	GAME     GameServerType = 2
	WORLD    GameServerType = 3
	LOGIN    GameServerType = 4
)

var ICodec = &tcp.DefaultCodec{}

var InnerClientMap = make(map[GameServerType]*ConnInnerClientContext)

func Connect(addr string) (tcp.Channel, error) {
	context, err := tcp.Dial("tcp", addr)
	if err != nil {
		//log.Infof("----- connect failed 3 s after reconnect %v", err.Error())
		return nil, err
	}
	return context, nil
}

func InnerClientConnect(serverType GameServerType, addr string, myServerType GameServerType) *ConnInnerClientContext {
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
	InnerClientMap[serverType] = gameInnerClient
	handShake := &protoGen.InnerServerHandShake{
		FromServerId:   common.BuildServerUid(int(serverType), 35),
		FromServerType: int32(myServerType),
	}

	header := &protoGen.InnerHead{
		Id: 0,
	}

	packet := &tcp.MsgPacket{MsgId: int16(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE), Header: header, Body: handShake}
	gameInnerClient.Send(packet)
	return gameInnerClient
}

func AddInnerClientConnect(key GameServerType, ctx *ConnInnerClientContext) {
	InnerClientMap[key] = ctx
}

func GetInnerClient(clientType GameServerType) *ConnInnerClientContext {
	return InnerClientMap[clientType]
}

type ConnInnerClientContext struct {
	Ctx      tcp.Channel
	Sid      int64
	ServerId int64 //this client from which server
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
	sendLen, err := client.Ctx.SendTo(encode)
	if err != nil {
		log.Error(err)
		return
	}
	if sendLen != len(encode) {
		log.Error(errors.New(fmt.Sprintf(" send = %d  total en = %d", sendLen, len(encode))))
	}
}

func (client *ConnInnerClientContext) SendInnerMsg(innerCode protoGen.InnerProtoCode, roleId int64, body proto.Message) {
	header := &protoGen.InnerHead{Id: roleId}
	packet := &tcp.MsgPacket{
		MsgId:  int16(innerCode),
		Header: header,
		Body:   body,
	}
	client.Send(packet)
}

func (client *ConnInnerClientContext) SendMsg(protoCode protoGen.ProtoCode, roleId int64, body proto.Message) {
	header := &protoGen.InnerHead{Id: roleId}
	packet := &tcp.MsgPacket{
		MsgId:  int16(protoCode),
		Header: header,
		Body:   body,
	}
	client.Send(packet)
}

func (client *ConnInnerClientContext) SendBytesMsg(body []byte) {
	client.Ctx.SendTo(body)
}
