package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(msg)
	}()
	msg := "Hello world"
}
