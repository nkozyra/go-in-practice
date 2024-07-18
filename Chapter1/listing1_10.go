package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseName(name string) string {
	reversed := make([]byte, 0)           // #E
	for i := len(name) - 1; i >= 0; i-- { // #F
		reversed = append(reversed, name[i])
	}
	return string(reversed) // #G
}

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)  // #A
	name, err := reader.ReadString('\n') // #A
	if err != nil {
		log.Fatal("could not read from stdin", err)
	}
	name = strings.TrimSpace(name) // #B

	reversed := reverseName(name)   // #C
	fmt.Println(reversed, ",olleH") // #D
}
