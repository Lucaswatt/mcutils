// Written from documentation from https://wiki.vg/Protocol#Data_types
// https://protobuf.dev/programming-guides/encoding/#varints is helpful too
// wiki.vg is honestly a lifesaver

package mcutils

import (
	"encoding/binary"
	"errors"
	"io"
	"unicode/utf8"
)

const SEGMENT_BITS = 0x7F
const CONTINUE_BIT = 0x80

// Reads a variable length integer from an io stream, returns an int32
func ReadVarInt(r io.Reader) (int32, error) {
	var value, position int32
	var currentByte byte

	for {
		err := binary.Read(r, binary.BigEndian, &currentByte)
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

// Reads a variable length long (int64) from an io stream, returns an int64
func ReadVarLong(r io.Reader) (int64, error) {
	var value int64
	var position int
	var currentByte byte

	for {
		err := binary.Read(r, binary.BigEndian, &currentByte)
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

// Writes a Variable length integer to an io stream
func WriteVarInt(w io.Writer, value int32) error {
	for {
		if (value &^ SEGMENT_BITS) == 0 {
			return binary.Write(w, binary.BigEndian, byte(value))
		}

		err := binary.Write(w, binary.BigEndian, byte((value&SEGMENT_BITS)|CONTINUE_BIT))
		if err != nil {
			return err
		}

		value >>= 7
	}
}

// Writes a Variable length long (int64) to an io stream
func WriteVarLong(w io.Writer, value int64) error {
	for {
		if (value &^ SEGMENT_BITS) == 0 {
			return binary.Write(w, binary.BigEndian, byte(value))
		}

		err := binary.Write(w, binary.BigEndian, byte((value&SEGMENT_BITS)|CONTINUE_BIT))
		if err != nil {
			return err
		}

		value >>= 7
	}
}

// Read a boolean value from an io stream
func ReadBoolean(r io.Reader) (bool, error) {
	var boolByte byte
	err := binary.Read(r, binary.BigEndian, &boolByte)
	if err != nil {
		return false, err
	}

	return boolByte != 0x00, nil
}

// Write a boolean value to an io stream
func WriteBoolean(w io.Writer, boolValue bool) error {
	var boolByte byte
	if boolValue {
		boolByte = 0x01
	} else {
		boolByte = 0x00
	}

	err := binary.Write(w, binary.BigEndian, boolByte)
	return err
}

// Read a signed 8 bit integer from a byte buffer (Byte)
func ReadInt8(r io.Reader) (int8, error) {
	var int8Value int8
	err := binary.Read(r, binary.BigEndian, &int8Value)
	return int8Value, err
}

// Write a signed 8 bit integer to a byte buffer (Byte)
func WriteInt8(w io.Writer, intValue int8) error {
	err := binary.Write(w, binary.BigEndian, intValue)
	return err
}

// Read an unsigned 8 bit integer from a byte buffer (Unsigned byte)
func ReadUint8(r io.Reader) (uint8, error) {
	var uint8Value uint8
	err := binary.Read(r, binary.BigEndian, &uint8Value)
	return uint8Value, err
}

// Write an unsigned 8 bit integer to a byte buffer (Unsigned byte)
func WriteUint8(w io.Writer, uintValue uint8) error {
	err := binary.Write(w, binary.BigEndian, uintValue)
	return err
}

// Read a Big-endian signed 16 bit integer from a byte buffer (Short)
func ReadInt16(r io.Reader) (int16, error) {
	var int16Value int16
	err := binary.Read(r, binary.BigEndian, &int16Value)
	return int16Value, err
}

// Write a Big-endian signed 16 bit integer to a byte buffer (Short)
func WriteInt16(w io.Writer, intValue int16) error {
	err := binary.Write(w, binary.BigEndian, intValue)
	return err
}

// Read a Big-endian unsigned 16 bit integer from a byte buffer (Unsigned short)
func ReadUint16(r io.Reader) (uint16, error) {
	var uint16Value uint16
	err := binary.Read(r, binary.BigEndian, &uint16Value)
	return uint16Value, err
}

// Write a Big-endian unsigned 16 bit integer to a byte buffer (Unsigned short)
func WriteUint16(w io.Writer, uintValue uint16) error {
	err := binary.Write(w, binary.BigEndian, uintValue)
	return err
}

// Read a Big-endian signed 32 bit integer from a byte buffer (Int)
func ReadInt32(r io.Reader) (int32, error) {
	var int32Value int32
	err := binary.Read(r, binary.BigEndian, &int32Value)
	return int32Value, err
}

// Write a Big-endian signed 32 bit integer to a byte buffer (Int)
func WriteInt32(w io.Writer, intValue int32) error {
	err := binary.Write(w, binary.BigEndian, intValue)
	return err
}

// Read a Big-endian signed 64 bit integer from a byte buffer (Long)
func ReadInt64(r io.Reader) (int64, error) {
	var int64Value int64
	err := binary.Read(r, binary.BigEndian, &int64Value)
	return int64Value, err
}

// Write a Big-endian signed 64 bit integer to a byte buffer (Long)
func WriteInt64(w io.Writer, intValue int64) error {
	err := binary.Write(w, binary.BigEndian, intValue)
	return err
}

// Read a Big-endian single-precision 32-bit IEEE 754 floating point number from a byte buffer (Float)
func ReadFloat32(r io.Reader) (float32, error) {
	var float32Value float32
	err := binary.Read(r, binary.BigEndian, &float32Value)
	return float32Value, err
}

// Write a Big-endian single-precision 32-bit IEEE 754 floating point number to a byte buffer (Float)
func WriteFloat32(w io.Writer, floatVlaue float32) error {
	err := binary.Write(w, binary.BigEndian, floatVlaue)
	return err
}

// Read a Big-endian double-precision 64-bit IEEE 754 floating point number from a byte buffer (Double)
func ReadFloat64(r io.Reader) (float64, error) {
	var float64Value float64
	err := binary.Read(r, binary.BigEndian, &float64Value)
	return float64Value, err
}

// Write a Big-endian double-precision 64-bit IEEE 754 floating point number to a byte buffer (Double)
func WriteFloat64(w io.Writer, floatVlaue float64) error {
	err := binary.Write(w, binary.BigEndian, floatVlaue)
	return err
}

// Read a unicode string from a byte buffer, assuming the string is prefixed by a VarInt with the length of the string
// Had to implement my own utf8 reading system because there is no easy way of reading runes directly from an io.Reader
func ReadString(r io.Reader) (string, error) {
	stringLength, err := ReadVarInt(r)

	if err != nil {
		return "", err
	}

	if stringLength < 0 || stringLength > 32767 {
		return "", errors.New("Invalid string length")
	}

	stringRunes := make([]rune, 0, stringLength)

	for i := 0; i < int(stringLength); i++ {
		var startByte byte
		binary.Read(r, binary.BigEndian, &startByte)

		var utf8ByteCount int
		if startByte&0x80 == 0 { // bitwise AND with 0b1000000 to isolate first bit
			utf8ByteCount = 1
		} else if startByte&0xE0 == 0xC0 {
			utf8ByteCount = 2
		} else if startByte&0xF0 == 0xE0 {
			utf8ByteCount = 3
		} else if startByte&0xF8 == 0xF0 {
			utf8ByteCount = 4
		} else {
			return "", errors.New("Error when reading utf8 rune")
		}

		remainingBytes := make([]byte, utf8ByteCount-1)
		binary.Read(r, binary.BigEndian, &remainingBytes)
		bytes := append([]byte{startByte}, remainingBytes...)
		currentRune, _ := utf8.DecodeRune(bytes)

		stringRunes = append(stringRunes, currentRune)
	}

	return string(stringRunes), nil
}

// Write a unicode string to a byte buffer according to the protocol specification
func WriteString(w io.Writer, str string) error {
	stringLength := utf8.RuneCountInString(str)
	WriteVarInt(w, int32(stringLength))

	_, err := io.WriteString(w, str)
	return err
}
