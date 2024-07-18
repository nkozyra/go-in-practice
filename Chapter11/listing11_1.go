package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.manning.com/") // # A
	b, err := io.ReadAll(res.Body)                 // # B
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s", b) // # C
}
