package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("templates/simple.html")) // # A
type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t.Execute(w, p) // # B
}
func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
