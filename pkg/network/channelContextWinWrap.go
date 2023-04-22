//go:build windows
// +build windows

package network

import (
	"bytes"
	"encoding/binary"
	ringbuff "gameSrv/pkg/buff"
	"gateway/pkg/log"
	"net"
)

type ChannelContextWin struct {
	conn           net.Conn
	recBuf         *ringbuff.RingBuffer
	handlerProcess GameEventHandler
	ctx            interface{}
}

// Context returns a user-defined context.
func (context *ChannelContextWin) Context() (ctx interface{}) {
	return context.ctx
}

// SetContext sets a user-defined context.
func (context *ChannelContextWin) SetContext(ctx interface{}) {
	context.ctx = ctx
}

// LocalAddr is the connection's local socket address.
func (context *ChannelContextWin) LocalAddr() (addr net.Addr) {
	return context.conn.LocalAddr()
}

// RemoteAddr is the connection's remote peer address.
func (context *ChannelContextWin) RemoteAddr() (addr net.Addr) {
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
func (context *ChannelContextWin) Read() (buf []byte) {
	body := make([]byte, 4096)
	readLen, err := context.conn.Read(body)
	if err != nil {
		return nil
		panic(err)
	}

	context.recBuf.Write(body[:readLen])
	recvBufLen := context.recBuf.Length()
	var msgLen int32
	bytebuffer := bytes.NewBuffer(context.recBuf.Bytes())
	binary.Read(bytebuffer, binary.BigEndian, &msgLen)
	//log.Infof("-------receive msg len = %d  readLen =%d", msgLen, readLen)
	if msgLen+4 <= int32(recvBufLen) {
		cmBody := make([]byte, msgLen+4)
		context.recBuf.Read(cmBody)
		//log.Infof("oooooooooo read bufLen =%d receivBufLen=%d  msgLen=%d leftlen=%d", readLen, recvBufLen, msgLen, context.recBuf.Length())
		context.handlerProcess.React(cmBody[4:], context)
	}
	return nil
}

// ResetBuffer resets the buffers, which means all data in inbound ring-buffer and event-loop-buffer will be evicted.
func (context *ChannelContextWin) ResetBuffer() {
	//context.conn.ResetBuffer()
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
func (context *ChannelContextWin) ReadN(n int) (size int, buf []byte) {
	//return context.conn.ReadN(n)
	return 0, nil
}

// ShiftN shifts "read" pointer in the internal buffers with the given length.
func (context *ChannelContextWin) ShiftN(n int) (size int) {
	return context.ShiftN(n)
}

// BufferLength returns the length of available data in the internal buffers.
func (context *ChannelContextWin) BufferLength() (size int) {
	//return context.recBuf.BufferLength()
	return 0
}

// ==================================== Concurrency-safe API's ====================================

// SendTo writes data for UDP sockets, it allows you to send data back to UDP socket in individual goroutines.
func (context *ChannelContextWin) SendTo(buf []byte) error {
	//return context.conn.SendTo(buf)
	return nil
}

// AsyncWrite writes one byte slice to peer asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (context *ChannelContextWin) AsyncWrite(buf []byte) error {
	length := len(buf)
	for length > 0 {
		writeLen, err := context.conn.Write(buf)
		if err != nil {
			log.Error(err)
			return err

		}
		length -= writeLen
	}
	return nil
}

// AsyncWritev writes multiple byte slices to peer asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (context *ChannelContextWin) AsyncWritev(bs [][]byte) error {
	//	return context.conn.AsyncWritev(bs)
	return nil
}

// Wake triggers a React event for the connection.
func (context *ChannelContextWin) Wake() error {
	//return context.conn.Wake()
	return nil
}

// Close closes the current connection.
func (context *ChannelContextWin) Close() error {
	context.handlerProcess.OnClosed(context, nil)
	return context.conn.Close()
}

//func (context *ChannelContextWin) Send(msgId int32, msg proto.Message) {
//	marshal, err := proto.Marshal(msg)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//	body := make([]byte, len(marshal)+8)
//	writeBuffer := bytes.NewBuffer(body)
//	writeBuffer.Reset()
//	binary.Write(writeBuffer, binary.BigEndian, msgId)
//	binary.Write(writeBuffer, binary.BigEndian, int32(len(marshal)))
//	binary.Write(writeBuffer, binary.BigEndian, marshal)
//	context.AsyncWrite(writeBuffer.Bytes())
//
//}
