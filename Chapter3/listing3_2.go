package main

import (
    "fmt"
)

type Animal struct { // #A
    name string
}

func (a Animal) speak() string { // #B
    switch a.name {
    case "cat":
        return "meow"
    case "dog":
        return "woof"
    default:
        return "nondescript animal noise?"
    }

}

func main() {
    a := Animal{
        name: "cat",
    }

    fmt.Println(a.speak()) // #C

    a.name = "dog"
    fmt.Println(a.speak())// #C

    a.name = "llama"
    fmt.Println(a.speak())#
}
