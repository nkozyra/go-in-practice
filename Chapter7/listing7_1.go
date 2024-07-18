package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("myfile.txt") // # A
	if err != nil {
		panic(err)
	}
	log.Println(string(data)) // # B
}
