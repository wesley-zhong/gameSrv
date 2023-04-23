package client

import (
	protoGen "gameSrv/proto"
	"google.golang.org/protobuf/proto"
)

type InnerMessage struct {
	InnerHeader *protoGen.InnerHead
	Body        proto.Message
}
