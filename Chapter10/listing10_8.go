package main

import (
	_ "embed" // # A
	"fmt"
	"log"
)

//go:embed files/example.html // # B
var myString string // # C

func main() {
	log.Println(fmt.Sprintf("embedded value: %s", myString)) // # D
}
