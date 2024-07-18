package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct { // #A
	Enabled bool   // #A
	Path    string // #A
} // #A
func main() {
	file, err := os.Open("config.json") // #B
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file) // #C
	conf := configuration{}          // #C
	err = decoder.Decode(&conf)      // #C
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
