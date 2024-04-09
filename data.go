// Written from documentation from https://wiki.vg/Protocol#Data_types
// https://protobuf.dev/programming-guides/encoding/#varints is helpful too
// wiki.vg is honestly a lifesaver

package mcutils

import (
	"bytes"
	"encoding/binary"
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
	var int8Value int8
	err := binary.Read(bbuf, binary.BigEndian, &int8Value)
	return int8Value, err
}

// Write a signed 8 bit integer to a byte buffer (Byte)
func WriteInt8(bbuf *bytes.Buffer, intValue int8) error {
	err := binary.Write(bbuf, binary.BigEndian, intValue)
	return err
}

// Read an unsigned 8 bit integer from a byte buffer (Unsigned byte)
func ReadUint8(bbuf *bytes.Buffer) (uint8, error) {
	var uint8Value uint8
	err := binary.Read(bbuf, binary.BigEndian, &uint8Value)
	return uint8Value, err
}

// Write an unsigned 8 bit integer to a byte buffer (Unsigned byte)
func WriteUint8(bbuf *bytes.Buffer, uintValue uint8) error {
	err := binary.Write(bbuf, binary.BigEndian, uintValue)
	return err
}

// Read a Big-endian signed 16 bit integer from a byte buffer (Short)
func ReadInt16(bbuf *bytes.Buffer) (int16, error) {
	var int16Value int16
	err := binary.Read(bbuf, binary.BigEndian, &int16Value)
	return int16Value, err
}

// Write a Big-endian signed 16 bit integer to a byte buffer (Short)
func WriteInt16(bbuf *bytes.Buffer, intValue int16) error {
	err := binary.Write(bbuf, binary.BigEndian, intValue)
	return err
}

// Read a Big-endian unsigned 16 bit integer from a byte buffer (Unsigned short)
func ReadUint16(bbuf *bytes.Buffer) (uint16, error) {
	var uint16Value uint16
	err := binary.Read(bbuf, binary.BigEndian, &uint16Value)
	return uint16Value, err
}

// Write a Big-endian unsigned 16 bit integer to a byte buffer (Unsigned short)
func WriteUint16(bbuf *bytes.Buffer, uintValue uint16) error {
	err := binary.Write(bbuf, binary.BigEndian, uintValue)
	return err
}

// Read a Big-endian signed 32 bit integer from a byte buffer (Int)
func ReadInt32(bbuf *bytes.Buffer) (int32, error) {
	var int32Value int32
	err := binary.Read(bbuf, binary.BigEndian, &int32Value)
	return int32Value, err
}

// Write a Big-endian signed 32 bit integer to a byte buffer (Int)
func WriteInt32(bbuf *bytes.Buffer, intValue int32) error {
	err := binary.Write(bbuf, binary.BigEndian, intValue)
	return err
}

// Read a Big-endian signed 64 bit integer from a byte buffer (Long)
func ReadInt64(bbuf *bytes.Buffer) (int64, error) {
	var int64Value int64
	err := binary.Read(bbuf, binary.BigEndian, &int64Value)
	return int64Value, err
}

// Write a Big-endian signed 64 bit integer to a byte buffer (Long)
func WriteInt64(bbuf *bytes.Buffer, intValue int64) error {
	err := binary.Write(bbuf, binary.BigEndian, intValue)
	return err
}
