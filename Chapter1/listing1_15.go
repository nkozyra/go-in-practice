package main

import ( // #B
	"fmt"      // #B
	"net/http" // #B
)

func hello(w http.ResponseWriter, r *http.Request) { // #C
	fmt.Fprint(w, "Hello, my name is Inigo Montoya") // #C
}

func main() { // #D
	http.HandleFunc("/", hello)                // #D
	http.ListenAndServe("localhost:4000", nil) // #D
}
