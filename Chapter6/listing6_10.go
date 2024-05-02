package main

import (
	"runtime/debug"
)

func main() {
	foo()
}
func foo() {
	bar()
}
func bar() {
	debug.PrintStack()
}
