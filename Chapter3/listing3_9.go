package main

import (
	"fmt"
	"unicode"
)

func filter[T any](items []T, fx func(T) bool) []T { // #A
	var filtered []T
	for _, v := range items {
		if fx(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	strings := []string{"My", "name", "is", "Inigo", "Montoya"}

	strings = filter[string](strings, func(s string) bool { // #B
		return unicode.IsUpper(rune(s[0]))
	})

	fmt.Println(strings)
}
