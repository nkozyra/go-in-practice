package main

import (
	"fmt"
)

func getStrings() (first string, second string) { // #A
	first = "Foo"  // #B
	second = "Bar" // #B
	return         // #C
}
func main() {
	n1, n2 := getStrings() // #D
	fmt.Println(n1, n2)
}
