package main

import (
	"log"
	"time"
)

func main() {
	msg := make(chan string) // #A
	done := make(chan bool)  // #B
	go send(msg, done) // #C
	for {
		select {
		case m := <-msg: // #D
			log.Println(m)
		case <-time.After(5 * time.Second): // #E
			done <- true
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done: // #G
			log.Println("Done")
			close(ch)
			return
		default:
			ch <- "hello" #F
			time.Sleep(500 * time.Millisecond) // #G
		}
	}
}
