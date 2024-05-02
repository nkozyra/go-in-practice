package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ltime)
	log.Println("Only show the time")

	log.SetFlags(log.Llongfile)
	log.Println("Show the full filename")

	log.SetFlags(log.LUTC | log.Lshortfile)
	log.Println("Display in UTC and use a short filename")
}
