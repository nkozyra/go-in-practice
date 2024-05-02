package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed files/example.html
var myString string

func main() {
	log.Println(fmt.Sprintf("embedded value: %s", myString))
}
