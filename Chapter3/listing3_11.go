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

type Numeric interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func filterPositive[T Numeric](items []T) []T {
	var filtered []T
	for _, v := range items {
		if v > 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	ints := []int8{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	ints = filterPositive[int8](ints)
	fmt.Println(ints)

	floats := []float32{-4.5, -3.5, -2.5, -1.5, 0.5, 1.5, 2.5, 3.5, 4.5}
	floats = filterPositive[float32](floats)
	fmt.Println(floats)
}
