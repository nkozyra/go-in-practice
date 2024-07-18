package main

import (
	"log"
	"strings"

	"github.com/bytedance/sonic" // # A
)

func main() {
	var receiver = []map[string]interface{}{} // # B
	jsonData := strings.NewReader(`
        [
            { "email": "inigo@montoya.example.com", "name": "Inigo Montoya" },
            { "email": "fezzik@example.com", "name": "Fezzik" },
        ]
    `)
	decoder := sonic.ConfigDefault.NewDecoder(jsonData) // # C
	decoder.Decode(&receiver)                           // # D
	for k := range receiver {
		log.Println(receiver[k]) // # E
	}
}
