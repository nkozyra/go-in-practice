package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	res := sum(a, b, c)
	fmt.Printf("Result: %f\n", res)
}

type myString string

func (s myString) ToFloat() (float64, error) {
	fl, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return fl, nil
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
			var str myString = myString(val.(string))
			value, err := str.ToFloat()
			if err != nil {
				panic(err)
			}
			res += value
		default:
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}
