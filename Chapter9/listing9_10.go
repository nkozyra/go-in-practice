package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html")) // # A
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t.ExecuteTemplate(w, "index.html", p) // # B
}
func main() {
	http.HandleFunc("/", displayPage) // # C
	http.ListenAndServe(":8080", nil) // # C
}
