package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", helloHandler)      // #A
	http.HandleFunc("/goodbye/", goodbyeHandler) // #A
	http.HandleFunc("/", homePageHandler)        // #A
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} // #B
}
func helloHandler(res http.ResponseWriter, req *http.Request) { // #C
	query := req.URL.Query()  // #D
	name := query.Get("name") // #D
	if name == "" {           // #D
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
func goodbyeHandler(res http.ResponseWriter, req *http.Request) { // #E
	path := req.URL.Path              // #F
	parts := strings.Split(path, "/") // #F
	name := parts[2]                  // #F
	if name == "" {                   // #F
		name = "Inigo Montoya" // #F
	} // #F
	fmt.Fprint(res, "Goodbye ", name)
}
func homePageHandler(res http.ResponseWriter, req *http.Request) { // #G
	if req.URL.Path != "/" { // #H
		http.NotFound(res, req) // #H
		return                  // #H
	} // #H
	fmt.Fprint(res, "The homepage.")
}
