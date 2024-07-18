package main

import (
	"context"
	"fmt"
	"net/http"
)

type config struct {
	HomepageDescription string
	Pageviews           int64
}

func main() {
	c := config{ // # A
		HomepageDescription: "my 1997-style personal web site",
		Pageviews:           0,
	}
	ctx, _ := context.WithCancel(context.Background()) // # B
	ctx = context.WithValue(ctx, "webConfig", c)       // # C

	http.HandleFunc("/home", homeHandler(ctx)) // # D
	http.HandleFunc("/guestbook", guestbookHandler(ctx))
	http.ListenAndServe(":8081", nil)
}

func homeHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config) // # D
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf("welcome to %s", myValue.HomepageDescription))
	}
}

func guestbookHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config) // # D
	return func(w http.ResponseWriter, r *http.Request) {
		myValue.Pageviews++
		fmt.Fprintln(w, fmt.Sprintf("welcome to my guestbook. hit counter since server restart: %v", myValue.Pageviews))
	}
}
