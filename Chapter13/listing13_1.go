package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	res := sum(a, b, c) // # A
	fmt.Printf("Result: %f\n", res)
}
func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v { // # B
		switch val.(type) { // # B
		case int: // # C
			res += float64(val.(int)) // # C
		case int64: // # C
			res += float64(val.(int64)) // # C
		case uint8: // # C
			res += float64(val.(uint8)) // # C
		case string: // # D
			a, err := strconv.ParseFloat(val.(string), 64) // # D
			if err != nil {                                // # D
				panic(err) // # D
			} // # D
			res += a // # D
		default: // # E
			fmt.Printf("Unsupported type %T. Ignoring.\n", val) // # E
		}
	}
	return res
}
