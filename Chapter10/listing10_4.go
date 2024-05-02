package main

import "net/http"

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)
	if err := http.ListenAndServe(":8080", pr); err != nil {
		panic(err)
	}
}
