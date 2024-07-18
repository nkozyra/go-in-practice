package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) { // #A
	num := 0
	for num >= 0 {
		num = <-c // #B
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) // #C
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	go printCount(c)      // #D
	for _, v := range a { // #E
		c <- v // #E
	}
	time.Sleep(time.Millisecond * 1) // #F
	fmt.Println("End of main")
}
