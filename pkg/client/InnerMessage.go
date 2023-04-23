package client

import (
	protoGen "gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

type InnerMessage struct {
	InnerHeader *protoGen.InnerHead
	Body        proto.Message
}
