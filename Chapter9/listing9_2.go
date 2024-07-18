package main

import (
	"html/template" // # A
	"net/http"
)

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{ // # B
		Title:   "An Example",                   // # B
		Content: "Have fun stormin’ da castle.", // # B
	} // # B
	t := template.Must(template.ParseFiles("templates/simple.html")) // # C
	t.Execute(w, p)                                                  // # D
}
func main() {
	http.HandleFunc("/", displayPage) // # E
	http.ListenAndServe(":8080", nil) // # E
}
