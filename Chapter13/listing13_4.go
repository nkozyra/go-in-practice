package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) { // # A
		fmt.Printf("%T is a stringer\n", b)
	}
	i := 123
	if isStringer(i) { // # B
		fmt.Printf("%T is a stringer\n", i)
	}
}
func isStringer(v interface{}) bool { // # C
	_, ok := v.(fmt.Stringer) // # C
	return ok
}
