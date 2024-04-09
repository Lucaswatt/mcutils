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

// Read a boolean value from a byte buffer
func ReadBoolean(bbuf *bytes.Buffer) (bool, error) {
	boolByte, err := bbuf.ReadByte()
	if err != nil {
		return false, err
	}

	if boolByte == 0x00 {
		return false, nil
	} else {
		return true, nil
	}
}

// Write a boolean value to a byte buffer
func WriteBoolean(bbuf *bytes.Buffer, boolValue bool) error {
	var boolByte byte
	if boolValue {
		boolByte = 0x01
	} else {
		boolByte = 0x00
	}

	err := bbuf.WriteByte(boolByte)
	return err
}

// Read a signed 8 bit integer from a byte buffer (Byte)
func ReadInt8(bbuf *bytes.Buffer) (int8, error) {
	intByte, err := bbuf.ReadByte()
	if err != nil {
		return 0, err
	}

	return int8(intByte), nil
}

// Write a signed 8 bit integer to a byte buffer (Byte)
func WriteInt8(bbuf *bytes.Buffer, intValue int8) error {
	err := bbuf.WriteByte(byte(intValue))
	return err
}

// Read an unsigned 8 bit integer from a byte buffer (Unsigned byte)
func ReadUint8(bbuf *bytes.Buffer) (uint8, error) {
	intByte, err := bbuf.ReadByte()
	if err != nil {
		return 0, err
	}

	return uint8(intByte), nil
}

// Write an unsigned 8 bit integer to a byte buffer (Unsigned byte)
func WriteUint8(bbuf bytes.Buffer, uintValue uint8) error {
	err := bbuf.WriteByte(uintValue)
	return err
}
