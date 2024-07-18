package main

import (
	"html/template"
	"net/http"
)

type comment struct { // # A
	Username string
	Text     string
}

type Page struct {
	Title, Content string
	Comments       []comment // # A
}

func routeComments(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
		Comments: []comment{
			{Username: "Bill", Text: "Looks like a good example."},
			{Username: "Jill", Text: "I really enjoyed this article."},
			{Username: "Phil", Text: "I don’t like to read."},
		}, // # B
	}
	t := template.Must(template.ParseFiles("templates/list.html"))
	if err := t.Execute(w, p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/comments", routeComments)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
