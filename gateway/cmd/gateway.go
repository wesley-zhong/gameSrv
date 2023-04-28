package main

import (
	"fmt"
	"gameSrv/gateway/controller"
	"gameSrv/gateway/networkHandler"
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
		gnet.WithTCPNoDelay(0),
		gnet.WithTicker(true),
		gnet.WithCodec(network.NewInnerLengthFieldBasedFrameCodecEx()))

	controller.Init()

	controller := &networkHandler.ServerNetWork{}
	network.ServerStart(9001, controller)
}
