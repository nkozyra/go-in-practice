package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")                                                 // # A
	if err := http.ListenAndServe(":8080", http.FileServer(dir)); err != nil { // # B
		panic(err)
	}
}
