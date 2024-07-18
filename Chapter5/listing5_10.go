package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch) // #A
	for {       // #A
		select {
		case m, ok := <-ch: // #B
			if !ok { // #C
				log.Println("Channel closed.")
				return
			}
			log.Println("Got message:", m)
		case <-timeout: // #D
			log.Println("Time out")
			return
		default: // #E
			log.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func send(ch chan bool) { // #F
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	log.Println("Sent and closed")
}
