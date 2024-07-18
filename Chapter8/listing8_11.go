package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type comment struct {
	username   string
	text       string
	dateString string
}

var comments []comment

var validUsers = map[string]string{ // # A
	"bill": "abc123",
}

func login(username, password string) bool { // # B
	if validPassword, ok := validUsers[username]; ok {
		return validPassword == password
	}
	return false
}

func main() {
	http.HandleFunc("GET /comments", getComments)
	http.HandleFunc("POST /comments", postComments)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func postComments(w http.ResponseWriter, r *http.Request) {
	username, password, auth := r.BasicAuth() // # C

	if !auth || !login(username, password) { // # D
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	commentText, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	comments = append(comments, comment{username: username, text: string(commentText), dateString: time.Now().Format(time.RFC3339)})
	w.WriteHeader(http.StatusOK)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("%s (%s) - @%s\n", comments[i].text, comments[i].dateString, comments[i].username)
	}
	fmt.Fprintln(w, fmt.Sprintf("Comments: \n%s", commentBody))
}
