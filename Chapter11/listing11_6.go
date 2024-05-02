package main

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}
func main() {
	http.HandleFunc("/", displayError)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
