package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"
	"github.com/panjf2000/gnet/v2"
	"google.golang.org/protobuf/proto"
	"sync"
)

// DefaultCodec Protocol format:
//
// * 0     4           2
// * +-----------+-----------------------+
// * |body len   |  msgId
// * +-----------+-----------+-----------+
// * |                                   |
// * +                                   +
// * |           body bytes              |
// * +                                   +
// * |            ... ...                |
// * +-----------------------------------+

const (
	bodySize      = 4
	MaxPackageLen = 16 * 1024
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	byteSlicePool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 256)
		},
	}
)

type ICodec interface {
	Decode(conn gnet.Conn) ([]byte, error)
	//InnerEncode between server's msg encode
	InnerEncode(packet *MsgPacket) ([]byte, error)
	//Encode  user client to server
	Encode(packet *MsgPacket) ([]byte, error)
	UnPacket(msgId int16, body *bytes.Buffer, method *protoMethod[int64]) *MsgPacket
}

var ErrIncompletePacket = errors.New("incomplete packet")

type DefaultCodec struct {
}

type MsgPacket struct {
	MsgId  int16
	Header proto.Message
	Body   proto.Message
}

func (codec *DefaultCodec) Decode(c gnet.Conn) ([]byte, error) {
	// Protocol format: [4 bytes: msg len][2 bytes: msgId][2 bytes: header len][header bytes][body bytes]
	totalSizeOffset := 4
	buf, err := c.Peek(totalSizeOffset)
	if err != nil {
		return nil, ErrIncompletePacket
	}

	msgLen := binary.BigEndian.Uint32(buf[:totalSizeOffset])
	if msgLen > MaxPackageLen {
		return nil, fmt.Errorf("packet too large: %d, max allowed: %d", msgLen, MaxPackageLen)
	}

	totalPacketSize := totalSizeOffset + int(msgLen)
	if c.InboundBuffered() < totalPacketSize {
		return nil, ErrIncompletePacket
	}

	buf, _ = c.Peek(totalPacketSize)
	_, _ = c.Discard(totalPacketSize)

	return buf, nil
}

func (code *DefaultCodec) InnerEncode(packet *MsgPacket) (sendBody []byte, err error) {
	headerLen := 0
	var header []byte
	if packet.Header != nil {
		header, err = proto.Marshal(packet.Header)
		if err != nil {
			log.Error(err)
			return
		}
		headerLen = len(header)
	}
	bodyLen := 0
	if packet.Body != nil {
		sendBody, err = proto.Marshal(packet.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		bodyLen = len(sendBody)
	}

	msgLen := 2 //msgId
	msgLen += 2 // header len
	msgLen += headerLen
	msgLen += bodyLen

	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	buffer.Reset()

	binary.Write(buffer, binary.BigEndian, int32(msgLen))
	binary.Write(buffer, binary.BigEndian, packet.MsgId)
	binary.Write(buffer, binary.BigEndian, int16(headerLen))
	if headerLen > 0 {
		buffer.Write(header)
	}

	if bodyLen > 0 {
		buffer.Write(sendBody)
	}
	// Return a copy of the buffer data
	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, nil
}

func (codec *DefaultCodec) Encode(packet *MsgPacket) (sendBody []byte, err error) {
	if packet.Header != nil {
		log.Error(errors.New("Encode method packet.Header should be nil "))
		return
	}
	bodyLen := 0
	if packet.Body != nil {
		sendBody, err = proto.Marshal(packet.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		bodyLen = len(sendBody)
	}

	msgLen := 2 //msgId
	msgLen += bodyLen

	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	buffer.Reset()

	binary.Write(buffer, binary.BigEndian, int32(msgLen))
	binary.Write(buffer, binary.BigEndian, packet.MsgId)
	if bodyLen > 0 {
		buffer.Write(sendBody)
	}
	// Return a copy of the buffer data
	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, nil
}

func (codec *DefaultCodec) UnPacket(msgId int16, body *bytes.Buffer, method *protoMethod[int64]) *MsgPacket {
	var shortHeadLen int16
	binary.Read(body, binary.BigEndian, shortHeadLen)
	header := &protoGen.InnerHead{}

	var headerBody []byte
	if shortHeadLen > 0 {
		headerBody = make([]byte, shortHeadLen)
		readLen, err := body.Read(headerBody)
		if err != nil {
			log.Errorf("Failed to read header body: %v", err)
			return nil
		}

		if readLen != int(shortHeadLen) {
			log.Errorf("Read %d bytes but headLen = %d", readLen, shortHeadLen)
			return nil
		}
		err = proto.Unmarshal(headerBody, header)
		if err != nil {
			log.Errorf("Failed to unmarshal header: %v", err)
			return nil
		}
	}

	bodyBytes := body.Bytes()
	msg := method.param.ProtoReflect().New().Interface()
	err := proto.Unmarshal(bodyBytes, msg)
	if err != nil {
		log.Errorf("Failed to unmarshal body: %v", err)
		return nil
	}
	return &MsgPacket{MsgId: msgId, Header: header, Body: msg}
}
