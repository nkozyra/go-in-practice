package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dest, err := os.Create("test2.txt")
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		panic(err)
	}

}
