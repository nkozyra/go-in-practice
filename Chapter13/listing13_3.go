package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyInt int64

func main() {
	//…var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	var d MyInt = 1
	res := sum(a, b, c, d)
	fmt.Printf("Result: %f\n", res)
}
func sum(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		ref := reflect.ValueOf(val) // # A
		switch ref.Kind() {         // # B
		case reflect.Int, reflect.Int64: // # C
			res += float64(ref.Int()) // # D
		case reflect.Uint8:
			res += float64(ref.Uint()) // # D
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64) // # D
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
