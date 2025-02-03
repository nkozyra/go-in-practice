package main

import "net/http"

func main() {
	usersRouter := http.NewServeMux()
	commentsRouter := http.NewServeMux()

	usersRouter.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("New user form"))
	})

	commentsRouter.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("New comment form"))
	})

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/users/", http.StripPrefix("/users", usersRouter))
	mainRouter.Handle("/comments/", http.StripPrefix("/comments", commentsRouter))
	if err := http.ListenAndServe(":8085", mainRouter); err != nil {
		panic(err)
	}
}
