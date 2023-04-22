//go:build darwin || linux
// +build darwin linux

package network

import (
	"time"

	"github.com/panjf2000/gnet"
)

type Client struct {
}

type gnetHandler struct {
	server      gnet.Server
	gameHandler GameEventHandler
}

func (handler *gnetHandler) OnInitComplete(server gnet.Server) (action gnet.Action) {
	handler.server = server
	return 0
}

// OnShutdown fires when the server is being shut down, it is called right after
// all event-loops and connections are closed.
func (handler *gnetHandler) OnShutdown(server gnet.Server) {
	//handler.gameHandler
}

// OnOpened fires when a new connection has been opened.
// The Conn c has information about the connection such as it's local and remote address.
// The parameter out is the return value which is going to be sent back to the peer.
// It is usually not recommended to send large amounts of data back to the peer in OnOpened.
//
// Note that the bytes returned by OnOpened will be sent back to the peer without being encoded.
func (handler *gnetHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	opened, a := handler.gameHandler.OnOpened(c)
	return opened, gnet.Action(a)
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (handler *gnetHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	handler.gameHandler.OnClosed(c, err)
	return gnet.Close
}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (handler *gnetHandler) PreWrite(c gnet.Conn) {
	handler.gameHandler.PreWrite(c)
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (handler *gnetHandler) AfterWrite(c gnet.Conn, b []byte) {
	handler.gameHandler.AfterWrite(c, b)
}

// React fires when a socket receives data from the peer.
// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
// The parameter out is the return value which is going to be sent back to the peer.
//
// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after React() returns.
// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (handler *gnetHandler) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	react, a := handler.gameHandler.React(packet, c)
	return react, gnet.Action(a)
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (handler *gnetHandler) Tick() (delay time.Duration, action gnet.Action) {
	tick, a := handler.gameHandler.Tick()
	return tick, gnet.Action(a)
}

var gClient *gnet.Client

func ClientStart(handler GameEventHandler, opts ...gnet.Option) error {
	gnetHandler := &gnetHandler{gameHandler: handler}
	client, err := gnet.NewClient(gnetHandler, opts...)
	client.Start()
	gClient = client
	return err
}

func ClientStop() {
	gClient.Stop()

}

func Dial(network, address string) (ChannelContext, error) {
	return gClient.Dial(network, address)
}
