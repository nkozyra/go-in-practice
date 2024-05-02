package main

import (
	"html/template"
	"net/http"
)

type comment struct {
	Username string
	Text     string
}

type Page struct {
	Title, Content string
	Comments       []comment
}

var t = template.Must(template.ParseGlob("templates/*.html"))

func routeComments(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
		Comments: []comment{
			{Username: "Bill", Text: "Looks like a good example."},
			{Username: "Jill", Text: "I really enjoyed this article."},
			{Username: "Phil", Text: "I don’t like to read."},
		},
	}

	if err := t.ExecuteTemplate(w, "list.html", p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/comments", routeComments)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
