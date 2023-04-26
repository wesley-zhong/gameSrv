package main

import (
	"fmt"
	"gameSrv/game/controller"
	"gameSrv/game/networkHandler"
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
		gnet.WithTCPNoDelay(0))

	controller.Init()

	controller := &networkHandler.ServerNetWork{}
	network.ServerStart(9002, controller)
}
