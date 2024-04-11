package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteString(&bbuf, "hi")
	mcutils.WriteString(&bbuf, "world")

	fmt.Println(bbuf.Bytes())

	fmt.Println(mcutils.ReadString(&bbuf))
	fmt.Println(bbuf.Len())
	fmt.Println(mcutils.ReadString(&bbuf))
}
