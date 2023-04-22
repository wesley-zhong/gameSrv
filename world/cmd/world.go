package main

import (
	"fmt"
	"gameSvr/pkg/network"
	"gameSvr/world/controller"
	"gameSvr/world/networkHandler"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()

	// no need to connect other server  so do not need init
	//clientNetwork := networkHandler.ClientNetwork{}
	//network.ClientStart(&clientNetwork,
	//	gnet.WithMulticore(true),
	//	gnet.WithReusePort(true),
	//	gnet.WithTCPNoDelay(0))

	controller.Init()

	controller := &networkHandler.ServerNetWork{}
	network.ServerStart(9003, controller)
}
