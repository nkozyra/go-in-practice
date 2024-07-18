package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902") // # A
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()                     // # B
	f := log.Ldate | log.Lshortfile        // # C
	logger := log.New(conn, "example ", f) // # D
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.") // # E
}
