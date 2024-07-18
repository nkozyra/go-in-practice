package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	var b bytes.Buffer      // # A
	err := t.Execute(&b, p) // # A
	if err != nil {         // # B
		fmt.Fprint(w, "A error occured.") // # B
		return
	}
	b.WriteTo(w) // # C
}
func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
