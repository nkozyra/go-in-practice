package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(msg) // # A
	}()
	msg := "Hello world" // # B
}
