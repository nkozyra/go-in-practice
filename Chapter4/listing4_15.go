package main

import "fmt"

func main() {
	var msg string // #A
	defer func() {
		fmt.Println(msg) // #B
	}()
	msg = "Hello world" // #C
}
