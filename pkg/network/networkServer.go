package network

import (
	"github.com/panjf2000/gnet"
	log "github.com/panjf2000/gnet/pkg/logging"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"

	"strconv"
	"time"
)

type tcpServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (ts tcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	//log.Infof("conn =%s React", c.RemoteAddr())
	gEventHandler.React(frame, c)
	return nil, 0
}

func (ts tcpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Infof("game server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

// OnShutdown fires when the server is being shut down, it is called right after
// all event-loops and connections are closed.
func (ts tcpServer) OnShutdown(server gnet.Server) {
	log.Infof("server stop")

}

// OnOpened fires when a new connection has been opened.
// The Conn c has information about the connection such as it's local and remote address.
// The parameter out is the return value which is going to be sent back to the peer.
// It is usually not recommended to send large amounts of data back to the peer in OnOpened.
//
// Note that the bytes returned by OnOpened will be sent back to the peer without being encoded.
func (ts tcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	//log.Infof("conn =%s opened", c.RemoteAddr())
	out, act := gEventHandler.OnOpened(c)
	return out, gnet.Action(act)
}

// OnClosed fires when a connection has been closed.
// The parameter err is the last known connection error.
func (ts tcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return gnet.Action(gEventHandler.OnClosed(c, err))
}

// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
func (ts tcpServer) PreWrite(c gnet.Conn) {
	gEventHandler.PreWrite(c)
	return
}

// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
// you put the []byte returned from React() back to your memory pool.
func (ts tcpServer) AfterWrite(c gnet.Conn, b []byte) {
	gEventHandler.AfterWrite(c, b)
	return
}

// Tick fires immediately after the server starts and will fire again
// following the duration specified by the delay return value.
func (ts tcpServer) Tick() (delay time.Duration, action gnet.Action) {
	delay, intAction := gEventHandler.Tick()
	action = gnet.Action(intAction)
	return delay, action
}

var gEventHandler EventHandler

func ServerStartWithDeCode(port int32, eventHandler EventHandler, codec gnet.ICodec) {
	p := goroutine.Default()
	defer p.Release()

	gEventHandler = eventHandler

	ts := &tcpServer{pool: p}
	err := gnet.Serve(ts, "tcp://:"+strconv.Itoa(int(port)),
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTCPNoDelay(0),
		gnet.WithTicker(true),
		gnet.WithCodec(codec))
	if err != nil {
		log.Error(err)
	}
}
