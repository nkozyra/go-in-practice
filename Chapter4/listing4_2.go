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
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}
