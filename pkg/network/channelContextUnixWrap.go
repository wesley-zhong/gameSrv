//go:build darwin || linux
// +build darwin linux

package network

import (
	"net"

	"github.com/panjf2000/gnet"
)

type ChannelContextUnix struct {
	conn gnet.Conn
}

// Context returns a user-defined context.
func (context *ChannelContextUnix) Context() (ctx interface{}) {
	return context.conn.Context()
}

// SetContext sets a user-defined context.
func (context *ChannelContextUnix) SetContext(ctx interface{}) {
	context.conn.SetContext(ctx)
}

// LocalAddr is the connection's local socket address.
func (context *ChannelContextUnix) LocalAddr() (addr net.Addr) {
	return context.conn.LocalAddr()
}

// RemoteAddr is the connection's remote peer address.
func (context *ChannelContextUnix) RemoteAddr() (addr net.Addr) {
	return context.conn.RemoteAddr()
}

// Read reads all data from inbound ring-buffer without moving "read" pointer,
// which means it does not evict the data from buffers actually and those data will
// present in buffers until the ResetBuffer method is called.
//
// Note that the (buf []byte) returned by Read() is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop.
// If you have to use buf in a new goroutine, then you need to make a copy of buf and pass this copy
// to that new goroutine.
func (context *ChannelContextUnix) Read() (buf []byte) {
	return context.conn.Read()
}

// ResetBuffer resets the buffers, which means all data in inbound ring-buffer and event-loop-buffer will be evicted.
func (context *ChannelContextUnix) ResetBuffer() {
	context.conn.ResetBuffer()
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
func (context *ChannelContextUnix) ReadN(n int) (size int, buf []byte) {
	return context.conn.ReadN(n)
}

// ShiftN shifts "read" pointer in the internal buffers with the given length.
func (context *ChannelContextUnix) ShiftN(n int) (size int) {
	return context.ShiftN(n)
}

// BufferLength returns the length of available data in the internal buffers.
func (context *ChannelContextUnix) BufferLength() (size int) {
	return context.conn.BufferLength()
}

// ==================================== Concurrency-safe API's ====================================

// SendTo writes data for UDP sockets, it allows you to send data back to UDP socket in individual goroutines.
func (context *ChannelContextUnix) SendTo(buf []byte) error {
	return context.conn.SendTo(buf)
}

// AsyncWrite writes one byte slice to peer asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (context *ChannelContextUnix) AsyncWrite(buf []byte) error {
	return context.conn.AsyncWrite(buf)
}

// AsyncWritev writes multiple byte slices to peer asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (context *ChannelContextUnix) AsyncWritev(bs [][]byte) error {
	return context.conn.AsyncWritev(bs)
}

// Wake triggers a React event for the connection.
func (context *ChannelContextUnix) Wake() error {
	return context.conn.Wake()
}

// Close closes the current connection.
func (context *ChannelContextUnix) Close() error {
	return context.conn.Close()
}
