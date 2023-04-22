//go:build windows
// +build windows

package network

import (
	"fmt"
	ringbuff "gameSrv/pkg/buff"
	"net"
	"sync"

	"github.com/panjf2000/gnet"
)

var contextMap sync.Map

func init() {

}

var handlerProcess GameEventHandler

func ClientStart(handler GameEventHandler, options ...gnet.Option) error {
	handlerProcess = handler
	return nil
}

func Dial(network, address string) (ChannelContext, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Errorf("client dial err=%s", err.Error())
		return nil, err
	}
	ringBuf := ringbuff.NewRingBuff(2048)
	context := &ChannelContextWin{conn: conn, handlerProcess: handlerProcess, recBuf: ringBuf}
	handlerProcess.OnOpened(context)
	contextMap.Store(conn, context)
	go receiveMsg(context)
	return context, nil
}

func receiveMsg(context ChannelContext) {
	for {
		context.Read()
	}
}
