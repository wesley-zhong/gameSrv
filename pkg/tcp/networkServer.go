package tcp

import (
	"fmt"
	"gameSrv/pkg/log"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/pool/goroutine"
	"time"
)

type tcpServer struct {
	gnet.BuiltinEventEngine
	eng          gnet.Engine
	network      string
	addr         string
	multicore    bool
	connected    int32
	disconnected int32
	codec        ICodec
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (ts tcpServer) OnBoot(srv gnet.Engine) (action gnet.Action) {
	log.Infof("game server  started)\n")
	return
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (ts tcpServer) OnShutdown(eng gnet.Engine) {
	log.Infof("server stop")
}

// OnOpened fires when a new connection has been opened.
// The Conn c has information about the connection such as it's local and remote address.
// The parameter out is the return value which is going to be sent back to the peer.
// It is usually not recommended to send large amounts of data back to the peer in OnOpened.

func (ts tcpServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Infof("---OnOpen   conn =%s opened", c.RemoteAddr())
	context := &ChannelGnet{
		c,
		nil,
	}
	out, act := gEventHandler.OnOpened(context)
	c.SetContext(context)
	return out, gnet.Action(act)
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (ts tcpServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	channel := c.Context().(Channel)
	return gnet.Action(gEventHandler.OnClosed(channel, err))
}

// OnTraffic fires when a socket receives data from the peer.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (ts tcpServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	for {
		bytes, err := ts.codec.Decode(c)
		if err != nil {
			return 0
		}
		channel := c.Context().(Channel)
		gEventHandler.React(bytes, channel)
	}
	return gnet.None
}

func (ts tcpServer) OnTick() (delay time.Duration, action gnet.Action) {
	delay, intAction := gEventHandler.Tick()
	return delay, gnet.Action(intAction)
}

var gEventHandler EventHandler

func ServerStartWithDeCode(port int32, eventHandler EventHandler, codec ICodec) {
	p := goroutine.Default()
	defer p.Release()

	gEventHandler = eventHandler

	ss := &tcpServer{
		network:   "tcp",
		addr:      fmt.Sprintf(":%d", port),
		multicore: true,
		codec:     codec,
	}
	log.Infof("###### server  start: %v", port)
	err := gnet.Run(ss, ss.network+"://"+ss.addr, gnet.WithMulticore(true), gnet.WithReuseAddr(true), gnet.WithReusePort(true))
	log.Infof("server exits with error: %v", err)
	if err != nil {
		log.Error(err)
	}
}
