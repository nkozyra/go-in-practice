package main

import (
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	sleep := time.After(5 * time.Second)
	<-sleep
}
