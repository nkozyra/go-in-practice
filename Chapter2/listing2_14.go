package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	handleFunc := newHandler() // #A
	server := &http.Server{    // #A
		Addr:    ":8080",
		Handler: handleFunc,
	}

	ch := make(chan os.Signal, 1)            // #B
	signal.Notify(ch, os.Interrupt, os.Kill) // #B

	go func() {
		server.ListenAndServe() // #C
	}()

	time.Sleep(5 * time.Second) // #D
	<-ch                        // #D
	if err := server.Shutdown(nil); err != nil {
		panic(err) // #D
	}
}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) { // #E
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
