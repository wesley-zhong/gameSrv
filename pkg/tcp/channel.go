package tcp

import (
	"errors"
	"github.com/panjf2000/gnet/v2"
	"net"
	"time"
)

type Channel interface {
	Context() (ctx interface{})
	SetContext(ctx interface{})
	LocalAddr() (addr net.Addr)
	RemoteAddr() (addr net.Addr)
	ReadN(n int) (size int, buf []byte)
	ShiftN(n int) (int, error)
	BufferLength() (size int)
	SendTo(buf []byte) (int, error)
	AsyncWrite(buf []byte) error
	Wake() error
	Close() error
	SetDeadline(t time.Time) (err error)
	// SetReadDeadline implements net.Conn.
	SetReadDeadline(t time.Time) (err error)
	GetId() int
}

type ChannelGnet struct {
	conn gnet.Conn
	ctx  interface{}
}

// Context returns a user-defined context.
func (context *ChannelGnet) Context() (ctx interface{}) {
	//return context.conn.Context()
	return context.ctx
}

// SetContext sets a user-defined context.
func (context *ChannelGnet) SetContext(ctx interface{}) {
	//context.conn.SetContext(ctx)
	context.ctx = ctx
}

// LocalAddr is the connection's local socket address.
func (context *ChannelGnet) LocalAddr() (addr net.Addr) {
	return context.conn.LocalAddr()
}

// RemoteAddr is the connection's remote peer address.
func (context *ChannelGnet) RemoteAddr() (addr net.Addr) {
	return context.conn.RemoteAddr()
}

// ReadN reads bytes with the given length from inbound ring-buffer without moving "read" pointer,
// which means it will not evict the data from buffers until the ShiftN method is called,
// it reads data from the inbound ring-buffer and returns both bytes and the size of it.
// If the length of the available data is less than the given "n", ReadN will return all available data,
// so you should make use of the variable "size" returned by ReadN() to be aware of the exact length of the returned data.
//
// Note that the []byte buf returned by ReadN() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop.
// If you have to use buf in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (context *ChannelGnet) ReadN(n int) (size int, buf []byte) {
	//return context.conn.Read()
	return 0, nil
}

// ShiftN shifts "read" pointer in the gateway buffers with the given length.
func (context *ChannelGnet) ShiftN(n int) (int, error) {
	return context.conn.Discard(n)
}

// BufferLength returns the length of available data in the gateway buffers.
func (context *ChannelGnet) BufferLength() (size int) {
	return context.conn.InboundBuffered()
}

// ==================================== Concurrency-safe API's ====================================

// SendTo writes data for UDP sockets, it allows you to send data back to UDP socket in individual goroutines.
func (context *ChannelGnet) SendTo(buf []byte) (int, error) {
	if len(buf) == 4 {
		return 0, errors.New("Fucccccccccccccccccck")
	}
	return context.conn.Write(buf)
}

// AsyncWrite writes one byte slice to peer asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (context *ChannelGnet) AsyncWrite(buf []byte) error {
	return context.conn.AsyncWrite(buf, nil)
}

// instead of the event-loop goroutines.
func (context *ChannelGnet) AsyncWritev(bs [][]byte) error {
	return context.conn.AsyncWritev(bs, nil)
}

// Wake triggers a React event for the connection.
func (context *ChannelGnet) Wake() error {
	return context.conn.Wake(nil)
}

// Close closes the current connection.
func (context *ChannelGnet) Close() error {
	return context.conn.Close()
}

// SetDeadline implements net.Conn.
func (context *ChannelGnet) SetDeadline(t time.Time) (err error) {
	return context.conn.SetDeadline(t)
}

// SetReadDeadline implements net.Conn.
func (context *ChannelGnet) SetReadDeadline(t time.Time) (err error) {
	return context.conn.SetReadDeadline(t)
}

// SetWriteDeadline implements net.Conn.
func (context *ChannelGnet) SetWriteDeadline(t time.Time) (err error) {
	return context.conn.SetWriteDeadline(t)
}

func (context *ChannelGnet) GetId() int {
	return context.conn.Fd()
}
