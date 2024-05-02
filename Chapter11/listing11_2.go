package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("DELETE", "http://example.com/foo/bar", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res.Status)
}
