package main

import (
	"time"
)

func main() {
	time.Sleep(5 * time.Second)          // # A
	sleep := time.After(5 * time.Second) // # B
	<-sleep                              // # B
}
