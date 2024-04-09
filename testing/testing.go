package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteFloat32(&bbuf, 3.14159265358)
	mcutils.WriteFloat64(&bbuf, 3.14159265358)
	mcutils.WriteFloat32(&bbuf, 2198289)
	mcutils.WriteFloat64(&bbuf, 230497026223456.2)

	fmt.Println(bbuf.Bytes())

	fmt.Println(mcutils.ReadFloat32(&bbuf))
	fmt.Println(mcutils.ReadFloat64(&bbuf))
	fmt.Println(mcutils.ReadFloat32(&bbuf))
	fmt.Println(mcutils.ReadFloat64(&bbuf))
}
