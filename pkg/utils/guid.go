package utils

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"net"
	"sync"
	"time"
)

// UUID epoch: October 15, 1582 (RFC 4122)
const uuidEpoch = 122192928000000000

var (
	nodeID     [6]byte
	clockSeq  uint16
	once      sync.Once
	timeLast  uint64
)

func initNode() {
	// Get first available non-loopback interface MAC address
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) == 0 {
		// Fallback: generate random node ID with multicast bit set
		_, _ = rand.Read(nodeID[:])
		nodeID[0] |= 0x01 // Set multicast bit
		return
	}

	for _, iface := range interfaces {
		if len(iface.HardwareAddr) >= 6 && iface.Flags&net.FlagLoopback == 0 {
			copy(nodeID[:], iface.HardwareAddr[:6])
			return
		}
	}

	// Fallback: generate random node ID with multicast bit set
	_, _ = rand.Read(nodeID[:])
	nodeID[0] |= 0x01
}

func initClockSeq() {
	b := make([]byte, 2)
	_, _ = rand.Read(b)
	clockSeq = binary.BigEndian.Uint16(b)
}

// GenGUID generates a UUID v1 (time-based) and returns it without hyphens.
// The returned string is 32 hexadecimal characters.
// UUID v1 combines timestamp + clock sequence + node ID (MAC address),
// providing strong uniqueness guarantees in distributed systems.
func GenGUID() string {
	once.Do(func() {
		initNode()
		initClockSeq()
	})

	now := time.Now().UnixNano() / 100 // convert to 100-nanosecond units
	timestamp := uint64(now) + uuidEpoch

	var timestampNow uint64
	var clockSeqNow uint16

	for {
		timestampNow = atomicAddTime(timestamp)
		if timestampNow != 0 {
			break
		}
		time.Sleep(time.Microsecond)
	}

	clockSeqNow = clockSeq

	uuid := make([]byte, 16)

	// time_low (32 bits)
	uuid[0] = byte(timestampNow >> 24)
	uuid[1] = byte(timestampNow >> 16)
	uuid[2] = byte(timestampNow >> 8)
	uuid[3] = byte(timestampNow)

	// time_mid (16 bits)
	uuid[4] = byte(timestampNow >> 40)
	uuid[5] = byte(timestampNow >> 32)

	// time_hi_and_version (16 bits): version 1
	uuid[6] = byte(timestampNow>>56)&0x0F | 0x10
	uuid[7] = byte(timestampNow >> 48)

	// clock_seq_hi_and_reserved (8 bits): variant
	uuid[8] = byte(clockSeqNow>>8)&0x3F | 0x80

	// clock_seq_low (8 bits)
	uuid[9] = byte(clockSeqNow)

	// node (48 bits)
	copy(uuid[10:], nodeID[:])

	return hex.EncodeToString(uuid)
}

func atomicAddTime(timestamp uint64) uint64 {
	if timestamp <= timeLast {
		clockSeq++
	}
	timeLast = timestamp
	return timeLast
}