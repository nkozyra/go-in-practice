package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() { // #A
		if err := recover(); err != nil { // #A
			fmt.Printf("Trapped panic: %s (%T)\n", err, err) // #A
		} // #A
	}() // #A
	yikes() // #B
}
func yikes() {
	panic(errors.New("Something bad happened.")) // #C
}
