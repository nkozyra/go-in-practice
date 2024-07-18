package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // #A

	mux.HandleFunc("/hello", helloHandler)                // #B
	mux.HandleFunc("GET /goodbye/", goodbyeHandler)       // #B
	mux.HandleFunc("GET /goodbye/{name}", goodbyeHandler) // #B

	if err := http.ListenAndServe(":8084", mux); err != nil { // #C
		panic(err)
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()  // #D
	name := query.Get("name") // #D
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbyeHandler(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name") // #E
	if name == "" {               // #E
		name = "Inigo Montoya" // #F
	}
	fmt.Fprint(res, "Goodbye ", name) // #G
}
