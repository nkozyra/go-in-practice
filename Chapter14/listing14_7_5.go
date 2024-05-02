package main

import (
	"log"
	"reflect"
)

type MyType struct {
	Name string
}

func main() {
	myType := MyType{Name: "my name value"}

	myTypeStruct := reflect.TypeOf(myType)
	log.Println(myTypeStruct.MethodByName("GetName"))
}
