package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteString(&bbuf, "hi 你")
	mcutils.WriteString(&bbuf, "world  你好, नमस्ते, مرحبا")

	fmt.Println(bbuf.Bytes())

	fmt.Println(mcutils.ReadString(&bbuf))
	fmt.Println(mcutils.ReadString(&bbuf))
}
