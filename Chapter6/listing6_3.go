package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	username  string `json:"username"` // # A
	Email     string `json:"email"`
	firstName string // # A
}

func main() {

	m := user{
		username:  "manning_go",
		Email:     "email@example.com",
		firstName: "Joe",
	}

	out, err := json.Marshal(m) // # B
	if err != nil {
		panic("could not marshal")
	}

	fmt.Println(string(out))

}
