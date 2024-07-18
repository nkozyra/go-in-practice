package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)
	go send(msg) // # A
	for {        // # B
		select {
		case m := <-msg: // # C
			fmt.Println(m)
		case <-until: // # D
			close(msg)                         // # D
			time.Sleep(500 * time.Millisecond) // # D
			return                             // # D
		}
	}
}
func send(ch chan string) { // # E
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}
