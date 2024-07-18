package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ltime) // # A
	log.Println("Only show the time")

	log.SetFlags(log.Llongfile) // # B
	log.Println("Show the full filename")

	log.SetFlags(log.LUTC | log.Lshortfile) // # C
	log.Println("Display in UTC and use a short filename")
}
