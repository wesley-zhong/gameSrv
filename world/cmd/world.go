package main

import (
	"fmt"
	"gameSrv/pkg/network"
	"gameSrv/world/controller"
	"gameSrv/world/networkHandler"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	// no need to connect other server  so do not need init
	//clientNetwork := networkHandler.ClientEventHandler{}
	//network.ClientStart(&clientNetwork,
	//	gnet.WithMulticore(true),
	//	gnet.WithReusePort(true),
	//	gnet.WithTCPNoDelay(0))

	controller.Init()

	controller := &networkHandler.ServerEventHandler{}
	network.ServerStartWithDeCode(9003, controller, network.NewInnerLengthFieldBasedFrameCodecEx())
}
