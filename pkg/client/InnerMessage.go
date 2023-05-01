package client

import (
	"gameSrv/protoGen"
	"google.golang.org/protobuf/proto"
)

type InnerMessage struct {
	InnerHeader *protoGen.InnerHead
	Body        proto.Message
}
