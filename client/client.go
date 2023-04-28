package main

import (
	"fmt"
	"gameSrv/client/controller"
	"gameSrv/client/networkHandler"
	"gameSrv/pkg/network"
	"github.com/panjf2000/gnet"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	clientNetwork := networkHandler.ClientNetwork{}
	network.ClientStart(&clientNetwork,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))

	controller.Init()
}
