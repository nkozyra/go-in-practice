package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/readme.txt")
}
