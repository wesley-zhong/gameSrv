package network

import (
	"encoding/binary"
	"github.com/panjf2000/gnet"
)

type LengthFieldBasedFrameCodecEx struct {
	codeC *gnet.LengthFieldBasedFrameCodec
}

func NewLengthFieldBasedFrameCodecEx() *LengthFieldBasedFrameCodecEx {
	return &LengthFieldBasedFrameCodecEx{codeC: gnet.NewLengthFieldBasedFrameCodec(gnet.EncoderConfig{
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}, gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   4,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 0})}
}

func (cc *LengthFieldBasedFrameCodecEx) Encode(c gnet.Conn, buf []byte) (out []byte, err error) {
	return buf, nil
}

func (cc *LengthFieldBasedFrameCodecEx) Decode(c gnet.Conn) ([]byte, error) {
	return cc.codeC.Decode(c)
}
