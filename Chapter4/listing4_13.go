package main

import "fmt"

func main() {
	defer goodbye()             // #A
	fmt.Println("Hello world.") // #B
}
func goodbye() {
	fmt.Println("Goodbye")
}
