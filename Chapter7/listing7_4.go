package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {
	file, err := os.Open("structured.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file) // # A
	scan.Split(bufio.ScanLines)    // # B
	lineJSON := make(map[string]interface{})
	for scan.Scan() { // # C
		if err := json.Unmarshal([]byte(scan.Text()), &lineJSON); err != nil {
			log.Println(err)
		} else {
			log.Println(lineJSON["level"])
		}
	}
}
