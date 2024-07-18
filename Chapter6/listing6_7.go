package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("logging.log", os.O_RDWR|os.O_CREATE, 0755) // # A
	if err != nil {
		panic("could not open log file")
	}
	log.SetOutput(file)
	log.SetFlags(log.LUTC | log.Lshortfile)
	log.Println("Display in UTC and use a short filename")
}
