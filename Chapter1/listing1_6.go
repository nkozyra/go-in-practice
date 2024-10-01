package main

import (
	"fmt"
	"time"
)

func count() { // #A
	for i := 0; i < 5; i++ { // #A
		fmt.Println(i)                   // #A
		time.Sleep(time.Millisecond * 5) // #A
	}
}
func main() {
	go count() // #B
	time.Sleep(time.Millisecond * 20)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 10)
}
