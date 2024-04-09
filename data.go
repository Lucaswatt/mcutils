package mcutils

import (
	"bytes"
	"errors"
)

const SEGMENT_BITS = 0x7F
const CONTINUE_BIT = 0x80

func readVarInt(bbuf bytes.Buffer) (int32, error) {
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

func readVarLong(bbuf bytes.Buffer) (int64, error) {
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

func writeVarInt(bbuf bytes.Buffer, value int32) {
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

func writeVarLong(bbuf bytes.Buffer, value int64) {
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
