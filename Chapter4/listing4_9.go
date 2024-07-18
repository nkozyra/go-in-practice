package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrTimeout = errors.New("The request timed out") // #A

func SendRequest(req string) (string, error) {
	return "", fmt.Errorf("we got an error: %w", ErrTimeout) // #B
}

func main() {
	if _, err := SendRequest("Hello"); err != nil {
		if err == ErrTimeout {
			log.Println("we got a timeout error") // #C
		} else {
			log.Println("we got some other error") // #D
		}
	}
}
