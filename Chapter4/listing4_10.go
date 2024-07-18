package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")   // #A
	_, err := precheckDivide(1, 0) // #A
	if err != nil {                // #A
		fmt.Printf("Error: %s\n", err) // #A
	} // #A
	fmt.Println("Divide 2 by 0") // #B
	divide(2, 0)                 // #B
}
func precheckDivide(a, b int) (int, error) { // #C
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}
func divide(a, b int) int { // #D
	return a / b
}
