package network

import "time"

type (
	GameEventHandler interface {
		// OnInitComplete fires when the server is ready for accepting connections.

		// OnOpened fires when a new connection has been opened.
		// The Conn c has information about the connection such as it's local and remote address.
		// The parameter out is the return value which is going to be sent back to the peer.
		// It is usually not recommended to send large amounts of data back to the peer in OnOpened.
		//
		// Note that the bytes returned by OnOpened will be sent back to the peer without being encoded.
		OnOpened(c ChannelContext) (out []byte, action int)

		// OnClosed fires when a connection has been closed.
		// The parameter err is the last known connection error.
		OnClosed(c ChannelContext, err error) (action int)

		// PreWrite fires just before a packet is written to the peer socket, this event function is usually where
		// you put some code of logging/counting/reporting or any fore operations before writing data to the peer.
		PreWrite(c ChannelContext)

		// AfterWrite fires right after a packet is written to the peer socket, this event function is usually where
		// you put the []byte returned from React() back to your memory pool.
		AfterWrite(c ChannelContext, b []byte)

		// React fires when a socket receives data from the peer.
		// Call c.Read() or c.ReadN(n) of Conn c to read incoming data from the peer.
		// The parameter out is the return value which is going to be sent back to the peer.
		//
		// Note that the parameter packet returned from React() is not allowed to be passed to a new goroutine,
		// as this []byte will be reused within event-loop after React() returns.
		// If you have to use packet in a new goroutine, then you need to make a copy of buf and pass this copy
		// to that new goroutine.
		React(packet []byte, c ChannelContext) (out []byte, action int)

		// Tick fires immediately after the server starts and will fire again
		// following the duration specified by the delay return value.
		Tick() (delay time.Duration, action int)
	}

	// EventServer is a built-in implementation of EventHandler which sets up each method with a default implementation,
	// you can compose it with your own implementation of EventHandler when you don't want to implement all methods
	// in EventHandler.
)
