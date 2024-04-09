package main

import (
	"bytes"
	"fmt"
	"github.com/Lucaswatt/mcutils"
)

func main() {
	var bbuf bytes.Buffer

	mcutils.WriteString(&bbuf, "Hello world! Here are some unicode characters: 你好, नमस्ते, مرحبا")
	mcutils.WriteString(&bbuf, "this is a second string. Hopefully it works! extra unicode characters for good measure: 你好, नमस्ते, مرحبا")

	fmt.Println(bbuf.Bytes())

	fmt.Println(mcutils.ReadString(&bbuf))
	fmt.Println(mcutils.ReadString(&bbuf))
}
