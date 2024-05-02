package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const MAX_TIMEOUTS = 5

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	response, err := SendRequest("Hello")

	if err == ErrTimeout {
		timeouts := 0
		for err == ErrTimeout {
			timeouts++
			fmt.Println("Timeout. Retrying.")
			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts!")
			}
			response, err = SendRequest("Hello")
		}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
func SendRequest(req string) (string, error) {
	switch rand.Intn(3) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
