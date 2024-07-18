package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("DELETE", "http://example.com/foo/bar", nil) // # A
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req) // # B
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res.Status) // # C
}
