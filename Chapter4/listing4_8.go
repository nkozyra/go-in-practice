package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const MAX_TIMEOUTS = 5

var ErrTimeout = errors.New("The request timed out")     // #A
var ErrRejected = errors.New("The request was rejected") // #B

func init() {
	rand.Seed(time.Now().UnixNano()) // #C
}

func main() {
	response, err := SendRequest("Hello") // #D

	if errors.Is(err, ErrTimeout) {
		timeouts := 0
		for errors.Is(err, ErrTimeout) { // #E
			timeouts++
			fmt.Println("Timeout. Retrying.") // #E
			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts!")
			}
			response, err = SendRequest("Hello") // #E
		}
	}

	if err != nil { // #F
		fmt.Println(err) // #F
	} else { // #G
		fmt.Println(response) // #G
	} // #G
}
func SendRequest(req string) (string, error) { // #H
	switch rand.Intn(3) % 3 { // #I
	case 0: // #J
		return "Success", nil // #I
	case 1: // #J
		return "", ErrRejected // #I
	default: // #J
		return "", ErrTimeout // #K
	} // #K
}
