package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteBoolean(&bbuf, true)
	mcutils.WriteBoolean(&bbuf, false)

	mcutils.WriteInt8(&bbuf, 125)
	mcutils.WriteInt8(&bbuf, -74)

	mcutils.WriteUint8(&bbuf, 246)
	mcutils.WriteUint8(&bbuf, 52)

	mcutils.WriteInt16(&bbuf, -12385)
	mcutils.WriteInt16(&bbuf, 32147)

	mcutils.WriteUint16(&bbuf, 2358)
	mcutils.WriteUint16(&bbuf, 62481)

	mcutils.WriteInt32(&bbuf, -324859203)
	mcutils.WriteInt32(&bbuf, 2000231752)

	mcutils.WriteInt64(&bbuf, -8223372036854775808)
	mcutils.WriteInt64(&bbuf, 8724372036854775808)

	fmt.Println(bbuf.Bytes())

	fmt.Println(mcutils.ReadBoolean(&bbuf))
	fmt.Println(mcutils.ReadBoolean(&bbuf))

	fmt.Println(mcutils.ReadInt8(&bbuf))
	fmt.Println(mcutils.ReadInt8(&bbuf))

	fmt.Println(mcutils.ReadUint8(&bbuf))
	fmt.Println(mcutils.ReadUint8(&bbuf))

	fmt.Println(mcutils.ReadInt16(&bbuf))
	fmt.Println(mcutils.ReadInt16(&bbuf))

	fmt.Println(mcutils.ReadUint16(&bbuf))
	fmt.Println(mcutils.ReadUint16(&bbuf))

	fmt.Println(mcutils.ReadInt32(&bbuf))
	fmt.Println(mcutils.ReadInt32(&bbuf))

	fmt.Println(mcutils.ReadInt64(&bbuf))
	fmt.Println(mcutils.ReadInt64(&bbuf))
}
