package main

import (
	"fmt"
)

type character struct {
	name string
}

func (ch *character) fixName() {
	ch.name = "Inigo Montoya"
}

func main() {
	ch := new(character)
	ch.name = "Prince Humperdinck"
	fmt.Println("my name is", ch.name)

	ch.fixName()
	fmt.Println("just kidding, my name is", ch.name)
}
