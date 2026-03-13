package client

import (
	"gameSrv/pkg/global"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"google.golang.org/protobuf/proto"
)

// sendMsg sends a protobuf message to the specified server type
func sendMsg(serverType global.GameServerType, pid int64, msgCode int16, body proto.Message) {
	c := getInnerClient(serverType)
	if c == nil {
		log.Warnf("client not connected, type=%d", serverType)
		return
	}
	c.SendMsg(msgCode, pid, body)
}

// sendBytes sends a byte slice message to the specified server type
func sendBytes(serverType global.GameServerType, body []byte) {
	c := getInnerClient(serverType)
	if c == nil {
		log.Warnf("client not connected, type=%d", serverType)
		return
	}
	c.sendBytesMsg(body)
}

// Gateway functions
func SendInnerToGateway(pid int64, msgCode protoGen.InnerProtoCode, body proto.Message) {
	sendMsg(global.GATE_WAY, pid, int16(msgCode), body)
}

func SendToGateway(pid int64, msgId protoGen.ProtoCode, body proto.Message) {
	sendMsg(global.GATE_WAY, pid, int16(msgId), body)
}

func SendPckToGateway(body []byte) {
	sendBytes(global.GATE_WAY, body)
}

// GameServer functions
func SendInnerToGameServer(pid int64, msgId protoGen.InnerProtoCode, body proto.Message) {
	sendMsg(global.GAME, pid, int16(msgId), body)
}

func SendToGameServer(pid int64, msgId protoGen.ProtoCode, body proto.Message) {
	sendMsg(global.GAME, pid, int16(msgId), body)
}

func SendPckToGameServer(body []byte) {
	sendBytes(global.GAME, body)
}

// Router functions
func SendMsgToRouterServer(pid int64, innerCode protoGen.InnerProtoCode, body proto.Message) {
	sendMsg(global.ROUTER, pid, int16(innerCode), body)
}

func SendPckToRouterServer(body []byte) {
	sendBytes(global.ROUTER, body)
}
