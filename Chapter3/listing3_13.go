package main

import (
	"fmt"

	"pkg.go.dev/golang.org/x/exp/constraints"
)

type Smallint int8

func doubler[T constraints.Integer](value T) T {
	return value * T(2)
}

func main() {
	var four Smallint = 4
	fmt.Println(doubler(four))
}
