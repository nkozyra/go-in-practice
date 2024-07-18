package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt") // # A
	if err != nil {
		panic(err)
	}
	defer file.Close() // # B

	file.WriteString("test") // # C
}
