package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteVarInt(&bbuf, 124)
	mcutils.WriteVarLong(&bbuf, 2934857090382947509)
	mcutils.WriteVarInt(&bbuf, 25565)

	fmt.Println(bbuf.Bytes())
	fmt.Println(mcutils.ReadVarInt(&bbuf))
	fmt.Println(mcutils.ReadVarLong(&bbuf))
	fmt.Println(mcutils.ReadVarInt(&bbuf))
}
