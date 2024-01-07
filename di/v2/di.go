package main

import "bytes"

func Greet(b *bytes.Buffer, s string) {
	b.WriteString("Hello, ")
	b.WriteString(s)
}
