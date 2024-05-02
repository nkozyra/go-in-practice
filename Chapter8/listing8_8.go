package main

import (
	"fmt"
	"log"
	"net/http"
)

type customMuxRoute struct {
	pattern string
	handler http.Handler
}

type customMux struct {
	routes map[string][]customMuxRoute
}

func (c *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(c.routes)
	h := c.match(r.URL.Path, r.Method)
	h.ServeHTTP(w, r)
}

func (c *customMux) addRoute(pattern, method string, handler http.HandlerFunc) {
	c.routes[method] = append(c.routes[method], customMuxRoute{pattern, handler})
}

func (c *customMux) match(pattern, method string) http.Handler {
	if _, ok := c.routes[method]; !ok {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		})
	}

	for _, v := range c.routes[method] {
		if v.pattern == pattern {
			return v.handler
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}

func newCustomMux() *customMux {
	return &customMux{
		routes: make(map[string][]customMuxRoute),
	}
}

func main() {
	myMux := newCustomMux()
	myMux.addRoute("/comments", "GET", commentsGetHandler)
	myMux.addRoute("/comments", "POST", commentsPostHandler)

	if err := http.ListenAndServe(":8085", myMux); err != nil {
		panic(err)
	}
}

func commentsGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Comments home page!")
}

func commentsPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Comments post page!")
}
