package main

import (
	"fmt"
	"strconv"
)

type MyInt int64

func main() {
	//…
	var d MyInt = 1
	res := sum(a, b, c, d)
	fmt.Printf("Result: %f\n", res)
}
func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case uint8:
			res += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}
