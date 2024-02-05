package tcp

import (
	"time"

	"github.com/panjf2000/gnet/v2"
)

type Client struct {
}

type gnetHandler struct {
	server      gnet.BuiltinEventEngine
	gameHandler EventHandler
	codec       ICodec
}

func (handler *gnetHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {

	return gnet.None
}

// OnShutdown fires when the server is being shut down, it is called right after
// all event-loops and connections are closed.
func (handler *gnetHandler) OnShutdown(server gnet.Engine) {
	//handler.gameHandler
}

// OnOpen fires when a new connection has been opened.
//
// The Conn c has information about the connection such as its local and remote addresses.
// The parameter out is the return value which is going to be sent back to the peer.
// Sending large amounts of data back to the peer in OnOpen is usually not recommended.
func (handler *gnetHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	context := &ChannelContextGnet{c}
	opened, a := handler.gameHandler.OnOpened(context)
	return opened, gnet.Action(a)
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (handler *gnetHandler) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	context := &ChannelContextGnet{c}
	handler.gameHandler.OnClosed(context, err)
	return gnet.Close
}

// OnTraffic fires when a socket receives data from the peer.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (handler *gnetHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	bytes, err := handler.codec.Decode(c)
	if err != nil {
		return 0
	}
	context := &ChannelContextGnet{c}
	a := handler.gameHandler.React(bytes, context)
	return gnet.Action(a)
}

// OnTick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (handler *gnetHandler) OnTick() (delay time.Duration, action gnet.Action) {
	tick, a := handler.gameHandler.Tick()
	return tick, gnet.Action(a)
}

var gClient *gnet.Client

func ClientStart(handler EventHandler, opts ...gnet.Option) error {
	gnetHandler := &gnetHandler{gameHandler: handler, codec: &DefaultCodec{}}
	client, err := gnet.NewClient(gnetHandler, opts...)
	client.Start()
	gClient = client
	return err
}

func ClientStop() {
	gClient.Stop()

}

func ClientInited() bool {
	return gClient != nil
}

func Dial(network, address string) (ChannelContext, error) {
	conn, err := gClient.Dial(network, address)
	if err != nil {
		return &ChannelContextGnet{}, err
	}
	return &ChannelContextGnet{conn}, nil
}
