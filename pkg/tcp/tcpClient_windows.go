package tcp

import (
	"fmt"
	ringbuff "gameSrv/pkg/buff"
	"net"
	"sync"
	"time"

	"github.com/panjf2000/gnet"
)

var contextMap sync.Map

func init() {

}

var handlerProcess EventHandler

func ClientStart(handler EventHandler, options ...gnet.Option) error {
	if handlerProcess != nil {
		return nil
	}

	handlerProcess = handler
	go timerTick(handler)
	return nil
}
func ClientInited() bool {
	return handlerProcess != nil
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

func timerTick(handler EventHandler) {
	timer := time.NewTimer(3 * time.Second)
	for {
		timer.Reset(3 * time.Second)
		select {
		case <-timer.C:
			handler.Tick()
		}
	}
}
