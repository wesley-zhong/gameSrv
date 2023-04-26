package main

import (
	"encoding/binary"
	"fmt"
	"gameSrv/pkg/network"
	"gameSrv/world/controller"
	"gameSrv/world/networkHandler"
	"github.com/panjf2000/gnet"
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
	network.ServerStartWithDeCode(9003, controller, gnet.NewLengthFieldBasedFrameCodec(gnet.EncoderConfig{
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}, gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 4}))
}
