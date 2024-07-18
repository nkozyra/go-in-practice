package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a number for fizzbuzz")
	fmt.Scanln(&input) // # A

	numInput, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		panic("that's not a number!")
	}

	result := fizzbuzz(numInput) // # C
	fmt.Println("result:", result)
}

func fizzbuzz(n int64) string {
	var fizzbuzzes []string
	for i := int64(0); i < n; i++ { // # B
		v := ""
		isThree := i%3 == 0
		isFive := i%5 == 0
		if isThree && isFive {
			v = "Fizz Buzz"
		} else if isThree {
			v = "Fizz"
		} else if isFive {
			v = "Buzz "
		} else {
			v = fmt.Sprintf("%d", i)
		}
		fizzbuzzes = append(fizzbuzzes, v)
	}

	return strings.Join(fizzbuzzes, " ")
}
