package client

import (
	protoGen "gameSvr/proto"
	"google.golang.org/protobuf/proto"
)

type InnerMessage struct {
	InnerHeader *protoGen.InnerHead
	Body        proto.Message
}
