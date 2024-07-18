package main

import (
	"fmt"
)

type Numeric interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

type Smallint int8

func doubler[T Numeric](value T) T {
	return value * 2
}

func main() {
	var four Smallint = 4
	fmt.Println(doubler(four))
}
