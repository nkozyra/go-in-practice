package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")
	if err := http.ListenAndServe(":8082", http.FileServer(dir)); err != nil {
		panic(err)
	}
}
