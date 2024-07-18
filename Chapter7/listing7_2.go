package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("structured.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, err := file.Stat() // # A
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("File: name is %s, mode is %v, size is %d. Is directory: %v", info.Name(), info.Mode(), info.Size(), info.IsDir())) // # B
}
