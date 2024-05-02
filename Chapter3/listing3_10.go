package main

import (
	"fmt"
)

func filter[T any](items []T, fx func(T) bool) []T {
	var filtered []T
	for _, v := range items {
		if fx(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	ints = filter[int](ints, func(i int) bool {
		if i%3 == 0 {
			return true
		}
		return false
	})

	fmt.Println(ints)
}
