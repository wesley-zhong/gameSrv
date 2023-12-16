package tcp

import (
	"encoding/binary"
	"errors"
	"gameSrv/pkg/log"
	"github.com/panjf2000/gnet"
	gerrors "github.com/panjf2000/gnet/pkg/errors"
)

var MaxPackageLen = 2048 * 2048

// EncoderConfig config for encoder.
type EncoderConfig struct {
	// ByteOrder is the ByteOrder of the length field.
	ByteOrder binary.ByteOrder
	// LengthFieldLength is the length of the length field.
	LengthFieldLength int
	// LengthAdjustment is the compensation value to add to the value of the length field
	LengthAdjustment int
	// LengthIncludesLengthFieldLength is true, the length of the prepended length field is added to the value of
	// the prepended length field
	LengthIncludesLengthFieldLength bool
}

// DecoderConfig config for decoder.
type DecoderConfig struct {
	// ByteOrder is the ByteOrder of the length field.
	ByteOrder binary.ByteOrder
	// LengthFieldOffset is the offset of the length field
	LengthFieldOffset int
	// LengthFieldLength is the length of the length field
	LengthFieldLength int
	// LengthAdjustment is the compensation value to add to the value of the length field
	LengthAdjustment int
	// InitialBytesToStrip is the number of first bytes to strip out from the decoded frame
	InitialBytesToStrip int
	PackageMaxLen       uint64
}

type LengthFieldBasedFrameCodecEx struct {
	encoderConfig *EncoderConfig
	decoderConfig *DecoderConfig
}

func (cc *LengthFieldBasedFrameCodecEx) Encode(c gnet.Conn, buf []byte) (out []byte, err error) {
	return buf, nil
}

func (cc *LengthFieldBasedFrameCodecEx) Decode(c gnet.Conn) ([]byte, error) {
	var (
		in     innerBuffer
		header []byte
		err    error
	)
	in = c.Read()
	if cc.decoderConfig.LengthFieldOffset > 0 { // discard header(offset)
		header, err = in.readN(cc.decoderConfig.LengthFieldOffset)
		if err != nil {
			return nil, gerrors.ErrIncompletePacket
		}
	}

	lenBuf, frameLength, err := cc.getUnadjustedFrameLength(&in)
	if err != nil {
		if errors.Is(err, gerrors.ErrUnsupportedLength) {
			return nil, err
		}
		return nil, gerrors.ErrIncompletePacket
	}
	if frameLength > cc.decoderConfig.PackageMaxLen {
		err := errors.New("XXXXXX msg len too large  error ")
		log.Warnf("XXXXXX %s msg len too large  error maxLen=%d receiveLen=%d should be closed", c.RemoteAddr(), MaxPackageLen, frameLength)
		c.Close()
		return nil, err
	}

	// real message length
	msgLength := int(frameLength) + cc.decoderConfig.LengthAdjustment
	msg, err := in.readN(msgLength)
	if err != nil {
		return nil, gerrors.ErrIncompletePacket
	}

	fullMessage := make([]byte, len(header)+len(lenBuf)+msgLength)
	copy(fullMessage, header)
	copy(fullMessage[len(header):], lenBuf)
	copy(fullMessage[len(header)+len(lenBuf):], msg)
	c.ShiftN(len(fullMessage))
	return fullMessage[cc.decoderConfig.InitialBytesToStrip:], nil
}

type innerBuffer []byte

func (in *innerBuffer) readN(n int) (buf []byte, err error) {
	if n == 0 {
		return nil, nil
	}

	if n < 0 {
		return nil, errors.New("negative length is invalid")
	} else if n > len(*in) {
		return nil, errors.New("exceeding buffer length")
	}
	buf = (*in)[:n]
	*in = (*in)[n:]
	return
}

func (cc *LengthFieldBasedFrameCodecEx) getUnadjustedFrameLength(in *innerBuffer) ([]byte, uint64, error) {
	switch cc.decoderConfig.LengthFieldLength {
	case 2:
		lenBuf, err := in.readN(2)
		if err != nil {
			return nil, 0, gerrors.ErrUnexpectedEOF
		}
		return lenBuf, uint64(cc.decoderConfig.ByteOrder.Uint16(lenBuf)), nil
	case 4:
		lenBuf, err := in.readN(4)
		if err != nil {
			return nil, 0, gerrors.ErrUnexpectedEOF
		}
		return lenBuf, uint64(cc.decoderConfig.ByteOrder.Uint32(lenBuf)), nil
	default:
		return nil, 0, gerrors.ErrUnsupportedLength
	}
}

// NewLengthFieldBasedFrameCodecEx for client
func NewLengthFieldBasedFrameCodecEx() *LengthFieldBasedFrameCodecEx {
	return &LengthFieldBasedFrameCodecEx{&EncoderConfig{
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}, &DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   4,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 0,
		PackageMaxLen:       uint64(MaxPackageLen)}}
}

// NewInnerLengthFieldBasedFrameCodecEx for inner client
func NewInnerLengthFieldBasedFrameCodecEx() *LengthFieldBasedFrameCodecEx {
	return &LengthFieldBasedFrameCodecEx{&EncoderConfig{
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}, &DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 4,
		PackageMaxLen:       uint64(MaxPackageLen),
	}}
}
