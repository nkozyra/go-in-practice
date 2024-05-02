package main

import (
	"fmt"
)

type Cat struct {
	wearsBow bool
}

type Dog struct {
	canFetch bool
}

type AnimalType interface {
	Cat | Dog
}

type Animal[T AnimalType] struct {
	value       T
	AnimalNoise func() string
}

func (a Animal[T]) Speak() {
	fmt.Println(fmt.Sprintf("we got a %T", a.value))
	fmt.Println(a.AnimalNoise())
}

func main() {
	catAnimal := Animal[Cat]{value: Cat{}, AnimalNoise: func() string { return "meow!" }}
	dogAnimal := Animal[Dog]{value: Dog{}, AnimalNoise: func() string { return "woof!" }}

	catAnimal.Speak()
	dogAnimal.Speak()
}
