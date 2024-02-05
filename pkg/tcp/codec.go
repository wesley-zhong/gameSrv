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

type ICodec interface {
	Decode(conn gnet.Conn) ([]byte, error)
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
	bodyOffset := bodySize
	buf, _ := c.Peek(bodyOffset)
	if len(buf) < bodyOffset {
		return nil, ErrIncompletePacket
	}

	bodyLen := binary.BigEndian.Uint32(buf[:bodyOffset])
	msgLen := bodyOffset + int(bodyLen)
	if c.InboundBuffered() < msgLen {
		return nil, ErrIncompletePacket
	}
	buf, _ = c.Peek(msgLen)
	_, _ = c.Discard(msgLen)

	return buf[:msgLen], nil
}

func (code *DefaultCodec) Encode(packet *MsgPacket) (sendBody []byte, err error) {
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

	buffer := &bytes.Buffer{}
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
	return buffer.Bytes(), nil
}

func (codec *DefaultCodec) UnPacket(msgId int16, body *bytes.Buffer, method *protoMethod[int64]) *MsgPacket {
	var shortHeadLen int16
	binary.Read(body, binary.BigEndian, shortHeadLen)
	header := &protoGen.InnerHead{}

	headerBody := make([]byte, shortHeadLen)
	readLen, err := body.Read(headerBody)
	if err != nil {
		return nil
	}

	if readLen != int(shortHeadLen) {
		log.Error(errors.New(fmt.Sprintf(" read =%d but headLen =%d", readLen, shortHeadLen)))
		return nil
	}
	err = proto.Unmarshal(headerBody, header)
	if err != nil {
		log.Error(err)
		return nil
	}

	bodyBytes := body.Bytes()
	msg := method.param.ProtoReflect().New().Interface()
	err = proto.Unmarshal(bodyBytes, msg)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &MsgPacket{MsgId: msgId, Header: header, Body: msg}
}
