package main

import (
	"fmt"
)

func getStrings() (string, string) {
	return "Foo", "Bar"
}
func main() {
	n1, n2 := getStrings()
	fmt.Println(n1, n2)
	n3, _ := getStrings()
	fmt.Println(n3)
}
