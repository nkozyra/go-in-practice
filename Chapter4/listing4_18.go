package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}
func listen() {
	listener, err := net.Listen("tcp", ":1026") // # A
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for { // # B
		conn, err := listener.Accept() // # B
		if err != nil {                // # B
			fmt.Println("Error accepting connection") // # B
			continue                                  // # B
		} // # B
		go handle(conn) // # C
	}
}
func handle(conn net.Conn) { // # D
	reader := bufio.NewReader(conn)     // # D
	data, err := reader.ReadBytes('\n') // # D
	if err != nil {
		fmt.Println("Failed to read from socket.") // # E
		conn.Close()                               // # E
	}
	response(data, conn) // # F
}
func response(data []byte, conn net.Conn) { // # G
	defer func() { // # G
		conn.Close() // # G
	}() // # G
	conn.Write(data) // # G
} // # G
