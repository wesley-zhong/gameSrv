package network

import (
	"net"
)

type (
	ChannelContext interface {
		// Context returns a user-defined context.
		Context() (ctx interface{})

		// SetContext sets a user-defined context.
		SetContext(ctx interface{})

		// LocalAddr is the connection's local socket address.
		LocalAddr() (addr net.Addr)

		// RemoteAddr is the connection's remote peer address.
		RemoteAddr() (addr net.Addr)

		// Read reads all data from inbound ring-buffer without moving "read" pointer,
		// which means it does not evict the data from buffers actually and those data will
		// present in buffers until the ResetBuffer method is called.
		//
		// Note that the (buf []byte) returned by Read() is not allowed to be passed to a new goroutine,
		// as this []byte will be reused within event-loop.
		// If you have to use buf in a new goroutine, then you need to make a copy of buf and pass this copy
		// to that new goroutine.
		Read() (buf []byte)

		// ResetBuffer resets the buffers, which means all data in inbound ring-buffer and event-loop-buffer will be evicted.
		ResetBuffer()

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
		ReadN(n int) (size int, buf []byte)

		// ShiftN shifts "read" pointer in the internal buffers with the given length.
		ShiftN(n int) (size int)

		// BufferLength returns the length of available data in the internal buffers.
		BufferLength() (size int)

		// ==================================== Concurrency-safe API's ====================================

		// SendTo writes data for UDP sockets, it allows you to send data back to UDP socket in individual goroutines.
		SendTo(buf []byte) error

		// AsyncWrite writes one byte slice to peer asynchronously, usually you would call it in individual goroutines
		// instead of the event-loop goroutines.
		AsyncWrite(buf []byte) error

		// AsyncWritev writes multiple byte slices to peer asynchronously, usually you would call it in individual goroutines
		// instead of the event-loop goroutines.
		AsyncWritev(bs [][]byte) error

		// Wake triggers a React event for the connection.
		Wake() error

		// Close closes the current connection.
		Close() error
	}
)
