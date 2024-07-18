package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handler)                         // # A
	if err := http.ListenAndServe(":8080", nil); err != nil { // # B
		panic("could not start server", err)
	}
}
