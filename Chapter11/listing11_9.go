package main

import (
	"encoding/json"
	"fmt"
)

type Person struct { // # A
	Name string `json:"name"` // # A
} // # A

var JSON = `{ // # B
  "name": "Miracle Max" // # B
}` // # B

func main() {
	var p Person                            // # C
	err := json.Unmarshal([]byte(JSON), &p) // # D
	if err != nil {                         // # E
		fmt.Println(err) // # E
		return           // # E
	} // # E
	fmt.Println(p) // # F
}
