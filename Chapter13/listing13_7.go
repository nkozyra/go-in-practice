package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Walking a simple integer")
	var one MyInt = 1 // # A
	walk(one, 0)      // # A
	fmt.Println("Walking a simple struct")
	two := struct{ Name string }{"foo"} // # B
	walk(two, 0)                        // # B
	fmt.Println("Walking a struct with struct fields")
	p := &Person{ // # C
		Name:    &Name{"Count", "Tyrone", "Rugen"},        // # C
		Address: &Address{"Humperdink Castle", "Florian"}, // # C
	} // # C
	walk(p, 0) // # C
}

type MyInt int
type Person struct {
	Name    *Name
	Address *Address
}
type Name struct {
	Title, First, Last string
}
type Address struct {
	Street, Region string
}

func walk(u interface{}, depth int) { // # D
	val := reflect.Indirect(reflect.ValueOf(u)) // # E
	t := val.Type()                             // # F
	tabs := strings.Repeat("\t", depth+1)       // # G
	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct { // # H
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)                        // # I
			fieldVal := reflect.Indirect(val.Field(i)) // # I
			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is type %q (%s)\n",
				tabs, field.Name, field.Type, fieldVal.Kind())
			if fieldVal.Kind() == reflect.Struct { // # J
				walk(fieldVal.Interface(), depth+1) // # J
			} // # J
		}
	}
}
