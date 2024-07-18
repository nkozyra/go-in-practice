package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("Type anything below for up to 30 seconds")
	go echo(os.Stdin, os.Stdout) // # A
	time.Sleep(30 * time.Second) // # B
	fmt.Println("Timed out.")    // # C
	os.Exit(0)                   // # D
}
func echo(in io.Reader, out io.Writer) { // # E
	io.Copy(out, in) // # F
}
