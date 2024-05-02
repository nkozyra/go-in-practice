package main

import (
	"log"

	"github.com/nkozyra/ch11"
)

func main() {
	if err := ch11.ConnectToGRPC(); err != nil {
		log.Fatal(err)
	}
}
