package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type comment struct { // # A
	text       string
	dateString string
}

var comments []comment

func getComments(w http.ResponseWriter, r *http.Request) { // # B
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("%s (%s)\n", comments[i].text, comments[i].dateString)
	}
	fmt.Fprintln(w, fmt.Sprintf("Comments: \n%s", commentBody))
}

func postComments(w http.ResponseWriter, r *http.Request) { // # B
	commentText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	comments = append(comments, comment{text: string(commentText), dateString: time.Now().Format(time.RFC3339)})
	w.WriteHeader(http.StatusOK)
}

func main() {

	http.HandleFunc("GET /comments", getComments)
	http.HandleFunc("POST /comments", postComments)

	if err := http.ListenAndServe(":8004", nil); err != nil {
		panic(err)
	}
}
