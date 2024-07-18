package main

import (
	"fmt"
)

type Animal interface { // #A
	speak()
}

type Cat struct{} // #B
func (c Cat) speak() {
	fmt.Println("meow") // #D
}

func NewCat() *Cat { // #C
	return &Cat{}
}

type Dog struct{} // #B
func (d Dog) speak() {
	fmt.Println("woof") // #D
}

func NewDog() *Dog { // #C
	return &Dog{}
}

type Llama struct{} // #B

func NewLlama() *Llama { // #C
	return &Llama{}
}

func main() {

	var a Animal

	c := NewCat()
	a = c
	a.speak()

	d := NewDog()
	a = d
	a.speak()

	l := NewLlama()
	a = l
}
