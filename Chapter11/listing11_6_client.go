package main

import "net/http"

func init() {
}

type retryableRequest struct {
	url          string
	maxRetries   int
	delaySeconds int
}

func requestUntil(url string, ops retryableRequest) (http.Response, error) {
	for i := 0; i < ops.maxRetries; i++ {
		// if succeeds, early return
	}
	return nil, nil
}

func main() {
	if err := requestUntil(requestUntil({
		url:          "http://localhost:8080",
		maxRetries:   3,
		delaySeconds: 5,
	})); err != nil {
		panic(err)
	}
}
