package main

import (
	"fmt"
)

func getStrings() (string, string) { //#A
	return "Foo", "Bar" // #B
}
func main() {
	n1, n2 := getStrings() // #C
	fmt.Println(n1, n2)    // #C
	n3, _ := getStrings()  // #D
	fmt.Println(n3)        // #D
}
