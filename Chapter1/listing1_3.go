package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80") // #A
	if err != nil {                               // #B
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")           // #C
	status, err := bufio.NewReader(conn).ReadString('\n') // #D
	if err != nil {                                       // #E
		log.Fatal(err)
	}
	log.Println(status) // #F
}
