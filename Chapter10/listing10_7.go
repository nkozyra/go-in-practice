package main

import (
	"embed"
	"net/http"
)

//go:embed files
var f embed.FS

func main() {
	if err := http.ListenAndServe(":8088", http.FileServer(http.FS(f))); err != nil {
		panic(err)
	}
}
