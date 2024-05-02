package main

import (
	"fmt"
)

type Processes struct {
	Total    int     `ini:"total"`
	Running  int     `ini:"running"`
	Sleeping int     `ini:"sleeping"`
	Threads  int     `ini:"threads"`
	Load     float64 `ini:"load"`
}

func main() {
	fmt.Println("Write a struct to output:")
	proc := &Processes{
		Total:    23,
		Running:  3,
		Sleeping: 20,
		Threads:  34,
		Load:     1.8,
	}
	data, err := Marshal(proc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println("Read the data back into a struct")
	proc2 := &Processes{}
	if err := Unmarshal(data, proc2); err != nil {
		panic(err)
	}
	fmt.Printf("Struct: %", proc2)
}
