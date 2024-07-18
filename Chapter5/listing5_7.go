package main

import (
	"os"
	"time"
)

func main() {
	echo := make(chan []byte) // # A
	go readStdin(echo)        // # B
	for {
		select { // # C
		case buf := <-echo: // # D
			os.Stdout.Write(buf) // # D
		case <-time.After(30 * time.Second): // # E
			break // # E
		} // # D
	}
}
func readStdin(out chan<- []byte) { // # F
	for { // # F
		data := make([]byte, 1024)  // # F
		l, _ := os.Stdin.Read(data) // # F
		if l > 0 {
			out <- data // # G
		}
	}
}
