// Written from documentation from https://wiki.vg/Protocol#Data_types
// https://protobuf.dev/programming-guides/encoding/#varints is helpful too
// wiki.vg is honestly a lifesaver

package mcutils

import (
	"bytes"
	"errors"
)

const SEGMENT_BITS = 0x7F
const CONTINUE_BIT = 0x80

// Takes a reference to a byte buffer object and returns an int32
func ReadVarInt(bbuf *bytes.Buffer) (int32, error) {
	var value, position int32
	var currentByte byte
	var err error

	for {
		currentByte, err = bbuf.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= int32(currentByte&SEGMENT_BITS) << position

		if (currentByte & CONTINUE_BIT) == 0 {
			break
		}

		position += 7

		if position >= 32 {
			return 0, errors.New("VarInt is too big")
		}
	}

	return value, nil
}

// Takes a reference to a byte buffer and returns an int64
func ReadVarLong(bbuf *bytes.Buffer) (int64, error) {
	var value int64
	var position int
	var currentByte byte
	var err error

	for {
		currentByte, err = bbuf.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= int64(currentByte&SEGMENT_BITS) << position

		if (currentByte & CONTINUE_BIT) == 0 {
			break
		}

		position += 7

		if position >= 64 {
			return 0, errors.New("VarLong is too big")
		}
	}

	return value, nil
}

// Writes a VarInt to a byte buffer, using a reference
func WriteVarInt(bbuf *bytes.Buffer, value int32) {
	for {
		if (value &^ SEGMENT_BITS) == 0 {
			bbuf.WriteByte(byte(value))
			return
		}

		bbuf.WriteByte(byte((value & SEGMENT_BITS) | CONTINUE_BIT))

		// Not tested, might not work as expected. The sign bit needs to be shifted with the rest of the number rather than being left alone.
		value >>= 7
	}
}

// Writes a VarLong to a byte buffer, using a reference
func WriteVarLong(bbuf *bytes.Buffer, value int64) {
	for {
		if (value &^ SEGMENT_BITS) == 0 {
			bbuf.WriteByte(byte(value))
			return
		}

		bbuf.WriteByte(byte((value & SEGMENT_BITS) | CONTINUE_BIT))

		// Not tested, might not work as expected. The sign bit needs to be shifted with the rest of the number rather than being left alone.
		value >>= 7
	}
}
