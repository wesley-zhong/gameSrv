package client

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gameSrv/pkg/common"
	"gameSrv/pkg/log"
	"gameSrv/pkg/network"
	"gameSrv/protoGen"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"
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
)

var InnerClientMap = make(map[GameServerType]*ConnInnerClientContext)

func InnerClientConnect(serverType GameServerType, addr string, myServerType GameServerType) *ConnInnerClientContext {
	if !network.ClientInited() {
		log.Error(errors.New(" XXXXXXXX  net work client not init ，pleaser init first！"))
		return nil
	}

connect:
	context, err := network.Dial("tcp", addr)
	if err != nil {
		log.Infof("----- connect failed 3 s after reconnect %v", err.Error())
		time.Sleep(3 * time.Second)
		goto connect
	}

	gameInnerClient := NewInnerClientContext(context)
	InnerClientMap[serverType] = gameInnerClient
	handShake := &protoGen.InnerServerHandShake{
		FromServerId:   common.BuildServerUid(int(serverType), 35),
		FromServerType: int32(myServerType),
	}

	header := &protoGen.InnerHead{
		Id:         0,
		SendType:   0,
		ProtoCode:  int32(protoGen.InnerProtoCode_INNER_SERVER_HAND_SHAKE),
		CallbackId: 0,
	}

	innerMessage := &InnerMessage{
		InnerHeader: header,
		Body:        handShake,
	}
	gameInnerClient.Send(innerMessage)
	return gameInnerClient
}

func AddInnerClientConnect(key GameServerType, ctx *ConnInnerClientContext) {
	InnerClientMap[key] = ctx
}

func GetInnerClient(clientType GameServerType) *ConnInnerClientContext {
	return InnerClientMap[clientType]
}

type ConnInnerClientContext struct {
	Ctx      network.ChannelContext
	Sid      int64
	ServerId int64 //this client from which server
}

func NewInnerClientContext(context network.ChannelContext) *ConnInnerClientContext {
	c := &ConnInnerClientContext{Ctx: context, Sid: genSid()}
	context.SetContext(c)
	return c
}

func (client *ConnInnerClientContext) Send(msg *InnerMessage) {
	header, err := proto.Marshal(msg.InnerHeader)
	if err != nil {
		log.Error(err)
	}

	body, err := proto.Marshal(msg.Body)
	headerLen := len(header)
	bodyLen := 0
	if body != nil {
		bodyLen = len(body)
	}

	msgLen := headerLen + bodyLen + 4
	buffer := &bytes.Buffer{}
	buffer.Reset()

	binary.Write(buffer, binary.BigEndian, int32(msgLen))
	binary.Write(buffer, binary.BigEndian, int32(headerLen))
	binary.Write(buffer, binary.BigEndian, header)
	if bodyLen > 0 {
		buffer.Write(body)
	}
	client.Ctx.AsyncWrite(buffer.Bytes())
}

func (client *ConnInnerClientContext) SendInnerMsgProtoCode(innerCode protoGen.InnerProtoCode, roleId int64, msg proto.Message) {
	client.SendInnerMsg(int32(innerCode), roleId, msg)
}

func (client *ConnInnerClientContext) SendInnerMsg(protoCode int32, roleId int64, msg proto.Message) {
	head := &protoGen.InnerHead{
		Id:         roleId,
		SendType:   0,
		ProtoCode:  protoCode,
		CallbackId: 0,
	}
	header, err := proto.Marshal(head)
	if err != nil {
		log.Error(err)
	}

	body, err := proto.Marshal(msg)

	headerLen := len(header)
	bodyLen := 0
	if body != nil {
		bodyLen = len(body)
	}

	msgLen := headerLen + bodyLen + 4
	buffer := &bytes.Buffer{}
	buffer.Reset()

	binary.Write(buffer, binary.BigEndian, int32(msgLen))
	binary.Write(buffer, binary.BigEndian, int32(headerLen))
	binary.Write(buffer, binary.BigEndian, header)
	if bodyLen > 0 {
		buffer.Write(body)
	}
	client.Ctx.AsyncWrite(buffer.Bytes())
}

// ConnClientContext ====================================================================================
type ConnClientContext struct {
	Ctx network.ChannelContext
	Sid int64
}

// NewClientContext - ------ user client -------------------
func NewClientContext(context network.ChannelContext) *ConnClientContext {
	return &ConnClientContext{Ctx: context, Sid: genSid()}
}

func ClientConnect(addr string) *ConnClientContext {
	context, err := network.Dial("tcp", addr)
	if err != nil {
		log.Error(err)
		return nil
	}
	clientContext := NewClientContext(context)
	context.SetContext(clientContext)
	return clientContext

}

func (client *ConnClientContext) SendMsg(code protoGen.ProtoCode, message proto.Message) {
	client.Send(int32(code), message)
}

func (client *ConnClientContext) Send(msgId int32, msg proto.Message) {
	buffer := &bytes.Buffer{}
	buffer.Reset()
	binary.Write(buffer, binary.BigEndian, msgId)
	marshal, err := proto.Marshal(msg)
	if err != nil {
		log.Error(err)
		return
	}
	bodyLen := len(marshal)
	binary.Write(buffer, binary.BigEndian, int32(bodyLen))
	binary.Write(buffer, binary.BigEndian, marshal)
	client.Ctx.AsyncWrite(buffer.Bytes())
}
