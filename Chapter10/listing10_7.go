package main

import (
	"embed" // # A
	"net/http"
)

//go:embed files // # B
var f embed.FS // # C

func main() {
	if err := http.ListenAndServe(":8088", http.FileServer(http.FS(f))); err != nil { // # D
		panic(err)
	}
}
