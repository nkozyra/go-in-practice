package main

import (
	"html/template"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML> 
<html>
  <head>
    <meta charset="utf-8">
    <title>Date Example</title>
  </head>
  <body>
       <p>{{.Date | dateFormat "Jan 2, 2006"}}</p> 
  </body>
</html>`

var funcMap = template.FuncMap{ // # C
	"dateFormat": dateFormat, // # C
} // # C

func dateFormat(layout string, d time.Time) string { // # D
	return d.Format(layout) // # D
} // # D

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")         // # E
	t.Funcs(funcMap)                  // # F
	t.Parse(tpl)                      // # G
	data := struct{ Date time.Time }{ // # H
		Date: time.Now(), // # H
	} // # H
	t.Execute(res, data) // # I
}

func main() {
	http.HandleFunc("/", serveTemplate) // # J
	http.ListenAndServe(":8080", nil)   // # J
}
