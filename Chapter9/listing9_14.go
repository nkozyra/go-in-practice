package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template // # A
func init() {
	t = make(map[string]*template.Template)                              // # B
	temp := template.Must(template.ParseFiles("base.html", "user.html")) // # C
	t["user.html"] = temp                                                // # C
	temp = template.Must(template.ParseFiles("base.html", "page.html"))  // # C
	t["page.html"] = temp                                                // # C
}

type Page struct { // # D
	Title, Content string // # D
}                  // # D
type User struct { // # D
	Username, Name string // # D
} // # D
func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{ // # E
		Title:   "An Example",                   // # E
		Content: "Have fun stormin’ da castle.", // # E
	} // # E
	t["page.html"].ExecuteTemplate(w, "base", p) // # F
}
func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "swordsmith",
		Name:     "Inigo Montoya",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}
func main() {
	http.HandleFunc("/user", displayUser) // # G
	http.HandleFunc("/", displayPage)     // # G
	http.ListenAndServe(":8080", nil)     // # G
}
