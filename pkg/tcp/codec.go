package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gameSrv/pkg/log"
	"gameSrv/protoGen"

	"github.com/panjf2000/gnet/v2"
	"google.golang.org/protobuf/proto"
	"sync"
)

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
)

type ICodec interface {
	Decode(conn gnet.Conn) ([]byte, error)
	InnerEncode(packet *MsgPacket) ([]byte, error)
	Encode(packet *MsgPacket) ([]byte, error)
	UnPacket(msgId int16, body *bytes.Buffer, method *protoMethod[int64]) *MsgPacket
}

var ErrIncompletePacket = errors.New("incomplete packet")

type DefaultCodec struct{}

type MsgPacket struct {
	MsgId  int16
	Header proto.Message
	Body   proto.Message
}

func (codec *DefaultCodec) Decode(c gnet.Conn) ([]byte, error) {
	const headerSize = 4

	buf, err := c.Peek(headerSize)
	if err != nil {
		return nil, ErrIncompletePacket
	}

	msgLen := binary.BigEndian.Uint32(buf[:headerSize])
	if msgLen > MaxPackageLen {
		return nil, errors.New("packet too large")
	}

	totalPacketSize := headerSize + int(msgLen)
	if c.InboundBuffered() < totalPacketSize {
		return nil, ErrIncompletePacket
	}

	buf, _ = c.Peek(totalPacketSize)
	_, _ = c.Discard(totalPacketSize)
	return buf, nil
}

func (codec *DefaultCodec) InnerEncode(packet *MsgPacket) ([]byte, error) {
	var header []byte
	var headerLen int
	if packet.Header != nil {
		var err error
		header, err = proto.Marshal(packet.Header)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		headerLen = len(header)
	}

	var body []byte
	var bodyLen int
	if packet.Body != nil {
		var err error
		body, err = proto.Marshal(packet.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		bodyLen = len(body)
	}

	msgLen := 2 + 2 + headerLen + bodyLen

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	binary.Write(buf, binary.BigEndian, int32(msgLen))
	binary.Write(buf, binary.BigEndian, packet.MsgId)
	binary.Write(buf, binary.BigEndian, int16(headerLen))

	if headerLen > 0 {
		buf.Write(header)
	}
	if bodyLen > 0 {
		buf.Write(body)
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	return result, nil
}

func (codec *DefaultCodec) Encode(packet *MsgPacket) ([]byte, error) {
	if packet.Header != nil {
		err := errors.New("Encode: Header should be nil for client messages")
		log.Error(err)
		return nil, err
	}

	var body []byte
	var bodyLen int
	if packet.Body != nil {
		var err error
		body, err = proto.Marshal(packet.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		bodyLen = len(body)
	}

	msgLen := 2 + bodyLen

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	binary.Write(buf, binary.BigEndian, int32(msgLen))
	binary.Write(buf, binary.BigEndian, packet.MsgId)

	if bodyLen > 0 {
		buf.Write(body)
	}

	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	return result, nil
}

func (codec *DefaultCodec) UnPacket(msgId int16, body *bytes.Buffer, method *protoMethod[int64]) *MsgPacket {
	var headerLen int16
	if err := binary.Read(body, binary.BigEndian, &headerLen); err != nil {
		log.Errorf("read header length error: %v", err)
		return nil
	}

	var header protoGen.InnerHead
	if headerLen > 0 {
		headerBytes := make([]byte, headerLen)
		if _, err := body.Read(headerBytes); err != nil {
			log.Errorf("read header body error: %v", err)
			return nil
		}

		if err := proto.Unmarshal(headerBytes, &header); err != nil {
			log.Errorf("unmarshal header error: %v", err)
			return nil
		}
	}

	bodyBytes := body.Bytes()
	msg := method.param.ProtoReflect().New().Interface()

	if err := proto.Unmarshal(bodyBytes, msg); err != nil {
		log.Errorf("unmarshal body error: %v", err)
		return nil
	}

	return &MsgPacket{
		MsgId:  msgId,
		Header: &header,
		Body:   msg,
	}
}
