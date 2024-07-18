package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]                             // #A
	if result, err := Concat(args...); err != nil { // #B
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result) // #C
	}
}
