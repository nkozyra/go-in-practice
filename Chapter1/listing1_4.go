package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/") // #A
	if err != nil {
		log.Fatal("could not retrieve example.com", err) // #B
	}
	defer resp.Body.Close() // #C

	body, err := io.ReadAll(resp.Body) // #D
	if err != nil {
		log.Fatal("could not read body", err) // #B
	}
	log.Println(string(body)) //#E
}
