package main

import (
	"html/template"
	"log"
	"net/http"
)

type comment struct { // #A
	Username string
	Text     string
}

type Page struct {
	Title, Content string
	Comments       []comment // #A
}

var t = template.New("templates")

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
	if err := t.ExecuteTemplate(w, "list.html", p); err != nil { // #B
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func init() {
	_, err := t.ParseGlob("templates/*.html") // #C
	if err != nil {
		log.Fatal("Error loading templates:" + err.Error())
	}
}

func main() {
	http.HandleFunc("/comments", routeComments)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
