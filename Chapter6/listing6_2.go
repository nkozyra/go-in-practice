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
	c := config{
		HomepageDescription: "my 1997-style personal web site",
		Pageviews:           0,
	}
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "webConfig", c)

	http.HandleFunc("/home", homeHandler(ctx))
	http.HandleFunc("/guestbook", guestbookHandler(ctx))
	http.ListenAndServe(":8081", nil)
}

func homeHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf("welcome to %s", myValue.HomepageDescription))
	}
}

func guestbookHandler(ctx context.Context) http.HandlerFunc {
	myValue := ctx.Value("webConfig").(config)
	return func(w http.ResponseWriter, r *http.Request) {
		myValue.Pageviews++
		fmt.Fprintln(w, fmt.Sprintf("welcome to my guestbook. hit counter since server restart: %v", myValue.Pageviews))
	}
}
