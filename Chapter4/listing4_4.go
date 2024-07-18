package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}
	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]
	result, _ := Concat(args...) // #A
	fmt.Printf("Concatenated string: '%s'\n", result)
}
