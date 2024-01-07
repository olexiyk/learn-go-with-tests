package main

import (
	"bytes"
	"fmt"
)

func Countdown(buffer *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(buffer, i)
	}
	fmt.Fprint(buffer, "Go!")
}
