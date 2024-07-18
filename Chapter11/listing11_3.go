package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cc := &http.Client{Timeout: time.Second}     // # A
	res, err := cc.Get("http://www.manning.com") // # B
	if err != nil {                              // # C
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s", b)
}
