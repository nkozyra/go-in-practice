package main

import (
	"log"
	"time"
)

func main() {
	lock := make(chan bool, 1) // #A
	for i := 1; i < 7; i++ {   // #B
		go worker(i, lock) // #B
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	log.Printf("%d wants the lock\n", id)
	lock <- true                        // #C
	log.Printf("%d has the lock\n", id) // #D
	<-lock                              // #D
	log.Printf("%d is releasing the lock\n", id)
}
