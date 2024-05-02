package main

import (
	"log"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)
	defer close(done)
	until := time.After(5 * time.Second)
	go send(msg, done)
	// for {
	// 	select {
	// 	case m := <-msg:
	// 		fmt.Println("Message:", m)
	// 	case <-until:
	// 		done <- true
	// 		time.Sleep(500 * time.Millisecond)
	// 		return
	// 	}
	// }
	for {
		select {
		case <-until:
			done <- true
			return
		}
	}
	for v := range until {
		log.Println(v)
	}
	log.Println("hey")
}
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			close(ch)
			return
		default:
			ch <- "message from send"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
