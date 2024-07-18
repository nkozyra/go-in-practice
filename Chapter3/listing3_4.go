package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct { // #A
	Name           string  `json:"animal_name"`           // #B
	ScientificName string  `json:"scientific_name"`       // #B
	Weight         float32 `json:"animal_average_weight"` // #B
}

func main() {
	a := Animal{
		Name:           "cat",
		ScientificName: "Felis catus",
		Weight:         10.5,
	} // #C

	output, err := json.Marshal(a)
	if err != nil {
		panic("couldn't encode json")
	}
	fmt.Println(string(output)) // #D
}
