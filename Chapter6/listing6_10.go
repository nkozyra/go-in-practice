package main

import (
	"runtime/debug"
)

func main() { // # A
	foo()
}
func foo() { // # A
	bar()
}
func bar() { // # A
	debug.PrintStack() // # B
}
