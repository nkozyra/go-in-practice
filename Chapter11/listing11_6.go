package main

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden) // # A
}
func main() {
	http.HandleFunc("/", displayError) // # B
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
