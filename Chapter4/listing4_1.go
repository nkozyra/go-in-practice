package main

import (
	"errors" // #A
	"strings"
)

func Concat(parts ...string) (string, error) { // #B
	if len(parts) == 0 {
		return "", errors.New("No strings supplied") // #C
	} // #C
	return strings.Join(parts, " "), nil // #D
}
