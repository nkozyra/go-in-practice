package main

import (
	"fmt"
	"io"
	"reflect"
)

type Name struct { // # A
	First, Last string
}

func (n *Name) String() string { // # A
	return n.First + " " + n.Last
}
func main() { // # A
	n := &Name{First: "Inigo", Last: "Montoya"}
	stringer := (*fmt.Stringer)(nil) // # B
	implements(n, stringer)          // # C
	writer := (*io.Writer)(nil)      // # D
	implements(n, writer)            // # E
}
func implements(concrete interface{}, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem() // # F
	v := reflect.ValueOf(concrete)         // # G
	t := v.Type()                          // # G
	if t.Implements(iface) {               // # H
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("%T is not a %s\n", concrete, iface.Name())
	return false
}
