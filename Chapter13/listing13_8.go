package main

import (
	"log"
	"strings"

	"github.com/bytedance/sonic"
)

func main() {
	var receiver = []map[string]interface{}{}
	jsonData := strings.NewReader(`
		[
			{ "email": "inigo@montoya.example.com", "name": "Inigo Montoya" },
			{ "email": "fezzik@example.com", "name": "Fezzik" },
		]
	`)
	decoder := sonic.ConfigDefault.NewDecoder(jsonData)
	decoder.Decode(&receiver)
	for k := range receiver {
		log.Println(receiver[k])
	}
}
