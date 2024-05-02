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

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	lineJSON := make(map[string]interface{})
	for scan.Scan() {
		if err := json.Unmarshal([]byte(scan.Text()), &lineJSON); err != nil {
			log.Println(err)
		} else {
			log.Println(lineJSON["level"])
		}
	}
}
