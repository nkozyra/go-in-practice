package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}
func foo() {
	bar()
}
func bar() {
	buf := make([]byte, 1024)        // # A
	runtime.Stack(buf, false)        // # B
	fmt.Printf("Trace:\n %s\n", buf) // # C
}
