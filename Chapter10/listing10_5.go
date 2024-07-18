package main

import (
	"fmt"
	"net/http"

	fs "github.com/Masterminds/go-fileserver" // # A
)

func main() {
	fs.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) { // # B
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // # B
		fmt.Fprintln(w, "The requested page could not be found.")   // # B
	} // # B

	dir := http.Dir("./files")                       // # C
	http.ListenAndServe(":8080", fs.FileServer(dir)) // # D
}
