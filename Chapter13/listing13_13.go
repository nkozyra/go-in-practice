package main

import ( // # A

	"fmt"
)

type Processes struct { // # B
	Total    int     `ini:"total"`
	Running  int     `ini:"running"`
	Sleeping int     `ini:"sleeping"`
	Threads  int     `ini:"threads"`
	Load     float64 `ini:"load"`
}

func main() {
	fmt.Println("Write a struct to output:")
	proc := &Processes{ // # C
		Total:    23,  // # C
		Running:  3,   // # C
		Sleeping: 20,  // # C
		Threads:  34,  // # C
		Load:     1.8, // # C
	} // # C
	data, err := Marshal(proc) // # D
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // # E
	fmt.Println("Read the data back into a struct")
	proc2 := &Processes{}                          // # F
	if err := Unmarshal(data, proc2); err != nil { // # F
		panic(err)
	}
	fmt.Printf("Struct: %// # v", proc2) // # G
}
