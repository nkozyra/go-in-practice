package main
import (
     "bytes"
     "html/template"
     "net/http"
)
var t *template.Template // # A
var qc template.HTML // # A
func init() {
     t = template.Must(template.ParseFiles("index.html", "quote.html")) // # B
}
type Page struct { // # C
     Title   string // # C
     Content template.HTML // # C
} // # C
type Quote struct { // # C
     Quote, Name string // # C
} // # C
func main() {
     q := &Quote{ // # D
            Quote: `You keep using that word. I do not think // # D
                    it means what you think it means.`, // # D
            Person: “Inigo Montoya”, // # D
     }
     var b bytes.Buffer  // # E
     t.ExecuteTemplate(&b, “quote.html”, q)
     qc = template.HTML(b.String())  // # F
     http.HandleFunc(“/”, displayPage) // # G
     http.ListenAndServe(":8080", nil) // # G
}
func displayPage(w http.ResponseWriter, r *http.Request) {
     p := &Page{ // # H
            Title:   "A User", // # H
            Content: qc,  // # H
     } // # H
     t.ExecuteTemplate(w, "index.html", p) // # I
}
