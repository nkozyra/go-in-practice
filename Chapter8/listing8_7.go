package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type comment struct {
	username   string
	text       string
	dateString string
}

var comments []comment

var validUsers = map[string]string{
	"bill": "abc123",
}

func login(username, password string) bool {
	if validPassword, ok := validUsers[username]; ok {
		return validPassword == password
	}
	return false
}

func main() {
	http.HandleFunc("/comments", routeComments)

	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

func routeComments(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postComments(w, r)
		return
	}
	if r.Method == http.MethodGet {
		getComments(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func postComments(w http.ResponseWriter, r *http.Request) {
	username, password, auth := r.BasicAuth()

	if !auth || !login(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println(password)
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
